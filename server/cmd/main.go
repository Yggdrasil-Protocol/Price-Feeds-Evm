package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	pricefeed "github.com/Yggdrasil-Protocol/EVM-Worker/PriceFeedContract" // Update this import path
)

// Config holds the application configuration
type Config struct {
	RPCUrl           string
	PrivateKey       string
	ContractAddress  string
	UpdateInterval   string
	MetricsPort      string
	HealthCheckPort  string
	LogLevel         string
	PriceAPIEndpoint string
}

// Global variables
var (
	logger *zap.Logger
	config Config

	updateCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "price_feed_updates_total",
			Help: "Total number of price feed updates",
		},
	)

	updateDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "price_feed_update_duration_seconds",
			Help:    "Duration of price feed updates in seconds",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func main() {
	// Load configuration
	if err := loadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	if err := initLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Initialize Prometheus metrics
	prometheus.MustRegister(updateCounter)
	prometheus.MustRegister(updateDuration)

	// Connect to Ethereum client
	client, err := ethclient.Dial(config.RPCUrl)
	if err != nil {
		logger.Fatal("Failed to connect to the Ethereum client", zap.Error(err))
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		logger.Fatal("Failed to load private key", zap.Error(err))
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Fatal("Failed to get public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	logger.Info("Loaded account", zap.String("address", fromAddress.Hex()))

	address := common.HexToAddress(config.ContractAddress)
	instance, err := pricefeed.NewPriceFeedContract(address, client)
	if err != nil {
		logger.Fatal("Failed to instantiate contract", zap.Error(err))
	}

	// Create a new cron scheduler
	c := cron.New(cron.WithSeconds())

	// Schedule the updatePriceFeed function
	_, err = c.AddFunc(config.UpdateInterval, func() {
		if err := updatePriceFeed(client, instance, privateKey, fromAddress); err != nil {
			logger.Error("Failed to update price feed", zap.Error(err))
		}
	})
	if err != nil {
		logger.Fatal("Failed to schedule updatePriceFeed", zap.Error(err))
	}

	// Start the cron scheduler
	c.Start()

	// Start metrics server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		logger.Info("Starting metrics server", zap.String("port", config.MetricsPort))
		if err := http.ListenAndServe(":"+config.MetricsPort, nil); err != nil {
			logger.Error("Metrics server failed", zap.Error(err))
		}
	}()

	// Start health check server
	go func() {
		http.HandleFunc("/healthz", healthCheckHandler)
		logger.Info("Starting health check server", zap.String("port", config.HealthCheckPort))
		if err := http.ListenAndServe(":"+config.HealthCheckPort, nil); err != nil {
			logger.Error("Health check server failed", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Graceful shutdown
	logger.Info("Shutting down gracefully...")
	ctx := c.Stop()
	<-ctx.Done()
	logger.Info("Shutdown complete")
}

func updatePriceFeed(client *ethclient.Client, instance *pricefeed.PriceFeedContract, privateKey *ecdsa.PrivateKey, fromAddress common.Address) error {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		updateDuration.Observe(duration)
	}()

	// Fetch latest prices from API
	prices, err := fetchPrices()
	if err != nil {
		return fmt.Errorf("failed to fetch prices: %v", err)
	}

	// Prepare the data for updatePrice function
	assets := make([][32]byte, len(prices))
	decimals := make([]uint8, len(prices))
	priceValues := make([]*big.Int, len(prices))

	for i, price := range prices {
		assets[i] = nameToKeccak256(price.Asset)
		decimals[i] = price.Decimals
		priceValues[i] = price.Price
	}

	// Get the nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}

	// Create the transaction signer
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // adjust based on your contract's requirements
	auth.GasPrice = gasPrice

	// Call the updatePrice function
	tx, err := instance.UpdatePrice(auth, assets, decimals, priceValues)
	if err != nil {
		return fmt.Errorf("failed to update price: %v", err)
	}

	logger.Info("Transaction sent", zap.String("hash", tx.Hash().Hex()))

	// Wait for the transaction to be mined
	receipt, err := waitForTransaction(context.Background(), client, tx.Hash())
	if err != nil {
		return fmt.Errorf("failed to get transaction receipt: %v", err)
	}

	logger.Info("Transaction mined", zap.Uint64("block", receipt.BlockNumber.Uint64()))
	updateCounter.Inc()

	return nil
}

func nameToKeccak256(name string) [32]byte {
	hash := crypto.Keccak256([]byte(name))
	var result [32]byte
	copy(result[:], hash)
	return result
}

func waitForTransaction(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err != nil {
			if err == ethereum.NotFound {
				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				case <-time.After(time.Second):
					continue
				}
			}
			return nil, fmt.Errorf("failed to get transaction receipt: %v", err)
		}
		if receipt != nil {
			return receipt, nil
		}
	}
}

func loadConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	config = Config{
		RPCUrl:           os.Getenv("RPC_URL"),
		PrivateKey:       os.Getenv("PRIVATE_KEY"),
		ContractAddress:  os.Getenv("CONTRACT_ADDRESS"),
		UpdateInterval:   os.Getenv("UPDATE_INTERVAL"),
		MetricsPort:      os.Getenv("METRICS_PORT"),
		HealthCheckPort:  os.Getenv("HEALTH_CHECK_PORT"),
		LogLevel:         os.Getenv("LOG_LEVEL"),
		PriceAPIEndpoint: os.Getenv("PRICE_API_ENDPOINT"),
	}

	if config.RPCUrl == "" || config.PrivateKey == "" || config.ContractAddress == "" {
		return fmt.Errorf("missing required environment variables")
	}

	if config.UpdateInterval == "" {
		config.UpdateInterval = "*/5 * * * * *" // Default to every 5 seconds
	}

	if config.MetricsPort == "" {
		config.MetricsPort = "9090" // Default metrics port
	}

	if config.HealthCheckPort == "" {
		config.HealthCheckPort = "8080" // Default health check port
	}

	if config.LogLevel == "" {
		config.LogLevel = "info" // Default log level
	}

	return nil
}

func initLogger() error {
	logConfig := zap.NewProductionConfig()
	logLevel, err := zap.ParseAtomicLevel(config.LogLevel)
	if err != nil {
		return fmt.Errorf("invalid log level: %v", err)
	}
	logConfig.Level = logLevel

	var logErr error
	logger, logErr = logConfig.Build()
	if logErr != nil {
		return fmt.Errorf("failed to initialize logger: %v", logErr)
	}
	return nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

type PriceData struct {
	Asset    string
	Decimals uint8
	Price    *big.Int
}

func fetchPrices() ([]PriceData, error) {
	resp, err := http.Get(config.PriceAPIEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []struct {
		Asset    string `json:"asset"`
		Decimals uint8  `json:"decimals"`
		Price    string `json:"price"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	prices := make([]PriceData, len(data))
	for i, item := range data {
		price, ok := new(big.Int).SetString(item.Price, 10)
		if !ok {
			return nil, fmt.Errorf("invalid price for asset %s: %s", item.Asset, item.Price)
		}
		prices[i] = PriceData{
			Asset:    item.Asset,
			Decimals: item.Decimals,
			Price:    price,
		}
	}

	return prices, nil
}

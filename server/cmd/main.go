package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Yggdrasil-Protocol/PriceFeedServer/PriceFeedContract"
)

const (
	rpcURL             = "https://rpc.ankr.com/polygon_amoy"
	privateKeyHex      = "49f841619c9ba5edaf2a5eb7aa8c146a5649b4b02aac462dccf3d02e990fb662"
	contractAddressHex = "0xE5C5CE5f8EDbF88E1663b4ADE74ED13d369c89B9"
	workerCount        = 5
)

var (
	client   *ethclient.Client
	contract *PriceFeedContract.PriceFeedContract
	auth     *bind.TransactOpts
	nonceMu  sync.Mutex
	nonce    uint64
	chainID  *big.Int
)

type PriceFeed struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
}

func main() {
	var err error
	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	chainID, err = client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddressHex)
	contract, err := PriceFeedContract.NewPriceFeedContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	auth, err := getAuthTransactor(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to create auth transactor: %v", err)
	}

	// Initialize nonce
	nonce, err = client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatalf("Failed to get initial nonce: %v", err)
	}

	updateMultiplePriceFeeds(contract, auth, preparePriceFeeds())
}

func preparePriceFeeds() []PriceFeedContract.IPriceFeedPrice {
	return []PriceFeedContract.IPriceFeedPrice{
		{Pair: "HOT-USD", Price: big.NewInt(100000), Decimals: big.NewInt(8)},
		{Pair: "ZIL-USD", Price: big.NewInt(2000000), Decimals: big.NewInt(8)},
		{Pair: "BAT-USD", Price: big.NewInt(25000000), Decimals: big.NewInt(8)},
		{Pair: "RVN-USD", Price: big.NewInt(2000000), Decimals: big.NewInt(8)},
		{Pair: "THETA-USD", Price: big.NewInt(80000000), Decimals: big.NewInt(8)},
		// Add more feeds as needed
	}
}

func updateMultiplePriceFeeds(contract *PriceFeedContract.PriceFeedContract, auth *bind.TransactOpts, priceFeeds []PriceFeedContract.IPriceFeedPrice) {
	workChan := make(chan PriceFeedContract.IPriceFeedPrice, len(priceFeeds))
	resultChan := make(chan struct{}, len(priceFeeds))

	// Start workers
	for i := 0; i < workerCount; i++ {
		go worker(contract, auth, workChan, resultChan)
	}

	// Distribute work
	for _, feed := range priceFeeds {
		workChan <- feed
	}
	close(workChan)

	// Wait for results
	for i := 0; i < len(priceFeeds); i++ {
		<-resultChan
	}

	fmt.Println("All price feeds updated.")
}

func worker(contract *PriceFeedContract.PriceFeedContract, auth *bind.TransactOpts, workChan <-chan PriceFeedContract.IPriceFeedPrice, resultChan chan<- struct{}) {
	for feed := range workChan {
		updatePriceFeed(contract, auth, feed)
		resultChan <- struct{}{}
	}
}

func updatePriceFeed(contract *PriceFeedContract.PriceFeedContract, auth *bind.TransactOpts, feed PriceFeedContract.IPriceFeedPrice) {
	contractAbi, err := abi.JSON(strings.NewReader(PriceFeedContract.PriceFeedContractABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	callData, err := contractAbi.Pack("updatePriceFeed", feed)
	if err != nil {
		log.Fatalf("Failed to pack data for updatePriceFeed: %v", err)
	}

	nonceMu.Lock()
	defer nonceMu.Unlock()

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddressHex)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasPrice,
		GasFeeCap: auth.GasPrice,
		Gas:       auth.GasLimit,
		To:        &contractAddress,
		Value:     big.NewInt(0),
		Data:      callData,
	})
	nonce++

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Price feed update transaction sent: %s\n", signedTx.Hash().Hex())
	waitForTx(signedTx.Hash(), feed)
}

func waitForTx(txHash common.Hash, feed PriceFeedContract.IPriceFeedPrice) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == ethereum.NotFound {
			fmt.Println("Transaction not yet mined. Waiting...")
			continue
		}
		if err != nil {
			log.Printf("Error getting transaction receipt: %v", err)
			return
		}
		fmt.Printf("Transaction mined in block: %d\n", receipt.BlockNumber)
		fmt.Printf("Updated %s: Price = %s, Decimals = %s\n", feed.Pair, feed.Price.String(), feed.Decimals.String())
		return
	}
}

func getAuthTransactor(privateKeyHex string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	auth.GasLimit = uint64(3000000) // Increased gas limit for multiple calls
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	return auth, nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron/v3"

	"github.com/Yggdrasil-Protocol/PriceFeedServer/PriceFeedContract"
)

const (
	rpcURL             = "https://rpc.ankr.com/polygon_amoy"
	privateKeyHex      = "49f841619c9ba5edaf2a5eb7aa8c146a5649b4b02aac462dccf3d02e990fb662"
	contractAddressHex = "0xE5C5CE5f8EDbF88E1663b4ADE74ED13d369c89B9"
)

var (
	client   *ethclient.Client
	contract *PriceFeedContract.PriceFeedContract
	auth     *bind.TransactOpts
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

	contractAddress := common.HexToAddress(contractAddressHex)
	contract, err := PriceFeedContract.NewPriceFeedContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	auth, err := getAuthTransactor(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to create auth transactor: %v", err)
	}

	c := cron.New()
	_, err = c.AddFunc("@every 15s", func() {
		priceFeeds := preparePriceFeeds()
		updatePriceFeedsAtomically(contract, auth, priceFeeds)
	})
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	c.Start()
	select {} // Block forever
}

func preparePriceFeeds() []PriceFeed {
	return []PriceFeed{
		{Pair: "HOT-USD", Price: big.NewInt(100000), Decimals: big.NewInt(8)},
		{Pair: "BTC-USD", Price: big.NewInt(3200000), Decimals: big.NewInt(8)},
		{Pair: "ETH-USD", Price: big.NewInt(250000), Decimals: big.NewInt(8)},
		{Pair: "ADA-USD", Price: big.NewInt(150000), Decimals: big.NewInt(8)},
		{Pair: "LINK-USD", Price: big.NewInt(2300), Decimals: big.NewInt(8)},
		{Pair: "DOT-USD", Price: big.NewInt(30000), Decimals: big.NewInt(8)},
		{Pair: "UNI-USD", Price: big.NewInt(2800), Decimals: big.NewInt(8)},
		{Pair: "AAVE-USD", Price: big.NewInt(40000), Decimals: big.NewInt(8)},
		{Pair: "MATIC-USD", Price: big.NewInt(15000), Decimals: big.NewInt(8)},
		{Pair: "SOL-USD", Price: big.NewInt(40000), Decimals: big.NewInt(8)},
		{Pair: "XRP-USD", Price: big.NewInt(7500), Decimals: big.NewInt(8)},
		{Pair: "LTC-USD", Price: big.NewInt(180000), Decimals: big.NewInt(8)},
		{Pair: "BCH-USD", Price: big.NewInt(220000), Decimals: big.NewInt(8)},
		{Pair: "DOGE-USD", Price: big.NewInt(300), Decimals: big.NewInt(8)},
		{Pair: "AVAX-USD", Price: big.NewInt(120000), Decimals: big.NewInt(8)},
		{Pair: "SHIB-USD", Price: big.NewInt(800), Decimals: big.NewInt(8)},
		{Pair: "ATOM-USD", Price: big.NewInt(8500), Decimals: big.NewInt(8)},
		{Pair: "FIL-USD", Price: big.NewInt(4500), Decimals: big.NewInt(8)},
		{Pair: "XLM-USD", Price: big.NewInt(1100), Decimals: big.NewInt(8)},
		{Pair: "TRX-USD", Price: big.NewInt(900), Decimals: big.NewInt(8)},
		{Pair: "XMR-USD", Price: big.NewInt(160000), Decimals: big.NewInt(8)},
		{Pair: "EOS-USD", Price: big.NewInt(4000), Decimals: big.NewInt(8)},
		{Pair: "NEO-USD", Price: big.NewInt(3500), Decimals: big.NewInt(8)},
		{Pair: "XTZ-USD", Price: big.NewInt(3000), Decimals: big.NewInt(8)},
		{Pair: "MANA-USD", Price: big.NewInt(800), Decimals: big.NewInt(8)},
		{Pair: "SAND-USD", Price: big.NewInt(700), Decimals: big.NewInt(8)},
		{Pair: "THETA-USD", Price: big.NewInt(650), Decimals: big.NewInt(8)},
		{Pair: "FTM-USD", Price: big.NewInt(200), Decimals: big.NewInt(8)},
		{Pair: "HBAR-USD", Price: big.NewInt(300), Decimals: big.NewInt(8)},
		{Pair: "VET-USD", Price: big.NewInt(150), Decimals: big.NewInt(8)},
	}
}

func updatePriceFeedsAtomically(contract *PriceFeedContract.PriceFeedContract, auth *bind.TransactOpts, priceFeeds []PriceFeed) {
	contractAbi, err := abi.JSON(strings.NewReader(PriceFeedContract.PriceFeedContractABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// Prepare the encoded data for each updatePriceFeed call
	var allCallData []byte
	feed := priceFeeds
	callData, err := contractAbi.Pack("updatePriceFeed", feed)
	if err != nil {
		log.Fatalf("Failed to pack data for updatePriceFeed: %v", err)
	}
	allCallData = append(allCallData, callData...)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddressHex)

	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), auth.GasLimit, gasPrice, allCallData)

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Atomic update transaction sent: %s\n", signedTx.Hash().Hex())

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		log.Fatalf("Failed to wait for transaction to be mined: %v", err)
	}

	fmt.Printf("Transaction mined in block: %d\n", receipt.BlockNumber)
}

func getAuthTransactor(privateKeyHex string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
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

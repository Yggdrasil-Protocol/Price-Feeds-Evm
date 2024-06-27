package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/DeonLabs/PriceFeedServer/PriceFeedContract"
)

const (
	rpcURL             = "https://rpc.ankr.com/polygon_amoy"
	privateKeyHex      = "49f841619c9ba5edaf2a5eb7aa8c146a5649b4b02aac462dccf3d02e990fb662"
	contractAddressHex = "0xE5C5CE5f8EDbF88E1663b4ADE74ED13d369c89B9"
)

var client *ethclient.Client

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

	// Example 1: Get multiple price feeds concurrently
	getPriceFeeds(contract, []string{"ETH-USD", "BTC-USD", "MATIC-USD"})

	// Example 2: Update a price feed
	updatePriceFeed(contract, auth)

	// Example 3: Request multiple price feeds and wait for events
	requestPriceFeeds(contract, auth, []string{"ETH-USD", "BTC-USD", "MATIC-USD"})
}

func getPriceFeeds(contract *PriceFeedContract.PriceFeedContract, pairs []string) {
	var wg sync.WaitGroup
	results := make(chan struct {
		pair string
		feed struct {
			Pair     string
			Price    *big.Int
			Decimals *big.Int
		}
		err error
	}, len(pairs))

	for _, pair := range pairs {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			feed, err := contract.Feed(&bind.CallOpts{}, p)
			results <- struct {
				pair string
				feed struct {
					Pair     string
					Price    *big.Int
					Decimals *big.Int
				}
				err error
			}{p, feed, err}
		}(pair)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.err != nil {
			log.Printf("Failed to get feed for %s: %v\n", result.pair, result.err)
		} else {
			fmt.Printf("Feed for %s: Price = %s, Decimals = %s\n", result.pair, result.feed.Price.String(), result.feed.Decimals.String())
		}
	}
}

func waitForTx(txHash common.Hash) (*types.Receipt, error) {
	ctx := context.Background()
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err != nil {
			if err == ethereum.NotFound {
				time.Sleep(time.Second)
				continue
			}
			return nil, err
		}
		return receipt, nil
	}
}

func updatePriceFeed(contract *PriceFeedContract.PriceFeedContract, auth *bind.TransactOpts) {
	updatePrice := PriceFeedContract.IPriceFeedPrice{
		Pair:     "MATIC-USD",
		Price:    big.NewInt(200000000000), // $2000.00000000
		Decimals: big.NewInt(8),
	}
	tx, err := contract.UpdatePriceFeed(auth, updatePrice)
	if err != nil {
		log.Fatalf("Failed to update price feed: %v", err)
	}
	fmt.Printf("Price feed update transaction sent: %s\n", tx.Hash().Hex())

	receipt, err := waitForTx(tx.Hash())
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}
	fmt.Printf("Price feed update transaction mined in block: %d\n", receipt.BlockNumber)
}

func requestPriceFeeds(contract *PriceFeedContract.PriceFeedContract, auth *bind.TransactOpts, pairs []string) {
	request := PriceFeedContract.IPriceFeedRequest{
		Pair: pairs,
	}
	tx, err := contract.RequestPriceFeed(auth, request)
	if err != nil {
		log.Fatalf("Failed to request price feed: %v", err)
	}
	fmt.Printf("Price feed request transaction sent: %s\n", tx.Hash().Hex())

	receipt, err := waitForTx(tx.Hash())
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}
	fmt.Printf("Price feed request transaction mined in block: %d\n", receipt.BlockNumber)

	watchPriceFeedRequested(contract, auth.From, receipt.BlockNumber)
}

func watchPriceFeedRequested(contract *PriceFeedContract.PriceFeedContract, requester common.Address, fromBlock *big.Int) {
	events := make(chan *PriceFeedContract.PriceFeedContractPriceFeedRequested)
	startBlock := fromBlock.Uint64()
	sub, err := contract.WatchPriceFeedRequested(&bind.WatchOpts{Start: &startBlock}, events, []common.Address{requester})
	if err != nil {
		log.Fatalf("Failed to watch for PriceFeedRequested events: %v", err)
	}
	defer sub.Unsubscribe()

	select {
	case event := <-events:
		fmt.Printf("PriceFeedRequested event:\n")
		fmt.Printf("  Requester: %s\n", event.Requester.Hex())
		for i, pair := range event.Pairs {
			fmt.Printf("  Pair: %s, Price: %s, Decimals: %s\n", pair, event.Prices[i].String(), event.Decimals[i].String())
		}
	case err := <-sub.Err():
		log.Fatalf("Error in event subscription: %v", err)
	case <-time.After(2 * time.Minute):
		log.Println("Timeout waiting for PriceFeedRequested event")
	}
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

	auth.GasLimit = uint64(300000) // Consider using client.EstimateGas in production
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	return auth, err
}

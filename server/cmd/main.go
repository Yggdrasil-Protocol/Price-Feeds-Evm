package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Yggdrasil-Protocol/EVM-Worker/PriceFeedContract"
)

const (
	rpcURL             = "https://rpc.ankr.com/polygon_amoy"
	privateKeyHex      = "49f841619c9ba5edaf2a5eb7aa8c146a5649b4b02aac462dccf3d02e990fb662"
	contractAddressHex = "0x4e7169Dfa03B8ec4862315B5a56711C84775530f"
)

var (
	client   *ethclient.Client
	contract *PriceFeedContract.PriceFeedContract
	auth     *bind.TransactOpts
)

func main() {
	// Connect to the Ethereum network
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

}

func updatePriceFeeds() {
	assets := [][32]byte{
		[32]byte{0x01},
		[32]byte{0x02},
	}
	decimals := []uint8{8, 8}
	prices := []*big.Int{
		big.NewInt(100000000), // $1.00
		big.NewInt(200000000), // $2.00
	}

	// Create the message hash
	messageHash := crypto.Keccak256(
		crypto.Keccak256(assets...),
		crypto.Keccak256(common.BigToHash(prices[0]).Bytes(), common.BigToHash(prices[1]).Bytes()),
		crypto.Keccak256([]byte{decimals[0]}, []byte{decimals[1]}),
	)

	// Sign the message hash
	signature, err := crypto.Sign(messageHash, privateKey)
	if err != nil {
		log.Printf("Failed to sign message: %v", err)
		return
	}

	// Get the latest nonce for the from address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Failed to get nonce: %v", err)
		return
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Failed to get gas price: %v", err)
		return
	}

	// Create the transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1)) // Replace 1 with your chain ID
	if err != nil {
		log.Printf("Failed to create transactor: %v", err)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Call the updatePrice function
	tx, err := contract.UpdatePrice(auth, assets, decimals, prices, signature)
	if err != nil {
		log.Printf("Failed to update price: %v", err)
		// returnate the price feed array in  a single transaction

		log.Printf("Price update transaction sent: %s", tx.Hash().Hex())

		// Wait for the transaction to be mined
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			log.Printf("Failed to get transaction receipt: %v", err)
			return
		}

		log.Printf("Price update transaction mined in block %d", receipt.BlockNumber.Uint64())
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

	auth.GasLimit = uint64(3000000) // Increased gas limit for multiple calls
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	return auth, nil
}

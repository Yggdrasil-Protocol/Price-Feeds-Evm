package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	pricefeed "github.com/Yggdrasil-Protocol/EVM-Worker/PriceFeed" // Update this import path
)

const (
	rpcURL          = "https://rpc-amoy.polygon.technology/"
	privateKey      = "49f841619c9ba5edaf2a5eb7aa8c146a5649b4b02aac462dccf3d02e990fb662"
	contractAddress = "0xF7E614226b190C7b8b3f1d410D2f5C6b103073bA"
)

func main() {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Load the private key
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Get the public address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to get public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Create a new instance of the contract
	address := common.HexToAddress(contractAddress)
	instance, err := pricefeed.NewPriceFeed(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	// Prepare the data for updatePrice function
	assetNames := []string{
		"BTC", "ETH", "LINK", "ADA", "DOT",
		"XRP", "SOL", "DOGE", "UNI", "AAVE",
	}
	assets := make([][32]byte, len(assetNames))
	for i, name := range assetNames {
		assets[i] = nameToKeccak256(name)
	}

	decimals := []uint8{8, 18, 18, 6, 10, 6, 9, 8, 18, 18}
	prices := []*big.Int{
		big.NewInt(3000000000000),       // BTC: 30,000 USD with 8 decimals
		big.NewInt(20000000000000000),   // ETH: 2,000 USD with 18 decimals
		big.NewInt(25000000000000000),   // LINK: 25 USD with 18 decimals
		big.NewInt(500000),              // ADA: 0.5 USD with 6 decimals
		big.NewInt(2000000000),          // DOT: 20 USD with 10 decimals
		big.NewInt(700000),              // XRP: 0.7 USD with 6 decimals
		big.NewInt(40000000000),         // SOL: 40 USD with 9 decimals
		big.NewInt(7000000000),          // DOGE: 0.07 USD with 8 decimals
		big.NewInt(5000000000000000000), // UNI: 5 USD with 18 decimals
		big.NewInt(80000000000000000),   // AAVE: 80 USD with 18 decimals
	}

	// Generate the proper signature

	if err != nil {
		log.Fatalf("Failed to generate signature: %v", err)
	}

	// Get the nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	// Increase gas price by 20% to ensure quick mining
	// gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(120))
	// gasPrice = new(big.Int).Div(gasPrice, big.NewInt(100))

	// Create the transaction signer
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // adjust based on your contract's requirements
	auth.GasPrice = gasPrice

	// Check if we have enough balance
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	gasLimit := new(big.Int).SetUint64(auth.GasLimit)
	gasCost := new(big.Int).Mul(gasPrice, gasLimit)
	if balance.Cmp(gasCost) < 0 {
		log.Fatalf("Insufficient funds for gas. Have %s, need %s", balance.String(), gasCost.String())
	}

	// Call the updatePrice function
	tx, err := instance.UpdatePrice(auth, assets, decimals, prices)
	if err != nil {
		log.Fatalf("Failed to update price: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := waitForTransaction(context.Background(), client, tx.Hash())
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}

	fmt.Printf("Transaction mined in block: %d\n", receipt.BlockNumber)

	// Read back the updated prices
	for i, asset := range assets {
		price, err := instance.Prices(&bind.CallOpts{}, asset)
		if err != nil {
			log.Printf("Failed to get price for asset %s: %v", assetNames[i], err)
			continue
		}
		fmt.Printf("Updated price for %s (%x): %s\n", assetNames[i], asset, price.String())
	}
}

func nameToKeccak256(name string) [32]byte {
	hash := crypto.Keccak256([]byte(name))
	var result [32]byte
	copy(result[:], hash)
	return result
}

func generateSignature(privateKey *ecdsa.PrivateKey, assets [][32]byte, decimals []uint8, prices []*big.Int) ([]byte, error) {
	// Encode the data
	encodedAssets, err := abi.NewType("bytes32[]", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ABI type for assets: %v", err)
	}
	encodedPrices, err := abi.NewType("uint256[]", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ABI type for prices: %v", err)
	}
	encodedDecimals, err := abi.NewType("uint8[]", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ABI type for decimals: %v", err)
	}

	arguments := abi.Arguments{
		{Type: encodedAssets},
		{Type: encodedPrices},
		{Type: encodedDecimals},
	}

	packedData, err := arguments.Pack(assets, prices, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to pack data: %v", err)
	}

	// Hash the packed data
	messageHash := crypto.Keccak256Hash(packedData)

	// Create the Ethereum signed message hash
	ethSignedMessageHash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n32%s", messageHash.Bytes())))

	// Sign the Ethereum signed message hash
	signature, err := crypto.Sign(ethSignedMessageHash.Bytes(), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %v", err)
	}

	// Convert V from 0/1 to 27/28
	signature[64] += 27

	return signature, nil
}
func waitForTransaction(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err != nil {
			if err == ethereum.NotFound {
				// Transaction not yet mined, wait and retry
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

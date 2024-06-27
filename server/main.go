package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client represents a client to interact with the PriceFeed contract
type Client struct {
	client   *ethclient.Client
	contract *PriceFeed
	auth     *bind.TransactOpts
	abi      abi.ABI
}

// NewClient creates a new instance of Client
func NewClient(ethURL, contractAddress string, auth *bind.TransactOpts) (*Client, error) {
	client, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	address := common.HexToAddress(contractAddress)
	contract, err := NewPriceFeed(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate PriceFeed contract: %w", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(PriceFeedABI))

	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	return &Client{
		client:   client,
		contract: contract,
		auth:     auth,
		abi:      contractAbi,
	}, nil
}

// GetFeed retrieves the price feed for a given pair
func (c *Client) GetFeed(ctx context.Context, pair string) (string, *big.Int, *big.Int, error) {
	feed, err := c.contract.Feed(&bind.CallOpts{Context: ctx}, pair)
	if err != nil {
		return "", nil, nil, fmt.Errorf("failed to get feed: %w", err)
	}
	return feed.Pair, feed.Price, feed.Decimals, nil
}

// RequestPriceFeed requests price feed for multiple pairs
func (c *Client) RequestPriceFeed(ctx context.Context, pairs []string) (*types.Transaction, error) {
	request := IPriceFeedRequest{Pair: pairs}
	opts := &bind.TransactOpts{
		From:    c.auth.From,
		Signer:  c.auth.Signer,
		Context: ctx,
	}
	tx, err := c.contract.RequestPriceFeed(opts, request)
	if err != nil {
		return nil, fmt.Errorf("failed to request price feed: %w", err)
	}
	return tx, nil
}

func (c *Client) GetRequestPriceFeedResult(ctx context.Context, tx *types.Transaction) (*IPriceFeedPriceResponse, error) {
	receipt, err := bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("transaction failed")
	}

	for _, log := range receipt.Logs {
		event := new(PriceFeedPriceFeedRequested)
		err = c.abi.UnpackIntoInterface(event, "PriceFeedRequested", log.Data)
		if err != nil {
			continue // This log is not the event we're looking for
		}

		// Parse indexed parameters
		if len(log.Topics) > 1 {
			event.Requester = common.HexToAddress(log.Topics[1].Hex())
		}

		return &IPriceFeedPriceResponse{
			Prices:   event.Prices,
			Decimals: event.Decimals,
		}, nil
	}

	return nil, fmt.Errorf("PriceFeedRequested event not found in transaction logs")
}

// UpdatePriceFeed updates the price feed for a given pair
func (c *Client) UpdatePriceFeed(ctx context.Context, pair string, price *big.Int, decimals *big.Int) error {
	priceData := IPriceFeedPrice{
		Pair:     pair,
		Price:    price,
		Decimals: decimals,
	}
	opts := &bind.TransactOpts{
		From:    c.auth.From,
		Signer:  c.auth.Signer,
		Context: ctx,
	}
	_, err := c.contract.UpdatePriceFeed(opts, priceData)
	if err != nil {
		return fmt.Errorf("failed to update price feed: %w", err)
	}
	return nil
}

// GetOwner retrieves the current owner of the contract
func (c *Client) GetOwner(ctx context.Context) (common.Address, error) {
	owner, err := c.contract.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get owner: %w", err)
	}
	return owner, nil
}

// TransferOwnership transfers the ownership of the contract to a new address
func (c *Client) TransferOwnership(ctx context.Context, newOwner common.Address) error {
	opts := &bind.TransactOpts{
		From:    c.auth.From,
		Signer:  c.auth.Signer,
		Context: ctx,
	}
	_, err := c.contract.TransferOwnership(opts, newOwner)
	if err != nil {
		return fmt.Errorf("failed to transfer ownership: %w", err)
	}
	return nil
}

// WatchPriceFeedUpdated sets up a watch for PriceFeedUpdated events
func (c *Client) WatchPriceFeedUpdated(ctx context.Context) (<-chan *PriceFeedPriceFeedUpdated, <-chan error) {
	events := make(chan *PriceFeedPriceFeedUpdated)
	errors := make(chan error)

	go func() {
		defer close(events)
		defer close(errors)

		watchOpts := &bind.WatchOpts{Context: ctx}
		eventChan := make(chan *PriceFeedPriceFeedUpdated)
		sub, err := c.contract.WatchPriceFeedUpdated(watchOpts, eventChan)
		if err != nil {
			errors <- fmt.Errorf("failed to watch PriceFeedUpdated events: %w", err)
			return
		}
		defer sub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				return
			case err := <-sub.Err():
				errors <- fmt.Errorf("subscription error: %w", err)
				return
			case event := <-eventChan:
				events <- event
			}
		}
	}()

	return events, errors
}

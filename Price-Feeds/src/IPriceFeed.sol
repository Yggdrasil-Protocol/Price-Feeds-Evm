// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

/**
 * @title IPriceFeed
 * @dev Interface for the PriceFeed contract.
 */
interface IPriceFeed {

    
    /**
     * @dev Struct to store price data.
     * @param pair The currency pair (e.g., "ETH/USD").
     * @param price The price of the pair.
     * @param decimals The number of decimals for the price.
     */
    struct Price { 
        string pair; 
        uint256 price; 
        uint256 decimals; 
    }

    /**
     * @dev Struct to store request data.
     * @param pair The array of currency pairs.
     */
    struct Request {
        string[] pair;
    }

    /**
     * @dev Struct to store price response data.
     * @param prices The array of prices for the requested pairs.
     * @param decimals The array of decimals for the requested pairs.
     */
    struct PriceResponse { 
        uint256[] prices; 
        uint256[] decimals; 
    } 

    /**
     * @dev Updates the price feed for a given pair.
     * @param price The price data including pair, price, and decimals.
     */
    function updatePriceFeed(Price calldata price) external;

    /**
     * @dev Publishes a new price feed for a given pair.
     * @param price The price data including pair, price, and decimals.
     */
    function publishPriceFeed(Price calldata price) external;
    
    /**
     * @dev Requests the price feed for a given set of pairs.
     * @param request The request data containing pairs.
     * @return response The response data including prices and decimals for the requested pairs.
     */
    function requestPriceFeed(Request memory request) external payable returns (PriceResponse memory response);
}

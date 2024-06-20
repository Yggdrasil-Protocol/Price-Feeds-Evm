// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import "../lib/openzeppelin-contracts/contracts/security/ReentrancyGuard.sol";
import "./IPriceFeed.sol";
import "../lib/openzeppelin-contracts/contracts/utils/math/SafeMath.sol";


/**
 * @title PriceFeed
 * @dev Manages price feeds that can be updated, published, and requested.
 */
contract PriceFeed is Ownable, IPriceFeed, ReentrancyGuard {
    using SafeMath for uint256;

    event PriceFeedUpdated(string pair, uint256 price, uint256 decimals);
    event PriceFeedPublished(string pair, uint256 price, uint256 decimals);
    event PriceFeedRequested(address indexed requester, string[] pairs, uint256[] prices, uint256[] decimals);

    mapping(string pair => Price price ) public Feed;

  /**
     * @dev Updates the price feed for a given pair.
     * @param price The price data including pair, price, and decimals.
     */
    function updatePriceFeed(Price calldata price) external override onlyOwner {
        require(bytes(price.pair).length > 0, "PriceFeed: Pair cannot be empty");
        require(price.price > 0, "PriceFeed: Price must be greater than zero");
        require(price.decimals <= 18, "PriceFeed: Invalid decimals");

        Feed[price.pair] = price;
        emit PriceFeedUpdated(price.pair, price.price, price.decimals);
    }

    /**
     * @dev Publishes a new price feed for a given pair.
     * @param price The price data including pair, price, and decimals.
     */
    function publishPriceFeed(Price calldata price) external override onlyOwner {
        require(bytes(price.pair).length > 0, "PriceFeed: Pair cannot be empty");
        require(price.price > 0, "PriceFeed: Price must be greater than zero");
        require(price.decimals <= 18, "PriceFeed: Invalid decimals");

        Feed[price.pair] = price;
        emit PriceFeedPublished(price.pair, price.price, price.decimals);
    }

    /**
     * @dev Requests the price feed for a given set of pairs.
     * @param request The request data containing pairs.
     * @return response The response data including prices and decimals for the requested pairs.
     */
    function requestPriceFeed(Request memory request) external payable override nonReentrant returns (PriceResponse memory response) {
        require(request.pair.length > 0, "PriceFeed: No pairs requested");
        require(msg.value >= 0.000001 ether, "PriceFeed: Insufficient funds");

        uint256[] memory prices = new uint256[](request.pair.length);
        uint256[] memory decimals = new uint256[](request.pair.length);

        for (uint256 i = 0; i < request.pair.length; i++) {
            Price memory priceData = Feed[request.pair[i]];
            require(bytes(priceData.pair).length > 0, "PriceFeed: Price not found");
            prices[i] = priceData.price;
            decimals[i] = priceData.decimals;
        }

        emit PriceFeedRequested(msg.sender, request.pair, prices, decimals);

        return PriceResponse(prices, decimals);
    }
}
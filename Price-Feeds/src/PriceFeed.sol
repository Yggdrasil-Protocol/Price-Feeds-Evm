// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import "../lib/openzeppelin-contracts-upgradeable/contracts/security/ReentrancyGuardUpgradeable.sol";
import "./IPriceFeed.sol";
import "../lib/openzeppelin-contracts/contracts/utils/math/SafeMath.sol";
import "../lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/UUPSUpgradeable.sol";

/**
 * @title PriceFeed
 * @dev Manages price feeds that can be updated, published, and requested.
 */
contract PriceFeed is Initializable, IPriceFeed, ReentrancyGuardUpgradeable, UUPSUpgradeable, OwnableUpgradeable {
    using SafeMath for uint256;

    event PriceFeedUpdated(string pair, uint256 price, uint8 decimals);
    event PriceFeedRequested(address indexed requester, string[] pairs, uint256[] prices, uint8[] decimals);

  
    mapping(bytes32 => Price) public Feed;

    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init();
        __ReentrancyGuard_init();
        __UUPSUpgradeable_init();
    }

    /**
     * @dev Updates the price feed for a given pair.
     * @param price The price data including pair, price, and decimals.
     */
    function updatePriceFeed(Price calldata price) external override onlyOwner {
        require(bytes(price.pair).length > 0, "PriceFeed: Pair cannot be empty");
        require(price.price > 0, "PriceFeed: Price must be greater than zero");
        require(price.decimals <= 18, "PriceFeed: Invalid decimals");

        bytes32 pairHash = keccak256(bytes(price.pair));
        Feed[pairHash] = price;

        emit PriceFeedUpdated(price.pair, price.price, price.decimals);
    }

    /**
     * @dev Updates multiple price feeds in a single transaction.
     * @param prices Array of price data including pairs, prices, and decimals.
     */
    function updateMultiplePriceFeeds(Price[] calldata prices) external onlyOwner {
        uint256 length = prices.length;

        for (uint256 i = 0; i < length; ) {
            Price calldata price = prices[i];

            // Validate input data
            require(bytes(price.pair).length > 0, "PriceFeed: Pair cannot be empty");
            require(price.price > 0, "PriceFeed: Price must be greater than zero");
            require(price.decimals <= 18, "PriceFeed: Invalid decimals");

            // Store price data in storage efficiently
            bytes32 pairHash = keccak256(bytes(price.pair));
            Feed[pairHash] = price;

            emit PriceFeedUpdated(price.pair, price.price, price.decimals);

            unchecked { i++; }
        }
    }

    /**
     * @dev Requests the price feed for a given set of pairs.
     * @param request The request data containing pairs.
     * @return response The response data including prices and decimals for the requested pairs.
     */
    function requestPriceFeed(Request memory request) external payable override nonReentrant returns (PriceResponse memory response) {
        uint256 length = request.pair.length;
        require(length > 0, "PriceFeed: No pairs requested");
        require(msg.value >= 0.000001 ether, "PriceFeed: Insufficient funds");

        uint256[] memory prices = new uint256[](length);
        uint8[] memory decimals = new uint8[](length);

        for (uint256 i = 0; i < length; ) {
            bytes32 pairHash = keccak256(bytes(request.pair[i]));
            Price memory priceData = Feed[pairHash];
            require(bytes(priceData.pair).length > 0, "PriceFeed: Price not found");
            prices[i] = priceData.price;
            decimals[i] = priceData.decimals;

            unchecked { i++; }
        }

        emit PriceFeedRequested(msg.sender, request.pair, prices, decimals);

        return PriceResponse(prices, decimals);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/PriceFeed.sol";
import "../src/IPriceFeed.sol";

/**
 * @title PriceFeedTest
 * @dev Comprehensive test suite for the PriceFeed contract using Foundry.
 */
contract PriceFeedTest is Test {
    IPriceFeed private priceFeed;
    address private owner;

    function setUp() public {
        owner = address(this);
        priceFeed = new PriceFeed();
        Ownable(address(priceFeed)).transferOwnership(owner);
    }

    /**
     * @dev Tests that the price feed can be updated correctly.
     */
    function testUpdatePriceFeed() public {
        IPriceFeed.Price memory price = IPriceFeed.Price("ETH/USD", 3000 * 10**18, 18);
        priceFeed.updatePriceFeed(price);

        (uint256 updatedPrice, uint256 decimals) = priceFeed.Feed("ETH/USD");
        assertEq(updatedPrice, 3000 * 10**18);
        assertEq(decimals, 18);
    }

    /**
     * @dev Tests that the price feed can be published correctly.
     */
    function testPublishPriceFeed() public {
        IPriceFeed.Price memory price = IPriceFeed.Price("BTC/USD", 40000 * 10**18, 18);
        priceFeed.publishPriceFeed(price);

        (uint256 publishedPrice, uint256 decimals) = priceFeed.Feed("BTC/USD");
        assertEq(publishedPrice, 40000 * 10**18);
        assertEq(decimals, 18);
    }

    /**
     * @dev Tests that the price feed can be requested and returned correctly.
     */
    function testRequestPriceFeed() public {
        IPriceFeed.Price memory price1 = IPriceFeed.Price("ETH/USD", 3000 * 10**18, 18);
        IPriceFeed.Price memory price2 = IPriceFeed.Price("BTC/USD", 40000 * 10**18, 18);
        priceFeed.updatePriceFeed(price1);
        priceFeed.updatePriceFeed(price2);

        string[] memory pairs = new string[](2);
        pairs[0] = "ETH/USD";
        pairs[1] = "BTC/USD";
        IPriceFeed.Request memory request = IPriceFeed.Request(pairs);

        vm.deal(address(this), 1 ether);
        IPriceFeed.PriceResponse memory response = priceFeed.requestPriceFeed{value: 0.000001 ether}(request);

        assertEq(response.prices[0], 3000 * 10**18);
        assertEq(response.decimals[0], 18);
        assertEq(response.prices[1], 40000 * 10**18);
        assertEq(response.decimals[1], 18);
    }

    /**
     * @dev Tests that an insufficient fund error is thrown when requesting a price feed with low ether value.
     */
    function testFailInsufficientFunds() public {
        string[] memory pairs = new string[](1);
        pairs[0] = "ETH/USD";
        IPriceFeed.Request memory request = IPriceFeed.Request(pairs);
        priceFeed.requestPriceFeed{value: 0.0000001 ether}(request);
    }

    /**
     * @dev Tests that an invalid price error is thrown when updating a price feed with an invalid price.
     */
    function testFailInvalidPrice() public {
        IPriceFeed.Price memory invalidPrice = IPriceFeed.Price("ETH/USD", type(uint256).max, 18);
        priceFeed.updatePriceFeed(invalidPrice);
    }

    /**
     * @dev Tests that an invalid decimals error is thrown when updating a price feed with invalid decimals.
     */
    function testFailInvalidDecimals() public {
        IPriceFeed.Price memory invalidPrice = IPriceFeed.Price("ETH/USD", 3000 * 10**18, 20);
        priceFeed.updatePriceFeed(invalidPrice);
    }

    /**
     * @dev Tests that only the owner can update the price feed.
     */
    function testOnlyOwnerCanUpdatePriceFeed() public {
        vm.prank(address(0x1234));
        IPriceFeed.Price memory price = IPriceFeed.Price("ETH/USD", 3000 * 10**18, 18);
        vm.expectRevert("Ownable: caller is not the owner");
        priceFeed.updatePriceFeed(price);
    }

    /**
     * @dev Tests that only the owner can publish the price feed.
     */
    function testOnlyOwnerCanPublishPriceFeed() public {
        vm.prank(address(0x1234));
        IPriceFeed.Price memory price = IPriceFeed.Price("BTC/USD", 40000 * 10**18, 18);
        vm.expectRevert("Ownable: caller is not the owner");
        priceFeed.publishPriceFeed(price);
    }

    /**
     * @dev Tests that the contract handles multiple price updates correctly.
     */
    function testMultiplePriceUpdates() public {
        IPriceFeed.Price memory price1 = IPriceFeed.Price("ETH/USD", 3000 * 10**18, 18);
        IPriceFeed.Price memory price2 = IPriceFeed.Price("BTC/USD", 40000 * 10**18, 18);
        priceFeed.updatePriceFeed(price1);
        priceFeed.updatePriceFeed(price2);

        (uint256 updatedPrice1, uint256 decimals1) = priceFeed.Feed("ETH/USD");
        (uint256 updatedPrice2, uint256 decimals2) = priceFeed.Feed("BTC/USD");
        assertEq(updatedPrice1, 3000 * 10**18);
        assertEq(decimals1, 18);
        assertEq(updatedPrice2, 40000 * 10**18);
        assertEq(decimals2, 18);
    }

    /**
     * @dev Tests that the contract handles multiple price requests correctly.
     */
    function testMultiplePriceRequests() public {
        IPriceFeed.Price memory price1 = IPriceFeed.Price("ETH/USD", 3000 * 10**18, 18);
        IPriceFeed.Price memory price2 = IPriceFeed.Price("BTC/USD", 40000 * 10**18, 18);
        priceFeed.updatePriceFeed(price1);
        priceFeed.updatePriceFeed(price2);

        string[] memory pairs = new string[](2);
        pairs[0] = "ETH/USD";
        pairs[1] = "BTC/USD";
        IPriceFeed.Request memory request = IPriceFeed.Request(pairs);

        vm.deal(address(this), 1 ether);
        IPriceFeed.PriceResponse memory response1 = priceFeed.requestPriceFeed{value: 0.000001 ether}(request);
        IPriceFeed.PriceResponse memory response2 = priceFeed.requestPriceFeed{value: 0.000001 ether}(request);

        assertEq(response1.prices[0], 3000 * 10**18);
        assertEq(response1.decimals[0], 18);
        assertEq(response1.prices[1], 40000 * 10**18);
        assertEq(response1.decimals[1], 18);

        assertEq(response2.prices[0], 3000 * 10**18);
        assertEq(response2.decimals[0], 18);
        assertEq(response2.prices[1], 40000 * 10**18);
        assertEq(response2.decimals[1], 18);
    }

    /**
     * @dev Tests that the contract reverts when a non-existent price feed is requested.
     */
    function testFailNonExistentPriceFeed() public {
        string[] memory pairs = new string[](1);
        pairs[0] = "NONEXISTENT/PAIR";
        IPriceFeed.Request memory request = IPriceFeed.Request(pairs);
        priceFeed.requestPriceFeed{value: 0.000001 ether}(request);
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/PriceFeed.sol";
import "../src/IPriceFeed.sol";

contract PriceFeedTest is Test {
    PriceFeed private priceFeed;
    address private owner;

    function setUp() public {
        owner = address(this);
        priceFeed = new PriceFeed();
        priceFeed.transferOwnership(owner);
    }

    function testUpdatePriceFeed() public {
        IPriceFeed.Price memory newPrice = IPriceFeed.Price({
            pair: "ETH/USD",
            price: 2000,
            decimals: 2
        });

        vm.prank(owner);
        priceFeed.updatePriceFeed(newPrice);

        assertEq(priceFeed.Feed("ETH/USD"), 2000 / 10 ** 2);
    }

    function testPublishPriceFeed() public {
        IPriceFeed.Price memory newPrice = IPriceFeed.Price({
            pair: "BTC/USD",
            price: 30000,
            decimals: 2
        });

        vm.prank(owner);
        priceFeed.publishPriceFeed(newPrice);

        assertEq(priceFeed.Feed("BTC/USD"), 30000 / 10 ** 2);
    }

    function testRequestPriceFeed() public {
        IPriceFeed.Price memory ethPrice = IPriceFeed.Price({
            pair: "ETH/USD",
            price: 2000,
            decimals: 2
        });

        IPriceFeed.Price memory btcPrice = IPriceFeed.Price({
            pair: "BTC/USD",
            price: 30000,
            decimals: 2
        });

        vm.prank(owner);
        priceFeed.updatePriceFeed(ethPrice);
        vm.prank(owner);
        priceFeed.updatePriceFeed(btcPrice);

        string[] memory pairs = new string[](2);
        pairs[0] = "ETH/USD";
        pairs[1] = "BTC/USD";

        uint256[] memory decimals = new uint256[](2);
        decimals[0] = 2;
        decimals[1] = 2;

        IPriceFeed.Request memory request = IPriceFeed.Request({
            pair: pairs,
            decimals: decimals
        });

        uint256 fee = 0.000001 ether;

        uint256[] memory prices = priceFeed.requestPriceFeed{value: fee}(request);

        uint256[] memory expectedPrices = new uint256[](2);
        expectedPrices[0] = 2000 / 10 ** 2;
        expectedPrices[1] = 30000 / 10 ** 2;

        for (uint256 i = 0; i < prices.length; i++) {
            assertEq(prices[i], expectedPrices[i]);
        }
    }

    function testInsufficientFunds() public {
        IPriceFeed.Request memory request = IPriceFeed.Request({
            pair: new string[](1),
            decimals: new uint256[](1)
        });
    
        request.pair[0] = "ETH/USD";
    
        vm.expectRevert("PriceFeed: Insufficient funds");
        priceFeed.requestPriceFeed{value: 0.0000001 ether}(request);
    }
}

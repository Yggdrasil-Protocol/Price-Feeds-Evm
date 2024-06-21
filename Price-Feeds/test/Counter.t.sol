// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../lib/forge-std/src/Test.sol";
import "../src/PriceFeed.sol";
import "../src/IPriceFeed.sol";

contract PriceFeedTest is Test {
    PriceFeed private priceFeed;
    address private owner;

    function setUp() public {
        owner = address(this);
        priceFeed = new PriceFeed();
        Ownable(address(priceFeed)).transferOwnership(owner);
    }

    function testUpdatePriceFeed() public {
        IPriceFeed.Price memory price = IPriceFeed.Price("ETH/USD", 3000 * 10**18, 18);
        priceFeed.updatePriceFeed(price);

        // Accessing the updated price
        ( string memory pair , uint256 uPrice , uint256 decimals ) = priceFeed.Feed("ETH/USD");
        assertEq(uPrice , 3000 * 10**18);
        assertEq(decimals, 18);
    }

      function testPublishPriceFeed() public {
        IPriceFeed.Price memory price = IPriceFeed.Price("ETH/USD", 3000 * 10**18, 18);
        priceFeed.publishPriceFeed(price);

        // Accessing the updated price
        ( string memory pair , uint256 uPrice , uint256 decimals ) = priceFeed.Feed("ETH/USD");
        assertEq(uPrice , 3000 * 10**18);
        assertEq(decimals, 18);
    }


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

        // Asserting the response prices and decimals
        assertEq(response.prices[0], 3000 * 10**18);
        assertEq(response.decimals[0], 18);
        assertEq(response.prices[1], 40000 * 10**18);
        assertEq(response.decimals[1], 18);
    }

    function testUpdatePriceFeedWithEmptyPair() public {
        IPriceFeed.Price memory price = IPriceFeed.Price("", 3000 * 10**18, 18);

        (bool success, ) = address(priceFeed).call(abi.encodeWithSelector(priceFeed.updatePriceFeed.selector, price));
        assertEq(success, false, "Update should fail with empty pair");
    }

    function testUpdatePriceFeedWithZeroPrice() public {
        IPriceFeed.Price memory price = IPriceFeed.Price("ETH/USD", 0, 18);

        (bool success, ) = address(priceFeed).call(abi.encodeWithSelector(priceFeed.updatePriceFeed.selector, price));
        assertEq(success, false, "Update should fail with zero price");
    }

    function testRequestPriceFeedWithInsufficientFunds() public {
        IPriceFeed.Price memory price1 = IPriceFeed.Price("ETH/USD", 3000 * 10**18, 18);
        priceFeed.updatePriceFeed(price1);

        string[] memory pairs = new string[](1);
        pairs[0] = "ETH/USD";
        IPriceFeed.Request memory request = IPriceFeed.Request(pairs);

        // Sending less than required ether
        (bool success, ) = address(priceFeed).call{value: 0.0000001 ether}(abi.encodeWithSelector(priceFeed.requestPriceFeed.selector, request));
        assertEq(success, false, "Request should fail with insufficient funds");
    }

    function testRequestPriceFeedWithNonExistingPair() public {
        string[] memory pairs = new string[](1);
        pairs[0] = "ETH/USD";
        IPriceFeed.Request memory request = IPriceFeed.Request(pairs);

        // Requesting for a pair that doesn't exist
        (bool success, ) = address(priceFeed).call{value: 0.000001 ether}(abi.encodeWithSelector(priceFeed.requestPriceFeed.selector, request));
        assertEq(success, false, "Request should fail for non-existing pair");
    }
}
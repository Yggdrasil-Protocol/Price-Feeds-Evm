// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/PriceFeed.sol";
import "../src/PriceFeedProxy.sol";
import "../src/IPriceFeed.sol";

contract PriceFeedProxyTest is Test {
    PriceFeed public priceFeed;
    PriceFeedProxy public proxy;
    address public owner;
    address public user;

    IPriceFeed.Price public ethPrice;
    IPriceFeed.Price public btcPrice;
    IPriceFeed.Request public request;

    struct PriceResponse {
        uint256[] prices;
        uint256[] decimals;
    }

    function setUp() public {
        owner = address(this);
        user = address(1);
        priceFeed = new PriceFeed();
        proxy = new PriceFeedProxy(address(priceFeed));

        // Transfer ownership of PriceFeed to Proxy
        priceFeed.transferOwnership(address(proxy));

        // Set up test data
        ethPrice = IPriceFeed.Price({pair: "ETH/USD", price: 2000, decimals: 18});
        btcPrice = IPriceFeed.Price({pair: "BTC/USD", price: 30000, decimals: 18});
        string[] memory pairs = new string[](2);
        pairs[0] = "ETH/USD";
        pairs[1] = "BTC/USD";
        request = IPriceFeed.Request({pair: pairs});
    }

    function testUpdatePriceFeed() public {
        vm.prank(owner);
        proxy.updatePriceFeed(ethPrice);

        (string memory pair , uint256 priceValue, uint256 decimals) = priceFeed.Feed("ETH/USD");
        assertEq(priceValue, 2000);
        assertEq(decimals, 18);
    }

    function testPublishPriceFeed() public {
        vm.prank(owner);
        proxy.publishPriceFeed(btcPrice);

        (string memory pair  , uint256 priceValue, uint256 decimals) = priceFeed.Feed("BTC/USD");
        assertEq(priceValue, 30000);
        assertEq(decimals, 18);
    }

    function testRequestPriceFeed() public {
        vm.prank(owner);
        proxy.updatePriceFeed(ethPrice);
        proxy.updatePriceFeed(btcPrice);

        vm.prank(user);
        vm.deal(user, 1 ether); // Ensure user has enough balance to send ether
        (bool success, bytes memory data) = address(proxy).call{value: 0.000001 ether}(abi.encodeWithSelector(proxy.requestPriceFeed.selector, request));
        require(success, "Request Price Feed call failed");

        PriceResponse memory response = abi.decode(data, (PriceResponse));
        assertEq(response.prices[0], 2000);
        assertEq(response.decimals[0], 18);
        assertEq(response.prices[1], 30000);
        assertEq(response.decimals[1], 18);
    }

    function testUpdatePriceFeedAddress() public {
        PriceFeed newPriceFeed = new PriceFeed();
        
        vm.prank(owner);
        proxy.updatePriceFeedAddress(address(newPriceFeed));
        
        assertEq(proxy.priceFeedAddress(), address(newPriceFeed));
    }

    function testFailUpdatePriceFeedNonOwner() public {
        vm.prank(user);
        proxy.updatePriceFeed(ethPrice);
    }

    function testFailPublishPriceFeedNonOwner() public {
        vm.prank(user);
        proxy.publishPriceFeed(btcPrice);
    }

    function testFailRequestPriceFeedInsufficientFunds() public {
        vm.prank(user);
        (bool success, ) = address(proxy).call{value: 0.0000001 ether}(abi.encodeWithSelector(proxy.requestPriceFeed.selector, request));
        require(!success, "Request Price Feed should fail with insufficient funds");
    }
}

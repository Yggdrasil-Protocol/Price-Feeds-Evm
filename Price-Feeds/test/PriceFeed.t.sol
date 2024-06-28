// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/PriceFeed.sol";
import "../src/PriceFeedProxy.sol";

contract PriceFeedTest is Test {
    PriceFeed public priceFeed;
    PriceFeedProxy public proxy;
    address public owner;

    function setUp() public {
        owner = address(this);

        // Deploy PriceFeed logic contract
        PriceFeed logic = new PriceFeed();
        bytes memory data = abi.encodeWithSignature("initialize()");

        // Deploy proxy with logic and data
        proxy = new PriceFeedProxy(address(logic), data);
        priceFeed = PriceFeed(address(proxy));

        // Verify that the owner is set correctly
        assertEq(priceFeed.owner(), owner);
    }

    function testUpdatePriceFeed() public {
        IPriceFeed.Price memory price = IPriceFeed.Price({
            pair: "ETH/USD",
            price: 3000 * 10 ** 18,
            decimals: 18
        });

        bytes32 priceHash = keccak256(abi.encodePacked(price.pair));
        // Update price feed
        priceFeed.updatePriceFeed(price);

        // Fetch the updated price feed
        (string memory pair, uint256 priceValue, uint256 decimals) = priceFeed
            .Feed(priceHash);

        console.log(pair);
        console.log(priceValue);
        console.log(decimals);

        assertEq(priceValue, 3000 * 10 ** 18);
        assertEq(decimals, 18);
    }

    function testFailUpdatePriceFeedByNonOwner() public {
        vm.prank(address(0x1));
        IPriceFeed.Price memory price = IPriceFeed.Price({
            pair: "ETH/USD",
            price: 3000 * 10 ** 18,
            decimals: 18
        });

        // Attempt to update price feed by non-owner
        priceFeed.updatePriceFeed(price);
    }

    function testRequestPriceFeed() public {
        IPriceFeed.Price memory price1 = IPriceFeed.Price({
            pair: "ETH/USD",
            price: 3000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price memory price2 = IPriceFeed.Price({
            pair: "BTC/USD",
            price: 40000 * 10 ** 18,
            decimals: 18
        });

        // Update price feeds
        priceFeed.updatePriceFeed(price1);
        priceFeed.updatePriceFeed(price2);

        // Request price feeds
        IPriceFeed.Request memory request = IPriceFeed.Request({
            pair: new string[](2)
        });
        request.pair[0] = "ETH/USD";
        request.pair[1] = "BTC/USD";

        uint256 requestFee = 0.000001 ether;
        IPriceFeed.PriceResponse memory response = priceFeed.requestPriceFeed{
            value: requestFee
        }(request);

        // Verify the response
        assertEq(response.prices[0], 3000 * 10 ** 18);
        assertEq(response.decimals[0], 18);
        assertEq(response.prices[1], 40000 * 10 ** 18);
        assertEq(response.decimals[1], 18);
    }

    function testFailRequestPriceFeedWithoutPayment() public {
        IPriceFeed.Request memory request = IPriceFeed.Request({
            pair: new string[](1)
        });
        request.pair[0] = "ETH/USD";

        // Attempt to request price feed without payment
        priceFeed.requestPriceFeed(request);
    }

    function testFailRequestPriceFeedNonExistentPair() public {
        IPriceFeed.Request memory request = IPriceFeed.Request({
            pair: new string[](1)
        });
        request.pair[0] = "NONEXISTENT/PAIR";

        uint256 requestFee = 0.000001 ether;

        // Attempt to request price feed for a non-existent pair
        priceFeed.requestPriceFeed{value: requestFee}(request);
    }

    function testMultipleUpdates() public {
        IPriceFeed.Price memory price = IPriceFeed.Price({
            pair: "ETH/USD",
            price: 3000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price memory price2 = IPriceFeed.Price({
            pair: "SOL/USD",
            price: 70000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price memory price3 = IPriceFeed.Price({
            pair: "BTC/USD",
            price: 300000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price memory price4 = IPriceFeed.Price({
            pair: "MTC/USD",
            price: 6900000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price memory price5 = IPriceFeed.Price({
            pair: "RTC/USD",
            price: 1000000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price[] memory prices = new IPriceFeed.Price[](5);
        prices[0] = price;
        prices[1] = price2;
        prices[2] = price3;
        prices[3] = price4;
        prices[4] = price5;

        priceFeed.updateMultiplePriceFeeds(prices);

        (string memory pair, uint256 priceValue, uint256 decimals) = priceFeed
            .Feed(keccak256(abi.encodePacked("ETH/USD")));
        (
            string memory pair2,
            uint256 priceValue2,
            uint256 decimals2
        ) = priceFeed.Feed(keccak256(abi.encodePacked("SOL/USD")));
        (
            string memory pair3,
            uint256 priceValue3,
            uint256 decimals3
        ) = priceFeed.Feed(keccak256(abi.encodePacked("BTC/USD")));
        (
            string memory pair4,
            uint256 priceValue4,
            uint256 decimals4
        ) = priceFeed.Feed(keccak256(abi.encodePacked("MTC/USD")));
        (
            string memory pair5,
            uint256 priceValue5,
            uint256 decimals5
        ) = priceFeed.Feed(keccak256(abi.encodePacked("RTC/USD")));

        assertEq(pair, "ETH/USD");
        assertEq(priceValue, 3000 * 10 ** 18);
        assertEq(decimals, 18);

        assertEq(pair2, "SOL/USD");
        assertEq(priceValue2, 70000 * 10 ** 18);
        assertEq(decimals2, 18);

        assertEq(pair3, "BTC/USD");
        assertEq(priceValue3, 300000 * 10 ** 18);
        assertEq(decimals3, 18);

        assertEq(pair4, "MTC/USD");
        assertEq(priceValue4, 6900000 * 10 ** 18);
        assertEq(decimals4, 18);

        assertEq(pair5, "RTC/USD");
        assertEq(priceValue5, 1000000 * 10 ** 18);
        assertEq(decimals5, 18);
    }

    function testMultipleUpdatesInASingleUpdate() public {
        IPriceFeed.Price memory price = IPriceFeed.Price({
            pair: "ETH/USD",
            price: 3000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price memory price2 = IPriceFeed.Price({
            pair: "SOL/USD",
            price: 70000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price memory price3 = IPriceFeed.Price({
            pair: "BTC/USD",
            price: 300000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price memory price4 = IPriceFeed.Price({
            pair: "MTC/USD",
            price: 6900000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price memory price5 = IPriceFeed.Price({
            pair: "RTC/USD",
            price: 1000000 * 10 ** 18,
            decimals: 18
        });

        IPriceFeed.Price[] memory prices = new IPriceFeed.Price[](5);
        prices[0] = price;
        prices[1] = price2;
        prices[2] = price3;
        prices[3] = price4;
        prices[4] = price5;

        for (uint8 i = 0; i < 5; i++) {
            priceFeed.updatePriceFeed(prices[i]);
        }

        (string memory pair, uint256 priceValue, uint256 decimals) = priceFeed
            .Feed(keccak256(abi.encodePacked("ETH/USD")));
        (
            string memory pair2,
            uint256 priceValue2,
            uint256 decimals2
        ) = priceFeed.Feed(keccak256(abi.encodePacked("SOL/USD")));
        (
            string memory pair3,
            uint256 priceValue3,
            uint256 decimals3
        ) = priceFeed.Feed(keccak256(abi.encodePacked("BTC/USD")));
        (
            string memory pair4,
            uint256 priceValue4,
            uint256 decimals4
        ) = priceFeed.Feed(keccak256(abi.encodePacked("MTC/USD")));

        (
            string memory pair5,
            uint256 priceValue5,
            uint256 decimals5
        ) = priceFeed.Feed(keccak256(abi.encodePacked("RTC/USD")));
      

        assertEq(pair, "ETH/USD");
        assertEq(priceValue, 3000 * 10 ** 18);
        assertEq(decimals, 18);

        assertEq(pair2, "SOL/USD");
        assertEq(priceValue2, 70000 * 10 ** 18);
        assertEq(decimals2, 18);

        assertEq(pair3, "BTC/USD");
        assertEq(priceValue3, 300000 * 10 ** 18);
        assertEq(decimals3, 18);

        assertEq(pair4, "MTC/USD");
        assertEq(priceValue4, 6900000 * 10 ** 18);
        assertEq(decimals4, 18);

        assertEq(pair5, "RTC/USD");
        assertEq(priceValue5, 1000000 * 10 ** 18);
        assertEq(decimals5, 18);

       
    }
}

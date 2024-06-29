// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/PriceFeed.sol";
import "../src/PriceFeedProxy.sol";
import "../lib/openzeppelin-contracts-upgradeable/contracts/utils/cryptography/ECDSAUpgradeable.sol";
import "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";

contract PriceFeedTest is Test {
    PriceFeed priceFeedImplementation;
    PriceFeedProxy priceFeedProxy;
    PriceFeed priceFeed;

    address owner = address(1);
    address user = address(2);
    address trustedSigner = address(3);

    bytes32 asset1 = keccak256(abi.encodePacked("ASSET1"));
    uint256 price1 = 1000;
    uint8 decimals1 = 18;

    bytes32 asset2 = keccak256(abi.encodePacked("ASSET2"));
    uint256 price2 = 2000;
    uint8 decimals2 = 18;

    function setUp() public {
        bytes memory data = abi.encodeWithSelector(priceFeedImplementation.initialize.selector, 100);
        priceFeedProxy = new PriceFeedProxy(address(priceFeedImplementation), data);
        priceFeed = PriceFeed(payable(address(priceFeedProxy)));

        // Set owner
        vm.prank(owner);
        priceFeed.transferOwnership(owner);

        // Set trusted signer
        vm.prank(owner);
        priceFeed.setTrustedSigner(trustedSigner);
    }

    function testInitialize() public {
        assertEq(priceFeed.feePerAsset(), 100);
        assertEq(priceFeed.trustedSigner(), trustedSigner);
    }

    function testSetTrustedSigner() public {
        vm.prank(owner);
        priceFeed.setTrustedSigner(user);
        assertEq(priceFeed.trustedSigner(), user);
    }

    function testSetFeePerAsset() public {
        vm.prank(owner);
        priceFeed.setFeePerAsset(200);
        assertEq(priceFeed.feePerAsset(), 200);
    }

    function testUpdatePrice() public {
        bytes32[] memory assets = new bytes32[](2);
        assets[0] = asset1;
        assets[1] = asset2;

        uint8[] memory decimals = new uint8[](2);
        decimals[0] = decimals1;
        decimals[1] = decimals2;

        uint256[] memory prices = new uint256[](2);
        prices[0] = price1;
        prices[1] = price2;

       

        vm.prank(trustedSigner);
        priceFeed.updatePrice(assets, decimals, prices);

        
    }

    function testRequestPrices() public {
        bytes32[] memory assets = new bytes32[](2);
        assets[0] = asset1;
        assets[1] = asset2;

        uint8[] memory decimals = new uint8[](2);
        decimals[0] = decimals1;
        decimals[1] = decimals2;

        uint256[] memory prices = new uint256[](2);
        prices[0] = price1;
        prices[1] = price2;

     

        vm.prank(trustedSigner);
        priceFeed.updatePrice(assets, decimals, prices);

        function(uint8[] memory, uint256[] memory) external callback = this.testCallback;
        bytes memory callData = abi.encodeWithSelector(callback.selector, decimals, prices);

        vm.prank(user);
        priceFeed.requestPrices{value: 200}(assets, callback);
    }

    function testCallback(uint8[] memory _decimals, uint256[] memory _prices) public {
        assertEq(_decimals[0], decimals1);
        assertEq(_prices[0], price1);
        assertEq(_decimals[1], decimals2);
        assertEq(_prices[1], price2);
    }

    function testWithdraw() public {
        vm.deal(address(priceFeed), 1000);

        vm.prank(owner);
        priceFeed.withdraw();

        assertEq(address(priceFeed).balance, 0);
        assertEq(owner.balance, 1000);
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/PriceFeed.sol";

contract PriceFeedTest is Test {
    PriceFeed public priceFeed;
    address public owner;
    address public trustedSigner;
    uint256 public initialFeePerAsset;

    function setUp() public {
        owner = address(this);
        trustedSigner = makeAddr("trustedSigner");
        initialFeePerAsset = 0.01 ether;
        
        priceFeed = new PriceFeed(initialFeePerAsset);
        priceFeed.setTrustedSigner(trustedSigner);
    }

    function testInitialState() public {
        assertEq(priceFeed.owner(), owner);
        assertEq(priceFeed.trustedSigner(), trustedSigner);
        assertEq(priceFeed.feePerAsset(), initialFeePerAsset);
    }

    function testSetTrustedSigner() public {
        address newSigner = makeAddr("newSigner");
        priceFeed.setTrustedSigner(newSigner);
        assertEq(priceFeed.trustedSigner(), newSigner);
    }

    function testSetTrustedSignerRevertOnZeroAddress() public {
        vm.expectRevert(PriceFeed.ZeroAddress.selector);
        priceFeed.setTrustedSigner(address(0));
    }

    function testSetFeePerAsset() public {
        uint256 newFee = 0.02 ether;
        priceFeed.setFeePerAsset(newFee);
        assertEq(priceFeed.feePerAsset(), newFee);
    }

    function testUpdatePrice() public {
        bytes32[] memory assets = new bytes32[](2);
        assets[0] = keccak256("ETH");
        assets[1] = keccak256("BTC");

        uint8[] memory decimals = new uint8[](2);
        decimals[0] = 18;
        decimals[1] = 8;

        uint256[] memory prices = new uint256[](2);
        prices[0] = 2000 * 10**18; // $2000 for ETH
        prices[1] = 30000 * 10**8; // $30000 for BTC

        bytes32 messageHash = keccak256(abi.encodePacked(assets, prices, decimals));
        bytes32 ethSignedMessageHash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(uint256(keccak256(abi.encodePacked("trustedSigner"))), ethSignedMessageHash);
        bytes memory signature = abi.encodePacked(r, s, v);

        priceFeed.updatePrice(assets, decimals, prices, signature);

        assertEq(priceFeed.prices(assets[0]), prices[0]);
        assertEq(priceFeed.prices(assets[1]), prices[1]);
    }

    function testUpdatePriceInvalidSignature() public {
        bytes32[] memory assets = new bytes32[](1);
        assets[0] = keccak256("ETH");

        uint8[] memory decimals = new uint8[](1);
        decimals[0] = 18;

        uint256[] memory prices = new uint256[](1);
        prices[0] = 2000 * 10**18;

        bytes memory invalidSignature = new bytes(65);

        vm.expectRevert(PriceFeed.InvalidSignature.selector);
        priceFeed.updatePrice(assets, decimals, prices, invalidSignature);
    }

    function testRequestPrices() public {
        // First, update some prices
        bytes32[] memory assets = new bytes32[](2);
        assets[0] = keccak256("ETH");
        assets[1] = keccak256("BTC");

        uint8[] memory decimals = new uint8[](2);
        decimals[0] = 18;
        decimals[1] = 8;

        uint256[] memory prices = new uint256[](2);
        prices[0] = 2000 * 10**18;
        prices[1] = 30000 * 10**8;

        bytes32 messageHash = keccak256(abi.encodePacked(assets, prices, decimals));
        bytes32 ethSignedMessageHash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(uint256(keccak256(abi.encodePacked("trustedSigner"))), ethSignedMessageHash);
        bytes memory signature = abi.encodePacked(r, s, v);

        priceFeed.updatePrice(assets, decimals, prices, signature);

        // Now request prices
        uint256 requiredFee = priceFeed.feePerAsset() * assets.length;
        MockPriceFeedReceiver receiver = new MockPriceFeedReceiver();

        vm.deal(address(this), requiredFee);
        priceFeed.requestPrices{value: requiredFee}(assets, receiver.receivePrice);

        assertEq(receiver.receivedPrices(0), prices[0]);
        assertEq(receiver.receivedPrices(1), prices[1]);
        assertEq(receiver.receivedDecimals(0), decimals[0]);
        assertEq(receiver.receivedDecimals(1), decimals[1]);
    }

    function testRequestPricesInsufficientFee() public {
        bytes32[] memory assets = new bytes32[](2);
        assets[0] = keccak256("ETH");
        assets[1] = keccak256("BTC");

        uint256 requiredFee = priceFeed.feePerAsset() * assets.length;
        MockPriceFeedReceiver receiver = new MockPriceFeedReceiver();

        vm.expectRevert(PriceFeed.TransferFailed.selector);
        priceFeed.requestPrices{value: requiredFee - 1}(assets, receiver.receivePrice);
    }

    function testRequestPricesTooManyAssets() public {
        bytes32[] memory assets = new bytes32[](101); // MAX_ASSETS_PER_REQUEST is 100
        for (uint i = 0; i < 101; i++) {
            assets[i] = keccak256(abi.encodePacked(i));
        }

        MockPriceFeedReceiver receiver = new MockPriceFeedReceiver();

        vm.expectRevert(abi.encodeWithSelector(PriceFeed.TooManyAssets.selector, 101, 100));
        priceFeed.requestPrices(assets, receiver.receivePrice);
    }

    function testWithdraw() public {
        // Send some Ether to the contract
        vm.deal(address(priceFeed), 1 ether);

        uint256 initialBalance = address(this).balance;
        priceFeed.withdraw();
        uint256 finalBalance = address(this).balance;

        assertEq(finalBalance - initialBalance, 1 ether);
        assertEq(address(priceFeed).balance, 0);
    }


    // Needed to receive Ether when calling withdraw()
    receive() external payable {}
}


    // Helper contract to test price feed callback
    contract MockPriceFeedReceiver  {
        uint256[] public receivedPrices;
        uint8[] public receivedDecimals;

        function receivePrice(uint256[] memory _prices, uint8[] memory _decimals) external override {
            receivedPrices = _prices;
            receivedDecimals = _decimals;
        }
    }
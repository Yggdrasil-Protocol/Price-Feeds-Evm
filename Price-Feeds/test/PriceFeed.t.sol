// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/PriceFeed.sol";
import "../src/IPriceFeedReciever.sol";
import "forge-std/console.sol";
import "../lib/openzeppelin-contracts/contracts/utils/math/SafeMath.sol";

contract MockConsumer {
    uint256[] public resprices;
    PriceFeed public priceFeed;

    function _saveFeeds(uint256[] calldata prices) external {
        resprices = prices;
    }

    function requestPrices(
        PriceFeed _priceFeed,
        bytes32[] calldata assets
    ) external {
        priceFeed = _priceFeed;
        _priceFeed.requestPrices(assets, this._saveFeeds);
    }
}

contract PriceFeedTest is Test {
    PriceFeed public priceFeed;
    address public trustedSigner;
    uint256 public initialFeePerAsset;
    using SafeMath for uint256;

    function setUp() public {
        trustedSigner = vm.addr(0x1);
        initialFeePerAsset = 0.01 ether;
        priceFeed = new PriceFeed(trustedSigner, initialFeePerAsset);
    }

    function testConstructor() public {
        assertEq(priceFeed.trustedSigner(), trustedSigner);
        assertEq(priceFeed.feePerAsset(), initialFeePerAsset);
    }

    function testSetTrustedSigner() public {
        address newSigner = vm.addr(2);
        vm.prank(priceFeed.owner());
        priceFeed.setTrustedSigner(newSigner);
        assertEq(priceFeed.trustedSigner(), newSigner);
    }

    function testSetTrustedSignerZeroAddress() public {
        vm.prank(priceFeed.owner());
        vm.expectRevert(PriceFeed.ZeroAddress.selector);
        priceFeed.setTrustedSigner(address(0));
    }

    function testSetFeePerAsset() public {
        uint256 newFee = 0.02 ether;
        vm.prank(priceFeed.owner());
        priceFeed.setFeePerAsset(newFee);
        assertEq(priceFeed.feePerAsset(), newFee);
    }

    function testUpdatePrice() public {
        bytes32[] memory assets = new bytes32[](100);
        assets[0] = keccak256("ETH");
        assets[1] = keccak256("BTC");
        assets[2] = keccak256("LINK");
        assets[3] = keccak256("AAVE");
        assets[4] = keccak256("UNI");
        assets[5] = keccak256("YFI");
        assets[6] = keccak256("SUSHI");
        assets[7] = keccak256("COMP");
        assets[8] = keccak256("MKR");
        assets[9] = keccak256("SNX");

        assets[10] = keccak256("ADA");
        assets[11] = keccak256("DOT");
        assets[12] = keccak256("SOL");
        assets[13] = keccak256("MATIC");
        assets[14] = keccak256("ATOM");
        assets[15] = keccak256("FTM");
        assets[16] = keccak256("AVAX");
        assets[17] = keccak256("LUNA");
        assets[18] = keccak256("ALGO");
        assets[19] = keccak256("VET");

        assets[20] = keccak256("THETA");
        assets[21] = keccak256("XTZ");
        assets[22] = keccak256("TRX");
        assets[23] = keccak256("NEO");
        assets[24] = keccak256("XMR");
        assets[25] = keccak256("EOS");
        assets[26] = keccak256("IOTA");
        assets[27] = keccak256("KSM");
        assets[28] = keccak256("ZEC");
        assets[29] = keccak256("BCH");

        assets[30] = keccak256("LTC");
        assets[31] = keccak256("ZIL");
        assets[32] = keccak256("ENJ");
        assets[33] = keccak256("MANA");
        assets[34] = keccak256("GRT");
        assets[35] = keccak256("BAL");
        assets[36] = keccak256("CRV");
        assets[37] = keccak256("1INCH");
        assets[38] = keccak256("RUNE");
        assets[39] = keccak256("BAT");

        assets[40] = keccak256("ANKR");
        assets[41] = keccak256("STORJ");
        assets[42] = keccak256("UMA");
        assets[43] = keccak256("BAND");
        assets[44] = keccak256("REN");
        assets[45] = keccak256("CVC");
        assets[46] = keccak256("RSR");
        assets[47] = keccak256("OCEAN");
        assets[48] = keccak256("MLN");
        assets[49] = keccak256("LRC");

        assets[50] = keccak256("KNC");
        assets[51] = keccak256("REP");
        assets[52] = keccak256("REQ");
        assets[53] = keccak256("STMX");
        assets[54] = keccak256("UOS");
        assets[55] = keccak256("GNO");
        assets[56] = keccak256("PERP");
        assets[57] = keccak256("SRM");
        assets[58] = keccak256("DYDX");
        assets[59] = keccak256("FTT");

        assets[60] = keccak256("CAKE");
        assets[61] = keccak256("XVS");
        assets[62] = keccak256("BEL");
        assets[63] = keccak256("MDX");
        assets[64] = keccak256("BAKE");
        assets[65] = keccak256("ALPHA");
        assets[66] = keccak256("DF");
        assets[67] = keccak256("HNT");
        assets[68] = keccak256("CTK");
        assets[69] = keccak256("DIA");

        assets[70] = keccak256("ORN");
        assets[71] = keccak256("LINA");
        assets[72] = keccak256("SPI");
        assets[73] = keccak256("UBT");
        assets[74] = keccak256("WNXM");
        assets[75] = keccak256("AKRO");
        assets[76] = keccak256("BELT");
        assets[77] = keccak256("DEGO");
        assets[78] = keccak256("AUDIO");
        assets[79] = keccak256("TRB");

        assets[80] = keccak256("AKT");
        assets[81] = keccak256("CHR");
        assets[82] = keccak256("PHA");
        assets[83] = keccak256("BNT");
        assets[84] = keccak256("ELF");
        assets[85] = keccak256("ANT");
        assets[86] = keccak256("POLY");
        assets[87] = keccak256("BOND");
        assets[88] = keccak256("DODO");
        assets[89] = keccak256("QNT");

        assets[90] = keccak256("RLC");
        assets[91] = keccak256("CFX");
        assets[92] = keccak256("REEF");
        assets[93] = keccak256("OXT");
        assets[94] = keccak256("COTI");
        assets[95] = keccak256("SKL");
        assets[96] = keccak256("GHST");
        assets[97] = keccak256("MIR");
        assets[98] = keccak256("FEI");
        assets[99] = keccak256("RGT");

        uint256[] memory prices = new uint256[](100);
        prices[0] = 2000 ether;
        prices[1] = 50000 ether;
        prices[2] = 30 ether;
        prices[3] = 300 ether;
        prices[4] = 20 ether;
        prices[5] = 30000 ether;
        prices[6] = 10 ether;
        prices[7] = 300 ether;
        prices[8] = 2000 ether;
        prices[9] = 100 ether;

        prices[10] = 1.5 ether;
        prices[11] = 15 ether;
        prices[12] = 40 ether;
        prices[13] = 1.2 ether;
        prices[14] = 10 ether;
        prices[15] = 0.25 ether;
        prices[16] = 25 ether;
        prices[17] = 80 ether;
        prices[18] = 0.9 ether;
        prices[19] = 0.12 ether;

        prices[20] = 6 ether;
        prices[21] = 2.5 ether;
        prices[22] = 0.07 ether;
        prices[23] = 35 ether;
        prices[24] = 200 ether;
        prices[25] = 4 ether;
        prices[26] = 1 ether;
        prices[27] = 300 ether;
        prices[28] = 100 ether;
        prices[29] = 500 ether;

        prices[30] = 150 ether;
        prices[31] = 0.08 ether;
        prices[32] = 1.5 ether;
        prices[33] = 0.5 ether;
        prices[34] = 0.75 ether;
        prices[35] = 12 ether;
        prices[36] = 2.5 ether;
        prices[37] = 3 ether;
        prices[38] = 8 ether;
        prices[39] = 0.25 ether;

        prices[40] = 0.1 ether;
        prices[41] = 0.5 ether;
        prices[42] = 10 ether;
        prices[43] = 7 ether;
        prices[44] = 0.5 ether;
        prices[45] = 0.03 ether;
        prices[46] = 0.02 ether;
        prices[47] = 0.7 ether;
        prices[48] = 60 ether;
        prices[49] = 0.25 ether;

        prices[50] = 1 ether;
        prices[51] = 15 ether;
        prices[52] = 0.5 ether;
        prices[53] = 0.02 ether;
        prices[54] = 0.3 ether;
        prices[55] = 100 ether;
        prices[56] = 3 ether;
        prices[57] = 6 ether;
        prices[58] = 12 ether;
        prices[59] = 25 ether;

        prices[60] = 2 ether;
        prices[61] = 20 ether;
        prices[62] = 1 ether;
        prices[63] = 2 ether;
        prices[64] = 1.5 ether;
        prices[65] = 0.25 ether;
        prices[66] = 0.3 ether;
        prices[67] = 15 ether;
        prices[68] = 3 ether;
        prices[69] = 2 ether;

        prices[70] = 5 ether;
        prices[71] = 0.15 ether;
        prices[72] = 50 ether;
        prices[73] = 0.8 ether;
        prices[74] = 50 ether;
        prices[75] = 0.01 ether;
        prices[76] = 0.3 ether;
        prices[77] = 0.25 ether;
        prices[78] = 1.5 ether;
        prices[79] = 30 ether;

        prices[80] = 2 ether;
        prices[81] = 0.5 ether;
        prices[82] = 1 ether;
        prices[83] = 2.5 ether;
        prices[84] = 0.5 ether;
        prices[85] = 4 ether;
        prices[86] = 0.03 ether;
        prices[87] = 40 ether;
        prices[88] = 1 ether;
        prices[89] = 70 ether;

        prices[90] = 3 ether;
        prices[91] = 0.25 ether;
        prices[92] = 0.02 ether;
        prices[93] = 0.5 ether;
        prices[94] = 0.05 ether;
        prices[95] = 0.08 ether;
        prices[96] = 0.9 ether;
        prices[97] = 1 ether;
        prices[98] = 1 ether;
        prices[99] = 15 ether;
        bytes32 messageHash = keccak256(abi.encodePacked(assets, prices));

        // Create the message hash

        // Create the Ethereum Signed Message hash
        bytes32 ethSignedMessageHash = keccak256(
            abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash)
        );

        // Sign the message
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(0x1, ethSignedMessageHash);

        // Concatenate the signature components
        bytes memory signature = abi.encodePacked(r, s, v);

        // Update the price
        priceFeed.updatePrice(assets, prices, signature);

        // Verify the price was updated
        assertEq(priceFeed.prices(assets[0] = keccak256("AAVE")), prices[3]);
    }

    // function testUpdatePriceInvalidSignature() public {
    //     bytes32 asset = keccak256("ETH");
    //     uint256 price = 2000 ether;
    //     uint256 timestamp = block.timestamp;

    //     bytes32 messageHash = keccak256(
    //         abi.encodePacked(asset, price, timestamp)
    //     );
    //     (uint8 v, bytes32 r, bytes32 s) = vm.sign(2, messageHash); // Wrong signer
    //     bytes memory signature = abi.encodePacked(r, s, v);

    //     vm.expectRevert(PriceFeed.InvalidSignature.selector);
    //     priceFeed.updatePrice(asset, price, signature);
    // }

    // function testRequestPrices() public {
    //     setUp();
    //     updatePriceHelper(keccak256("ETH"), 2000 ether);
    //     updatePriceHelper(keccak256("BTC"), 50000 ether);
    //     updatePriceHelper(keccak256("LINK"), 30 ether);

    //     bytes32[] memory assets = new bytes32[](3);
    //     assets[0] = keccak256("ETH");
    //     assets[1] = keccak256("BTC");
    //     assets[2] = keccak256("LINK");

    //     MockConsumer consumer = new MockConsumer();

    //     consumer.requestPrices(priceFeed, assets);

    //     assertEq(consumer.resprices(0), 2000 ether);
    //     assertEq(consumer.resprices(1), 50000 ether);
    //     assertEq(consumer.resprices(2), 30 ether);

    // }

    function testWithdraw() public {
        uint256 amount = 1 ether;
        vm.deal(address(priceFeed), amount);

        uint256 initialBalance = priceFeed.owner().balance;
        vm.prank(priceFeed.owner());
        priceFeed.withdraw();
        uint256 finalBalance = priceFeed.owner().balance;

        assertEq(finalBalance - initialBalance, amount);
        assertEq(address(priceFeed).balance, 0);
    }

    // function updatePriceHelper(bytes32 asset, uint256 price) internal {
    //     bytes32 messageHash = keccak256(abi.encodePacked(asset, price));
    //     bytes32 ethSignedMessageHash = keccak256(
    //         abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash)
    //     );
    //     (uint8 v, bytes32 r, bytes32 s) = vm.sign(0x1, ethSignedMessageHash);
    //     bytes memory signature = abi.encodePacked(r, s, v);
    //     priceFeed.updatePrice(asset, price, signature);
    // }
}

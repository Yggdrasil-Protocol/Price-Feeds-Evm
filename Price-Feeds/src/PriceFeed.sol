// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../lib/openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";
import "../lib/openzeppelin-contracts/contracts/security/ReentrancyGuard.sol";
import "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import "../lib/openzeppelin-contracts/contracts/utils/math/SafeMath.sol";
import "./IPriceFeedReciever.sol";

contract PriceFeed is ReentrancyGuard, Ownable {
    // Function to receive Ether. msg.data must be empty
    receive() external payable {}

    // Fallback function is called when msg.data is not empty
    fallback() external payable {}

    using ECDSA for bytes32;
    using SafeMath for uint256;

    address public trustedSigner;
    uint256 public feePerAsset;
    mapping(bytes32 => uint256) public prices;
    mapping(bytes32 => uint8) public decimals;

    uint256 public constant MAX_ASSETS_PER_REQUEST = 100;

    event PriceUpdated(bytes32 indexed asset, uint256 price, uint8 decimal);
    event PricesRequested(address indexed requester, bytes32[] assets);
    event FeeUpdated(uint256 newFeePerAsset);
    event TrustedSignerUpdated(address newSigner);

    error InvalidSignature();
    error InsufficientFee(uint256 required, uint256 provided);
    error TooManyAssets(uint256 provided, uint256 maximum);
    error PriceNotAvailable(bytes32 asset);

    error ZeroAddress();
    error TransferFailed();

    constructor(uint256 _initialFeePerAsset) {
        trustedSigner = owner();
        feePerAsset = _initialFeePerAsset;
    }

    function setTrustedSigner(address _newSigner) external onlyOwner {
        if (_newSigner == address(0)) revert ZeroAddress();
        trustedSigner = _newSigner;
        emit TrustedSignerUpdated(_newSigner);
    }

    function setFeePerAsset(uint256 _newFeePerAsset) external onlyOwner {
        feePerAsset = _newFeePerAsset;
        emit FeeUpdated(_newFeePerAsset);
    }

    function updatePrice(
        bytes32[] calldata _assets,
        uint8[] calldata _decimals,
        uint256[] calldata _prices,
        bytes memory _signature
    ) external {
        bytes32 messageHash = keccak256(
            abi.encodePacked(_assets, _prices, _decimals)
        );
        bytes32 ethSignedMessageHash = ECDSA.toEthSignedMessageHash(
            messageHash
        );
        address signer = ECDSA.recover(ethSignedMessageHash, _signature);

        if (signer != trustedSigner) revert InvalidSignature();

        uint256 length = _assets.length;

        for (uint256 i = 0; i < length; ) {
            bytes32 asset = _assets[i];
            uint8 decimal = _decimals[i];
            uint256 price = _prices[i];
            prices[asset] = price;
            emit PriceUpdated(asset, price, decimal);

            unchecked {
                i++;
            }
        }
    }

    function requestPrices(
        bytes32[] calldata _assets,
        function(uint256[] memory, uint8[] memory) external _callback
    ) external payable nonReentrant {
        if (_assets.length > MAX_ASSETS_PER_REQUEST)
            revert TooManyAssets(_assets.length, MAX_ASSETS_PER_REQUEST);

        // emit PricesRequested(msg.sender, _assets);
        uint8[] memory requestedDecimals = new uint8[](_assets.length);
        uint256[] memory requestedPrices = new uint256[](_assets.length);

        for (uint256 i = 0; i < _assets.length; ) {
            bytes32 asset = _assets[i];
            uint256 price = prices[asset];
            uint8 decimal = decimals[asset];

            if (price == 0) revert PriceNotAvailable(asset);
            requestedPrices[i] = price;
            requestedDecimals[i] = decimal;

            unchecked {
                i++;
            }
        }

        if (msg.value < feePerAsset.mul(_assets.length))
            revert InsufficientFee(feePerAsset.mul(_assets.length), msg.value);

        _callback(requestedPrices, requestedDecimals);
    }

    function withdraw() public onlyOwner {
        // get the amount of Ether stored in this contract
        uint256 amount = address(this).balance;

        // send all Ether to owner
        (bool success, ) = owner().call{value: amount}("");
        require(success, "Failed to send Ether");
    }
}

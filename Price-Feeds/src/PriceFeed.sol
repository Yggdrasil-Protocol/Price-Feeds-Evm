// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

// import "../lib/openzeppelin-contracts-upgradeable/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
// import "../lib/openzeppelin-contracts-upgradeable/contracts/security/PausableUpgradeable.sol";
// import "../lib/openzeppelin-contracts-upgradeable/contracts/utils/math/SafeMathUpgradeable.sol";
// import "../lib/openzeppelin-contracts-upgradeable/contracts/utils/cryptography/ECDSAUpgradeable.sol";
// import "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
// import "../lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/UUPSUpgradeable.sol";

import "../lib/openzeppelin-contracts/contracts/security/ReentrancyGuard.sol";
import "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import "../lib/openzeppelin-contracts/contracts/security/Pausable.sol";
import "../lib/openzeppelin-contracts/contracts/utils/math/SafeMath.sol";
import "../lib/openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";



/**
 * @title PriceFeed
 * @dev This contract allows for the updating and requesting of asset prices, managed by a trusted signer.
 */
contract PriceFeed is  ReentrancyGuard, Ownable, Pausable {
    /// @dev Function to receive Ether. `msg.data` must be empty.
    receive() external payable {}

    /// @dev Fallback function is called when `msg.data` is not empty.
    fallback() external payable {}

    using ECDSA for bytes32;
    using SafeMath for uint256;
    using SafeMath for uint8;

    /// @notice Fee charged per asset request.
    uint256 public feePerAsset;

    /// @notice Mapping of asset identifiers to their prices.
    mapping(bytes32 => uint256) private prices;

    /// @notice Mapping of asset identifiers to their decimals.
    mapping(bytes32 => uint8) private decimals;

    /// @notice Maximum number of assets allowed per request.
    uint256 private constant MAX_ASSETS_PER_REQUEST = 100;

    /// @notice Event emitted when an asset price is updated.
    /// @param asset The identifier of the asset.
    /// @param price The new price of the asset.
    /// @param decimal The number of decimals of the asset price.
    event PriceUpdated(bytes32 indexed asset, uint256 price, uint8 decimal);

    /// @notice Event emitted when prices are requested.
    /// @param requester The address of the requester.
    /// @param assets The array of asset identifiers.
    event PricesRequested(address indexed requester, bytes32[] assets);

    /// @notice Event emitted when the fee per asset is updated.
    /// @param newFeePerAsset The new fee per asset.
    event FeeUpdated(uint256 newFeePerAsset);

    

    /// @dev Error thrown when the provided fee is insufficient.
    /// @param required The required fee amount.
    /// @param provided The provided fee amount.
    error InsufficientFee(uint256 required, uint256 provided);

    /// @dev Error thrown when too many assets are requested.
    /// @param provided The number of assets provided.
    /// @param maximum The maximum number of assets allowed.
    error TooManyAssets(uint256 provided, uint256 maximum);

    /// @dev Error thrown when the price for an asset is not available.
    /// @param asset The identifier of the asset.
    error PriceNotAvailable(bytes32 asset);

    /// @dev Error thrown when a zero address is provided.
    error ZeroAddress();

    /// @dev Error thrown when a transfer fails.
    error TransferFailed();

    /// @dev Constructor that disables initializers.
    constructor(uint256 _feePerAsset) {
        feePerAsset = _feePerAsset; 
        
    }

    /// @notice Initializes the contract. Should be called only once.
    

    /**
     * @notice Sets a new fee per asset for price requests.
     * @dev Only the owner can call this function.
     * @param _newFeePerAsset The new fee to be charged per asset request.
     */
    function setFeePerAsset(uint256 _newFeePerAsset) external onlyOwner {
        feePerAsset = _newFeePerAsset;
        emit FeeUpdated(_newFeePerAsset);
    }

    /**
     * @notice Updates the prices of multiple assets.
     */
    function updatePrice(
        bytes32[] calldata _assets,
        uint8[] calldata _decimals,
        uint256[] calldata _prices
    ) external onlyOwner whenNotPaused {
        uint256 length = _assets.length;

        for (uint256 i = 0; i < length;) {
            bytes32 asset = _assets[i];
            uint8 decimal = _decimals[i];
            uint256 price = _prices[i];
            decimals[asset] = decimal;
            prices[asset] = price;
            emit PriceUpdated(asset, price, decimal);

            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Requests the prices of multiple assets.
     * @dev The function requires a fee to be paid and calls a callback function with the prices and decimals.
     * @param _assets The array of asset identifiers.
     * @param _callback The callback function to be called with the prices and decimals.
     */
    function requestPrices(
        bytes32[] calldata _assets,
        function(uint8[] memory, uint256[] memory) external _callback
    ) external payable nonReentrant whenNotPaused {
        if (_assets.length > MAX_ASSETS_PER_REQUEST)
            revert TooManyAssets(_assets.length, MAX_ASSETS_PER_REQUEST);

        uint8[] memory requestedDecimals = new uint8[](_assets.length);
        uint256[] memory requestedPrices = new uint256[](_assets.length);
        uint256 fees = feePerAsset.mul(_assets.length);

        if (msg.value < fees) revert InsufficientFee(fees, msg.value);

        for (uint256 i = 0; i < _assets.length;) {
            bytes32 asset = _assets[i];
            uint256 price = prices[asset];
            uint8 decimal = decimals[asset];

            if (price == 0) revert PriceNotAvailable(asset);
            requestedDecimals[i] = decimal;
            requestedPrices[i] = price;

            unchecked {
                i++;
            }
        }

        _callback(requestedDecimals, requestedPrices);
    }

    /**
     * @notice Withdraws all Ether from the contract to the owner's address.
     * @dev Only the owner can call this function.
     */
    function withdraw() public onlyOwner {
        uint256 amount = address(this).balance;
        (bool success,) = owner().call{value: amount}("");
        if (!success) revert TransferFailed();
    }

    /**
     * @notice Pauses the contract.
     * @dev Only the owner can call this function.
     */
    function pause() public whenNotPaused onlyOwner {
        _pause();
    }

    /**
     * @notice Unpauses the contract.
     * @dev Only the owner can call this function.
     */
    function unpause() public whenPaused onlyOwner {
        _unpause();
    }

    /**
     * @notice Pauses the contract.
     * @dev Internal function to pause the contract.
     */
    function _pause() internal override whenNotPaused {
       Pausable._pause();
        emit Paused(_msgSender());
    }

    /**
     * @notice Unpauses the contract.
     * @dev Internal function to unpause the contract.
     */
    function _unpause() internal override whenPaused {
        Pausable._unpause();
        emit Unpaused(_msgSender());
    }

    /**
     * @notice Gets the price of an asset.
     * @param _asset The identifier of the asset.
     * @return The price of the asset.
     */
    function getPrice(bytes32 _asset) public view returns (uint256) {
        return prices[_asset];
    }

    /**
     * @notice Gets the decimals of an asset.
     * @param _asset The identifier of the asset.
     * @return The decimals of the asset.
     */
    function getDecimal(bytes32 _asset) public onlyOwner view returns (uint8) {
        return decimals[_asset];
    }

   
}

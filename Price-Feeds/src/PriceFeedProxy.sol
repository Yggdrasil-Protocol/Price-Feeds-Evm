// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import "../lib/openzeppelin-contracts/contracts/security/ReentrancyGuard.sol";
import "./IPriceFeed.sol";
import "../lib/openzeppelin-contracts/contracts/utils/math/SafeMath.sol";

/**
 * @title PriceFeedProxy
 * @dev Proxy contract for interacting with the PriceFeed contract.
 */
contract PriceFeedProxy is Ownable, ReentrancyGuard , IPriceFeed{
    using SafeMath for uint256;

    address public priceFeedAddress;

    event PriceFeedAddressUpdated(address indexed oldAddress, address indexed newAddress);

    constructor(address _priceFeedAddress) {
        require(_priceFeedAddress != address(0), "PriceFeedProxy: Price feed address cannot be zero");
        priceFeedAddress = _priceFeedAddress;
    }

    function updatePriceFeedAddress(address _newPriceFeedAddress) external onlyOwner {
        require(_newPriceFeedAddress != address(0), "PriceFeedProxy: New price feed address cannot be zero");
        address oldAddress = priceFeedAddress;
        priceFeedAddress = _newPriceFeedAddress;
        emit PriceFeedAddressUpdated(oldAddress, _newPriceFeedAddress);
    }

    /**
     * @dev Fallback function to delegate calls to the PriceFeed contract.
     */
    fallback() external payable {
        _delegate(priceFeedAddress);
    }

    receive() external payable {
        // Receive function to accept ether transfers
    }

    function _delegate(address _impl) internal {
        assembly {
            // Copy msg.data. We take full control of memory in this inline assembly
            // block because it will not return to Solidity code.
            calldatacopy(0, 0, calldatasize())

            // Call the implementation.
            // out and outsize are 0 because we don't know the size yet.
            let result := delegatecall(gas(), _impl, 0, calldatasize(), 0, 0)

            // Copy the returned data.
            returndatacopy(0, 0, returndatasize())

            switch result
            // delegatecall returns 0 on error.
            case 0 { revert(0, returndatasize()) }
            default { return(0, returndatasize()) }
        }
    }

    function updatePriceFeed(Price calldata price) external onlyOwner {
        (bool success, bytes memory data) = priceFeedAddress.delegatecall(
            abi.encodeWithSignature("updatePriceFeed((string,uint256,uint256))", price)
        );
        require(success, "PriceFeedProxy: updatePriceFeed call failed");
    }

    function publishPriceFeed(Price calldata price) external onlyOwner {
        (bool success, bytes memory data) = priceFeedAddress.delegatecall(
            abi.encodeWithSignature("publishPriceFeed((string,uint256,uint256))", price)
        );
        require(success, "PriceFeedProxy: publishPriceFeed call failed");
    }

    function requestPriceFeed(Request memory request) external payable nonReentrant returns (PriceResponse memory response) {
        (bool success, bytes memory data) = priceFeedAddress.delegatecall(
            abi.encodeWithSignature("requestPriceFeed((string[]))", request)
        );
        require(success, "PriceFeedProxy: requestPriceFeed call failed");
        response = abi.decode(data, (PriceResponse));
    }
}

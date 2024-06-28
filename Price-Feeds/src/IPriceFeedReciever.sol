// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;


interface IPriceReceiver {
    function receivePrices(bytes32[] calldata assets, uint256[] calldata prices) external;
}
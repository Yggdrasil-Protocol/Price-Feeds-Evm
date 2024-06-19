// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

interface IPriceFeed{
    struct Price { 
        string pair ; 
        uint256 price ; 
        uint256 decimals ; 
    }

    struct Request {
        string[] pair;
        uint256[] decimals;
    }

    function updatePriceFeed(Price calldata price ) external;

    function publishPriceFeed(Price calldata price ) external;
    
    function requestPriceFeed(Request memory request) external payable returns (uint256[] memory price);
}
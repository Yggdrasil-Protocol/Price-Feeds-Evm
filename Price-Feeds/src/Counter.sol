// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import "./IPriceFeed.sol";

contract PriceFeed is Ownable, IPriceFeed  {
    event PriceFeedUpdated();
    event PriceFeedPublished();
    event PriceFeedRequested();

    struct Price { 
        string pair ; 
        uint256 price ; 
        uint256 decimals ; 
    }

    struct Request {
        string[] pair;
        uint256[] decimals;
    }

    mapping (string pair => uint256 price ) Feed  ; 

    function updatePriceFeed(Price calldata price ) public onlyOwner {

    }

    function publishPriceFeed(Price calldata price ) public onlyOwner {

    }

    function requesttPriceFeed(Request calldata request) public payable {
    
    }
}

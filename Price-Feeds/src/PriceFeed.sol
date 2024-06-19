// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import "./IPriceFeed.sol";

contract PriceFeed is Ownable, IPriceFeed  {

    event PriceFeedUpdated();
    event PriceFeedPublished();
    event PriceFeedRequested();


    mapping (string pair => uint256 price ) public Feed  ; 

    function updatePriceFeed(Price calldata price ) public onlyOwner {
        Feed[price.pair] = price.price / (10 ** price.decimals)  ;
        emit PriceFeedUpdated() ;
    }

    function publishPriceFeed(Price calldata price ) public onlyOwner {
        Feed[price.pair] = price.price /(10 ** price.decimals)  ;
        emit PriceFeedPublished() ; 
    }

   function requestPriceFeed(Request memory request) public payable returns (uint256[] memory price) {
    if (msg.value < 0.000001 ether) { 
        revert("PriceFeed: Insufficient funds"); 
    }

    price = new uint256[](request.pair.length);

    for (uint i = 0; i < request.pair.length; i++) {
        price[i] = Feed[request.pair[i]];
    }

    emit PriceFeedRequested();

    return price;
}
}

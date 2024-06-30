// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/PriceFeed.sol";

contract DeployPriceFeed is Script {
    function setUp() public {}


    function run() public {
        // uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        

        // Start broadcasting transactions
        vm.startBroadcast();
        

        // Deploy the logic contract
        PriceFeed priceFeed = new PriceFeed( 0.001 ether ) ; 

      


        // Log the addresses of the deployed contracts
        console.log("PriceFeed logic contract deployed at:", address(priceFeed));


        // Stop broadcasting transactions
        vm.stopBroadcast();
    }
}

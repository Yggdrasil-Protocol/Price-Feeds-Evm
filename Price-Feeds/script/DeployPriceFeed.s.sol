// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import  "../src/PriceFeed.sol";

contract DeployPriceFeed is Script {
    function run() external {
        // Start broadcasting transactions
        vm.startBroadcast();

        // uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        // vm.startBroadcast(deployerPrivateKey);
        // Deploy the PriceFeed contract
        PriceFeed priceFeed = new PriceFeed();

        // Optionally, transfer ownership if needed
        // priceFeed.transferOwnership(<new_owner_address>);

        // End broadcasting transactions
        vm.stopBroadcast();

        // Log the address of the deployed contract
        console.log("PriceFeed deployed to:", address(priceFeed));
    }
}

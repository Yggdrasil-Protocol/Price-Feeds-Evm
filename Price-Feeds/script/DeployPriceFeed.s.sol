// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/PriceFeed.sol";
import "../src/PriceFeedProxy.sol";

contract DeployPriceFeed is Script {
    function setUp() public {}


    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        

        // Start broadcasting transactions
        vm.startBroadcast((deployerPrivateKey));
        

        // Deploy the logic contract
        PriceFeed priceFeedLogic = new PriceFeed() ; 

        // Deploy the proxy contract wi th initialization
        PriceFeedProxy proxy = new PriceFeedProxy(
            address(priceFeedLogic),
            abi.encodeWithSignature("initialize().selector" , 0.001 ether )
        );

        // Log the addresses of the deployed contracts
        console.log("PriceFeed logic contract deployed at:", address(priceFeedLogic));
        console.log("PriceFeed proxy contract deployed at:", address(proxy));

        // Stop broadcasting transactions
        vm.stopBroadcast();
    }
}

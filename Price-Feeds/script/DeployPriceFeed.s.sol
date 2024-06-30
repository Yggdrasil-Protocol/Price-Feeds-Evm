// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/PriceFeed.sol";
import "../src/PriceFeedProxy.sol";

contract DeployPriceFeed is Script {
    function setUp() public {}


    function run() public {
        // uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        

        // Start broadcasting transactions
        vm.startBroadcast(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
        

        // Deploy the logic contract
        PriceFeed priceFeedLogic = new PriceFeed() ; 

        // Deploy the proxy contract wi th initialization
        PriceFeedProxy proxy = new PriceFeedProxy(
            address(priceFeedLogic),
            abi.encodeWithSignature("initialize().selector" , 0.00001 ether )
        );


        // Log the addresses of the deployed contracts
        console.log("PriceFeed logic contract deployed at:", address(priceFeedLogic));
        console.log("PriceFeed proxy contract deployed at:", address(proxy));

        // Stop broadcasting transactions
        vm.stopBroadcast();
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/PriceFeed.sol";
import "../src/PriceFeedProxy.sol";

contract DeployPriceFeed is Script {
    function setUp() public {}


    function run() public {
        // uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        // string memory rpcUrl = vm.envString("RPC_URL");

        // Start broadcasting transactions
        vm.startBroadcast(0x49f841619c9ba5edaf2a5eb7aa8c146a5649b4b02aac462dccf3d02e990fb662);

        // Deploy the logic contract
        PriceFeed priceFeedLogic = new PriceFeed(  0.01 ether);

        // Deploy the proxy contract with initialization
        PriceFeedProxy proxy = new PriceFeedProxy(
            address(priceFeedLogic),
            abi.encodeWithSignature("initialize()")
        );

        // Log the addresses of the deployed contracts
        console.log("PriceFeed logic contract deployed at:", address(priceFeedLogic));
        console.log("PriceFeed proxy contract deployed at:", address(proxy));

        // Stop broadcasting transactions
        vm.stopBroadcast();
    }
}

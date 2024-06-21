import Web3 from 'web3';
import dotenv from 'dotenv';
import { AbiItem } from 'web3-utils';

const abi: AbiItem[] = [
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "previousOwner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "OwnershipTransferred",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "string",
          "name": "pair",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "price",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "decimals",
          "type": "uint256"
        }
      ],
      "name": "PriceFeedPublished",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "requester",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "string[]",
          "name": "pairs",
          "type": "string[]"
        },
        {
          "indexed": false,
          "internalType": "uint256[]",
          "name": "prices",
          "type": "uint256[]"
        },
        {
          "indexed": false,
          "internalType": "uint256[]",
          "name": "decimals",
          "type": "uint256[]"
        }
      ],
      "name": "PriceFeedRequested",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "string",
          "name": "pair",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "price",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "decimals",
          "type": "uint256"
        }
      ],
      "name": "PriceFeedUpdated",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "pair",
          "type": "string"
        }
      ],
      "name": "Feed",
      "outputs": [
        {
          "internalType": "string",
          "name": "pair",
          "type": "string"
        },
        {
          "internalType": "uint256",
          "name": "price",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "decimals",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "owner",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "string",
              "name": "pair",
              "type": "string"
            },
            {
              "internalType": "uint256",
              "name": "price",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "decimals",
              "type": "uint256"
            }
          ],
          "internalType": "struct IPriceFeed.Price",
          "name": "price",
          "type": "tuple"
        }
      ],
      "name": "publishPriceFeed",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "renounceOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "string[]",
              "name": "pair",
              "type": "string[]"
            }
          ],
          "internalType": "struct IPriceFeed.Request",
          "name": "request",
          "type": "tuple"
        }
      ],
      "name": "requestPriceFeed",
      "outputs": [
        {
          "components": [
            {
              "internalType": "uint256[]",
              "name": "prices",
              "type": "uint256[]"
            },
            {
              "internalType": "uint256[]",
              "name": "decimals",
              "type": "uint256[]"
            }
          ],
          "internalType": "struct IPriceFeed.PriceResponse",
          "name": "response",
          "type": "tuple"
        }
      ],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "string",
              "name": "pair",
              "type": "string"
            },
            {
              "internalType": "uint256",
              "name": "price",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "decimals",
              "type": "uint256"
            }
          ],
          "internalType": "struct IPriceFeed.Price",
          "name": "price",
          "type": "tuple"
        }
      ],
      "name": "updatePriceFeed",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]
  ;

  dotenv.config();

  const { RPC_URL, PRIVATE_KEY, CONTRACT_ADDRESS } = process.env;
  
  if (!RPC_URL || !PRIVATE_KEY || !CONTRACT_ADDRESS) {
    throw new Error("Please define RPC_URL, PRIVATE_KEY, and CONTRACT_ADDRESS in your .env file");
  }
  
  const web3 = new Web3(new Web3.providers.HttpProvider(RPC_URL));
  const account = web3.eth.accounts.privateKeyToAccount(PRIVATE_KEY);
  web3.eth.accounts.wallet.add(account);
  
  const contract = new web3.eth.Contract(abi, CONTRACT_ADDRESS);
  
  interface Price {
    pair: string;
    price: number;
    decimals: number;
  }
  
  interface PriceResponse {
    prices: number[];
    decimals: number[];
  }

  interface Request {
    pair: string[];
  }
  
  async function updatePriceFeed(price: Price) {
    const data = contract.methods.updatePriceFeed(price).encodeABI();
  
    const tx = {
      to: CONTRACT_ADDRESS,
      from: account.address,
      data,
      gas: 200000,
      gasPrice: await web3.eth.getGasPrice(),  // Get current gas price
    };
  
    const signedTx = await web3.eth.accounts.signTransaction(tx,"49f841619c9ba5edaf2a5eb7aa8c146a5649b4b02aac462dccf3d02e990fb662");
    const receipt = await web3.eth.sendSignedTransaction(signedTx.rawTransaction!);
  
    console.log(`Updated price feed for ${price.pair} to ${price.price} with ${price.decimals} decimals. Transaction hash: ${receipt.transactionHash}`);
  }
  
  async function publishPriceFeed(price: Price) {
    const data = contract.methods.publishPriceFeed(price).encodeABI();
  
    const tx = {
      to: CONTRACT_ADDRESS,
      from: account.address,
      data,
      gas: 200000,
      gasPrice: await web3.eth.getGasPrice(),  // Get current gas price
    };
  
    const signedTx = await web3.eth.accounts.signTransaction(tx, "49f841619c9ba5edaf2a5eb7aa8c146a5649b4b02aac462dccf3d02e990fb662");
    const receipt = await web3.eth.sendSignedTransaction(signedTx.rawTransaction!);
  
    console.log(`Published price feed for ${price.pair} with price ${price.price} and ${price.decimals} decimals. Transaction hash: ${receipt.transactionHash}`);
  }
  
  async function requestPriceFeed(request: Request) {
    const data = contract.methods.requestPriceFeed(request).encodeABI();
  
    const tx = {
      to: CONTRACT_ADDRESS,
      from: account.address,
      data,
      value: web3.utils.toWei('0.000001' , 'ether'),
      gas: 200000,
      gasPrice: await web3.eth.getGasPrice(),  // Get current gas price
    };
  
    const signedTx = await web3.eth.accounts.signTransaction(tx, "49f841619c9ba5edaf2a5eb7aa8c146a5649b4b02aac462dccf3d02e990fb662");
    const receipt = await web3.eth.sendSignedTransaction(signedTx.rawTransaction!);

    let res ; 

    const event = receipt.events?.PriceFeedRequested;
    if (event) {
      const { pairs, prices , decimals } = event.returnValues;
    // @ts-ignore
    const response: PriceResponse = { prices, decimals};
    res = response
    console.log('Received price response:', response.prices[0])
    }

    console.log(res?.prices[0])
    return res?.prices[0] ; 




  }
  
  // Example usage
  const priceExample: Price = { pair: 'ETH/USD', price: 2000, decimals: 2 };
  const requestExample: Request = { pair: ['ETH/USD'] };
  
  updatePriceFeed(priceExample)
    .then(() => publishPriceFeed(priceExample))
    .then(() => requestPriceFeed(requestExample))
    .catch(console.error);
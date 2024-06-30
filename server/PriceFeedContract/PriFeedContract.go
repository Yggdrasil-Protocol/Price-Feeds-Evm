// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PriceFeedContract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PriceFeedContractMetaData contains all meta data concerning the PriceFeedContract contract.
var PriceFeedContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePerAsset\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"}],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"}],\"name\":\"PriceNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"TooManyAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeePerAsset\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"assets\",\"type\":\"bytes32[]\"}],\"name\":\"PricesRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"feePerAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_asset\",\"type\":\"bytes32\"}],\"name\":\"getDecimal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_asset\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_assets\",\"type\":\"bytes32[]\"},{\"internalType\":\"function(uint8[],uint256[])external\",\"name\":\"_callback\",\"type\":\"function\"}],\"name\":\"requestPrices\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeePerAsset\",\"type\":\"uint256\"}],\"name\":\"setFeePerAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_assets\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_decimals\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516119013803806119018339818101604052810190610031919061017c565b60015f8190555061005461004961007b60201b60201c565b61008260201b60201c565b5f600160146101000a81548160ff02191690831515021790555080600281905550506101a7565b5f33905090565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f80fd5b5f819050919050565b61015b81610149565b8114610165575f80fd5b50565b5f8151905061017681610152565b92915050565b5f6020828403121561019157610190610145565b5b5f61019e84828501610168565b91505092915050565b61174d806101b45f395ff3fe6080604052600436106100c5575f3560e01c8063715018a61161007e5780638da5cb5b116100585780638da5cb5b1461021a578063b2ad4c2e14610244578063e18e07cf14610260578063f2fde38b14610288576100cc565b8063715018a6146101b25780638456cb59146101c85780638a1b2d5c146101de576100cc565b806331d98b3f146100ce5780633ccfd60b1461010a5780633f4ba83a146101205780635c975abb146101365780636eefe8fa146101605780636f1518f81461018a576100cc565b366100cc57005b005b3480156100d9575f80fd5b506100f460048036038101906100ef9190610cbd565b6102b0565b6040516101019190610d00565b60405180910390f35b348015610115575f80fd5b5061011e6102ca565b005b34801561012b575f80fd5b50610134610380565b005b348015610141575f80fd5b5061014a61039a565b6040516101579190610d33565b60405180910390f35b34801561016b575f80fd5b506101746103b0565b6040516101819190610d00565b60405180910390f35b348015610195575f80fd5b506101b060048036038101906101ab9190610d76565b6103b6565b005b3480156101bd575f80fd5b506101c66103ff565b005b3480156101d3575f80fd5b506101dc610412565b005b3480156101e9575f80fd5b5061020460048036038101906101ff9190610cbd565b61042c565b6040516102119190610dbc565b60405180910390f35b348015610225575f80fd5b5061022e61045a565b60405161023b9190610e14565b60405180910390f35b61025e60048036038101906102599190610f51565b610482565b005b34801561026b575f80fd5b506102866004803603810190610281919061105d565b61073d565b005b348015610293575f80fd5b506102ae60048036038101906102a99190611137565b610856565b005b5f60035f8381526020019081526020015f20549050919050565b6102d26108d8565b5f4790505f6102df61045a565b73ffffffffffffffffffffffffffffffffffffffff16826040516103029061118f565b5f6040518083038185875af1925050503d805f811461033c576040519150601f19603f3d011682016040523d82523d5f602084013e610341565b606091505b505090508061037c576040517f90b8ec1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b610388610956565b6103906108d8565b61039861099f565b565b5f600160149054906101000a900460ff16905090565b60025481565b6103be6108d8565b806002819055507f8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76816040516103f49190610d00565b60405180910390a150565b6104076108d8565b6104105f6109ef565b565b61041a610ab2565b6104226108d8565b61042a610afc565b565b5f6104356108d8565b60045f8381526020019081526020015f205f9054906101000a900460ff169050919050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61048a610b4c565b610492610ab2565b60648484905011156104e1578383905060646040517f399c20d10000000000000000000000000000000000000000000000000000000081526004016104d89291906111a3565b60405180910390fd5b5f8484905067ffffffffffffffff8111156104ff576104fe6111ca565b5b60405190808252806020026020018201604052801561052d5781602001602082028036833780820191505090505b5090505f8585905067ffffffffffffffff81111561054e5761054d6111ca565b5b60405190808252806020026020018201604052801561057c5781602001602082028036833780820191505090505b5090505f61059887879050600254610b9990919063ffffffff16565b9050803410156105e15780346040517fa458261b0000000000000000000000000000000000000000000000000000000081526004016105d89291906111a3565b60405180910390fd5b5f5b878790508110156106dd575f888883818110610602576106016111f7565b5b9050602002013590505f60035f8381526020019081526020015f205490505f60045f8481526020019081526020015f205f9054906101000a900460ff1690505f820361068557826040517ff69cc81000000000000000000000000000000000000000000000000000000000815260040161067c9190611233565b60405180910390fd5b80878581518110610699576106986111f7565b5b602002602001019060ff16908160ff1681525050818685815181106106c1576106c06111f7565b5b60200260200101818152505083806001019450505050506105e3565b50848484846040518363ffffffff1660e01b81526004016106ff9291906113ba565b5f604051808303815f87803b158015610716575f80fd5b505af1158015610728573d5f803e3d5ffd5b50505050505050610737610bae565b50505050565b6107456108d8565b61074d610ab2565b5f8686905090505f5b8181101561084c575f888883818110610772576107716111f7565b5b9050602002013590505f87878481811061078f5761078e6111f7565b5b90506020020160208101906107a49190611419565b90505f8686858181106107ba576107b96111f7565b5b9050602002013590508160045f8581526020019081526020015f205f6101000a81548160ff021916908360ff1602179055508060035f8581526020019081526020015f2081905550827f8c62fe32113aae6ed87fdea7c7da35e9d2b99790bf4f577b0869fe9bbe12d58f8284604051610834929190611444565b60405180910390a28380600101945050505050610756565b5050505050505050565b61085e6108d8565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c3906114eb565b60405180910390fd5b6108d5816109ef565b50565b6108e0610bb7565b73ffffffffffffffffffffffffffffffffffffffff166108fe61045a565b73ffffffffffffffffffffffffffffffffffffffff1614610954576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094b90611553565b60405180910390fd5b565b61095e61039a565b61099d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610994906115bb565b60405180910390fd5b565b6109a7610956565b6109af610bbe565b7f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6109d8610bb7565b6040516109e59190610e14565b60405180910390a1565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610aba61039a565b15610afa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af190611623565b60405180910390fd5b565b610b04610ab2565b610b0c610c20565b7f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610b35610bb7565b604051610b429190610e14565b60405180910390a1565b60025f5403610b90576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b879061168b565b60405180910390fd5b60025f81905550565b5f8183610ba691906116d6565b905092915050565b60015f81905550565b5f33905090565b610bc6610956565b5f600160146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610c09610bb7565b604051610c169190610e14565b60405180910390a1565b610c28610ab2565b60018060146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610c6b610bb7565b604051610c789190610e14565b60405180910390a1565b5f80fd5b5f80fd5b5f819050919050565b610c9c81610c8a565b8114610ca6575f80fd5b50565b5f81359050610cb781610c93565b92915050565b5f60208284031215610cd257610cd1610c82565b5b5f610cdf84828501610ca9565b91505092915050565b5f819050919050565b610cfa81610ce8565b82525050565b5f602082019050610d135f830184610cf1565b92915050565b5f8115159050919050565b610d2d81610d19565b82525050565b5f602082019050610d465f830184610d24565b92915050565b610d5581610ce8565b8114610d5f575f80fd5b50565b5f81359050610d7081610d4c565b92915050565b5f60208284031215610d8b57610d8a610c82565b5b5f610d9884828501610d62565b91505092915050565b5f60ff82169050919050565b610db681610da1565b82525050565b5f602082019050610dcf5f830184610dad565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610dfe82610dd5565b9050919050565b610e0e81610df4565b82525050565b5f602082019050610e275f830184610e05565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112610e4e57610e4d610e2d565b5b8235905067ffffffffffffffff811115610e6b57610e6a610e31565b5b602083019150836020820283011115610e8757610e86610e35565b5b9250929050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082169050919050565b5f610ec382610e8e565b9050919050565b610ed381610eb9565b8114610edd575f80fd5b50565b5f81359050610eee81610eca565b92915050565b5f8160201c9050919050565b5f8160401c9050919050565b5f80610f1783610f00565b925063ffffffff83169050610f2b83610ef4565b9150915091565b5f80610f46610f418585610ee0565b610f0c565b915091509250929050565b5f805f8060408587031215610f6957610f68610c82565b5b5f85013567ffffffffffffffff811115610f8657610f85610c86565b5b610f9287828801610e39565b94509450506020610fa587828801610f32565b925092505092959194509250565b5f8083601f840112610fc857610fc7610e2d565b5b8235905067ffffffffffffffff811115610fe557610fe4610e31565b5b60208301915083602082028301111561100157611000610e35565b5b9250929050565b5f8083601f84011261101d5761101c610e2d565b5b8235905067ffffffffffffffff81111561103a57611039610e31565b5b60208301915083602082028301111561105657611055610e35565b5b9250929050565b5f805f805f806060878903121561107757611076610c82565b5b5f87013567ffffffffffffffff81111561109457611093610c86565b5b6110a089828a01610e39565b9650965050602087013567ffffffffffffffff8111156110c3576110c2610c86565b5b6110cf89828a01610fb3565b9450945050604087013567ffffffffffffffff8111156110f2576110f1610c86565b5b6110fe89828a01611008565b92509250509295509295509295565b61111681610df4565b8114611120575f80fd5b50565b5f813590506111318161110d565b92915050565b5f6020828403121561114c5761114b610c82565b5b5f61115984828501611123565b91505092915050565b5f81905092915050565b50565b5f61117a5f83611162565b91506111858261116c565b5f82019050919050565b5f6111998261116f565b9150819050919050565b5f6040820190506111b65f830185610cf1565b6111c36020830184610cf1565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b61122d81610c8a565b82525050565b5f6020820190506112465f830184611224565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61127e81610da1565b82525050565b5f61128f8383611275565b60208301905092915050565b5f602082019050919050565b5f6112b18261124c565b6112bb8185611256565b93506112c683611266565b805f5b838110156112f65781516112dd8882611284565b97506112e88361129b565b9250506001810190506112c9565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61133581610ce8565b82525050565b5f611346838361132c565b60208301905092915050565b5f602082019050919050565b5f61136882611303565b611372818561130d565b935061137d8361131d565b805f5b838110156113ad578151611394888261133b565b975061139f83611352565b925050600181019050611380565b5085935050505092915050565b5f6040820190508181035f8301526113d281856112a7565b905081810360208301526113e6818461135e565b90509392505050565b6113f881610da1565b8114611402575f80fd5b50565b5f81359050611413816113ef565b92915050565b5f6020828403121561142e5761142d610c82565b5b5f61143b84828501611405565b91505092915050565b5f6040820190506114575f830185610cf1565b6114646020830184610dad565b9392505050565b5f82825260208201905092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6114d560268361146b565b91506114e08261147b565b604082019050919050565b5f6020820190508181035f830152611502816114c9565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f61153d60208361146b565b915061154882611509565b602082019050919050565b5f6020820190508181035f83015261156a81611531565b9050919050565b7f5061757361626c653a206e6f74207061757365640000000000000000000000005f82015250565b5f6115a560148361146b565b91506115b082611571565b602082019050919050565b5f6020820190508181035f8301526115d281611599565b9050919050565b7f5061757361626c653a20706175736564000000000000000000000000000000005f82015250565b5f61160d60108361146b565b9150611618826115d9565b602082019050919050565b5f6020820190508181035f83015261163a81611601565b9050919050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c005f82015250565b5f611675601f8361146b565b915061168082611641565b602082019050919050565b5f6020820190508181035f8301526116a281611669565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6116e082610ce8565b91506116eb83610ce8565b92508282026116f981610ce8565b915082820484148315176117105761170f6116a9565b5b509291505056fea2646970667358221220e9bd2699337c8ff04e82dfa66d40e6075348fd837d1753a31bde05f4a789d4af64736f6c634300081a0033",
}

// PriceFeedContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedContractMetaData.ABI instead.
var PriceFeedContractABI = PriceFeedContractMetaData.ABI

// PriceFeedContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PriceFeedContractMetaData.Bin instead.
var PriceFeedContractBin = PriceFeedContractMetaData.Bin

// DeployPriceFeedContract deploys a new Ethereum contract, binding an instance of PriceFeedContract to it.
func DeployPriceFeedContract(auth *bind.TransactOpts, backend bind.ContractBackend, _feePerAsset *big.Int) (common.Address, *types.Transaction, *PriceFeedContract, error) {
	parsed, err := PriceFeedContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceFeedContractBin), backend, _feePerAsset)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceFeedContract{PriceFeedContractCaller: PriceFeedContractCaller{contract: contract}, PriceFeedContractTransactor: PriceFeedContractTransactor{contract: contract}, PriceFeedContractFilterer: PriceFeedContractFilterer{contract: contract}}, nil
}

// PriceFeedContract is an auto generated Go binding around an Ethereum contract.
type PriceFeedContract struct {
	PriceFeedContractCaller     // Read-only binding to the contract
	PriceFeedContractTransactor // Write-only binding to the contract
	PriceFeedContractFilterer   // Log filterer for contract events
}

// PriceFeedContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedContractSession struct {
	Contract     *PriceFeedContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PriceFeedContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedContractCallerSession struct {
	Contract *PriceFeedContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PriceFeedContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedContractTransactorSession struct {
	Contract     *PriceFeedContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PriceFeedContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedContractRaw struct {
	Contract *PriceFeedContract // Generic contract binding to access the raw methods on
}

// PriceFeedContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedContractCallerRaw struct {
	Contract *PriceFeedContractCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedContractTransactorRaw struct {
	Contract *PriceFeedContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeedContract creates a new instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContract(address common.Address, backend bind.ContractBackend) (*PriceFeedContract, error) {
	contract, err := bindPriceFeedContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContract{PriceFeedContractCaller: PriceFeedContractCaller{contract: contract}, PriceFeedContractTransactor: PriceFeedContractTransactor{contract: contract}, PriceFeedContractFilterer: PriceFeedContractFilterer{contract: contract}}, nil
}

// NewPriceFeedContractCaller creates a new read-only instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContractCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedContractCaller, error) {
	contract, err := bindPriceFeedContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractCaller{contract: contract}, nil
}

// NewPriceFeedContractTransactor creates a new write-only instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedContractTransactor, error) {
	contract, err := bindPriceFeedContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractTransactor{contract: contract}, nil
}

// NewPriceFeedContractFilterer creates a new log filterer instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedContractFilterer, error) {
	contract, err := bindPriceFeedContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractFilterer{contract: contract}, nil
}

// bindPriceFeedContract binds a generic wrapper to an already deployed contract.
func bindPriceFeedContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedContract *PriceFeedContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedContract.Contract.PriceFeedContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedContract *PriceFeedContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.PriceFeedContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedContract *PriceFeedContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.PriceFeedContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedContract *PriceFeedContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedContract *PriceFeedContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedContract *PriceFeedContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.contract.Transact(opts, method, params...)
}

// FeePerAsset is a free data retrieval call binding the contract method 0x6eefe8fa.
//
// Solidity: function feePerAsset() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) FeePerAsset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "feePerAsset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerAsset is a free data retrieval call binding the contract method 0x6eefe8fa.
//
// Solidity: function feePerAsset() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) FeePerAsset() (*big.Int, error) {
	return _PriceFeedContract.Contract.FeePerAsset(&_PriceFeedContract.CallOpts)
}

// FeePerAsset is a free data retrieval call binding the contract method 0x6eefe8fa.
//
// Solidity: function feePerAsset() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) FeePerAsset() (*big.Int, error) {
	return _PriceFeedContract.Contract.FeePerAsset(&_PriceFeedContract.CallOpts)
}

// GetDecimal is a free data retrieval call binding the contract method 0x8a1b2d5c.
//
// Solidity: function getDecimal(bytes32 _asset) view returns(uint8)
func (_PriceFeedContract *PriceFeedContractCaller) GetDecimal(opts *bind.CallOpts, _asset [32]byte) (uint8, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "getDecimal", _asset)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDecimal is a free data retrieval call binding the contract method 0x8a1b2d5c.
//
// Solidity: function getDecimal(bytes32 _asset) view returns(uint8)
func (_PriceFeedContract *PriceFeedContractSession) GetDecimal(_asset [32]byte) (uint8, error) {
	return _PriceFeedContract.Contract.GetDecimal(&_PriceFeedContract.CallOpts, _asset)
}

// GetDecimal is a free data retrieval call binding the contract method 0x8a1b2d5c.
//
// Solidity: function getDecimal(bytes32 _asset) view returns(uint8)
func (_PriceFeedContract *PriceFeedContractCallerSession) GetDecimal(_asset [32]byte) (uint8, error) {
	return _PriceFeedContract.Contract.GetDecimal(&_PriceFeedContract.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 _asset) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) GetPrice(opts *bind.CallOpts, _asset [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "getPrice", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 _asset) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) GetPrice(_asset [32]byte) (*big.Int, error) {
	return _PriceFeedContract.Contract.GetPrice(&_PriceFeedContract.CallOpts, _asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 _asset) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) GetPrice(_asset [32]byte) (*big.Int, error) {
	return _PriceFeedContract.Contract.GetPrice(&_PriceFeedContract.CallOpts, _asset)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractSession) Owner() (common.Address, error) {
	return _PriceFeedContract.Contract.Owner(&_PriceFeedContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractCallerSession) Owner() (common.Address, error) {
	return _PriceFeedContract.Contract.Owner(&_PriceFeedContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceFeedContract *PriceFeedContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceFeedContract *PriceFeedContractSession) Paused() (bool, error) {
	return _PriceFeedContract.Contract.Paused(&_PriceFeedContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceFeedContract *PriceFeedContractCallerSession) Paused() (bool, error) {
	return _PriceFeedContract.Contract.Paused(&_PriceFeedContract.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceFeedContract *PriceFeedContractSession) Pause() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Pause(&_PriceFeedContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Pause() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Pause(&_PriceFeedContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeedContract *PriceFeedContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeedContract *PriceFeedContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RenounceOwnership(&_PriceFeedContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RenounceOwnership(&_PriceFeedContract.TransactOpts)
}

// RequestPrices is a paid mutator transaction binding the contract method 0xb2ad4c2e.
//
// Solidity: function requestPrices(bytes32[] _assets, function _callback) payable returns()
func (_PriceFeedContract *PriceFeedContractTransactor) RequestPrices(opts *bind.TransactOpts, _assets [][32]byte, _callback [24]byte) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "requestPrices", _assets, _callback)
}

// RequestPrices is a paid mutator transaction binding the contract method 0xb2ad4c2e.
//
// Solidity: function requestPrices(bytes32[] _assets, function _callback) payable returns()
func (_PriceFeedContract *PriceFeedContractSession) RequestPrices(_assets [][32]byte, _callback [24]byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RequestPrices(&_PriceFeedContract.TransactOpts, _assets, _callback)
}

// RequestPrices is a paid mutator transaction binding the contract method 0xb2ad4c2e.
//
// Solidity: function requestPrices(bytes32[] _assets, function _callback) payable returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) RequestPrices(_assets [][32]byte, _callback [24]byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RequestPrices(&_PriceFeedContract.TransactOpts, _assets, _callback)
}

// SetFeePerAsset is a paid mutator transaction binding the contract method 0x6f1518f8.
//
// Solidity: function setFeePerAsset(uint256 _newFeePerAsset) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) SetFeePerAsset(opts *bind.TransactOpts, _newFeePerAsset *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "setFeePerAsset", _newFeePerAsset)
}

// SetFeePerAsset is a paid mutator transaction binding the contract method 0x6f1518f8.
//
// Solidity: function setFeePerAsset(uint256 _newFeePerAsset) returns()
func (_PriceFeedContract *PriceFeedContractSession) SetFeePerAsset(_newFeePerAsset *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.SetFeePerAsset(&_PriceFeedContract.TransactOpts, _newFeePerAsset)
}

// SetFeePerAsset is a paid mutator transaction binding the contract method 0x6f1518f8.
//
// Solidity: function setFeePerAsset(uint256 _newFeePerAsset) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) SetFeePerAsset(_newFeePerAsset *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.SetFeePerAsset(&_PriceFeedContract.TransactOpts, _newFeePerAsset)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeedContract *PriceFeedContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.TransferOwnership(&_PriceFeedContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.TransferOwnership(&_PriceFeedContract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceFeedContract *PriceFeedContractSession) Unpause() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Unpause(&_PriceFeedContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Unpause(&_PriceFeedContract.TransactOpts)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0xe18e07cf.
//
// Solidity: function updatePrice(bytes32[] _assets, uint8[] _decimals, uint256[] _prices) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) UpdatePrice(opts *bind.TransactOpts, _assets [][32]byte, _decimals []uint8, _prices []*big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "updatePrice", _assets, _decimals, _prices)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0xe18e07cf.
//
// Solidity: function updatePrice(bytes32[] _assets, uint8[] _decimals, uint256[] _prices) returns()
func (_PriceFeedContract *PriceFeedContractSession) UpdatePrice(_assets [][32]byte, _decimals []uint8, _prices []*big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpdatePrice(&_PriceFeedContract.TransactOpts, _assets, _decimals, _prices)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0xe18e07cf.
//
// Solidity: function updatePrice(bytes32[] _assets, uint8[] _decimals, uint256[] _prices) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) UpdatePrice(_assets [][32]byte, _decimals []uint8, _prices []*big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpdatePrice(&_PriceFeedContract.TransactOpts, _assets, _decimals, _prices)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceFeedContract *PriceFeedContractSession) Withdraw() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Withdraw(&_PriceFeedContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Withdraw(&_PriceFeedContract.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PriceFeedContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PriceFeedContract *PriceFeedContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Fallback(&_PriceFeedContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Fallback(&_PriceFeedContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PriceFeedContract *PriceFeedContractSession) Receive() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Receive(&_PriceFeedContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Receive() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Receive(&_PriceFeedContract.TransactOpts)
}

// PriceFeedContractFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the PriceFeedContract contract.
type PriceFeedContractFeeUpdatedIterator struct {
	Event *PriceFeedContractFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeedContractFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedContractFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedContractFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractFeeUpdated represents a FeeUpdated event raised by the PriceFeedContract contract.
type PriceFeedContractFeeUpdated struct {
	NewFeePerAsset *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newFeePerAsset)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*PriceFeedContractFeeUpdatedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractFeeUpdatedIterator{contract: _PriceFeedContract.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newFeePerAsset)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedContractFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractFeeUpdated)
				if err := _PriceFeedContract.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeUpdated is a log parse operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newFeePerAsset)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseFeeUpdated(log types.Log) (*PriceFeedContractFeeUpdated, error) {
	event := new(PriceFeedContractFeeUpdated)
	if err := _PriceFeedContract.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceFeedContract contract.
type PriceFeedContractOwnershipTransferredIterator struct {
	Event *PriceFeedContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeedContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractOwnershipTransferred represents a OwnershipTransferred event raised by the PriceFeedContract contract.
type PriceFeedContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceFeedContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractOwnershipTransferredIterator{contract: _PriceFeedContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceFeedContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractOwnershipTransferred)
				if err := _PriceFeedContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseOwnershipTransferred(log types.Log) (*PriceFeedContractOwnershipTransferred, error) {
	event := new(PriceFeedContractOwnershipTransferred)
	if err := _PriceFeedContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PriceFeedContract contract.
type PriceFeedContractPausedIterator struct {
	Event *PriceFeedContractPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeedContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedContractPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPaused represents a Paused event raised by the PriceFeedContract contract.
type PriceFeedContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPaused(opts *bind.FilterOpts) (*PriceFeedContractPausedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPausedIterator{contract: _PriceFeedContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPaused) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPaused)
				if err := _PriceFeedContract.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePaused(log types.Log) (*PriceFeedContractPaused, error) {
	event := new(PriceFeedContractPaused)
	if err := _PriceFeedContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the PriceFeedContract contract.
type PriceFeedContractPriceUpdatedIterator struct {
	Event *PriceFeedContractPriceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeedContractPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPriceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedContractPriceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedContractPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPriceUpdated represents a PriceUpdated event raised by the PriceFeedContract contract.
type PriceFeedContractPriceUpdated struct {
	Asset   [32]byte
	Price   *big.Int
	Decimal uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x8c62fe32113aae6ed87fdea7c7da35e9d2b99790bf4f577b0869fe9bbe12d58f.
//
// Solidity: event PriceUpdated(bytes32 indexed asset, uint256 price, uint8 decimal)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPriceUpdated(opts *bind.FilterOpts, asset [][32]byte) (*PriceFeedContractPriceUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "PriceUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPriceUpdatedIterator{contract: _PriceFeedContract.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x8c62fe32113aae6ed87fdea7c7da35e9d2b99790bf4f577b0869fe9bbe12d58f.
//
// Solidity: event PriceUpdated(bytes32 indexed asset, uint256 price, uint8 decimal)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPriceUpdated, asset [][32]byte) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "PriceUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPriceUpdated)
				if err := _PriceFeedContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceUpdated is a log parse operation binding the contract event 0x8c62fe32113aae6ed87fdea7c7da35e9d2b99790bf4f577b0869fe9bbe12d58f.
//
// Solidity: event PriceUpdated(bytes32 indexed asset, uint256 price, uint8 decimal)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePriceUpdated(log types.Log) (*PriceFeedContractPriceUpdated, error) {
	event := new(PriceFeedContractPriceUpdated)
	if err := _PriceFeedContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractPricesRequestedIterator is returned from FilterPricesRequested and is used to iterate over the raw logs and unpacked data for PricesRequested events raised by the PriceFeedContract contract.
type PriceFeedContractPricesRequestedIterator struct {
	Event *PriceFeedContractPricesRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeedContractPricesRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPricesRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedContractPricesRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedContractPricesRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPricesRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPricesRequested represents a PricesRequested event raised by the PriceFeedContract contract.
type PriceFeedContractPricesRequested struct {
	Requester common.Address
	Assets    [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPricesRequested is a free log retrieval operation binding the contract event 0xaa33eb4cfd36e0ee1d8d133a19a2620fb088f12c0141f8bc7e2d48c220be343f.
//
// Solidity: event PricesRequested(address indexed requester, bytes32[] assets)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPricesRequested(opts *bind.FilterOpts, requester []common.Address) (*PriceFeedContractPricesRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "PricesRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPricesRequestedIterator{contract: _PriceFeedContract.contract, event: "PricesRequested", logs: logs, sub: sub}, nil
}

// WatchPricesRequested is a free log subscription operation binding the contract event 0xaa33eb4cfd36e0ee1d8d133a19a2620fb088f12c0141f8bc7e2d48c220be343f.
//
// Solidity: event PricesRequested(address indexed requester, bytes32[] assets)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPricesRequested(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPricesRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "PricesRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPricesRequested)
				if err := _PriceFeedContract.contract.UnpackLog(event, "PricesRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePricesRequested is a log parse operation binding the contract event 0xaa33eb4cfd36e0ee1d8d133a19a2620fb088f12c0141f8bc7e2d48c220be343f.
//
// Solidity: event PricesRequested(address indexed requester, bytes32[] assets)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePricesRequested(log types.Log) (*PriceFeedContractPricesRequested, error) {
	event := new(PriceFeedContractPricesRequested)
	if err := _PriceFeedContract.contract.UnpackLog(event, "PricesRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PriceFeedContract contract.
type PriceFeedContractUnpausedIterator struct {
	Event *PriceFeedContractUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeedContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedContractUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractUnpaused represents a Unpaused event raised by the PriceFeedContract contract.
type PriceFeedContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PriceFeedContractUnpausedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractUnpausedIterator{contract: _PriceFeedContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PriceFeedContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractUnpaused)
				if err := _PriceFeedContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseUnpaused(log types.Log) (*PriceFeedContractUnpaused, error) {
	event := new(PriceFeedContractUnpaused)
	if err := _PriceFeedContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// IPriceFeedPrice is an auto generated low-level Go binding around an user-defined struct.
type IPriceFeedPrice struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
}

// IPriceFeedPriceResponse is an auto generated low-level Go binding around an user-defined struct.
type IPriceFeedPriceResponse struct {
	Prices   []*big.Int
	Decimals []*big.Int
}

// IPriceFeedRequest is an auto generated low-level Go binding around an user-defined struct.
type IPriceFeedRequest struct {
	Pair []string
}

// PriceFeedContractMetaData contains all meta data concerning the PriceFeedContract contract.
var PriceFeedContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"PriceFeedPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"pairs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"decimals\",\"type\":\"uint256[]\"}],\"name\":\"PriceFeedRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"PriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"}],\"name\":\"Feed\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"pair\",\"type\":\"string[]\"}],\"internalType\":\"structIPriceFeed.Request\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"requestPriceFeed\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"decimals\",\"type\":\"uint256[]\"}],\"internalType\":\"structIPriceFeed.PriceResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"name\":\"updatePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100dd565b5f54610100900460ff161561008c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146100db575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6080516119116101115f395f81816101d701528181610220015281816106080152818161064801526106d901526119115ff3fe60806040526004361061008f575f3560e01c80636d2fbecb116100575780636d2fbecb14610131578063715018a61461015f5780638129fc1c146101735780638da5cb5b14610187578063f2fde38b146101ae575f80fd5b80633659cfe61461009357806349c14ed9146100b45780634f1ef286146100dd57806352d1902d146100f057806368650ca214610112575b5f80fd5b34801561009e575f80fd5b506100b26100ad36600461112e565b6101cd565b005b6100c76100c2366004611229565b6102b3565b6040516100d49190611361565b60405180910390f35b6100b26100eb3660046113a2565b6105fe565b3480156100fb575f80fd5b506101046106cd565b6040519081526020016100d4565b34801561011d575f80fd5b506100b261012c3660046113ff565b61077e565b34801561013c575f80fd5b5061015061014b366004611435565b610922565b6040516100d493929190611494565b34801561016a575f80fd5b506100b26109d5565b34801561017e575f80fd5b506100b26109e8565b348015610192575f80fd5b5060c9546040516001600160a01b0390911681526020016100d4565b3480156101b9575f80fd5b506100b26101c836600461112e565b610afc565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361021e5760405162461bcd60e51b8152600401610215906114b8565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166102665f80516020611895833981519152546001600160a01b031690565b6001600160a01b03161461028c5760405162461bcd60e51b815260040161021590611504565b61029581610b72565b604080515f808252602082019092526102b091839190610b7a565b50565b60408051808201909152606080825260208201526102cf610ce9565b81515161031e5760405162461bcd60e51b815260206004820152601d60248201527f5072696365466565643a204e6f207061697273207265717565737465640000006044820152606401610215565b64e8d4a510003410156103735760405162461bcd60e51b815260206004820152601d60248201527f5072696365466565643a20496e73756666696369656e742066756e64730000006044820152606401610215565b8151515f906001600160401b0381111561038f5761038f611147565b6040519080825280602002602001820160405280156103b8578160200160208202803683370190505b5090505f835f0151516001600160401b038111156103d8576103d8611147565b604051908082528060200260200182016040528015610401578160200160208202803683370190505b5090505f5b84515181101561059c575f60fb865f0151838151811061042857610428611550565b602002602001015160405161043d919061157b565b90815260200160405180910390206040518060600160405290815f8201805461046590611586565b80601f016020809104026020016040519081016040528092919081815260200182805461049190611586565b80156104dc5780601f106104b3576101008083540402835291602001916104dc565b820191905f5260205f20905b8154815290600101906020018083116104bf57829003601f168201915b505050505081526020016001820154815260200160028201548152505090505f815f0151511161054e5760405162461bcd60e51b815260206004820152601a60248201527f5072696365466565643a205072696365206e6f7420666f756e640000000000006044820152606401610215565b806020015184838151811061056557610565611550565b602002602001018181525050806040015183838151811061058857610588611550565b602090810291909101015250600101610406565b50835160405133917f21fb481b338df488246d581cca0377c59b67cc32b0ba179ac6d2c2975432aed1916105d49190869086906115be565b60405180910390a260408051808201909152918252602082015290506105f960018055565b919050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036106465760405162461bcd60e51b8152600401610215906114b8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661068e5f80516020611895833981519152546001600160a01b031690565b6001600160a01b0316146106b45760405162461bcd60e51b815260040161021590611504565b6106bd82610b72565b6106c982826001610b7a565b5050565b5f306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461076c5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610215565b505f8051602061189583398151915290565b610786610d48565b5f610791828061163f565b9050116107e05760405162461bcd60e51b815260206004820152601f60248201527f5072696365466565643a20506169722063616e6e6f7420626520656d707479006044820152606401610215565b5f8160200135116108465760405162461bcd60e51b815260206004820152602a60248201527f5072696365466565643a205072696365206d7573742062652067726561746572604482015269207468616e207a65726f60b01b6064820152608401610215565b60128160400135111561089b5760405162461bcd60e51b815260206004820152601b60248201527f5072696365466565643a20496e76616c696420646563696d616c7300000000006044820152606401610215565b8060fb6108a8828061163f565b6040516108b6929190611688565b9081526040519081900360200190206108cf82826116e2565b507f3e51a618566783aa8f6e83296eb48686664aeb521acd6a7507a35fb5fd9ac74090506108fd828061163f565b8360200135846040013560405161091794939291906117e2565b60405180910390a150565b805160208183018101805160fb8252928201919093012091528054819061094890611586565b80601f016020809104026020016040519081016040528092919081815260200182805461097490611586565b80156109bf5780601f10610996576101008083540402835291602001916109bf565b820191905f5260205f20905b8154815290600101906020018083116109a257829003601f168201915b5050505050908060010154908060020154905083565b6109dd610d48565b6109e65f610da2565b565b5f54610100900460ff1615808015610a0657505f54600160ff909116105b80610a1f5750303b158015610a1f57505f5460ff166001145b610a825760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610215565b5f805460ff191660011790558015610aa3575f805461ff0019166101001790555b610aab610df3565b610ab3610e21565b610abb610e4f565b80156102b0575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610917565b610b04610d48565b6001600160a01b038116610b695760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610215565b6102b081610da2565b6102b0610d48565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610bb257610bad83610e75565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610c0c575060408051601f3d908101601f19168201909252610c0991810190611820565b60015b610c6f5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610215565b5f805160206118958339815191528114610cdd5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610215565b50610bad838383610f10565b600260015403610d3b5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610215565b6002600155565b60018055565b60c9546001600160a01b031633146109e65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610215565b60c980546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16610e195760405162461bcd60e51b815260040161021590611837565b6109e6610f3a565b5f54610100900460ff16610e475760405162461bcd60e51b815260040161021590611837565b6109e6610f69565b5f54610100900460ff166109e65760405162461bcd60e51b815260040161021590611837565b6001600160a01b0381163b610ee25760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610215565b5f8051602061189583398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b610f1983610f8f565b5f82511180610f255750805b15610bad57610f348383610fce565b50505050565b5f54610100900460ff16610f605760405162461bcd60e51b815260040161021590611837565b6109e633610da2565b5f54610100900460ff16610d425760405162461bcd60e51b815260040161021590611837565b610f9881610e75565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b6060610ff383836040518060600160405280602781526020016118b560279139610ffa565b9392505050565b60605f80856001600160a01b031685604051611016919061157b565b5f60405180830381855af49150503d805f811461104e576040519150601f19603f3d011682016040523d82523d5f602084013e611053565b606091505b50915091506110648683838761106e565b9695505050505050565b606083156110dc5782515f036110d5576001600160a01b0385163b6110d55760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610215565b50816110e6565b6110e683836110ee565b949350505050565b8151156110fe5781518083602001fd5b8060405162461bcd60e51b81526004016102159190611882565b80356001600160a01b03811681146105f9575f80fd5b5f6020828403121561113e575f80fd5b610ff382611118565b634e487b7160e01b5f52604160045260245ffd5b604051602081016001600160401b038111828210171561117d5761117d611147565b60405290565b604051601f8201601f191681016001600160401b03811182821017156111ab576111ab611147565b604052919050565b5f806001600160401b038411156111cc576111cc611147565b50601f8301601f19166020016111e181611183565b9150508281528383830111156111f5575f80fd5b828260208301375f602084830101529392505050565b5f82601f83011261121a575f80fd5b610ff3838335602085016111b3565b5f60208284031215611239575f80fd5b81356001600160401b0381111561124e575f80fd5b82016020818503121561125f575f80fd5b61126761115b565b81356001600160401b0381111561127c575f80fd5b80830192505084601f830112611290575f80fd5b81356001600160401b038111156112a9576112a9611147565b8060051b6112b960208201611183565b918252602081850181019290810190888411156112d4575f80fd5b6020860192505b838310156113195782356001600160401b038111156112f8575f80fd5b6113078a6020838a010161120b565b835250602092830192909101906112db565b845250919695505050505050565b5f8151808452602084019350602083015f5b82811015611357578151865260209586019590910190600101611339565b5093949350505050565b602081525f82516040602084015261137c6060840182611327565b90506020840151601f198483030160408501526113998282611327565b95945050505050565b5f80604083850312156113b3575f80fd5b6113bc83611118565b915060208301356001600160401b038111156113d6575f80fd5b8301601f810185136113e6575f80fd5b6113f5858235602084016111b3565b9150509250929050565b5f6020828403121561140f575f80fd5b81356001600160401b03811115611424575f80fd5b820160608185031215610ff3575f80fd5b5f60208284031215611445575f80fd5b81356001600160401b0381111561145a575f80fd5b6110e68482850161120b565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b606081525f6114a66060830186611466565b60208301949094525060400152919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b5f52603260045260245ffd5b5f81518060208401855e5f93019283525090919050565b5f610ff38284611564565b600181811c9082168061159a57607f821691505b6020821081036115b857634e487b7160e01b5f52602260045260245ffd5b50919050565b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b8281101561161557607f19878603018452611600858351611466565b945060209384019391909101906001016115e4565b50505050828103602084015261162b8186611327565b905082810360408401526110648185611327565b5f808335601e19843603018112611654575f80fd5b8301803591506001600160401b0382111561166d575f80fd5b602001915036819003821315611681575f80fd5b9250929050565b818382375f9101908152919050565b601f821115610bad57805f5260205f20601f840160051c810160208510156116bc5750805b601f840160051c820191505b818110156116db575f81556001016116c8565b5050505050565b8135601e198336030181126116f5575f80fd5b820180356001600160401b038111801561170d575f80fd5b81360360208401131561171e575f80fd5b5f905050611736816117308554611586565b85611697565b5f601f82116001811461176a575f83156117535750838201602001355b5f19600385901b1c1916600184901b1785556117c6565b5f85815260208120601f198516915b8281101561179b57602085880181013583559485019460019092019101611779565b50848210156117ba575f1960f88660031b161c19602085880101351681555b505060018360011b0185555b5050505060208201356001820155604090910135600290910155565b60608152836060820152838560808301375f608085830101525f6080601f19601f870116830101905083602083015282604083015295945050505050565b5f60208284031215611830575f80fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b602081525f610ff3602083018461146656fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212203a417be9a27d2801fb45fb5212dc47aded46496c438bfd05aadb8827db7ff5aa64736f6c634300081a0033",
}

// PriceFeedContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedContractMetaData.ABI instead.
var PriceFeedContractABI = PriceFeedContractMetaData.ABI

// PriceFeedContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PriceFeedContractMetaData.Bin instead.
var PriceFeedContractBin = PriceFeedContractMetaData.Bin

// DeployPriceFeedContract deploys a new Ethereum contract, binding an instance of PriceFeedContract to it.
func DeployPriceFeedContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceFeedContract, error) {
	parsed, err := PriceFeedContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceFeedContractBin), backend)
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

// Feed is a free data retrieval call binding the contract method 0x6d2fbecb.
//
// Solidity: function Feed(string pair) view returns(string pair, uint256 price, uint256 decimals)
func (_PriceFeedContract *PriceFeedContractCaller) Feed(opts *bind.CallOpts, pair string) (struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "Feed", pair)

	outstruct := new(struct {
		Pair     string
		Price    *big.Int
		Decimals *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pair = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Decimals = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Feed is a free data retrieval call binding the contract method 0x6d2fbecb.
//
// Solidity: function Feed(string pair) view returns(string pair, uint256 price, uint256 decimals)
func (_PriceFeedContract *PriceFeedContractSession) Feed(pair string) (struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
}, error) {
	return _PriceFeedContract.Contract.Feed(&_PriceFeedContract.CallOpts, pair)
}

// Feed is a free data retrieval call binding the contract method 0x6d2fbecb.
//
// Solidity: function Feed(string pair) view returns(string pair, uint256 price, uint256 decimals)
func (_PriceFeedContract *PriceFeedContractCallerSession) Feed(pair string) (struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
}, error) {
	return _PriceFeedContract.Contract.Feed(&_PriceFeedContract.CallOpts, pair)
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

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeedContract *PriceFeedContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeedContract *PriceFeedContractSession) ProxiableUUID() ([32]byte, error) {
	return _PriceFeedContract.Contract.ProxiableUUID(&_PriceFeedContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeedContract *PriceFeedContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PriceFeedContract.Contract.ProxiableUUID(&_PriceFeedContract.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PriceFeedContract *PriceFeedContractSession) Initialize() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Initialize(&_PriceFeedContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Initialize(&_PriceFeedContract.TransactOpts)
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

// RequestPriceFeed is a paid mutator transaction binding the contract method 0x49c14ed9.
//
// Solidity: function requestPriceFeed((string[]) request) payable returns((uint256[],uint256[]) response)
func (_PriceFeedContract *PriceFeedContractTransactor) RequestPriceFeed(opts *bind.TransactOpts, request IPriceFeedRequest) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "requestPriceFeed", request)
}

// RequestPriceFeed is a paid mutator transaction binding the contract method 0x49c14ed9.
//
// Solidity: function requestPriceFeed((string[]) request) payable returns((uint256[],uint256[]) response)
func (_PriceFeedContract *PriceFeedContractSession) RequestPriceFeed(request IPriceFeedRequest) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RequestPriceFeed(&_PriceFeedContract.TransactOpts, request)
}

// RequestPriceFeed is a paid mutator transaction binding the contract method 0x49c14ed9.
//
// Solidity: function requestPriceFeed((string[]) request) payable returns((uint256[],uint256[]) response)
func (_PriceFeedContract *PriceFeedContractTransactorSession) RequestPriceFeed(request IPriceFeedRequest) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RequestPriceFeed(&_PriceFeedContract.TransactOpts, request)
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

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0x68650ca2.
//
// Solidity: function updatePriceFeed((string,uint256,uint256) price) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) UpdatePriceFeed(opts *bind.TransactOpts, price IPriceFeedPrice) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "updatePriceFeed", price)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0x68650ca2.
//
// Solidity: function updatePriceFeed((string,uint256,uint256) price) returns()
func (_PriceFeedContract *PriceFeedContractSession) UpdatePriceFeed(price IPriceFeedPrice) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpdatePriceFeed(&_PriceFeedContract.TransactOpts, price)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0x68650ca2.
//
// Solidity: function updatePriceFeed((string,uint256,uint256) price) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) UpdatePriceFeed(price IPriceFeedPrice) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpdatePriceFeed(&_PriceFeedContract.TransactOpts, price)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PriceFeedContract *PriceFeedContractSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpgradeTo(&_PriceFeedContract.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpgradeTo(&_PriceFeedContract.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeedContract *PriceFeedContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeedContract *PriceFeedContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpgradeToAndCall(&_PriceFeedContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpgradeToAndCall(&_PriceFeedContract.TransactOpts, newImplementation, data)
}

// PriceFeedContractAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the PriceFeedContract contract.
type PriceFeedContractAdminChangedIterator struct {
	Event *PriceFeedContractAdminChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractAdminChanged)
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
		it.Event = new(PriceFeedContractAdminChanged)
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
func (it *PriceFeedContractAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractAdminChanged represents a AdminChanged event raised by the PriceFeedContract contract.
type PriceFeedContractAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PriceFeedContractAdminChangedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractAdminChangedIterator{contract: _PriceFeedContract.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedContractAdminChanged) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractAdminChanged)
				if err := _PriceFeedContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseAdminChanged(log types.Log) (*PriceFeedContractAdminChanged, error) {
	event := new(PriceFeedContractAdminChanged)
	if err := _PriceFeedContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the PriceFeedContract contract.
type PriceFeedContractBeaconUpgradedIterator struct {
	Event *PriceFeedContractBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractBeaconUpgraded)
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
		it.Event = new(PriceFeedContractBeaconUpgraded)
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
func (it *PriceFeedContractBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractBeaconUpgraded represents a BeaconUpgraded event raised by the PriceFeedContract contract.
type PriceFeedContractBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PriceFeedContractBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractBeaconUpgradedIterator{contract: _PriceFeedContract.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PriceFeedContractBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractBeaconUpgraded)
				if err := _PriceFeedContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseBeaconUpgraded(log types.Log) (*PriceFeedContractBeaconUpgraded, error) {
	event := new(PriceFeedContractBeaconUpgraded)
	if err := _PriceFeedContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PriceFeedContract contract.
type PriceFeedContractInitializedIterator struct {
	Event *PriceFeedContractInitialized // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractInitialized)
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
		it.Event = new(PriceFeedContractInitialized)
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
func (it *PriceFeedContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractInitialized represents a Initialized event raised by the PriceFeedContract contract.
type PriceFeedContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*PriceFeedContractInitializedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractInitializedIterator{contract: _PriceFeedContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PriceFeedContractInitialized) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractInitialized)
				if err := _PriceFeedContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseInitialized(log types.Log) (*PriceFeedContractInitialized, error) {
	event := new(PriceFeedContractInitialized)
	if err := _PriceFeedContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// PriceFeedContractPriceFeedPublishedIterator is returned from FilterPriceFeedPublished and is used to iterate over the raw logs and unpacked data for PriceFeedPublished events raised by the PriceFeedContract contract.
type PriceFeedContractPriceFeedPublishedIterator struct {
	Event *PriceFeedContractPriceFeedPublished // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractPriceFeedPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPriceFeedPublished)
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
		it.Event = new(PriceFeedContractPriceFeedPublished)
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
func (it *PriceFeedContractPriceFeedPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPriceFeedPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPriceFeedPublished represents a PriceFeedPublished event raised by the PriceFeedContract contract.
type PriceFeedContractPriceFeedPublished struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedPublished is a free log retrieval operation binding the contract event 0xf39d648fb8a9ac771561ab993f279c7eb13f1beb860d61c9cb3e787564d06803.
//
// Solidity: event PriceFeedPublished(string pair, uint256 price, uint256 decimals)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPriceFeedPublished(opts *bind.FilterOpts) (*PriceFeedContractPriceFeedPublishedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "PriceFeedPublished")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPriceFeedPublishedIterator{contract: _PriceFeedContract.contract, event: "PriceFeedPublished", logs: logs, sub: sub}, nil
}

// WatchPriceFeedPublished is a free log subscription operation binding the contract event 0xf39d648fb8a9ac771561ab993f279c7eb13f1beb860d61c9cb3e787564d06803.
//
// Solidity: event PriceFeedPublished(string pair, uint256 price, uint256 decimals)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPriceFeedPublished(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPriceFeedPublished) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "PriceFeedPublished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPriceFeedPublished)
				if err := _PriceFeedContract.contract.UnpackLog(event, "PriceFeedPublished", log); err != nil {
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

// ParsePriceFeedPublished is a log parse operation binding the contract event 0xf39d648fb8a9ac771561ab993f279c7eb13f1beb860d61c9cb3e787564d06803.
//
// Solidity: event PriceFeedPublished(string pair, uint256 price, uint256 decimals)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePriceFeedPublished(log types.Log) (*PriceFeedContractPriceFeedPublished, error) {
	event := new(PriceFeedContractPriceFeedPublished)
	if err := _PriceFeedContract.contract.UnpackLog(event, "PriceFeedPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractPriceFeedRequestedIterator is returned from FilterPriceFeedRequested and is used to iterate over the raw logs and unpacked data for PriceFeedRequested events raised by the PriceFeedContract contract.
type PriceFeedContractPriceFeedRequestedIterator struct {
	Event *PriceFeedContractPriceFeedRequested // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractPriceFeedRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPriceFeedRequested)
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
		it.Event = new(PriceFeedContractPriceFeedRequested)
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
func (it *PriceFeedContractPriceFeedRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPriceFeedRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPriceFeedRequested represents a PriceFeedRequested event raised by the PriceFeedContract contract.
type PriceFeedContractPriceFeedRequested struct {
	Requester common.Address
	Pairs     []string
	Prices    []*big.Int
	Decimals  []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedRequested is a free log retrieval operation binding the contract event 0x21fb481b338df488246d581cca0377c59b67cc32b0ba179ac6d2c2975432aed1.
//
// Solidity: event PriceFeedRequested(address indexed requester, string[] pairs, uint256[] prices, uint256[] decimals)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPriceFeedRequested(opts *bind.FilterOpts, requester []common.Address) (*PriceFeedContractPriceFeedRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "PriceFeedRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPriceFeedRequestedIterator{contract: _PriceFeedContract.contract, event: "PriceFeedRequested", logs: logs, sub: sub}, nil
}

// WatchPriceFeedRequested is a free log subscription operation binding the contract event 0x21fb481b338df488246d581cca0377c59b67cc32b0ba179ac6d2c2975432aed1.
//
// Solidity: event PriceFeedRequested(address indexed requester, string[] pairs, uint256[] prices, uint256[] decimals)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPriceFeedRequested(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPriceFeedRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "PriceFeedRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPriceFeedRequested)
				if err := _PriceFeedContract.contract.UnpackLog(event, "PriceFeedRequested", log); err != nil {
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

// ParsePriceFeedRequested is a log parse operation binding the contract event 0x21fb481b338df488246d581cca0377c59b67cc32b0ba179ac6d2c2975432aed1.
//
// Solidity: event PriceFeedRequested(address indexed requester, string[] pairs, uint256[] prices, uint256[] decimals)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePriceFeedRequested(log types.Log) (*PriceFeedContractPriceFeedRequested, error) {
	event := new(PriceFeedContractPriceFeedRequested)
	if err := _PriceFeedContract.contract.UnpackLog(event, "PriceFeedRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractPriceFeedUpdatedIterator is returned from FilterPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for PriceFeedUpdated events raised by the PriceFeedContract contract.
type PriceFeedContractPriceFeedUpdatedIterator struct {
	Event *PriceFeedContractPriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractPriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPriceFeedUpdated)
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
		it.Event = new(PriceFeedContractPriceFeedUpdated)
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
func (it *PriceFeedContractPriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPriceFeedUpdated represents a PriceFeedUpdated event raised by the PriceFeedContract contract.
type PriceFeedContractPriceFeedUpdated struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedUpdated is a free log retrieval operation binding the contract event 0x3e51a618566783aa8f6e83296eb48686664aeb521acd6a7507a35fb5fd9ac740.
//
// Solidity: event PriceFeedUpdated(string pair, uint256 price, uint256 decimals)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPriceFeedUpdated(opts *bind.FilterOpts) (*PriceFeedContractPriceFeedUpdatedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "PriceFeedUpdated")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPriceFeedUpdatedIterator{contract: _PriceFeedContract.contract, event: "PriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceFeedUpdated is a free log subscription operation binding the contract event 0x3e51a618566783aa8f6e83296eb48686664aeb521acd6a7507a35fb5fd9ac740.
//
// Solidity: event PriceFeedUpdated(string pair, uint256 price, uint256 decimals)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPriceFeedUpdated) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "PriceFeedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPriceFeedUpdated)
				if err := _PriceFeedContract.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
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

// ParsePriceFeedUpdated is a log parse operation binding the contract event 0x3e51a618566783aa8f6e83296eb48686664aeb521acd6a7507a35fb5fd9ac740.
//
// Solidity: event PriceFeedUpdated(string pair, uint256 price, uint256 decimals)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePriceFeedUpdated(log types.Log) (*PriceFeedContractPriceFeedUpdated, error) {
	event := new(PriceFeedContractPriceFeedUpdated)
	if err := _PriceFeedContract.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PriceFeedContract contract.
type PriceFeedContractUpgradedIterator struct {
	Event *PriceFeedContractUpgraded // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractUpgraded)
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
		it.Event = new(PriceFeedContractUpgraded)
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
func (it *PriceFeedContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractUpgraded represents a Upgraded event raised by the PriceFeedContract contract.
type PriceFeedContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PriceFeedContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractUpgradedIterator{contract: _PriceFeedContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PriceFeedContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractUpgraded)
				if err := _PriceFeedContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseUpgraded(log types.Log) (*PriceFeedContractUpgraded, error) {
	event := new(PriceFeedContractUpgraded)
	if err := _PriceFeedContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

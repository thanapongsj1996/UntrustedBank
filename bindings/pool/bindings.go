// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool

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
)

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tomToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jerryToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tomAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jerryAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tomIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jerryOut\",\"type\":\"uint256\"}],\"name\":\"swapTomforJerry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jerryAmount\",\"type\":\"uint256\"}],\"name\":\"getTomAmountByJerry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jerryIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tomOut\",\"type\":\"uint256\"}],\"name\":\"swapJerryforTom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tomAmount\",\"type\":\"uint256\"}],\"name\":\"getJerryAmountByTom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001db338038062001db3833981016040819052620000349162000186565b6040805180820182526005808252642a37b5b2b760d91b6020808401828152855180870190965292855284015281519192916200007491600391620000c3565b5080516200008a906004906020840190620000c3565b5050600580546001600160a01b039485166001600160a01b031991821617909155600680549390941692169190911790915550620001fb565b828054620000d190620001be565b90600052602060002090601f016020900481019282620000f5576000855562000140565b82601f106200011057805160ff191683800117855562000140565b8280016001018555821562000140579182015b828111156200014057825182559160200191906001019062000123565b506200014e92915062000152565b5090565b5b808211156200014e576000815560010162000153565b80516001600160a01b03811681146200018157600080fd5b919050565b600080604083850312156200019a57600080fd5b620001a58362000169565b9150620001b56020840162000169565b90509250929050565b600181811c90821680620001d357607f821691505b60208210811415620001f557634e487b7160e01b600052602260045260246000fd5b50919050565b611ba8806200020b6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806370a08231116100a2578063a0712d6811610071578063a0712d6814610235578063a457c2d714610248578063a9059cbb1461025b578063dd62ed3e1461026e578063fd196d55146102a757600080fd5b806370a08231146101de57806395d89b41146102075780639c8f9f231461020f5780639cd441da1461022257600080fd5b806322039111116100e9578063220391111461018357806323b872dd14610196578063313ce567146101a957806339509351146101b85780636a4f0fdf146101cb57600080fd5b806306fdde031461011b578063095ea7b3146101395780630a1217b41461015c57806318160ddd14610171575b600080fd5b6101236102ba565b60405161013091906118c4565b60405180910390f35b61014c610147366004611935565b61034c565b6040519015158152602001610130565b61016f61016a36600461195f565b610362565b005b6002545b604051908152602001610130565b610175610191366004611981565b6106bc565b61014c6101a436600461199a565b610868565b60405160128152602001610130565b61014c6101c6366004611935565b610912565b61016f6101d936600461195f565b61094e565b6101756101ec3660046119d6565b6001600160a01b031660009081526020819052604090205490565b610123610b24565b61016f61021d366004611981565b610b33565b61016f61023036600461195f565b610d88565b61016f610243366004611981565b611147565b61014c610256366004611935565b611154565b61014c610269366004611935565b6111ed565b61017561027c3660046119f1565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6101756102b5366004611981565b6111fa565b6060600380546102c990611a24565b80601f01602080910402602001604051908101604052809291908181526020018280546102f590611a24565b80156103425780601f1061031757610100808354040283529160200191610342565b820191906000526020600020905b81548152906001019060200180831161032557829003601f168201915b5050505050905090565b6000610359338484611393565b50600192915050565b6000821161038b5760405162461bcd60e51b815260040161038290611a5f565b60405180910390fd5b600081116103ab5760405162461bcd60e51b815260040161038290611a5f565b6005546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa1580156103f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061041d9190611a8d565b6006546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa15801561046d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104919190611a8d565b60055460405163a9059cbb60e01b8152336004820152602481018790529192506001600160a01b03169063a9059cbb906044016020604051808303816000875af11580156104e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105079190611aa6565b506006546040516323b872dd60e01b81526001600160a01b03909116906323b872dd9061053c90339087908a90600401611ac8565b6020604051808303816000875af115801561055b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057f9190611aa6565b506005546040516370a0823160e01b81526001600160a01b03858116600483015260009216906370a0823190602401602060405180830381865afa1580156105cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ef9190611a8d565b6006546040516370a0823160e01b81526001600160a01b038781166004830152929350600092909116906370a0823190602401602060405180830381865afa15801561063f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106639190611a8d565b905061066f8385611b02565b6106798284611b02565b10156106b35760405162461bcd60e51b8152602060048201526009602482015268696e76616c6964206b60b81b6044820152606401610382565b50505050505050565b600080821161070d5760405162461bcd60e51b815260206004820152601c60248201527f616d6f756e742073686f756c64206265206d6f7265207468616e2030000000006044820152606401610382565b6005546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa15801561075b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077f9190611a8d565b6006546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa1580156107cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107f39190611a8d565b90506000821180156108055750600081115b61084a5760405162461bcd60e51b8152602060048201526016602482015275696e73756666696369656e74206c697175696469747960501b6044820152606401610382565b816108558287611b02565b61085f9190611b21565b95945050505050565b60006108758484846114b8565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156108fa5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152608401610382565b6109078533858403611393565b506001949350505050565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610359918590610949908690611b43565b611393565b6000821161096e5760405162461bcd60e51b815260040161038290611a5f565b6000811161098e5760405162461bcd60e51b815260040161038290611a5f565b6005546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa1580156109dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a009190611a8d565b6006546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa158015610a50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a749190611a8d565b6005546040516323b872dd60e01b81529192506001600160a01b0316906323b872dd90610aa990339087908a90600401611ac8565b6020604051808303816000875af1158015610ac8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aec9190611aa6565b5060065460405163a9059cbb60e01b8152336004820152602481018690526001600160a01b039091169063a9059cbb9060440161053c565b6060600480546102c990611a24565b33600090815260208190526040902054811115610b625760405162461bcd60e51b815260040161038290611a5f565b6005546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa158015610bb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bd49190611a8d565b6006546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa158015610c24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c489190611a8d565b90506000610c5560025490565b9050600081610c648588611b02565b610c6e9190611b21565b9050600082610c7d8589611b02565b610c879190611b21565b9050610c933388611687565b60055460405163a9059cbb60e01b8152336004820152602481018490526001600160a01b039091169063a9059cbb906044016020604051808303816000875af1158015610ce4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d089190611aa6565b5060065460405163a9059cbb60e01b8152336004820152602481018390526001600160a01b039091169063a9059cbb906044016020604051808303816000875af1158015610d5a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7e9190611aa6565b5050505050505050565b6005546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa158015610dd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfa9190611a8d565b6006546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa158015610e4a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6e9190611a8d565b6005546040516323b872dd60e01b81529192506001600160a01b0316906323b872dd90610ea390339087908a90600401611ac8565b6020604051808303816000875af1158015610ec2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee69190611aa6565b506006546040516323b872dd60e01b81526001600160a01b03909116906323b872dd90610f1b90339087908990600401611ac8565b6020604051808303816000875af1158015610f3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5e9190611aa6565b506005546040516370a0823160e01b81526001600160a01b03858116600483015260009216906370a0823190602401602060405180830381865afa158015610faa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fce9190611a8d565b6006546040516370a0823160e01b81526001600160a01b038781166004830152929350600092909116906370a0823190602401602060405180830381865afa15801561101e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110429190611a8d565b905061104e8385611b02565b6110588284611b02565b10156110925760405162461bcd60e51b8152602060048201526009602482015268696e76616c6964206b60b81b6044820152606401610382565b600061109d60025490565b90506000816110ad5750876110e2565b6110df866110bb848c611b02565b6110c59190611b21565b866110d0858c611b02565b6110da9190611b21565b6117cd565b90505b600081116111325760405162461bcd60e51b815260206004820152601760248201527f6c70546f6b656e2073686f756c64206e6f7420626520300000000000000000006044820152606401610382565b61113c33826117e5565b505050505050505050565b61115133826117e5565b50565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156111d65760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610382565b6111e33385858403611393565b5060019392505050565b60006103593384846114b8565b600080821161124b5760405162461bcd60e51b815260206004820152601c60248201527f616d6f756e742073686f756c64206265206d6f7265207468616e2030000000006044820152606401610382565b6005546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa158015611299573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112bd9190611a8d565b6006546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa15801561130d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113319190611a8d565b90506000821180156113435750600081115b6113885760405162461bcd60e51b8152602060048201526016602482015275696e73756666696369656e74206c697175696469747960501b6044820152606401610382565b806108558387611b02565b6001600160a01b0383166113f55760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610382565b6001600160a01b0382166114565760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610382565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b03831661151c5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610382565b6001600160a01b03821661157e5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610382565b6001600160a01b038316600090815260208190526040902054818110156115f65760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610382565b6001600160a01b0380851660009081526020819052604080822085850390559185168152908120805484929061162d908490611b43565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161167991815260200190565b60405180910390a350505050565b6001600160a01b0382166116e75760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610382565b6001600160a01b0382166000908152602081905260409020548181101561175b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610382565b6001600160a01b038316600090815260208190526040812083830390556002805484929061178a908490611b5b565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016114ab565b60008183106117dc57816117de565b825b9392505050565b6001600160a01b03821661183b5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610382565b806002600082825461184d9190611b43565b90915550506001600160a01b0382166000908152602081905260408120805483929061187a908490611b43565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b600060208083528351808285015260005b818110156118f1578581018301518582016040015282016118d5565b81811115611903576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461193057600080fd5b919050565b6000806040838503121561194857600080fd5b61195183611919565b946020939093013593505050565b6000806040838503121561197257600080fd5b50508035926020909101359150565b60006020828403121561199357600080fd5b5035919050565b6000806000606084860312156119af57600080fd5b6119b884611919565b92506119c660208501611919565b9150604084013590509250925092565b6000602082840312156119e857600080fd5b6117de82611919565b60008060408385031215611a0457600080fd5b611a0d83611919565b9150611a1b60208401611919565b90509250929050565b600181811c90821680611a3857607f821691505b60208210811415611a5957634e487b7160e01b600052602260045260246000fd5b50919050565b602080825260149082015273696e73756666696369656e742062616c616e636560601b604082015260600190565b600060208284031215611a9f57600080fd5b5051919050565b600060208284031215611ab857600080fd5b815180151581146117de57600080fd5b6001600160a01b039384168152919092166020820152604081019190915260600190565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615611b1c57611b1c611aec565b500290565b600082611b3e57634e487b7160e01b600052601260045260246000fd5b500490565b60008219821115611b5657611b56611aec565b500190565b600082821015611b6d57611b6d611aec565b50039056fea2646970667358221220446ed256eab426aafd97a6172d6dcce7f6499e09f3e4545cdddfe63f83a1a20464736f6c634300080a0033",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// PoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolMetaData.Bin instead.
var PoolBin = PoolMetaData.Bin

// DeployPool deploys a new Ethereum contract, binding an instance of Pool to it.
func DeployPool(auth *bind.TransactOpts, backend bind.ContractBackend, tomToken common.Address, jerryToken common.Address) (common.Address, *types.Transaction, *Pool, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolBin), backend, tomToken, jerryToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pool *PoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pool *PoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pool.Contract.Allowance(&_Pool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pool *PoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pool.Contract.Allowance(&_Pool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pool *PoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pool *PoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pool *PoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pool *PoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pool *PoolSession) Decimals() (uint8, error) {
	return _Pool.Contract.Decimals(&_Pool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pool *PoolCallerSession) Decimals() (uint8, error) {
	return _Pool.Contract.Decimals(&_Pool.CallOpts)
}

// GetJerryAmountByTom is a free data retrieval call binding the contract method 0x22039111.
//
// Solidity: function getJerryAmountByTom(uint256 tomAmount) view returns(uint256)
func (_Pool *PoolCaller) GetJerryAmountByTom(opts *bind.CallOpts, tomAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getJerryAmountByTom", tomAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJerryAmountByTom is a free data retrieval call binding the contract method 0x22039111.
//
// Solidity: function getJerryAmountByTom(uint256 tomAmount) view returns(uint256)
func (_Pool *PoolSession) GetJerryAmountByTom(tomAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetJerryAmountByTom(&_Pool.CallOpts, tomAmount)
}

// GetJerryAmountByTom is a free data retrieval call binding the contract method 0x22039111.
//
// Solidity: function getJerryAmountByTom(uint256 tomAmount) view returns(uint256)
func (_Pool *PoolCallerSession) GetJerryAmountByTom(tomAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetJerryAmountByTom(&_Pool.CallOpts, tomAmount)
}

// GetTomAmountByJerry is a free data retrieval call binding the contract method 0xfd196d55.
//
// Solidity: function getTomAmountByJerry(uint256 jerryAmount) view returns(uint256)
func (_Pool *PoolCaller) GetTomAmountByJerry(opts *bind.CallOpts, jerryAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getTomAmountByJerry", jerryAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTomAmountByJerry is a free data retrieval call binding the contract method 0xfd196d55.
//
// Solidity: function getTomAmountByJerry(uint256 jerryAmount) view returns(uint256)
func (_Pool *PoolSession) GetTomAmountByJerry(jerryAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetTomAmountByJerry(&_Pool.CallOpts, jerryAmount)
}

// GetTomAmountByJerry is a free data retrieval call binding the contract method 0xfd196d55.
//
// Solidity: function getTomAmountByJerry(uint256 jerryAmount) view returns(uint256)
func (_Pool *PoolCallerSession) GetTomAmountByJerry(jerryAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetTomAmountByJerry(&_Pool.CallOpts, jerryAmount)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCallerSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolSession) Symbol() (string, error) {
	return _Pool.Contract.Symbol(&_Pool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolCallerSession) Symbol() (string, error) {
	return _Pool.Contract.Symbol(&_Pool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolSession) TotalSupply() (*big.Int, error) {
	return _Pool.Contract.TotalSupply(&_Pool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolCallerSession) TotalSupply() (*big.Int, error) {
	return _Pool.Contract.TotalSupply(&_Pool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 tomAmount, uint256 jerryAmount) returns()
func (_Pool *PoolTransactor) AddLiquidity(opts *bind.TransactOpts, tomAmount *big.Int, jerryAmount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addLiquidity", tomAmount, jerryAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 tomAmount, uint256 jerryAmount) returns()
func (_Pool *PoolSession) AddLiquidity(tomAmount *big.Int, jerryAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddLiquidity(&_Pool.TransactOpts, tomAmount, jerryAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 tomAmount, uint256 jerryAmount) returns()
func (_Pool *PoolTransactorSession) AddLiquidity(tomAmount *big.Int, jerryAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddLiquidity(&_Pool.TransactOpts, tomAmount, jerryAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pool *PoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Approve(&_Pool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Approve(&_Pool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Pool *PoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Pool *PoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DecreaseAllowance(&_Pool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Pool *PoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DecreaseAllowance(&_Pool.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pool *PoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pool *PoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.IncreaseAllowance(&_Pool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pool *PoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.IncreaseAllowance(&_Pool.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Pool *PoolTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Pool *PoolSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Pool *PoolTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpAmount) returns()
func (_Pool *PoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, lpAmount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "removeLiquidity", lpAmount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpAmount) returns()
func (_Pool *PoolSession) RemoveLiquidity(lpAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RemoveLiquidity(&_Pool.TransactOpts, lpAmount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpAmount) returns()
func (_Pool *PoolTransactorSession) RemoveLiquidity(lpAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RemoveLiquidity(&_Pool.TransactOpts, lpAmount)
}

// SwapJerryforTom is a paid mutator transaction binding the contract method 0x0a1217b4.
//
// Solidity: function swapJerryforTom(uint256 jerryIn, uint256 tomOut) returns()
func (_Pool *PoolTransactor) SwapJerryforTom(opts *bind.TransactOpts, jerryIn *big.Int, tomOut *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "swapJerryforTom", jerryIn, tomOut)
}

// SwapJerryforTom is a paid mutator transaction binding the contract method 0x0a1217b4.
//
// Solidity: function swapJerryforTom(uint256 jerryIn, uint256 tomOut) returns()
func (_Pool *PoolSession) SwapJerryforTom(jerryIn *big.Int, tomOut *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SwapJerryforTom(&_Pool.TransactOpts, jerryIn, tomOut)
}

// SwapJerryforTom is a paid mutator transaction binding the contract method 0x0a1217b4.
//
// Solidity: function swapJerryforTom(uint256 jerryIn, uint256 tomOut) returns()
func (_Pool *PoolTransactorSession) SwapJerryforTom(jerryIn *big.Int, tomOut *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SwapJerryforTom(&_Pool.TransactOpts, jerryIn, tomOut)
}

// SwapTomforJerry is a paid mutator transaction binding the contract method 0x6a4f0fdf.
//
// Solidity: function swapTomforJerry(uint256 tomIn, uint256 jerryOut) returns()
func (_Pool *PoolTransactor) SwapTomforJerry(opts *bind.TransactOpts, tomIn *big.Int, jerryOut *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "swapTomforJerry", tomIn, jerryOut)
}

// SwapTomforJerry is a paid mutator transaction binding the contract method 0x6a4f0fdf.
//
// Solidity: function swapTomforJerry(uint256 tomIn, uint256 jerryOut) returns()
func (_Pool *PoolSession) SwapTomforJerry(tomIn *big.Int, jerryOut *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SwapTomforJerry(&_Pool.TransactOpts, tomIn, jerryOut)
}

// SwapTomforJerry is a paid mutator transaction binding the contract method 0x6a4f0fdf.
//
// Solidity: function swapTomforJerry(uint256 tomIn, uint256 jerryOut) returns()
func (_Pool *PoolTransactorSession) SwapTomforJerry(tomIn *big.Int, jerryOut *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SwapTomforJerry(&_Pool.TransactOpts, tomIn, jerryOut)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Pool *PoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Transfer(&_Pool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Transfer(&_Pool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Pool *PoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferFrom(&_Pool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferFrom(&_Pool.TransactOpts, sender, recipient, amount)
}

// PoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pool contract.
type PoolApprovalIterator struct {
	Event *PoolApproval // Event containing the contract specifics and raw log

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
func (it *PoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolApproval)
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
		it.Event = new(PoolApproval)
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
func (it *PoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolApproval represents a Approval event raised by the Pool contract.
type PoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolApprovalIterator{contract: _Pool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolApproval)
				if err := _Pool.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) ParseApproval(log types.Log) (*PoolApproval, error) {
	event := new(PoolApproval)
	if err := _Pool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pool contract.
type PoolTransferIterator struct {
	Event *PoolTransfer // Event containing the contract specifics and raw log

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
func (it *PoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTransfer)
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
		it.Event = new(PoolTransfer)
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
func (it *PoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTransfer represents a Transfer event raised by the Pool contract.
type PoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolTransferIterator{contract: _Pool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTransfer)
				if err := _Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) ParseTransfer(log types.Log) (*PoolTransfer, error) {
	event := new(PoolTransfer)
	if err := _Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

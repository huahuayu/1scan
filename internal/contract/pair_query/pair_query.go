// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pair_query

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PairqueryABI is the input ABI used to generate the binding from.
const PairqueryABI = "[{\"inputs\":[{\"internalType\":\"contractFactory\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"getAllPairLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractFactory[]\",\"name\":\"_factorys\",\"type\":\"address[]\"}],\"name\":\"getAllPairLengthByFactorys\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPair[]\",\"name\":\"_pairs\",\"type\":\"address[]\"}],\"name\":\"getPairs\",\"outputs\":[{\"internalType\":\"address[3][]\",\"name\":\"\",\"type\":\"address[3][]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractFactory\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stop\",\"type\":\"uint256\"}],\"name\":\"getPairsByIndexRange\",\"outputs\":[{\"internalType\":\"address[3][]\",\"name\":\"\",\"type\":\"address[3][]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PairqueryBin is the compiled bytecode used for deploying new contracts.
var PairqueryBin = "0x608060405234801561001057600080fd5b50611028806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806301ac4fb21461005157806324740a2b146101b0578063ab2217e414610208578063b1a5e0e014610346575b600080fd5b6100c86004803603602081101561006757600080fd5b810190808035906020019064010000000081111561008457600080fd5b82018360208201111561009657600080fd5b803590602001918460208302840111640100000000831117156100b857600080fd5b909192939192939050505061041b565b6040518080602001806020018481526020018381038352868181518152602001915080516000925b818410156101395782846020026020010151600360200280838360005b8381101561012857808201518184015260208101905061010d565b5050505090500192600101926100f0565b925050508381038252858181518152602001915080516000925b8184101561019c5782846020026020010151600260200280838360005b8381101561018b578082015181840152602081019050610170565b505050509050019260010192610153565b925050509550505050505060405180910390f35b6101f2600480360360208110156101c657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610846565b6040518082815260200191505060405180910390f35b61025e6004803603606081101561021e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506108d0565b6040518080602001806020018481526020018381038352868181518152602001915080516000925b818410156102cf5782846020026020010151600360200280838360005b838110156102be5780820151818401526020810190506102a3565b505050509050019260010192610286565b925050508381038252858181518152602001915080516000925b818410156103325782846020026020010151600260200280838360005b83811015610321578082015181840152602081019050610306565b5050505090500192600101926102e9565b925050509550505050505060405180910390f35b6103bd6004803603602081101561035c57600080fd5b810190808035906020019064010000000081111561037957600080fd5b82018360208201111561038b57600080fd5b803590602001918460208302840111640100000000831117156103ad57600080fd5b9091929391929390505050610e6c565b6040518080602001838152602001828103825284818151815260200191508051906020019060200280838360005b838110156104065780820151818401526020810190506103eb565b50505050905001935050505060405180910390f35b60608060008085859050905060608167ffffffffffffffff8111801561044057600080fd5b5060405190808252806020026020018201604052801561047a57816020015b610467610fae565b81526020019060019003908161045f5790505b50905060608267ffffffffffffffff8111801561049657600080fd5b506040519080825280602002602001820160405280156104d057816020015b6104bd610fd0565b8152602001906001900390816104b55790505b50905060008090505b838110156108325760008989838181106104ef57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b15801561055457600080fd5b505afa158015610568573d6000803e3d6000fd5b505050506040513d602081101561057e57600080fd5b810190808051906020019092919050505084838151811061059b57fe5b60200260200101516000600381106105af57fe5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b15801561062c57600080fd5b505afa158015610640573d6000803e3d6000fd5b505050506040513d602081101561065657600080fd5b810190808051906020019092919050505084838151811061067357fe5b602002602001015160016003811061068757fe5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808483815181106106cb57fe5b60200260200101516002600381106106df57fe5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b15801561075c57600080fd5b505afa158015610770573d6000803e3d6000fd5b505050506040513d606081101561078657600080fd5b81019080805190602001909291908051906020019092919080519060200190929190505050826dffffffffffffffffffffffffffff169250816dffffffffffffffffffffffffffff169150508484815181106107de57fe5b60200260200101516000600281106107f257fe5b6020020185858151811061080257fe5b602002602001015160016002811061081657fe5b60200201828152508281525050505080806001019150506104d9565b508181439550955095505050509250925092565b60008173ffffffffffffffffffffffffffffffffffffffff1663574f2ba36040518163ffffffff1660e01b815260040160206040518083038186803b15801561088e57600080fd5b505afa1580156108a2573d6000803e3d6000fd5b505050506040513d60208110156108b857600080fd5b81019080805190602001909291905050509050919050565b6060806000808673ffffffffffffffffffffffffffffffffffffffff1663574f2ba36040518163ffffffff1660e01b815260040160206040518083038186803b15801561091c57600080fd5b505afa158015610930573d6000803e3d6000fd5b505050506040513d602081101561094657600080fd5b8101908080519060200190929190505050905080851115610965578094505b858510156109db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f73746172742063616e6e6f7420626520686967686572207468616e2073746f7081525060200191505060405180910390fd5b6000600187870301905060608167ffffffffffffffff811180156109fe57600080fd5b50604051908082528060200260200182016040528015610a3857816020015b610a25610fae565b815260200190600190039081610a1d5790505b50905060608267ffffffffffffffff81118015610a5457600080fd5b50604051908082528060200260200182016040528015610a8e57816020015b610a7b610fd0565b815260200190600190039081610a735790505b50905060008090505b83811015610e555760008b73ffffffffffffffffffffffffffffffffffffffff16631e3dd18b838d016040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610af457600080fd5b505afa158015610b08573d6000803e3d6000fd5b505050506040513d6020811015610b1e57600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015610b7757600080fd5b505afa158015610b8b573d6000803e3d6000fd5b505050506040513d6020811015610ba157600080fd5b8101908080519060200190929190505050848381518110610bbe57fe5b6020026020010151600060038110610bd257fe5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b158015610c4f57600080fd5b505afa158015610c63573d6000803e3d6000fd5b505050506040513d6020811015610c7957600080fd5b8101908080519060200190929190505050848381518110610c9657fe5b6020026020010151600160038110610caa57fe5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080848381518110610cee57fe5b6020026020010151600260038110610d0257fe5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b158015610d7f57600080fd5b505afa158015610d93573d6000803e3d6000fd5b505050506040513d6060811015610da957600080fd5b81019080805190602001909291908051906020019092919080519060200190929190505050826dffffffffffffffffffffffffffff169250816dffffffffffffffffffffffffffff16915050848481518110610e0157fe5b6020026020010151600060028110610e1557fe5b60200201858581518110610e2557fe5b6020026020010151600160028110610e3957fe5b6020020182815250828152505050508080600101915050610a97565b508181439650965096505050505093509350939050565b606060008084849050905060608167ffffffffffffffff81118015610e9057600080fd5b50604051908082528060200260200182016040528015610ebf5781602001602082028036833780820191505090505b50905060008090505b82811015610f9e57868682818110610edc57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663574f2ba36040518163ffffffff1660e01b815260040160206040518083038186803b158015610f3e57600080fd5b505afa158015610f52573d6000803e3d6000fd5b505050506040513d6020811015610f6857600080fd5b8101908080519060200190929190505050828281518110610f8557fe5b6020026020010181815250508080600101915050610ec8565b5080439350935050509250929050565b6040518060600160405280600390602082028036833780820191505090505090565b604051806040016040528060029060208202803683378082019150509050509056fea26469706673582212208f6fcc4ba8619d86e99fc2879b4392ae81d61fea04c33d5e776ee17560883c6a64736f6c63430006060033"

// DeployPairquery deploys a new Ethereum contract, binding an instance of Pairquery to it.
func DeployPairquery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairquery, error) {
	parsed, err := abi.JSON(strings.NewReader(PairqueryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairqueryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairquery{PairqueryCaller: PairqueryCaller{contract: contract}, PairqueryTransactor: PairqueryTransactor{contract: contract}, PairqueryFilterer: PairqueryFilterer{contract: contract}}, nil
}

// Pairquery is an auto generated Go binding around an Ethereum contract.
type Pairquery struct {
	PairqueryCaller     // Read-only binding to the contract
	PairqueryTransactor // Write-only binding to the contract
	PairqueryFilterer   // Log filterer for contract events
}

// PairqueryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairqueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairqueryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairqueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairqueryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairqueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairquerySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairquerySession struct {
	Contract     *Pairquery        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairqueryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairqueryCallerSession struct {
	Contract *PairqueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PairqueryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairqueryTransactorSession struct {
	Contract     *PairqueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PairqueryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairqueryRaw struct {
	Contract *Pairquery // Generic contract binding to access the raw methods on
}

// PairqueryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairqueryCallerRaw struct {
	Contract *PairqueryCaller // Generic read-only contract binding to access the raw methods on
}

// PairqueryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairqueryTransactorRaw struct {
	Contract *PairqueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairquery creates a new instance of Pairquery, bound to a specific deployed contract.
func NewPairquery(address common.Address, backend bind.ContractBackend) (*Pairquery, error) {
	contract, err := bindPairquery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairquery{PairqueryCaller: PairqueryCaller{contract: contract}, PairqueryTransactor: PairqueryTransactor{contract: contract}, PairqueryFilterer: PairqueryFilterer{contract: contract}}, nil
}

// NewPairqueryCaller creates a new read-only instance of Pairquery, bound to a specific deployed contract.
func NewPairqueryCaller(address common.Address, caller bind.ContractCaller) (*PairqueryCaller, error) {
	contract, err := bindPairquery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairqueryCaller{contract: contract}, nil
}

// NewPairqueryTransactor creates a new write-only instance of Pairquery, bound to a specific deployed contract.
func NewPairqueryTransactor(address common.Address, transactor bind.ContractTransactor) (*PairqueryTransactor, error) {
	contract, err := bindPairquery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairqueryTransactor{contract: contract}, nil
}

// NewPairqueryFilterer creates a new log filterer instance of Pairquery, bound to a specific deployed contract.
func NewPairqueryFilterer(address common.Address, filterer bind.ContractFilterer) (*PairqueryFilterer, error) {
	contract, err := bindPairquery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairqueryFilterer{contract: contract}, nil
}

// bindPairquery binds a generic wrapper to an already deployed contract.
func bindPairquery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairqueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairquery *PairqueryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairquery.Contract.PairqueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairquery *PairqueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairquery.Contract.PairqueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairquery *PairqueryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairquery.Contract.PairqueryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairquery *PairqueryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairquery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairquery *PairqueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairquery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairquery *PairqueryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairquery.Contract.contract.Transact(opts, method, params...)
}

// GetAllPairLength is a free data retrieval call binding the contract method 0x24740a2b.
//
// Solidity: function getAllPairLength(address _factory) view returns(uint256)
func (_Pairquery *PairqueryCaller) GetAllPairLength(opts *bind.CallOpts, _factory common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pairquery.contract.Call(opts, &out, "getAllPairLength", _factory)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllPairLength is a free data retrieval call binding the contract method 0x24740a2b.
//
// Solidity: function getAllPairLength(address _factory) view returns(uint256)
func (_Pairquery *PairquerySession) GetAllPairLength(_factory common.Address) (*big.Int, error) {
	return _Pairquery.Contract.GetAllPairLength(&_Pairquery.CallOpts, _factory)
}

// GetAllPairLength is a free data retrieval call binding the contract method 0x24740a2b.
//
// Solidity: function getAllPairLength(address _factory) view returns(uint256)
func (_Pairquery *PairqueryCallerSession) GetAllPairLength(_factory common.Address) (*big.Int, error) {
	return _Pairquery.Contract.GetAllPairLength(&_Pairquery.CallOpts, _factory)
}

// GetAllPairLengthByFactorys is a free data retrieval call binding the contract method 0xb1a5e0e0.
//
// Solidity: function getAllPairLengthByFactorys(address[] _factorys) view returns(uint256[], uint256)
func (_Pairquery *PairqueryCaller) GetAllPairLengthByFactorys(opts *bind.CallOpts, _factorys []common.Address) ([]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pairquery.contract.Call(opts, &out, "getAllPairLengthByFactorys", _factorys)

	if err != nil {
		return *new([]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAllPairLengthByFactorys is a free data retrieval call binding the contract method 0xb1a5e0e0.
//
// Solidity: function getAllPairLengthByFactorys(address[] _factorys) view returns(uint256[], uint256)
func (_Pairquery *PairquerySession) GetAllPairLengthByFactorys(_factorys []common.Address) ([]*big.Int, *big.Int, error) {
	return _Pairquery.Contract.GetAllPairLengthByFactorys(&_Pairquery.CallOpts, _factorys)
}

// GetAllPairLengthByFactorys is a free data retrieval call binding the contract method 0xb1a5e0e0.
//
// Solidity: function getAllPairLengthByFactorys(address[] _factorys) view returns(uint256[], uint256)
func (_Pairquery *PairqueryCallerSession) GetAllPairLengthByFactorys(_factorys []common.Address) ([]*big.Int, *big.Int, error) {
	return _Pairquery.Contract.GetAllPairLengthByFactorys(&_Pairquery.CallOpts, _factorys)
}

// GetPairs is a free data retrieval call binding the contract method 0x01ac4fb2.
//
// Solidity: function getPairs(address[] _pairs) view returns(address[3][], uint256[2][], uint256)
func (_Pairquery *PairqueryCaller) GetPairs(opts *bind.CallOpts, _pairs []common.Address) ([][3]common.Address, [][2]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pairquery.contract.Call(opts, &out, "getPairs", _pairs)

	if err != nil {
		return *new([][3]common.Address), *new([][2]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]common.Address)).(*[][3]common.Address)
	out1 := *abi.ConvertType(out[1], new([][2]*big.Int)).(*[][2]*big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPairs is a free data retrieval call binding the contract method 0x01ac4fb2.
//
// Solidity: function getPairs(address[] _pairs) view returns(address[3][], uint256[2][], uint256)
func (_Pairquery *PairquerySession) GetPairs(_pairs []common.Address) ([][3]common.Address, [][2]*big.Int, *big.Int, error) {
	return _Pairquery.Contract.GetPairs(&_Pairquery.CallOpts, _pairs)
}

// GetPairs is a free data retrieval call binding the contract method 0x01ac4fb2.
//
// Solidity: function getPairs(address[] _pairs) view returns(address[3][], uint256[2][], uint256)
func (_Pairquery *PairqueryCallerSession) GetPairs(_pairs []common.Address) ([][3]common.Address, [][2]*big.Int, *big.Int, error) {
	return _Pairquery.Contract.GetPairs(&_Pairquery.CallOpts, _pairs)
}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _factory, uint256 _start, uint256 _stop) view returns(address[3][], uint256[2][], uint256)
func (_Pairquery *PairqueryCaller) GetPairsByIndexRange(opts *bind.CallOpts, _factory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, [][2]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pairquery.contract.Call(opts, &out, "getPairsByIndexRange", _factory, _start, _stop)

	if err != nil {
		return *new([][3]common.Address), *new([][2]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]common.Address)).(*[][3]common.Address)
	out1 := *abi.ConvertType(out[1], new([][2]*big.Int)).(*[][2]*big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _factory, uint256 _start, uint256 _stop) view returns(address[3][], uint256[2][], uint256)
func (_Pairquery *PairquerySession) GetPairsByIndexRange(_factory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, [][2]*big.Int, *big.Int, error) {
	return _Pairquery.Contract.GetPairsByIndexRange(&_Pairquery.CallOpts, _factory, _start, _stop)
}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _factory, uint256 _start, uint256 _stop) view returns(address[3][], uint256[2][], uint256)
func (_Pairquery *PairqueryCallerSession) GetPairsByIndexRange(_factory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, [][2]*big.Int, *big.Int, error) {
	return _Pairquery.Contract.GetPairsByIndexRange(&_Pairquery.CallOpts, _factory, _start, _stop)
}

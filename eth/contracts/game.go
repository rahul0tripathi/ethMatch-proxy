// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// GamesigReq is an auto generated low-level Go binding around an user-defined struct.
type GamesigReq struct {
	Player        common.Address
	WinningAmount *big.Int
	InitalState   []byte
	FinalState    string
}

// Gameticket is an auto generated low-level Go binding around an user-defined struct.
type Gameticket struct {
	Id        string
	Player    common.Address
	Amount    *big.Int
	Signature []byte
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wantToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_gameTitle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxPlayers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minPlayers\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"LobbyGenerated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"LobbyResult\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structGame.ticket[]\",\"name\":\"tickets\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"lobbyId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"operatorsShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lobbyTimeout\",\"type\":\"uint256\"}],\"name\":\"createLobby\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createNewAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialDeposit\",\"type\":\"uint256\"}],\"name\":\"createNewAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSeason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositBalances\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"depositBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroyGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getLobbyPlayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_players\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEthGame\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"lobbies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"operatorsShare\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gameData\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"operatorRedeemed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPlayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minPlayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSeason\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"players\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"lobbyId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"finalState\",\"type\":\"string\"}],\"name\":\"redeemWinnings\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"winningAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initalState\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"finalState\",\"type\":\"string\"}],\"internalType\":\"structGame.sigReq[]\",\"name\":\"requests\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"lobbyId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gameData\",\"type\":\"string\"}],\"name\":\"submitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wantToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winnigAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// AccountBalances is a free data retrieval call binding the contract method 0x6ff96d17.
//
// Solidity: function accountBalances(address ) view returns(uint256)
func (_Contracts *ContractsCaller) AccountBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "accountBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalances is a free data retrieval call binding the contract method 0x6ff96d17.
//
// Solidity: function accountBalances(address ) view returns(uint256)
func (_Contracts *ContractsSession) AccountBalances(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.AccountBalances(&_Contracts.CallOpts, arg0)
}

// AccountBalances is a free data retrieval call binding the contract method 0x6ff96d17.
//
// Solidity: function accountBalances(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) AccountBalances(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.AccountBalances(&_Contracts.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts(address ) view returns(bool)
func (_Contracts *ContractsCaller) Accounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "accounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts(address ) view returns(bool)
func (_Contracts *ContractsSession) Accounts(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Accounts(&_Contracts.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts(address ) view returns(bool)
func (_Contracts *ContractsCallerSession) Accounts(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Accounts(&_Contracts.CallOpts, arg0)
}

// CurrentSeason is a free data retrieval call binding the contract method 0xbcb39621.
//
// Solidity: function currentSeason() view returns(uint256)
func (_Contracts *ContractsCaller) CurrentSeason(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "currentSeason")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSeason is a free data retrieval call binding the contract method 0xbcb39621.
//
// Solidity: function currentSeason() view returns(uint256)
func (_Contracts *ContractsSession) CurrentSeason() (*big.Int, error) {
	return _Contracts.Contract.CurrentSeason(&_Contracts.CallOpts)
}

// CurrentSeason is a free data retrieval call binding the contract method 0xbcb39621.
//
// Solidity: function currentSeason() view returns(uint256)
func (_Contracts *ContractsCallerSession) CurrentSeason() (*big.Int, error) {
	return _Contracts.Contract.CurrentSeason(&_Contracts.CallOpts)
}

// GameTitle is a free data retrieval call binding the contract method 0xb2033ce7.
//
// Solidity: function gameTitle() view returns(string)
func (_Contracts *ContractsCaller) GameTitle(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "gameTitle")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GameTitle is a free data retrieval call binding the contract method 0xb2033ce7.
//
// Solidity: function gameTitle() view returns(string)
func (_Contracts *ContractsSession) GameTitle() (string, error) {
	return _Contracts.Contract.GameTitle(&_Contracts.CallOpts)
}

// GameTitle is a free data retrieval call binding the contract method 0xb2033ce7.
//
// Solidity: function gameTitle() view returns(string)
func (_Contracts *ContractsCallerSession) GameTitle() (string, error) {
	return _Contracts.Contract.GameTitle(&_Contracts.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsSession) GetBalance() (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetBalance() (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts)
}

// GetLobbyPlayers is a free data retrieval call binding the contract method 0xc11137d1.
//
// Solidity: function getLobbyPlayers(string id) view returns(address[] _players)
func (_Contracts *ContractsCaller) GetLobbyPlayers(opts *bind.CallOpts, id string) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLobbyPlayers", id)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLobbyPlayers is a free data retrieval call binding the contract method 0xc11137d1.
//
// Solidity: function getLobbyPlayers(string id) view returns(address[] _players)
func (_Contracts *ContractsSession) GetLobbyPlayers(id string) ([]common.Address, error) {
	return _Contracts.Contract.GetLobbyPlayers(&_Contracts.CallOpts, id)
}

// GetLobbyPlayers is a free data retrieval call binding the contract method 0xc11137d1.
//
// Solidity: function getLobbyPlayers(string id) view returns(address[] _players)
func (_Contracts *ContractsCallerSession) GetLobbyPlayers(id string) ([]common.Address, error) {
	return _Contracts.Contract.GetLobbyPlayers(&_Contracts.CallOpts, id)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Contracts *ContractsCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Contracts *ContractsSession) Governance() (common.Address, error) {
	return _Contracts.Contract.Governance(&_Contracts.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Contracts *ContractsCallerSession) Governance() (common.Address, error) {
	return _Contracts.Contract.Governance(&_Contracts.CallOpts)
}

// IsEthGame is a free data retrieval call binding the contract method 0x6aecec14.
//
// Solidity: function isEthGame() view returns(bool)
func (_Contracts *ContractsCaller) IsEthGame(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isEthGame")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEthGame is a free data retrieval call binding the contract method 0x6aecec14.
//
// Solidity: function isEthGame() view returns(bool)
func (_Contracts *ContractsSession) IsEthGame() (bool, error) {
	return _Contracts.Contract.IsEthGame(&_Contracts.CallOpts)
}

// IsEthGame is a free data retrieval call binding the contract method 0x6aecec14.
//
// Solidity: function isEthGame() view returns(bool)
func (_Contracts *ContractsCallerSession) IsEthGame() (bool, error) {
	return _Contracts.Contract.IsEthGame(&_Contracts.CallOpts)
}

// Lobbies is a free data retrieval call binding the contract method 0x8f963349.
//
// Solidity: function lobbies(string ) view returns(uint256 pool, address operator, uint256 createdAt, uint256 expireAt, bool isValue, uint256 operatorsShare, string gameData, bool isCompleted, bool operatorRedeemed)
func (_Contracts *ContractsCaller) Lobbies(opts *bind.CallOpts, arg0 string) (struct {
	Pool             *big.Int
	Operator         common.Address
	CreatedAt        *big.Int
	ExpireAt         *big.Int
	IsValue          bool
	OperatorsShare   *big.Int
	GameData         string
	IsCompleted      bool
	OperatorRedeemed bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "lobbies", arg0)

	outstruct := new(struct {
		Pool             *big.Int
		Operator         common.Address
		CreatedAt        *big.Int
		ExpireAt         *big.Int
		IsValue          bool
		OperatorsShare   *big.Int
		GameData         string
		IsCompleted      bool
		OperatorRedeemed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pool = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ExpireAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsValue = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.OperatorsShare = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.GameData = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.IsCompleted = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.OperatorRedeemed = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Lobbies is a free data retrieval call binding the contract method 0x8f963349.
//
// Solidity: function lobbies(string ) view returns(uint256 pool, address operator, uint256 createdAt, uint256 expireAt, bool isValue, uint256 operatorsShare, string gameData, bool isCompleted, bool operatorRedeemed)
func (_Contracts *ContractsSession) Lobbies(arg0 string) (struct {
	Pool             *big.Int
	Operator         common.Address
	CreatedAt        *big.Int
	ExpireAt         *big.Int
	IsValue          bool
	OperatorsShare   *big.Int
	GameData         string
	IsCompleted      bool
	OperatorRedeemed bool
}, error) {
	return _Contracts.Contract.Lobbies(&_Contracts.CallOpts, arg0)
}

// Lobbies is a free data retrieval call binding the contract method 0x8f963349.
//
// Solidity: function lobbies(string ) view returns(uint256 pool, address operator, uint256 createdAt, uint256 expireAt, bool isValue, uint256 operatorsShare, string gameData, bool isCompleted, bool operatorRedeemed)
func (_Contracts *ContractsCallerSession) Lobbies(arg0 string) (struct {
	Pool             *big.Int
	Operator         common.Address
	CreatedAt        *big.Int
	ExpireAt         *big.Int
	IsValue          bool
	OperatorsShare   *big.Int
	GameData         string
	IsCompleted      bool
	OperatorRedeemed bool
}, error) {
	return _Contracts.Contract.Lobbies(&_Contracts.CallOpts, arg0)
}

// MaxPlayers is a free data retrieval call binding the contract method 0x4c2412a2.
//
// Solidity: function maxPlayers() view returns(uint256)
func (_Contracts *ContractsCaller) MaxPlayers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "maxPlayers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPlayers is a free data retrieval call binding the contract method 0x4c2412a2.
//
// Solidity: function maxPlayers() view returns(uint256)
func (_Contracts *ContractsSession) MaxPlayers() (*big.Int, error) {
	return _Contracts.Contract.MaxPlayers(&_Contracts.CallOpts)
}

// MaxPlayers is a free data retrieval call binding the contract method 0x4c2412a2.
//
// Solidity: function maxPlayers() view returns(uint256)
func (_Contracts *ContractsCallerSession) MaxPlayers() (*big.Int, error) {
	return _Contracts.Contract.MaxPlayers(&_Contracts.CallOpts)
}

// MinPlayers is a free data retrieval call binding the contract method 0x2770c895.
//
// Solidity: function minPlayers() view returns(uint256)
func (_Contracts *ContractsCaller) MinPlayers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "minPlayers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPlayers is a free data retrieval call binding the contract method 0x2770c895.
//
// Solidity: function minPlayers() view returns(uint256)
func (_Contracts *ContractsSession) MinPlayers() (*big.Int, error) {
	return _Contracts.Contract.MinPlayers(&_Contracts.CallOpts)
}

// MinPlayers is a free data retrieval call binding the contract method 0x2770c895.
//
// Solidity: function minPlayers() view returns(uint256)
func (_Contracts *ContractsCallerSession) MinPlayers() (*big.Int, error) {
	return _Contracts.Contract.MinPlayers(&_Contracts.CallOpts)
}

// OperatorBalances is a free data retrieval call binding the contract method 0xc54990b0.
//
// Solidity: function operatorBalances(address ) view returns(uint256)
func (_Contracts *ContractsCaller) OperatorBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "operatorBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorBalances is a free data retrieval call binding the contract method 0xc54990b0.
//
// Solidity: function operatorBalances(address ) view returns(uint256)
func (_Contracts *ContractsSession) OperatorBalances(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.OperatorBalances(&_Contracts.CallOpts, arg0)
}

// OperatorBalances is a free data retrieval call binding the contract method 0xc54990b0.
//
// Solidity: function operatorBalances(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) OperatorBalances(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.OperatorBalances(&_Contracts.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Contracts *ContractsCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Contracts *ContractsSession) Operators(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Operators(&_Contracts.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Contracts *ContractsCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Operators(&_Contracts.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0x2fffccb3.
//
// Solidity: function players(string , address ) view returns(uint256)
func (_Contracts *ContractsCaller) Players(opts *bind.CallOpts, arg0 string, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "players", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Players is a free data retrieval call binding the contract method 0x2fffccb3.
//
// Solidity: function players(string , address ) view returns(uint256)
func (_Contracts *ContractsSession) Players(arg0 string, arg1 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Players(&_Contracts.CallOpts, arg0, arg1)
}

// Players is a free data retrieval call binding the contract method 0x2fffccb3.
//
// Solidity: function players(string , address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) Players(arg0 string, arg1 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Players(&_Contracts.CallOpts, arg0, arg1)
}

// WantToken is a free data retrieval call binding the contract method 0xd23e0480.
//
// Solidity: function wantToken() view returns(address)
func (_Contracts *ContractsCaller) WantToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "wantToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WantToken is a free data retrieval call binding the contract method 0xd23e0480.
//
// Solidity: function wantToken() view returns(address)
func (_Contracts *ContractsSession) WantToken() (common.Address, error) {
	return _Contracts.Contract.WantToken(&_Contracts.CallOpts)
}

// WantToken is a free data retrieval call binding the contract method 0xd23e0480.
//
// Solidity: function wantToken() view returns(address)
func (_Contracts *ContractsCallerSession) WantToken() (common.Address, error) {
	return _Contracts.Contract.WantToken(&_Contracts.CallOpts)
}

// WinnigAmount is a free data retrieval call binding the contract method 0x09d802ab.
//
// Solidity: function winnigAmount(string , uint256 ) view returns(uint256)
func (_Contracts *ContractsCaller) WinnigAmount(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "winnigAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinnigAmount is a free data retrieval call binding the contract method 0x09d802ab.
//
// Solidity: function winnigAmount(string , uint256 ) view returns(uint256)
func (_Contracts *ContractsSession) WinnigAmount(arg0 string, arg1 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.WinnigAmount(&_Contracts.CallOpts, arg0, arg1)
}

// WinnigAmount is a free data retrieval call binding the contract method 0x09d802ab.
//
// Solidity: function winnigAmount(string , uint256 ) view returns(uint256)
func (_Contracts *ContractsCallerSession) WinnigAmount(arg0 string, arg1 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.WinnigAmount(&_Contracts.CallOpts, arg0, arg1)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_Contracts *ContractsTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addOperator", operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_Contracts *ContractsSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddOperator(&_Contracts.TransactOpts, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_Contracts *ContractsTransactorSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddOperator(&_Contracts.TransactOpts, operator)
}

// CreateLobby is a paid mutator transaction binding the contract method 0x36f8b25b.
//
// Solidity: function createLobby((string,address,uint256,bytes)[] tickets, string lobbyId, uint256 operatorsShare, uint256 lobbyTimeout) returns(bool)
func (_Contracts *ContractsTransactor) CreateLobby(opts *bind.TransactOpts, tickets []Gameticket, lobbyId string, operatorsShare *big.Int, lobbyTimeout *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createLobby", tickets, lobbyId, operatorsShare, lobbyTimeout)
}

// CreateLobby is a paid mutator transaction binding the contract method 0x36f8b25b.
//
// Solidity: function createLobby((string,address,uint256,bytes)[] tickets, string lobbyId, uint256 operatorsShare, uint256 lobbyTimeout) returns(bool)
func (_Contracts *ContractsSession) CreateLobby(tickets []Gameticket, lobbyId string, operatorsShare *big.Int, lobbyTimeout *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateLobby(&_Contracts.TransactOpts, tickets, lobbyId, operatorsShare, lobbyTimeout)
}

// CreateLobby is a paid mutator transaction binding the contract method 0x36f8b25b.
//
// Solidity: function createLobby((string,address,uint256,bytes)[] tickets, string lobbyId, uint256 operatorsShare, uint256 lobbyTimeout) returns(bool)
func (_Contracts *ContractsTransactorSession) CreateLobby(tickets []Gameticket, lobbyId string, operatorsShare *big.Int, lobbyTimeout *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateLobby(&_Contracts.TransactOpts, tickets, lobbyId, operatorsShare, lobbyTimeout)
}

// CreateNewAccount is a paid mutator transaction binding the contract method 0x0870035f.
//
// Solidity: function createNewAccount() payable returns()
func (_Contracts *ContractsTransactor) CreateNewAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createNewAccount")
}

// CreateNewAccount is a paid mutator transaction binding the contract method 0x0870035f.
//
// Solidity: function createNewAccount() payable returns()
func (_Contracts *ContractsSession) CreateNewAccount() (*types.Transaction, error) {
	return _Contracts.Contract.CreateNewAccount(&_Contracts.TransactOpts)
}

// CreateNewAccount is a paid mutator transaction binding the contract method 0x0870035f.
//
// Solidity: function createNewAccount() payable returns()
func (_Contracts *ContractsTransactorSession) CreateNewAccount() (*types.Transaction, error) {
	return _Contracts.Contract.CreateNewAccount(&_Contracts.TransactOpts)
}

// CreateNewAccount0 is a paid mutator transaction binding the contract method 0x85a237a9.
//
// Solidity: function createNewAccount(uint256 _initialDeposit) returns()
func (_Contracts *ContractsTransactor) CreateNewAccount0(opts *bind.TransactOpts, _initialDeposit *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createNewAccount0", _initialDeposit)
}

// CreateNewAccount0 is a paid mutator transaction binding the contract method 0x85a237a9.
//
// Solidity: function createNewAccount(uint256 _initialDeposit) returns()
func (_Contracts *ContractsSession) CreateNewAccount0(_initialDeposit *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateNewAccount0(&_Contracts.TransactOpts, _initialDeposit)
}

// CreateNewAccount0 is a paid mutator transaction binding the contract method 0x85a237a9.
//
// Solidity: function createNewAccount(uint256 _initialDeposit) returns()
func (_Contracts *ContractsTransactorSession) CreateNewAccount0(_initialDeposit *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateNewAccount0(&_Contracts.TransactOpts, _initialDeposit)
}

// DepositBalances is a paid mutator transaction binding the contract method 0x35747589.
//
// Solidity: function depositBalances() payable returns()
func (_Contracts *ContractsTransactor) DepositBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositBalances")
}

// DepositBalances is a paid mutator transaction binding the contract method 0x35747589.
//
// Solidity: function depositBalances() payable returns()
func (_Contracts *ContractsSession) DepositBalances() (*types.Transaction, error) {
	return _Contracts.Contract.DepositBalances(&_Contracts.TransactOpts)
}

// DepositBalances is a paid mutator transaction binding the contract method 0x35747589.
//
// Solidity: function depositBalances() payable returns()
func (_Contracts *ContractsTransactorSession) DepositBalances() (*types.Transaction, error) {
	return _Contracts.Contract.DepositBalances(&_Contracts.TransactOpts)
}

// DepositBalances0 is a paid mutator transaction binding the contract method 0x9087b69b.
//
// Solidity: function depositBalances(uint256 _amountIn) returns()
func (_Contracts *ContractsTransactor) DepositBalances0(opts *bind.TransactOpts, _amountIn *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositBalances0", _amountIn)
}

// DepositBalances0 is a paid mutator transaction binding the contract method 0x9087b69b.
//
// Solidity: function depositBalances(uint256 _amountIn) returns()
func (_Contracts *ContractsSession) DepositBalances0(_amountIn *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositBalances0(&_Contracts.TransactOpts, _amountIn)
}

// DepositBalances0 is a paid mutator transaction binding the contract method 0x9087b69b.
//
// Solidity: function depositBalances(uint256 _amountIn) returns()
func (_Contracts *ContractsTransactorSession) DepositBalances0(_amountIn *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositBalances0(&_Contracts.TransactOpts, _amountIn)
}

// DestroyGovernance is a paid mutator transaction binding the contract method 0x2aa05d36.
//
// Solidity: function destroyGovernance() returns()
func (_Contracts *ContractsTransactor) DestroyGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "destroyGovernance")
}

// DestroyGovernance is a paid mutator transaction binding the contract method 0x2aa05d36.
//
// Solidity: function destroyGovernance() returns()
func (_Contracts *ContractsSession) DestroyGovernance() (*types.Transaction, error) {
	return _Contracts.Contract.DestroyGovernance(&_Contracts.TransactOpts)
}

// DestroyGovernance is a paid mutator transaction binding the contract method 0x2aa05d36.
//
// Solidity: function destroyGovernance() returns()
func (_Contracts *ContractsTransactorSession) DestroyGovernance() (*types.Transaction, error) {
	return _Contracts.Contract.DestroyGovernance(&_Contracts.TransactOpts)
}

// GetOperatorsBalance is a paid mutator transaction binding the contract method 0x707112ca.
//
// Solidity: function getOperatorsBalance() returns(uint256)
func (_Contracts *ContractsTransactor) GetOperatorsBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "getOperatorsBalance")
}

// GetOperatorsBalance is a paid mutator transaction binding the contract method 0x707112ca.
//
// Solidity: function getOperatorsBalance() returns(uint256)
func (_Contracts *ContractsSession) GetOperatorsBalance() (*types.Transaction, error) {
	return _Contracts.Contract.GetOperatorsBalance(&_Contracts.TransactOpts)
}

// GetOperatorsBalance is a paid mutator transaction binding the contract method 0x707112ca.
//
// Solidity: function getOperatorsBalance() returns(uint256)
func (_Contracts *ContractsTransactorSession) GetOperatorsBalance() (*types.Transaction, error) {
	return _Contracts.Contract.GetOperatorsBalance(&_Contracts.TransactOpts)
}

// NewSeason is a paid mutator transaction binding the contract method 0x1974e0bc.
//
// Solidity: function newSeason() returns()
func (_Contracts *ContractsTransactor) NewSeason(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "newSeason")
}

// NewSeason is a paid mutator transaction binding the contract method 0x1974e0bc.
//
// Solidity: function newSeason() returns()
func (_Contracts *ContractsSession) NewSeason() (*types.Transaction, error) {
	return _Contracts.Contract.NewSeason(&_Contracts.TransactOpts)
}

// NewSeason is a paid mutator transaction binding the contract method 0x1974e0bc.
//
// Solidity: function newSeason() returns()
func (_Contracts *ContractsTransactorSession) NewSeason() (*types.Transaction, error) {
	return _Contracts.Contract.NewSeason(&_Contracts.TransactOpts)
}

// RedeemWinnings is a paid mutator transaction binding the contract method 0xf2213827.
//
// Solidity: function redeemWinnings(string lobbyId, string finalState) returns(bool)
func (_Contracts *ContractsTransactor) RedeemWinnings(opts *bind.TransactOpts, lobbyId string, finalState string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "redeemWinnings", lobbyId, finalState)
}

// RedeemWinnings is a paid mutator transaction binding the contract method 0xf2213827.
//
// Solidity: function redeemWinnings(string lobbyId, string finalState) returns(bool)
func (_Contracts *ContractsSession) RedeemWinnings(lobbyId string, finalState string) (*types.Transaction, error) {
	return _Contracts.Contract.RedeemWinnings(&_Contracts.TransactOpts, lobbyId, finalState)
}

// RedeemWinnings is a paid mutator transaction binding the contract method 0xf2213827.
//
// Solidity: function redeemWinnings(string lobbyId, string finalState) returns(bool)
func (_Contracts *ContractsTransactorSession) RedeemWinnings(lobbyId string, finalState string) (*types.Transaction, error) {
	return _Contracts.Contract.RedeemWinnings(&_Contracts.TransactOpts, lobbyId, finalState)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x9a9751d1.
//
// Solidity: function submitResults((address,uint256,bytes,string)[] requests, string lobbyId, string gameData) returns()
func (_Contracts *ContractsTransactor) SubmitResults(opts *bind.TransactOpts, requests []GamesigReq, lobbyId string, gameData string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submitResults", requests, lobbyId, gameData)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x9a9751d1.
//
// Solidity: function submitResults((address,uint256,bytes,string)[] requests, string lobbyId, string gameData) returns()
func (_Contracts *ContractsSession) SubmitResults(requests []GamesigReq, lobbyId string, gameData string) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitResults(&_Contracts.TransactOpts, requests, lobbyId, gameData)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x9a9751d1.
//
// Solidity: function submitResults((address,uint256,bytes,string)[] requests, string lobbyId, string gameData) returns()
func (_Contracts *ContractsTransactorSession) SubmitResults(requests []GamesigReq, lobbyId string, gameData string) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitResults(&_Contracts.TransactOpts, requests, lobbyId, gameData)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_Contracts *ContractsTransactor) TransferGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferGovernance", _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_Contracts *ContractsSession) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferGovernance(&_Contracts.TransactOpts, _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_Contracts *ContractsTransactorSession) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferGovernance(&_Contracts.TransactOpts, _governance)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_Contracts *ContractsTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_Contracts *ContractsSession) WithdrawBalance() (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawBalance(&_Contracts.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_Contracts *ContractsTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawBalance(&_Contracts.TransactOpts)
}

// ContractsLobbyGeneratedIterator is returned from FilterLobbyGenerated and is used to iterate over the raw logs and unpacked data for LobbyGenerated events raised by the Contracts contract.
type ContractsLobbyGeneratedIterator struct {
	Event *ContractsLobbyGenerated // Event containing the contract specifics and raw log

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
func (it *ContractsLobbyGeneratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLobbyGenerated)
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
		it.Event = new(ContractsLobbyGenerated)
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
func (it *ContractsLobbyGeneratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLobbyGeneratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLobbyGenerated represents a LobbyGenerated event raised by the Contracts contract.
type ContractsLobbyGenerated struct {
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLobbyGenerated is a free log retrieval operation binding the contract event 0xaf91b74b6066b038696c3c161cfc914d896dab8365fdcf2792fac3550cf385c6.
//
// Solidity: event LobbyGenerated(string id)
func (_Contracts *ContractsFilterer) FilterLobbyGenerated(opts *bind.FilterOpts) (*ContractsLobbyGeneratedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LobbyGenerated")
	if err != nil {
		return nil, err
	}
	return &ContractsLobbyGeneratedIterator{contract: _Contracts.contract, event: "LobbyGenerated", logs: logs, sub: sub}, nil
}

// WatchLobbyGenerated is a free log subscription operation binding the contract event 0xaf91b74b6066b038696c3c161cfc914d896dab8365fdcf2792fac3550cf385c6.
//
// Solidity: event LobbyGenerated(string id)
func (_Contracts *ContractsFilterer) WatchLobbyGenerated(opts *bind.WatchOpts, sink chan<- *ContractsLobbyGenerated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LobbyGenerated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLobbyGenerated)
				if err := _Contracts.contract.UnpackLog(event, "LobbyGenerated", log); err != nil {
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

// ParseLobbyGenerated is a log parse operation binding the contract event 0xaf91b74b6066b038696c3c161cfc914d896dab8365fdcf2792fac3550cf385c6.
//
// Solidity: event LobbyGenerated(string id)
func (_Contracts *ContractsFilterer) ParseLobbyGenerated(log types.Log) (*ContractsLobbyGenerated, error) {
	event := new(ContractsLobbyGenerated)
	if err := _Contracts.contract.UnpackLog(event, "LobbyGenerated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLobbyResultIterator is returned from FilterLobbyResult and is used to iterate over the raw logs and unpacked data for LobbyResult events raised by the Contracts contract.
type ContractsLobbyResultIterator struct {
	Event *ContractsLobbyResult // Event containing the contract specifics and raw log

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
func (it *ContractsLobbyResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLobbyResult)
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
		it.Event = new(ContractsLobbyResult)
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
func (it *ContractsLobbyResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLobbyResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLobbyResult represents a LobbyResult event raised by the Contracts contract.
type ContractsLobbyResult struct {
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLobbyResult is a free log retrieval operation binding the contract event 0x153fffcdd7178eec416a270a61b54a20db181360c9507267e9cecdce5aabad3e.
//
// Solidity: event LobbyResult(string id)
func (_Contracts *ContractsFilterer) FilterLobbyResult(opts *bind.FilterOpts) (*ContractsLobbyResultIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LobbyResult")
	if err != nil {
		return nil, err
	}
	return &ContractsLobbyResultIterator{contract: _Contracts.contract, event: "LobbyResult", logs: logs, sub: sub}, nil
}

// WatchLobbyResult is a free log subscription operation binding the contract event 0x153fffcdd7178eec416a270a61b54a20db181360c9507267e9cecdce5aabad3e.
//
// Solidity: event LobbyResult(string id)
func (_Contracts *ContractsFilterer) WatchLobbyResult(opts *bind.WatchOpts, sink chan<- *ContractsLobbyResult) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LobbyResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLobbyResult)
				if err := _Contracts.contract.UnpackLog(event, "LobbyResult", log); err != nil {
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

// ParseLobbyResult is a log parse operation binding the contract event 0x153fffcdd7178eec416a270a61b54a20db181360c9507267e9cecdce5aabad3e.
//
// Solidity: event LobbyResult(string id)
func (_Contracts *ContractsFilterer) ParseLobbyResult(log types.Log) (*ContractsLobbyResult, error) {
	event := new(ContractsLobbyResult)
	if err := _Contracts.contract.UnpackLog(event, "LobbyResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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
	_ = abi.ConvertType
)

// DSCEngineMetaData contains all meta data concerning the DSCEngine contract.
var DSCEngineMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"burnDsc\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateHealthFactor\",\"inputs\":[{\"name\":\"totalDscMinted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralValueInUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"depositCollateral\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateralAndMintDsc\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountDscToMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateralAndMintDscWithPermit\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountDscToMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateralWithPermit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccountCollateralValueInUsd\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalCollateralValueInUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccountInfo\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalDscMinted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralValueInUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollateralAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollateralBalanceOfUser\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollateralTokenPriceFeed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollateralTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDsc\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHealthFactor\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLiquidationBonus\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getLiquidationPrecision\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getLiquidationThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getMinHealthFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTokenAmountFromUsd\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdAmountInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUsdValue\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"tokenAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"priceFeedAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"dscAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtToCover\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintDsc\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCollateral\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCollateralForDsc\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountDscToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollateralDeposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralRedeemed\",\"inputs\":[{\"name\":\"redeemFrom\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"redeemTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DSCEngine_BreaksHealthFactor\",\"inputs\":[{\"name\":\"healthFactorValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"DSCEngine_HealthFactorNotImproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine_HealthFactorOk\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine_MintFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine_NeedsMoreThanZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine_TokenAddressesAndPriceFeedAddressesMustBeSameLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine_TokenNotAllowed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DSCEngine_TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// DSCEngineABI is the input ABI used to generate the binding from.
// Deprecated: Use DSCEngineMetaData.ABI instead.
var DSCEngineABI = DSCEngineMetaData.ABI

// DSCEngine is an auto generated Go binding around an Ethereum contract.
type DSCEngine struct {
	DSCEngineCaller     // Read-only binding to the contract
	DSCEngineTransactor // Write-only binding to the contract
	DSCEngineFilterer   // Log filterer for contract events
}

// DSCEngineCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSCEngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSCEngineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSCEngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSCEngineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSCEngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSCEngineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSCEngineSession struct {
	Contract     *DSCEngine        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSCEngineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSCEngineCallerSession struct {
	Contract *DSCEngineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DSCEngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSCEngineTransactorSession struct {
	Contract     *DSCEngineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DSCEngineRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSCEngineRaw struct {
	Contract *DSCEngine // Generic contract binding to access the raw methods on
}

// DSCEngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSCEngineCallerRaw struct {
	Contract *DSCEngineCaller // Generic read-only contract binding to access the raw methods on
}

// DSCEngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSCEngineTransactorRaw struct {
	Contract *DSCEngineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSCEngine creates a new instance of DSCEngine, bound to a specific deployed contract.
func NewDSCEngine(address common.Address, backend bind.ContractBackend) (*DSCEngine, error) {
	contract, err := bindDSCEngine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSCEngine{DSCEngineCaller: DSCEngineCaller{contract: contract}, DSCEngineTransactor: DSCEngineTransactor{contract: contract}, DSCEngineFilterer: DSCEngineFilterer{contract: contract}}, nil
}

// NewDSCEngineCaller creates a new read-only instance of DSCEngine, bound to a specific deployed contract.
func NewDSCEngineCaller(address common.Address, caller bind.ContractCaller) (*DSCEngineCaller, error) {
	contract, err := bindDSCEngine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSCEngineCaller{contract: contract}, nil
}

// NewDSCEngineTransactor creates a new write-only instance of DSCEngine, bound to a specific deployed contract.
func NewDSCEngineTransactor(address common.Address, transactor bind.ContractTransactor) (*DSCEngineTransactor, error) {
	contract, err := bindDSCEngine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSCEngineTransactor{contract: contract}, nil
}

// NewDSCEngineFilterer creates a new log filterer instance of DSCEngine, bound to a specific deployed contract.
func NewDSCEngineFilterer(address common.Address, filterer bind.ContractFilterer) (*DSCEngineFilterer, error) {
	contract, err := bindDSCEngine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSCEngineFilterer{contract: contract}, nil
}

// bindDSCEngine binds a generic wrapper to an already deployed contract.
func bindDSCEngine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DSCEngineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSCEngine *DSCEngineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DSCEngine.Contract.DSCEngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSCEngine *DSCEngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSCEngine.Contract.DSCEngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSCEngine *DSCEngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSCEngine.Contract.DSCEngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSCEngine *DSCEngineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DSCEngine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSCEngine *DSCEngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSCEngine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSCEngine *DSCEngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSCEngine.Contract.contract.Transact(opts, method, params...)
}

// CalculateHealthFactor is a free data retrieval call binding the contract method 0x01f72884.
//
// Solidity: function calculateHealthFactor(uint256 totalDscMinted, uint256 collateralValueInUsd) pure returns(uint256)
func (_DSCEngine *DSCEngineCaller) CalculateHealthFactor(opts *bind.CallOpts, totalDscMinted *big.Int, collateralValueInUsd *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "calculateHealthFactor", totalDscMinted, collateralValueInUsd)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateHealthFactor is a free data retrieval call binding the contract method 0x01f72884.
//
// Solidity: function calculateHealthFactor(uint256 totalDscMinted, uint256 collateralValueInUsd) pure returns(uint256)
func (_DSCEngine *DSCEngineSession) CalculateHealthFactor(totalDscMinted *big.Int, collateralValueInUsd *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.CalculateHealthFactor(&_DSCEngine.CallOpts, totalDscMinted, collateralValueInUsd)
}

// CalculateHealthFactor is a free data retrieval call binding the contract method 0x01f72884.
//
// Solidity: function calculateHealthFactor(uint256 totalDscMinted, uint256 collateralValueInUsd) pure returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) CalculateHealthFactor(totalDscMinted *big.Int, collateralValueInUsd *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.CalculateHealthFactor(&_DSCEngine.CallOpts, totalDscMinted, collateralValueInUsd)
}

// GetAccountCollateralValueInUsd is a free data retrieval call binding the contract method 0x545af4fe.
//
// Solidity: function getAccountCollateralValueInUsd(address user) view returns(uint256 totalCollateralValueInUsd)
func (_DSCEngine *DSCEngineCaller) GetAccountCollateralValueInUsd(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getAccountCollateralValueInUsd", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountCollateralValueInUsd is a free data retrieval call binding the contract method 0x545af4fe.
//
// Solidity: function getAccountCollateralValueInUsd(address user) view returns(uint256 totalCollateralValueInUsd)
func (_DSCEngine *DSCEngineSession) GetAccountCollateralValueInUsd(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetAccountCollateralValueInUsd(&_DSCEngine.CallOpts, user)
}

// GetAccountCollateralValueInUsd is a free data retrieval call binding the contract method 0x545af4fe.
//
// Solidity: function getAccountCollateralValueInUsd(address user) view returns(uint256 totalCollateralValueInUsd)
func (_DSCEngine *DSCEngineCallerSession) GetAccountCollateralValueInUsd(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetAccountCollateralValueInUsd(&_DSCEngine.CallOpts, user)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address user) view returns(uint256 totalDscMinted, uint256 collateralValueInUsd)
func (_DSCEngine *DSCEngineCaller) GetAccountInfo(opts *bind.CallOpts, user common.Address) (struct {
	TotalDscMinted       *big.Int
	CollateralValueInUsd *big.Int
}, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getAccountInfo", user)

	outstruct := new(struct {
		TotalDscMinted       *big.Int
		CollateralValueInUsd *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalDscMinted = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CollateralValueInUsd = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address user) view returns(uint256 totalDscMinted, uint256 collateralValueInUsd)
func (_DSCEngine *DSCEngineSession) GetAccountInfo(user common.Address) (struct {
	TotalDscMinted       *big.Int
	CollateralValueInUsd *big.Int
}, error) {
	return _DSCEngine.Contract.GetAccountInfo(&_DSCEngine.CallOpts, user)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address user) view returns(uint256 totalDscMinted, uint256 collateralValueInUsd)
func (_DSCEngine *DSCEngineCallerSession) GetAccountInfo(user common.Address) (struct {
	TotalDscMinted       *big.Int
	CollateralValueInUsd *big.Int
}, error) {
	return _DSCEngine.Contract.GetAccountInfo(&_DSCEngine.CallOpts, user)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x654c6295.
//
// Solidity: function getCollateralAmount(address token) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetCollateralAmount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getCollateralAmount", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x654c6295.
//
// Solidity: function getCollateralAmount(address token) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetCollateralAmount(token common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetCollateralAmount(&_DSCEngine.CallOpts, token)
}

// GetCollateralAmount is a free data retrieval call binding the contract method 0x654c6295.
//
// Solidity: function getCollateralAmount(address token) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetCollateralAmount(token common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetCollateralAmount(&_DSCEngine.CallOpts, token)
}

// GetCollateralBalanceOfUser is a free data retrieval call binding the contract method 0x31e92b83.
//
// Solidity: function getCollateralBalanceOfUser(address user, address collateralAddress) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetCollateralBalanceOfUser(opts *bind.CallOpts, user common.Address, collateralAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getCollateralBalanceOfUser", user, collateralAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralBalanceOfUser is a free data retrieval call binding the contract method 0x31e92b83.
//
// Solidity: function getCollateralBalanceOfUser(address user, address collateralAddress) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetCollateralBalanceOfUser(user common.Address, collateralAddress common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetCollateralBalanceOfUser(&_DSCEngine.CallOpts, user, collateralAddress)
}

// GetCollateralBalanceOfUser is a free data retrieval call binding the contract method 0x31e92b83.
//
// Solidity: function getCollateralBalanceOfUser(address user, address collateralAddress) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetCollateralBalanceOfUser(user common.Address, collateralAddress common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetCollateralBalanceOfUser(&_DSCEngine.CallOpts, user, collateralAddress)
}

// GetCollateralTokenPriceFeed is a free data retrieval call binding the contract method 0x1c08adda.
//
// Solidity: function getCollateralTokenPriceFeed(address token) view returns(address)
func (_DSCEngine *DSCEngineCaller) GetCollateralTokenPriceFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getCollateralTokenPriceFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateralTokenPriceFeed is a free data retrieval call binding the contract method 0x1c08adda.
//
// Solidity: function getCollateralTokenPriceFeed(address token) view returns(address)
func (_DSCEngine *DSCEngineSession) GetCollateralTokenPriceFeed(token common.Address) (common.Address, error) {
	return _DSCEngine.Contract.GetCollateralTokenPriceFeed(&_DSCEngine.CallOpts, token)
}

// GetCollateralTokenPriceFeed is a free data retrieval call binding the contract method 0x1c08adda.
//
// Solidity: function getCollateralTokenPriceFeed(address token) view returns(address)
func (_DSCEngine *DSCEngineCallerSession) GetCollateralTokenPriceFeed(token common.Address) (common.Address, error) {
	return _DSCEngine.Contract.GetCollateralTokenPriceFeed(&_DSCEngine.CallOpts, token)
}

// GetCollateralTokens is a free data retrieval call binding the contract method 0xb58eb63f.
//
// Solidity: function getCollateralTokens() view returns(address[])
func (_DSCEngine *DSCEngineCaller) GetCollateralTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getCollateralTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCollateralTokens is a free data retrieval call binding the contract method 0xb58eb63f.
//
// Solidity: function getCollateralTokens() view returns(address[])
func (_DSCEngine *DSCEngineSession) GetCollateralTokens() ([]common.Address, error) {
	return _DSCEngine.Contract.GetCollateralTokens(&_DSCEngine.CallOpts)
}

// GetCollateralTokens is a free data retrieval call binding the contract method 0xb58eb63f.
//
// Solidity: function getCollateralTokens() view returns(address[])
func (_DSCEngine *DSCEngineCallerSession) GetCollateralTokens() ([]common.Address, error) {
	return _DSCEngine.Contract.GetCollateralTokens(&_DSCEngine.CallOpts)
}

// GetDsc is a free data retrieval call binding the contract method 0xdeb8e018.
//
// Solidity: function getDsc() view returns(address)
func (_DSCEngine *DSCEngineCaller) GetDsc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getDsc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDsc is a free data retrieval call binding the contract method 0xdeb8e018.
//
// Solidity: function getDsc() view returns(address)
func (_DSCEngine *DSCEngineSession) GetDsc() (common.Address, error) {
	return _DSCEngine.Contract.GetDsc(&_DSCEngine.CallOpts)
}

// GetDsc is a free data retrieval call binding the contract method 0xdeb8e018.
//
// Solidity: function getDsc() view returns(address)
func (_DSCEngine *DSCEngineCallerSession) GetDsc() (common.Address, error) {
	return _DSCEngine.Contract.GetDsc(&_DSCEngine.CallOpts)
}

// GetHealthFactor is a free data retrieval call binding the contract method 0xfe6bcd7c.
//
// Solidity: function getHealthFactor(address user) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetHealthFactor(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getHealthFactor", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHealthFactor is a free data retrieval call binding the contract method 0xfe6bcd7c.
//
// Solidity: function getHealthFactor(address user) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetHealthFactor(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetHealthFactor(&_DSCEngine.CallOpts, user)
}

// GetHealthFactor is a free data retrieval call binding the contract method 0xfe6bcd7c.
//
// Solidity: function getHealthFactor(address user) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetHealthFactor(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetHealthFactor(&_DSCEngine.CallOpts, user)
}

// GetLiquidationBonus is a free data retrieval call binding the contract method 0x59aa9e72.
//
// Solidity: function getLiquidationBonus() pure returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetLiquidationBonus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getLiquidationBonus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationBonus is a free data retrieval call binding the contract method 0x59aa9e72.
//
// Solidity: function getLiquidationBonus() pure returns(uint256)
func (_DSCEngine *DSCEngineSession) GetLiquidationBonus() (*big.Int, error) {
	return _DSCEngine.Contract.GetLiquidationBonus(&_DSCEngine.CallOpts)
}

// GetLiquidationBonus is a free data retrieval call binding the contract method 0x59aa9e72.
//
// Solidity: function getLiquidationBonus() pure returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetLiquidationBonus() (*big.Int, error) {
	return _DSCEngine.Contract.GetLiquidationBonus(&_DSCEngine.CallOpts)
}

// GetLiquidationPrecision is a free data retrieval call binding the contract method 0x6c8102c0.
//
// Solidity: function getLiquidationPrecision() pure returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetLiquidationPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getLiquidationPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationPrecision is a free data retrieval call binding the contract method 0x6c8102c0.
//
// Solidity: function getLiquidationPrecision() pure returns(uint256)
func (_DSCEngine *DSCEngineSession) GetLiquidationPrecision() (*big.Int, error) {
	return _DSCEngine.Contract.GetLiquidationPrecision(&_DSCEngine.CallOpts)
}

// GetLiquidationPrecision is a free data retrieval call binding the contract method 0x6c8102c0.
//
// Solidity: function getLiquidationPrecision() pure returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetLiquidationPrecision() (*big.Int, error) {
	return _DSCEngine.Contract.GetLiquidationPrecision(&_DSCEngine.CallOpts)
}

// GetLiquidationThreshold is a free data retrieval call binding the contract method 0x4ae9b8bc.
//
// Solidity: function getLiquidationThreshold() pure returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetLiquidationThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getLiquidationThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationThreshold is a free data retrieval call binding the contract method 0x4ae9b8bc.
//
// Solidity: function getLiquidationThreshold() pure returns(uint256)
func (_DSCEngine *DSCEngineSession) GetLiquidationThreshold() (*big.Int, error) {
	return _DSCEngine.Contract.GetLiquidationThreshold(&_DSCEngine.CallOpts)
}

// GetLiquidationThreshold is a free data retrieval call binding the contract method 0x4ae9b8bc.
//
// Solidity: function getLiquidationThreshold() pure returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetLiquidationThreshold() (*big.Int, error) {
	return _DSCEngine.Contract.GetLiquidationThreshold(&_DSCEngine.CallOpts)
}

// GetMinHealthFactor is a free data retrieval call binding the contract method 0x8c1ae6c8.
//
// Solidity: function getMinHealthFactor() pure returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetMinHealthFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getMinHealthFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinHealthFactor is a free data retrieval call binding the contract method 0x8c1ae6c8.
//
// Solidity: function getMinHealthFactor() pure returns(uint256)
func (_DSCEngine *DSCEngineSession) GetMinHealthFactor() (*big.Int, error) {
	return _DSCEngine.Contract.GetMinHealthFactor(&_DSCEngine.CallOpts)
}

// GetMinHealthFactor is a free data retrieval call binding the contract method 0x8c1ae6c8.
//
// Solidity: function getMinHealthFactor() pure returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetMinHealthFactor() (*big.Int, error) {
	return _DSCEngine.Contract.GetMinHealthFactor(&_DSCEngine.CallOpts)
}

// GetTokenAmountFromUsd is a free data retrieval call binding the contract method 0xafea2e48.
//
// Solidity: function getTokenAmountFromUsd(address tokenAddress, uint256 usdAmountInWei) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetTokenAmountFromUsd(opts *bind.CallOpts, tokenAddress common.Address, usdAmountInWei *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getTokenAmountFromUsd", tokenAddress, usdAmountInWei)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenAmountFromUsd is a free data retrieval call binding the contract method 0xafea2e48.
//
// Solidity: function getTokenAmountFromUsd(address tokenAddress, uint256 usdAmountInWei) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetTokenAmountFromUsd(tokenAddress common.Address, usdAmountInWei *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.GetTokenAmountFromUsd(&_DSCEngine.CallOpts, tokenAddress, usdAmountInWei)
}

// GetTokenAmountFromUsd is a free data retrieval call binding the contract method 0xafea2e48.
//
// Solidity: function getTokenAmountFromUsd(address tokenAddress, uint256 usdAmountInWei) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetTokenAmountFromUsd(tokenAddress common.Address, usdAmountInWei *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.GetTokenAmountFromUsd(&_DSCEngine.CallOpts, tokenAddress, usdAmountInWei)
}

// GetUsdValue is a free data retrieval call binding the contract method 0xc660d112.
//
// Solidity: function getUsdValue(address token, uint256 amount) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetUsdValue(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getUsdValue", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsdValue is a free data retrieval call binding the contract method 0xc660d112.
//
// Solidity: function getUsdValue(address token, uint256 amount) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetUsdValue(token common.Address, amount *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.GetUsdValue(&_DSCEngine.CallOpts, token, amount)
}

// GetUsdValue is a free data retrieval call binding the contract method 0xc660d112.
//
// Solidity: function getUsdValue(address token, uint256 amount) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetUsdValue(token common.Address, amount *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.GetUsdValue(&_DSCEngine.CallOpts, token, amount)
}

// BurnDsc is a paid mutator transaction binding the contract method 0xf6876608.
//
// Solidity: function burnDsc(uint256 amount) returns()
func (_DSCEngine *DSCEngineTransactor) BurnDsc(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "burnDsc", amount)
}

// BurnDsc is a paid mutator transaction binding the contract method 0xf6876608.
//
// Solidity: function burnDsc(uint256 amount) returns()
func (_DSCEngine *DSCEngineSession) BurnDsc(amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.BurnDsc(&_DSCEngine.TransactOpts, amount)
}

// BurnDsc is a paid mutator transaction binding the contract method 0xf6876608.
//
// Solidity: function burnDsc(uint256 amount) returns()
func (_DSCEngine *DSCEngineTransactorSession) BurnDsc(amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.BurnDsc(&_DSCEngine.TransactOpts, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address token, uint256 amount) returns()
func (_DSCEngine *DSCEngineTransactor) DepositCollateral(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "depositCollateral", token, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address token, uint256 amount) returns()
func (_DSCEngine *DSCEngineSession) DepositCollateral(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateral(&_DSCEngine.TransactOpts, token, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address token, uint256 amount) returns()
func (_DSCEngine *DSCEngineTransactorSession) DepositCollateral(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateral(&_DSCEngine.TransactOpts, token, amount)
}

// DepositCollateralAndMintDsc is a paid mutator transaction binding the contract method 0xe90db8a3.
//
// Solidity: function depositCollateralAndMintDsc(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToMint) returns()
func (_DSCEngine *DSCEngineTransactor) DepositCollateralAndMintDsc(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToMint *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "depositCollateralAndMintDsc", tokenCollateralAddress, amountCollateral, amountDscToMint)
}

// DepositCollateralAndMintDsc is a paid mutator transaction binding the contract method 0xe90db8a3.
//
// Solidity: function depositCollateralAndMintDsc(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToMint) returns()
func (_DSCEngine *DSCEngineSession) DepositCollateralAndMintDsc(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToMint *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateralAndMintDsc(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral, amountDscToMint)
}

// DepositCollateralAndMintDsc is a paid mutator transaction binding the contract method 0xe90db8a3.
//
// Solidity: function depositCollateralAndMintDsc(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToMint) returns()
func (_DSCEngine *DSCEngineTransactorSession) DepositCollateralAndMintDsc(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToMint *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateralAndMintDsc(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral, amountDscToMint)
}

// DepositCollateralAndMintDscWithPermit is a paid mutator transaction binding the contract method 0x51e32bd9.
//
// Solidity: function depositCollateralAndMintDscWithPermit(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToMint, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DSCEngine *DSCEngineTransactor) DepositCollateralAndMintDscWithPermit(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToMint *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "depositCollateralAndMintDscWithPermit", tokenCollateralAddress, amountCollateral, amountDscToMint, deadline, v, r, s)
}

// DepositCollateralAndMintDscWithPermit is a paid mutator transaction binding the contract method 0x51e32bd9.
//
// Solidity: function depositCollateralAndMintDscWithPermit(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToMint, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DSCEngine *DSCEngineSession) DepositCollateralAndMintDscWithPermit(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToMint *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateralAndMintDscWithPermit(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral, amountDscToMint, deadline, v, r, s)
}

// DepositCollateralAndMintDscWithPermit is a paid mutator transaction binding the contract method 0x51e32bd9.
//
// Solidity: function depositCollateralAndMintDscWithPermit(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToMint, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DSCEngine *DSCEngineTransactorSession) DepositCollateralAndMintDscWithPermit(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToMint *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateralAndMintDscWithPermit(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral, amountDscToMint, deadline, v, r, s)
}

// DepositCollateralWithPermit is a paid mutator transaction binding the contract method 0xd221b81d.
//
// Solidity: function depositCollateralWithPermit(address token, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DSCEngine *DSCEngineTransactor) DepositCollateralWithPermit(opts *bind.TransactOpts, token common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "depositCollateralWithPermit", token, amount, deadline, v, r, s)
}

// DepositCollateralWithPermit is a paid mutator transaction binding the contract method 0xd221b81d.
//
// Solidity: function depositCollateralWithPermit(address token, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DSCEngine *DSCEngineSession) DepositCollateralWithPermit(token common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateralWithPermit(&_DSCEngine.TransactOpts, token, amount, deadline, v, r, s)
}

// DepositCollateralWithPermit is a paid mutator transaction binding the contract method 0xd221b81d.
//
// Solidity: function depositCollateralWithPermit(address token, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DSCEngine *DSCEngineTransactorSession) DepositCollateralWithPermit(token common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateralWithPermit(&_DSCEngine.TransactOpts, token, amount, deadline, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0x5274ac3f.
//
// Solidity: function initialize(address[] tokenAddresses, address[] priceFeedAddresses, address dscAddress) returns()
func (_DSCEngine *DSCEngineTransactor) Initialize(opts *bind.TransactOpts, tokenAddresses []common.Address, priceFeedAddresses []common.Address, dscAddress common.Address) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "initialize", tokenAddresses, priceFeedAddresses, dscAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x5274ac3f.
//
// Solidity: function initialize(address[] tokenAddresses, address[] priceFeedAddresses, address dscAddress) returns()
func (_DSCEngine *DSCEngineSession) Initialize(tokenAddresses []common.Address, priceFeedAddresses []common.Address, dscAddress common.Address) (*types.Transaction, error) {
	return _DSCEngine.Contract.Initialize(&_DSCEngine.TransactOpts, tokenAddresses, priceFeedAddresses, dscAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x5274ac3f.
//
// Solidity: function initialize(address[] tokenAddresses, address[] priceFeedAddresses, address dscAddress) returns()
func (_DSCEngine *DSCEngineTransactorSession) Initialize(tokenAddresses []common.Address, priceFeedAddresses []common.Address, dscAddress common.Address) (*types.Transaction, error) {
	return _DSCEngine.Contract.Initialize(&_DSCEngine.TransactOpts, tokenAddresses, priceFeedAddresses, dscAddress)
}

// Liquidate is a paid mutator transaction binding the contract method 0x26c01303.
//
// Solidity: function liquidate(address collateral, address user, uint256 debtToCover) returns()
func (_DSCEngine *DSCEngineTransactor) Liquidate(opts *bind.TransactOpts, collateral common.Address, user common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "liquidate", collateral, user, debtToCover)
}

// Liquidate is a paid mutator transaction binding the contract method 0x26c01303.
//
// Solidity: function liquidate(address collateral, address user, uint256 debtToCover) returns()
func (_DSCEngine *DSCEngineSession) Liquidate(collateral common.Address, user common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.Liquidate(&_DSCEngine.TransactOpts, collateral, user, debtToCover)
}

// Liquidate is a paid mutator transaction binding the contract method 0x26c01303.
//
// Solidity: function liquidate(address collateral, address user, uint256 debtToCover) returns()
func (_DSCEngine *DSCEngineTransactorSession) Liquidate(collateral common.Address, user common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.Liquidate(&_DSCEngine.TransactOpts, collateral, user, debtToCover)
}

// MintDsc is a paid mutator transaction binding the contract method 0xc9b7c327.
//
// Solidity: function mintDsc(uint256 amount) returns()
func (_DSCEngine *DSCEngineTransactor) MintDsc(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "mintDsc", amount)
}

// MintDsc is a paid mutator transaction binding the contract method 0xc9b7c327.
//
// Solidity: function mintDsc(uint256 amount) returns()
func (_DSCEngine *DSCEngineSession) MintDsc(amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.MintDsc(&_DSCEngine.TransactOpts, amount)
}

// MintDsc is a paid mutator transaction binding the contract method 0xc9b7c327.
//
// Solidity: function mintDsc(uint256 amount) returns()
func (_DSCEngine *DSCEngineTransactorSession) MintDsc(amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.MintDsc(&_DSCEngine.TransactOpts, amount)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineTransactor) RedeemCollateral(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "redeemCollateral", tokenCollateralAddress, amountCollateral)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineSession) RedeemCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.RedeemCollateral(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineTransactorSession) RedeemCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.RedeemCollateral(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// RedeemCollateralForDsc is a paid mutator transaction binding the contract method 0xf419ea9c.
//
// Solidity: function redeemCollateralForDsc(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToBurn) returns()
func (_DSCEngine *DSCEngineTransactor) RedeemCollateralForDsc(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToBurn *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "redeemCollateralForDsc", tokenCollateralAddress, amountCollateral, amountDscToBurn)
}

// RedeemCollateralForDsc is a paid mutator transaction binding the contract method 0xf419ea9c.
//
// Solidity: function redeemCollateralForDsc(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToBurn) returns()
func (_DSCEngine *DSCEngineSession) RedeemCollateralForDsc(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToBurn *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.RedeemCollateralForDsc(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral, amountDscToBurn)
}

// RedeemCollateralForDsc is a paid mutator transaction binding the contract method 0xf419ea9c.
//
// Solidity: function redeemCollateralForDsc(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToBurn) returns()
func (_DSCEngine *DSCEngineTransactorSession) RedeemCollateralForDsc(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToBurn *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.RedeemCollateralForDsc(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral, amountDscToBurn)
}

// DSCEngineCollateralDepositedIterator is returned from FilterCollateralDeposited and is used to iterate over the raw logs and unpacked data for CollateralDeposited events raised by the DSCEngine contract.
type DSCEngineCollateralDepositedIterator struct {
	Event *DSCEngineCollateralDeposited // Event containing the contract specifics and raw log

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
func (it *DSCEngineCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSCEngineCollateralDeposited)
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
		it.Event = new(DSCEngineCollateralDeposited)
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
func (it *DSCEngineCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSCEngineCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSCEngineCollateralDeposited represents a CollateralDeposited event raised by the DSCEngine contract.
type DSCEngineCollateralDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralDeposited is a free log retrieval operation binding the contract event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed user, address indexed token, uint256 indexed amount)
func (_DSCEngine *DSCEngineFilterer) FilterCollateralDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address, amount []*big.Int) (*DSCEngineCollateralDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _DSCEngine.contract.FilterLogs(opts, "CollateralDeposited", userRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &DSCEngineCollateralDepositedIterator{contract: _DSCEngine.contract, event: "CollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchCollateralDeposited is a free log subscription operation binding the contract event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed user, address indexed token, uint256 indexed amount)
func (_DSCEngine *DSCEngineFilterer) WatchCollateralDeposited(opts *bind.WatchOpts, sink chan<- *DSCEngineCollateralDeposited, user []common.Address, token []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _DSCEngine.contract.WatchLogs(opts, "CollateralDeposited", userRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSCEngineCollateralDeposited)
				if err := _DSCEngine.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
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

// ParseCollateralDeposited is a log parse operation binding the contract event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed user, address indexed token, uint256 indexed amount)
func (_DSCEngine *DSCEngineFilterer) ParseCollateralDeposited(log types.Log) (*DSCEngineCollateralDeposited, error) {
	event := new(DSCEngineCollateralDeposited)
	if err := _DSCEngine.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DSCEngineCollateralRedeemedIterator is returned from FilterCollateralRedeemed and is used to iterate over the raw logs and unpacked data for CollateralRedeemed events raised by the DSCEngine contract.
type DSCEngineCollateralRedeemedIterator struct {
	Event *DSCEngineCollateralRedeemed // Event containing the contract specifics and raw log

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
func (it *DSCEngineCollateralRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSCEngineCollateralRedeemed)
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
		it.Event = new(DSCEngineCollateralRedeemed)
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
func (it *DSCEngineCollateralRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSCEngineCollateralRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSCEngineCollateralRedeemed represents a CollateralRedeemed event raised by the DSCEngine contract.
type DSCEngineCollateralRedeemed struct {
	RedeemFrom common.Address
	RedeemTo   common.Address
	Token      common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralRedeemed is a free log retrieval operation binding the contract event 0xb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f6.
//
// Solidity: event CollateralRedeemed(address indexed redeemFrom, address indexed redeemTo, address token, uint256 amount)
func (_DSCEngine *DSCEngineFilterer) FilterCollateralRedeemed(opts *bind.FilterOpts, redeemFrom []common.Address, redeemTo []common.Address) (*DSCEngineCollateralRedeemedIterator, error) {

	var redeemFromRule []interface{}
	for _, redeemFromItem := range redeemFrom {
		redeemFromRule = append(redeemFromRule, redeemFromItem)
	}
	var redeemToRule []interface{}
	for _, redeemToItem := range redeemTo {
		redeemToRule = append(redeemToRule, redeemToItem)
	}

	logs, sub, err := _DSCEngine.contract.FilterLogs(opts, "CollateralRedeemed", redeemFromRule, redeemToRule)
	if err != nil {
		return nil, err
	}
	return &DSCEngineCollateralRedeemedIterator{contract: _DSCEngine.contract, event: "CollateralRedeemed", logs: logs, sub: sub}, nil
}

// WatchCollateralRedeemed is a free log subscription operation binding the contract event 0xb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f6.
//
// Solidity: event CollateralRedeemed(address indexed redeemFrom, address indexed redeemTo, address token, uint256 amount)
func (_DSCEngine *DSCEngineFilterer) WatchCollateralRedeemed(opts *bind.WatchOpts, sink chan<- *DSCEngineCollateralRedeemed, redeemFrom []common.Address, redeemTo []common.Address) (event.Subscription, error) {

	var redeemFromRule []interface{}
	for _, redeemFromItem := range redeemFrom {
		redeemFromRule = append(redeemFromRule, redeemFromItem)
	}
	var redeemToRule []interface{}
	for _, redeemToItem := range redeemTo {
		redeemToRule = append(redeemToRule, redeemToItem)
	}

	logs, sub, err := _DSCEngine.contract.WatchLogs(opts, "CollateralRedeemed", redeemFromRule, redeemToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSCEngineCollateralRedeemed)
				if err := _DSCEngine.contract.UnpackLog(event, "CollateralRedeemed", log); err != nil {
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

// ParseCollateralRedeemed is a log parse operation binding the contract event 0xb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f6.
//
// Solidity: event CollateralRedeemed(address indexed redeemFrom, address indexed redeemTo, address token, uint256 amount)
func (_DSCEngine *DSCEngineFilterer) ParseCollateralRedeemed(log types.Log) (*DSCEngineCollateralRedeemed, error) {
	event := new(DSCEngineCollateralRedeemed)
	if err := _DSCEngine.contract.UnpackLog(event, "CollateralRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

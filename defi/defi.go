// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package defi

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

// DefiMetaData contains all meta data concerning the Defi contract.
var DefiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minTokensBeforeSwap\",\"type\":\"uint256\"}],\"name\":\"MinTokensBeforeSwapUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSwapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensIntoLiqudity\",\"type\":\"uint256\"}],\"name\":\"SwapAndLiquify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SwapAndLiquifyEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_burnFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_charityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_liquidityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_taxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tAmount\",\"type\":\"uint256\"}],\"name\":\"deliver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableAllFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"excludeFromFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"excludeFromReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"geUnlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"includeInFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"includeInReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"deductTransferFee\",\"type\":\"bool\"}],\"name\":\"reflectionFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnFee\",\"type\":\"uint256\"}],\"name\":\"setBurnFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"charityFee\",\"type\":\"uint256\"}],\"name\":\"setChartityFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidityFee\",\"type\":\"uint256\"}],\"name\":\"setLiquidityFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setSwapAndLiquifyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taxFee\",\"type\":\"uint256\"}],\"name\":\"setTaxFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"setcharityWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapAndLiquifyEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rAmount\",\"type\":\"uint256\"}],\"name\":\"tokenFromReflection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DefiABI is the input ABI used to generate the binding from.
// Deprecated: Use DefiMetaData.ABI instead.
var DefiABI = DefiMetaData.ABI

// Defi is an auto generated Go binding around an Ethereum contract.
type Defi struct {
	DefiCaller     // Read-only binding to the contract
	DefiTransactor // Write-only binding to the contract
	DefiFilterer   // Log filterer for contract events
}

// DefiCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefiSession struct {
	Contract     *Defi             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefiCallerSession struct {
	Contract *DefiCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DefiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefiTransactorSession struct {
	Contract     *DefiTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefiRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefiRaw struct {
	Contract *Defi // Generic contract binding to access the raw methods on
}

// DefiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefiCallerRaw struct {
	Contract *DefiCaller // Generic read-only contract binding to access the raw methods on
}

// DefiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefiTransactorRaw struct {
	Contract *DefiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefi creates a new instance of Defi, bound to a specific deployed contract.
func NewDefi(address common.Address, backend bind.ContractBackend) (*Defi, error) {
	contract, err := bindDefi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Defi{DefiCaller: DefiCaller{contract: contract}, DefiTransactor: DefiTransactor{contract: contract}, DefiFilterer: DefiFilterer{contract: contract}}, nil
}

// NewDefiCaller creates a new read-only instance of Defi, bound to a specific deployed contract.
func NewDefiCaller(address common.Address, caller bind.ContractCaller) (*DefiCaller, error) {
	contract, err := bindDefi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefiCaller{contract: contract}, nil
}

// NewDefiTransactor creates a new write-only instance of Defi, bound to a specific deployed contract.
func NewDefiTransactor(address common.Address, transactor bind.ContractTransactor) (*DefiTransactor, error) {
	contract, err := bindDefi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefiTransactor{contract: contract}, nil
}

// NewDefiFilterer creates a new log filterer instance of Defi, bound to a specific deployed contract.
func NewDefiFilterer(address common.Address, filterer bind.ContractFilterer) (*DefiFilterer, error) {
	contract, err := bindDefi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefiFilterer{contract: contract}, nil
}

// bindDefi binds a generic wrapper to an already deployed contract.
func bindDefi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Defi *DefiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Defi.Contract.DefiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Defi *DefiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defi.Contract.DefiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Defi *DefiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Defi.Contract.DefiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Defi *DefiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Defi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Defi *DefiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Defi *DefiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Defi.Contract.contract.Transact(opts, method, params...)
}

// BurnFee is a free data retrieval call binding the contract method 0xc0b0fda2.
//
// Solidity: function _burnFee() view returns(uint256)
func (_Defi *DefiCaller) BurnFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "_burnFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnFee is a free data retrieval call binding the contract method 0xc0b0fda2.
//
// Solidity: function _burnFee() view returns(uint256)
func (_Defi *DefiSession) BurnFee() (*big.Int, error) {
	return _Defi.Contract.BurnFee(&_Defi.CallOpts)
}

// BurnFee is a free data retrieval call binding the contract method 0xc0b0fda2.
//
// Solidity: function _burnFee() view returns(uint256)
func (_Defi *DefiCallerSession) BurnFee() (*big.Int, error) {
	return _Defi.Contract.BurnFee(&_Defi.CallOpts)
}

// CharityFee is a free data retrieval call binding the contract method 0x40f8007a.
//
// Solidity: function _charityFee() view returns(uint256)
func (_Defi *DefiCaller) CharityFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "_charityFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CharityFee is a free data retrieval call binding the contract method 0x40f8007a.
//
// Solidity: function _charityFee() view returns(uint256)
func (_Defi *DefiSession) CharityFee() (*big.Int, error) {
	return _Defi.Contract.CharityFee(&_Defi.CallOpts)
}

// CharityFee is a free data retrieval call binding the contract method 0x40f8007a.
//
// Solidity: function _charityFee() view returns(uint256)
func (_Defi *DefiCallerSession) CharityFee() (*big.Int, error) {
	return _Defi.Contract.CharityFee(&_Defi.CallOpts)
}

// LiquidityFee is a free data retrieval call binding the contract method 0x6bc87c3a.
//
// Solidity: function _liquidityFee() view returns(uint256)
func (_Defi *DefiCaller) LiquidityFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "_liquidityFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityFee is a free data retrieval call binding the contract method 0x6bc87c3a.
//
// Solidity: function _liquidityFee() view returns(uint256)
func (_Defi *DefiSession) LiquidityFee() (*big.Int, error) {
	return _Defi.Contract.LiquidityFee(&_Defi.CallOpts)
}

// LiquidityFee is a free data retrieval call binding the contract method 0x6bc87c3a.
//
// Solidity: function _liquidityFee() view returns(uint256)
func (_Defi *DefiCallerSession) LiquidityFee() (*big.Int, error) {
	return _Defi.Contract.LiquidityFee(&_Defi.CallOpts)
}

// TaxFee is a free data retrieval call binding the contract method 0x3b124fe7.
//
// Solidity: function _taxFee() view returns(uint256)
func (_Defi *DefiCaller) TaxFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "_taxFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxFee is a free data retrieval call binding the contract method 0x3b124fe7.
//
// Solidity: function _taxFee() view returns(uint256)
func (_Defi *DefiSession) TaxFee() (*big.Int, error) {
	return _Defi.Contract.TaxFee(&_Defi.CallOpts)
}

// TaxFee is a free data retrieval call binding the contract method 0x3b124fe7.
//
// Solidity: function _taxFee() view returns(uint256)
func (_Defi *DefiCallerSession) TaxFee() (*big.Int, error) {
	return _Defi.Contract.TaxFee(&_Defi.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Defi *DefiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Defi *DefiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Defi.Contract.Allowance(&_Defi.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Defi *DefiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Defi.Contract.Allowance(&_Defi.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Defi *DefiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Defi *DefiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Defi.Contract.BalanceOf(&_Defi.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Defi *DefiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Defi.Contract.BalanceOf(&_Defi.CallOpts, account)
}

// CharityWallet is a free data retrieval call binding the contract method 0x7b208769.
//
// Solidity: function charityWallet() view returns(address)
func (_Defi *DefiCaller) CharityWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "charityWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CharityWallet is a free data retrieval call binding the contract method 0x7b208769.
//
// Solidity: function charityWallet() view returns(address)
func (_Defi *DefiSession) CharityWallet() (common.Address, error) {
	return _Defi.Contract.CharityWallet(&_Defi.CallOpts)
}

// CharityWallet is a free data retrieval call binding the contract method 0x7b208769.
//
// Solidity: function charityWallet() view returns(address)
func (_Defi *DefiCallerSession) CharityWallet() (common.Address, error) {
	return _Defi.Contract.CharityWallet(&_Defi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Defi *DefiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Defi *DefiSession) Decimals() (uint8, error) {
	return _Defi.Contract.Decimals(&_Defi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Defi *DefiCallerSession) Decimals() (uint8, error) {
	return _Defi.Contract.Decimals(&_Defi.CallOpts)
}

// GeUnlockTime is a free data retrieval call binding the contract method 0xb6c52324.
//
// Solidity: function geUnlockTime() view returns(uint256)
func (_Defi *DefiCaller) GeUnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "geUnlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeUnlockTime is a free data retrieval call binding the contract method 0xb6c52324.
//
// Solidity: function geUnlockTime() view returns(uint256)
func (_Defi *DefiSession) GeUnlockTime() (*big.Int, error) {
	return _Defi.Contract.GeUnlockTime(&_Defi.CallOpts)
}

// GeUnlockTime is a free data retrieval call binding the contract method 0xb6c52324.
//
// Solidity: function geUnlockTime() view returns(uint256)
func (_Defi *DefiCallerSession) GeUnlockTime() (*big.Int, error) {
	return _Defi.Contract.GeUnlockTime(&_Defi.CallOpts)
}

// IsExcludedFromFee is a free data retrieval call binding the contract method 0x5342acb4.
//
// Solidity: function isExcludedFromFee(address account) view returns(bool)
func (_Defi *DefiCaller) IsExcludedFromFee(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "isExcludedFromFee", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromFee is a free data retrieval call binding the contract method 0x5342acb4.
//
// Solidity: function isExcludedFromFee(address account) view returns(bool)
func (_Defi *DefiSession) IsExcludedFromFee(account common.Address) (bool, error) {
	return _Defi.Contract.IsExcludedFromFee(&_Defi.CallOpts, account)
}

// IsExcludedFromFee is a free data retrieval call binding the contract method 0x5342acb4.
//
// Solidity: function isExcludedFromFee(address account) view returns(bool)
func (_Defi *DefiCallerSession) IsExcludedFromFee(account common.Address) (bool, error) {
	return _Defi.Contract.IsExcludedFromFee(&_Defi.CallOpts, account)
}

// IsExcludedFromReward is a free data retrieval call binding the contract method 0x88f82020.
//
// Solidity: function isExcludedFromReward(address account) view returns(bool)
func (_Defi *DefiCaller) IsExcludedFromReward(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "isExcludedFromReward", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromReward is a free data retrieval call binding the contract method 0x88f82020.
//
// Solidity: function isExcludedFromReward(address account) view returns(bool)
func (_Defi *DefiSession) IsExcludedFromReward(account common.Address) (bool, error) {
	return _Defi.Contract.IsExcludedFromReward(&_Defi.CallOpts, account)
}

// IsExcludedFromReward is a free data retrieval call binding the contract method 0x88f82020.
//
// Solidity: function isExcludedFromReward(address account) view returns(bool)
func (_Defi *DefiCallerSession) IsExcludedFromReward(account common.Address) (bool, error) {
	return _Defi.Contract.IsExcludedFromReward(&_Defi.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Defi *DefiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Defi *DefiSession) Name() (string, error) {
	return _Defi.Contract.Name(&_Defi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Defi *DefiCallerSession) Name() (string, error) {
	return _Defi.Contract.Name(&_Defi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Defi *DefiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Defi *DefiSession) Owner() (common.Address, error) {
	return _Defi.Contract.Owner(&_Defi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Defi *DefiCallerSession) Owner() (common.Address, error) {
	return _Defi.Contract.Owner(&_Defi.CallOpts)
}

// ReflectionFromToken is a free data retrieval call binding the contract method 0x4549b039.
//
// Solidity: function reflectionFromToken(uint256 tAmount, bool deductTransferFee) view returns(uint256)
func (_Defi *DefiCaller) ReflectionFromToken(opts *bind.CallOpts, tAmount *big.Int, deductTransferFee bool) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "reflectionFromToken", tAmount, deductTransferFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReflectionFromToken is a free data retrieval call binding the contract method 0x4549b039.
//
// Solidity: function reflectionFromToken(uint256 tAmount, bool deductTransferFee) view returns(uint256)
func (_Defi *DefiSession) ReflectionFromToken(tAmount *big.Int, deductTransferFee bool) (*big.Int, error) {
	return _Defi.Contract.ReflectionFromToken(&_Defi.CallOpts, tAmount, deductTransferFee)
}

// ReflectionFromToken is a free data retrieval call binding the contract method 0x4549b039.
//
// Solidity: function reflectionFromToken(uint256 tAmount, bool deductTransferFee) view returns(uint256)
func (_Defi *DefiCallerSession) ReflectionFromToken(tAmount *big.Int, deductTransferFee bool) (*big.Int, error) {
	return _Defi.Contract.ReflectionFromToken(&_Defi.CallOpts, tAmount, deductTransferFee)
}

// SwapAndLiquifyEnabled is a free data retrieval call binding the contract method 0x4a74bb02.
//
// Solidity: function swapAndLiquifyEnabled() view returns(bool)
func (_Defi *DefiCaller) SwapAndLiquifyEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "swapAndLiquifyEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapAndLiquifyEnabled is a free data retrieval call binding the contract method 0x4a74bb02.
//
// Solidity: function swapAndLiquifyEnabled() view returns(bool)
func (_Defi *DefiSession) SwapAndLiquifyEnabled() (bool, error) {
	return _Defi.Contract.SwapAndLiquifyEnabled(&_Defi.CallOpts)
}

// SwapAndLiquifyEnabled is a free data retrieval call binding the contract method 0x4a74bb02.
//
// Solidity: function swapAndLiquifyEnabled() view returns(bool)
func (_Defi *DefiCallerSession) SwapAndLiquifyEnabled() (bool, error) {
	return _Defi.Contract.SwapAndLiquifyEnabled(&_Defi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Defi *DefiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Defi *DefiSession) Symbol() (string, error) {
	return _Defi.Contract.Symbol(&_Defi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Defi *DefiCallerSession) Symbol() (string, error) {
	return _Defi.Contract.Symbol(&_Defi.CallOpts)
}

// TokenFromReflection is a free data retrieval call binding the contract method 0x2d838119.
//
// Solidity: function tokenFromReflection(uint256 rAmount) view returns(uint256)
func (_Defi *DefiCaller) TokenFromReflection(opts *bind.CallOpts, rAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "tokenFromReflection", rAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenFromReflection is a free data retrieval call binding the contract method 0x2d838119.
//
// Solidity: function tokenFromReflection(uint256 rAmount) view returns(uint256)
func (_Defi *DefiSession) TokenFromReflection(rAmount *big.Int) (*big.Int, error) {
	return _Defi.Contract.TokenFromReflection(&_Defi.CallOpts, rAmount)
}

// TokenFromReflection is a free data retrieval call binding the contract method 0x2d838119.
//
// Solidity: function tokenFromReflection(uint256 rAmount) view returns(uint256)
func (_Defi *DefiCallerSession) TokenFromReflection(rAmount *big.Int) (*big.Int, error) {
	return _Defi.Contract.TokenFromReflection(&_Defi.CallOpts, rAmount)
}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_Defi *DefiCaller) TotalFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "totalFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_Defi *DefiSession) TotalFees() (*big.Int, error) {
	return _Defi.Contract.TotalFees(&_Defi.CallOpts)
}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_Defi *DefiCallerSession) TotalFees() (*big.Int, error) {
	return _Defi.Contract.TotalFees(&_Defi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Defi *DefiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Defi *DefiSession) TotalSupply() (*big.Int, error) {
	return _Defi.Contract.TotalSupply(&_Defi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Defi *DefiCallerSession) TotalSupply() (*big.Int, error) {
	return _Defi.Contract.TotalSupply(&_Defi.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Defi *DefiCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Defi *DefiSession) UniswapV2Pair() (common.Address, error) {
	return _Defi.Contract.UniswapV2Pair(&_Defi.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Defi *DefiCallerSession) UniswapV2Pair() (common.Address, error) {
	return _Defi.Contract.UniswapV2Pair(&_Defi.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Defi *DefiCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Defi *DefiSession) UniswapV2Router() (common.Address, error) {
	return _Defi.Contract.UniswapV2Router(&_Defi.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Defi *DefiCallerSession) UniswapV2Router() (common.Address, error) {
	return _Defi.Contract.UniswapV2Router(&_Defi.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Defi *DefiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Defi *DefiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Approve(&_Defi.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Defi *DefiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Approve(&_Defi.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Defi *DefiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Defi *DefiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.DecreaseAllowance(&_Defi.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Defi *DefiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.DecreaseAllowance(&_Defi.TransactOpts, spender, subtractedValue)
}

// Deliver is a paid mutator transaction binding the contract method 0x3bd5d173.
//
// Solidity: function deliver(uint256 tAmount) returns()
func (_Defi *DefiTransactor) Deliver(opts *bind.TransactOpts, tAmount *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "deliver", tAmount)
}

// Deliver is a paid mutator transaction binding the contract method 0x3bd5d173.
//
// Solidity: function deliver(uint256 tAmount) returns()
func (_Defi *DefiSession) Deliver(tAmount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Deliver(&_Defi.TransactOpts, tAmount)
}

// Deliver is a paid mutator transaction binding the contract method 0x3bd5d173.
//
// Solidity: function deliver(uint256 tAmount) returns()
func (_Defi *DefiTransactorSession) Deliver(tAmount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Deliver(&_Defi.TransactOpts, tAmount)
}

// DisableAllFees is a paid mutator transaction binding the contract method 0x741af87f.
//
// Solidity: function disableAllFees() returns()
func (_Defi *DefiTransactor) DisableAllFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "disableAllFees")
}

// DisableAllFees is a paid mutator transaction binding the contract method 0x741af87f.
//
// Solidity: function disableAllFees() returns()
func (_Defi *DefiSession) DisableAllFees() (*types.Transaction, error) {
	return _Defi.Contract.DisableAllFees(&_Defi.TransactOpts)
}

// DisableAllFees is a paid mutator transaction binding the contract method 0x741af87f.
//
// Solidity: function disableAllFees() returns()
func (_Defi *DefiTransactorSession) DisableAllFees() (*types.Transaction, error) {
	return _Defi.Contract.DisableAllFees(&_Defi.TransactOpts)
}

// ExcludeFromFee is a paid mutator transaction binding the contract method 0x437823ec.
//
// Solidity: function excludeFromFee(address account) returns()
func (_Defi *DefiTransactor) ExcludeFromFee(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "excludeFromFee", account)
}

// ExcludeFromFee is a paid mutator transaction binding the contract method 0x437823ec.
//
// Solidity: function excludeFromFee(address account) returns()
func (_Defi *DefiSession) ExcludeFromFee(account common.Address) (*types.Transaction, error) {
	return _Defi.Contract.ExcludeFromFee(&_Defi.TransactOpts, account)
}

// ExcludeFromFee is a paid mutator transaction binding the contract method 0x437823ec.
//
// Solidity: function excludeFromFee(address account) returns()
func (_Defi *DefiTransactorSession) ExcludeFromFee(account common.Address) (*types.Transaction, error) {
	return _Defi.Contract.ExcludeFromFee(&_Defi.TransactOpts, account)
}

// ExcludeFromReward is a paid mutator transaction binding the contract method 0x52390c02.
//
// Solidity: function excludeFromReward(address account) returns()
func (_Defi *DefiTransactor) ExcludeFromReward(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "excludeFromReward", account)
}

// ExcludeFromReward is a paid mutator transaction binding the contract method 0x52390c02.
//
// Solidity: function excludeFromReward(address account) returns()
func (_Defi *DefiSession) ExcludeFromReward(account common.Address) (*types.Transaction, error) {
	return _Defi.Contract.ExcludeFromReward(&_Defi.TransactOpts, account)
}

// ExcludeFromReward is a paid mutator transaction binding the contract method 0x52390c02.
//
// Solidity: function excludeFromReward(address account) returns()
func (_Defi *DefiTransactorSession) ExcludeFromReward(account common.Address) (*types.Transaction, error) {
	return _Defi.Contract.ExcludeFromReward(&_Defi.TransactOpts, account)
}

// IncludeInFee is a paid mutator transaction binding the contract method 0xea2f0b37.
//
// Solidity: function includeInFee(address account) returns()
func (_Defi *DefiTransactor) IncludeInFee(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "includeInFee", account)
}

// IncludeInFee is a paid mutator transaction binding the contract method 0xea2f0b37.
//
// Solidity: function includeInFee(address account) returns()
func (_Defi *DefiSession) IncludeInFee(account common.Address) (*types.Transaction, error) {
	return _Defi.Contract.IncludeInFee(&_Defi.TransactOpts, account)
}

// IncludeInFee is a paid mutator transaction binding the contract method 0xea2f0b37.
//
// Solidity: function includeInFee(address account) returns()
func (_Defi *DefiTransactorSession) IncludeInFee(account common.Address) (*types.Transaction, error) {
	return _Defi.Contract.IncludeInFee(&_Defi.TransactOpts, account)
}

// IncludeInReward is a paid mutator transaction binding the contract method 0x3685d419.
//
// Solidity: function includeInReward(address account) returns()
func (_Defi *DefiTransactor) IncludeInReward(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "includeInReward", account)
}

// IncludeInReward is a paid mutator transaction binding the contract method 0x3685d419.
//
// Solidity: function includeInReward(address account) returns()
func (_Defi *DefiSession) IncludeInReward(account common.Address) (*types.Transaction, error) {
	return _Defi.Contract.IncludeInReward(&_Defi.TransactOpts, account)
}

// IncludeInReward is a paid mutator transaction binding the contract method 0x3685d419.
//
// Solidity: function includeInReward(address account) returns()
func (_Defi *DefiTransactorSession) IncludeInReward(account common.Address) (*types.Transaction, error) {
	return _Defi.Contract.IncludeInReward(&_Defi.TransactOpts, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Defi *DefiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Defi *DefiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.IncreaseAllowance(&_Defi.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Defi *DefiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.IncreaseAllowance(&_Defi.TransactOpts, spender, addedValue)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 time) returns()
func (_Defi *DefiTransactor) Lock(opts *bind.TransactOpts, time *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "lock", time)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 time) returns()
func (_Defi *DefiSession) Lock(time *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Lock(&_Defi.TransactOpts, time)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 time) returns()
func (_Defi *DefiTransactorSession) Lock(time *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Lock(&_Defi.TransactOpts, time)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Defi *DefiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Defi *DefiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Defi.Contract.RenounceOwnership(&_Defi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Defi *DefiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Defi.Contract.RenounceOwnership(&_Defi.TransactOpts)
}

// SetBurnFeePercent is a paid mutator transaction binding the contract method 0xcea26958.
//
// Solidity: function setBurnFeePercent(uint256 burnFee) returns()
func (_Defi *DefiTransactor) SetBurnFeePercent(opts *bind.TransactOpts, burnFee *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "setBurnFeePercent", burnFee)
}

// SetBurnFeePercent is a paid mutator transaction binding the contract method 0xcea26958.
//
// Solidity: function setBurnFeePercent(uint256 burnFee) returns()
func (_Defi *DefiSession) SetBurnFeePercent(burnFee *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.SetBurnFeePercent(&_Defi.TransactOpts, burnFee)
}

// SetBurnFeePercent is a paid mutator transaction binding the contract method 0xcea26958.
//
// Solidity: function setBurnFeePercent(uint256 burnFee) returns()
func (_Defi *DefiTransactorSession) SetBurnFeePercent(burnFee *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.SetBurnFeePercent(&_Defi.TransactOpts, burnFee)
}

// SetChartityFeePercent is a paid mutator transaction binding the contract method 0x6d945814.
//
// Solidity: function setChartityFeePercent(uint256 charityFee) returns()
func (_Defi *DefiTransactor) SetChartityFeePercent(opts *bind.TransactOpts, charityFee *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "setChartityFeePercent", charityFee)
}

// SetChartityFeePercent is a paid mutator transaction binding the contract method 0x6d945814.
//
// Solidity: function setChartityFeePercent(uint256 charityFee) returns()
func (_Defi *DefiSession) SetChartityFeePercent(charityFee *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.SetChartityFeePercent(&_Defi.TransactOpts, charityFee)
}

// SetChartityFeePercent is a paid mutator transaction binding the contract method 0x6d945814.
//
// Solidity: function setChartityFeePercent(uint256 charityFee) returns()
func (_Defi *DefiTransactorSession) SetChartityFeePercent(charityFee *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.SetChartityFeePercent(&_Defi.TransactOpts, charityFee)
}

// SetLiquidityFeePercent is a paid mutator transaction binding the contract method 0x8ee88c53.
//
// Solidity: function setLiquidityFeePercent(uint256 liquidityFee) returns()
func (_Defi *DefiTransactor) SetLiquidityFeePercent(opts *bind.TransactOpts, liquidityFee *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "setLiquidityFeePercent", liquidityFee)
}

// SetLiquidityFeePercent is a paid mutator transaction binding the contract method 0x8ee88c53.
//
// Solidity: function setLiquidityFeePercent(uint256 liquidityFee) returns()
func (_Defi *DefiSession) SetLiquidityFeePercent(liquidityFee *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.SetLiquidityFeePercent(&_Defi.TransactOpts, liquidityFee)
}

// SetLiquidityFeePercent is a paid mutator transaction binding the contract method 0x8ee88c53.
//
// Solidity: function setLiquidityFeePercent(uint256 liquidityFee) returns()
func (_Defi *DefiTransactorSession) SetLiquidityFeePercent(liquidityFee *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.SetLiquidityFeePercent(&_Defi.TransactOpts, liquidityFee)
}

// SetRouterAddress is a paid mutator transaction binding the contract method 0x41cb87fc.
//
// Solidity: function setRouterAddress(address newRouter) returns()
func (_Defi *DefiTransactor) SetRouterAddress(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "setRouterAddress", newRouter)
}

// SetRouterAddress is a paid mutator transaction binding the contract method 0x41cb87fc.
//
// Solidity: function setRouterAddress(address newRouter) returns()
func (_Defi *DefiSession) SetRouterAddress(newRouter common.Address) (*types.Transaction, error) {
	return _Defi.Contract.SetRouterAddress(&_Defi.TransactOpts, newRouter)
}

// SetRouterAddress is a paid mutator transaction binding the contract method 0x41cb87fc.
//
// Solidity: function setRouterAddress(address newRouter) returns()
func (_Defi *DefiTransactorSession) SetRouterAddress(newRouter common.Address) (*types.Transaction, error) {
	return _Defi.Contract.SetRouterAddress(&_Defi.TransactOpts, newRouter)
}

// SetSwapAndLiquifyEnabled is a paid mutator transaction binding the contract method 0xc49b9a80.
//
// Solidity: function setSwapAndLiquifyEnabled(bool _enabled) returns()
func (_Defi *DefiTransactor) SetSwapAndLiquifyEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "setSwapAndLiquifyEnabled", _enabled)
}

// SetSwapAndLiquifyEnabled is a paid mutator transaction binding the contract method 0xc49b9a80.
//
// Solidity: function setSwapAndLiquifyEnabled(bool _enabled) returns()
func (_Defi *DefiSession) SetSwapAndLiquifyEnabled(_enabled bool) (*types.Transaction, error) {
	return _Defi.Contract.SetSwapAndLiquifyEnabled(&_Defi.TransactOpts, _enabled)
}

// SetSwapAndLiquifyEnabled is a paid mutator transaction binding the contract method 0xc49b9a80.
//
// Solidity: function setSwapAndLiquifyEnabled(bool _enabled) returns()
func (_Defi *DefiTransactorSession) SetSwapAndLiquifyEnabled(_enabled bool) (*types.Transaction, error) {
	return _Defi.Contract.SetSwapAndLiquifyEnabled(&_Defi.TransactOpts, _enabled)
}

// SetTaxFeePercent is a paid mutator transaction binding the contract method 0x061c82d0.
//
// Solidity: function setTaxFeePercent(uint256 taxFee) returns()
func (_Defi *DefiTransactor) SetTaxFeePercent(opts *bind.TransactOpts, taxFee *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "setTaxFeePercent", taxFee)
}

// SetTaxFeePercent is a paid mutator transaction binding the contract method 0x061c82d0.
//
// Solidity: function setTaxFeePercent(uint256 taxFee) returns()
func (_Defi *DefiSession) SetTaxFeePercent(taxFee *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.SetTaxFeePercent(&_Defi.TransactOpts, taxFee)
}

// SetTaxFeePercent is a paid mutator transaction binding the contract method 0x061c82d0.
//
// Solidity: function setTaxFeePercent(uint256 taxFee) returns()
func (_Defi *DefiTransactorSession) SetTaxFeePercent(taxFee *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.SetTaxFeePercent(&_Defi.TransactOpts, taxFee)
}

// SetcharityWallet is a paid mutator transaction binding the contract method 0x428c80e4.
//
// Solidity: function setcharityWallet(address newWallet) returns()
func (_Defi *DefiTransactor) SetcharityWallet(opts *bind.TransactOpts, newWallet common.Address) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "setcharityWallet", newWallet)
}

// SetcharityWallet is a paid mutator transaction binding the contract method 0x428c80e4.
//
// Solidity: function setcharityWallet(address newWallet) returns()
func (_Defi *DefiSession) SetcharityWallet(newWallet common.Address) (*types.Transaction, error) {
	return _Defi.Contract.SetcharityWallet(&_Defi.TransactOpts, newWallet)
}

// SetcharityWallet is a paid mutator transaction binding the contract method 0x428c80e4.
//
// Solidity: function setcharityWallet(address newWallet) returns()
func (_Defi *DefiTransactorSession) SetcharityWallet(newWallet common.Address) (*types.Transaction, error) {
	return _Defi.Contract.SetcharityWallet(&_Defi.TransactOpts, newWallet)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Defi *DefiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Defi *DefiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Transfer(&_Defi.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Defi *DefiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Transfer(&_Defi.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Defi *DefiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Defi *DefiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.TransferFrom(&_Defi.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Defi *DefiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.TransferFrom(&_Defi.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Defi *DefiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Defi *DefiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Defi.Contract.TransferOwnership(&_Defi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Defi *DefiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Defi.Contract.TransferOwnership(&_Defi.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Defi *DefiTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Defi *DefiSession) Unlock() (*types.Transaction, error) {
	return _Defi.Contract.Unlock(&_Defi.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Defi *DefiTransactorSession) Unlock() (*types.Transaction, error) {
	return _Defi.Contract.Unlock(&_Defi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Defi *DefiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defi.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Defi *DefiSession) Receive() (*types.Transaction, error) {
	return _Defi.Contract.Receive(&_Defi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Defi *DefiTransactorSession) Receive() (*types.Transaction, error) {
	return _Defi.Contract.Receive(&_Defi.TransactOpts)
}

// DefiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Defi contract.
type DefiApprovalIterator struct {
	Event *DefiApproval // Event containing the contract specifics and raw log

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
func (it *DefiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiApproval)
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
		it.Event = new(DefiApproval)
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
func (it *DefiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiApproval represents a Approval event raised by the Defi contract.
type DefiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Defi *DefiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DefiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Defi.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DefiApprovalIterator{contract: _Defi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Defi *DefiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DefiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Defi.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiApproval)
				if err := _Defi.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Defi *DefiFilterer) ParseApproval(log types.Log) (*DefiApproval, error) {
	event := new(DefiApproval)
	if err := _Defi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiMinTokensBeforeSwapUpdatedIterator is returned from FilterMinTokensBeforeSwapUpdated and is used to iterate over the raw logs and unpacked data for MinTokensBeforeSwapUpdated events raised by the Defi contract.
type DefiMinTokensBeforeSwapUpdatedIterator struct {
	Event *DefiMinTokensBeforeSwapUpdated // Event containing the contract specifics and raw log

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
func (it *DefiMinTokensBeforeSwapUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiMinTokensBeforeSwapUpdated)
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
		it.Event = new(DefiMinTokensBeforeSwapUpdated)
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
func (it *DefiMinTokensBeforeSwapUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiMinTokensBeforeSwapUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiMinTokensBeforeSwapUpdated represents a MinTokensBeforeSwapUpdated event raised by the Defi contract.
type DefiMinTokensBeforeSwapUpdated struct {
	MinTokensBeforeSwap *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMinTokensBeforeSwapUpdated is a free log retrieval operation binding the contract event 0x5948780118f41f7c4577ae4619d5cbd064057bd8562d9f7b7e60324053375c00.
//
// Solidity: event MinTokensBeforeSwapUpdated(uint256 minTokensBeforeSwap)
func (_Defi *DefiFilterer) FilterMinTokensBeforeSwapUpdated(opts *bind.FilterOpts) (*DefiMinTokensBeforeSwapUpdatedIterator, error) {

	logs, sub, err := _Defi.contract.FilterLogs(opts, "MinTokensBeforeSwapUpdated")
	if err != nil {
		return nil, err
	}
	return &DefiMinTokensBeforeSwapUpdatedIterator{contract: _Defi.contract, event: "MinTokensBeforeSwapUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTokensBeforeSwapUpdated is a free log subscription operation binding the contract event 0x5948780118f41f7c4577ae4619d5cbd064057bd8562d9f7b7e60324053375c00.
//
// Solidity: event MinTokensBeforeSwapUpdated(uint256 minTokensBeforeSwap)
func (_Defi *DefiFilterer) WatchMinTokensBeforeSwapUpdated(opts *bind.WatchOpts, sink chan<- *DefiMinTokensBeforeSwapUpdated) (event.Subscription, error) {

	logs, sub, err := _Defi.contract.WatchLogs(opts, "MinTokensBeforeSwapUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiMinTokensBeforeSwapUpdated)
				if err := _Defi.contract.UnpackLog(event, "MinTokensBeforeSwapUpdated", log); err != nil {
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

// ParseMinTokensBeforeSwapUpdated is a log parse operation binding the contract event 0x5948780118f41f7c4577ae4619d5cbd064057bd8562d9f7b7e60324053375c00.
//
// Solidity: event MinTokensBeforeSwapUpdated(uint256 minTokensBeforeSwap)
func (_Defi *DefiFilterer) ParseMinTokensBeforeSwapUpdated(log types.Log) (*DefiMinTokensBeforeSwapUpdated, error) {
	event := new(DefiMinTokensBeforeSwapUpdated)
	if err := _Defi.contract.UnpackLog(event, "MinTokensBeforeSwapUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Defi contract.
type DefiOwnershipTransferredIterator struct {
	Event *DefiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DefiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiOwnershipTransferred)
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
		it.Event = new(DefiOwnershipTransferred)
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
func (it *DefiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiOwnershipTransferred represents a OwnershipTransferred event raised by the Defi contract.
type DefiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Defi *DefiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DefiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Defi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DefiOwnershipTransferredIterator{contract: _Defi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Defi *DefiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Defi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiOwnershipTransferred)
				if err := _Defi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Defi *DefiFilterer) ParseOwnershipTransferred(log types.Log) (*DefiOwnershipTransferred, error) {
	event := new(DefiOwnershipTransferred)
	if err := _Defi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiSwapAndLiquifyIterator is returned from FilterSwapAndLiquify and is used to iterate over the raw logs and unpacked data for SwapAndLiquify events raised by the Defi contract.
type DefiSwapAndLiquifyIterator struct {
	Event *DefiSwapAndLiquify // Event containing the contract specifics and raw log

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
func (it *DefiSwapAndLiquifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiSwapAndLiquify)
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
		it.Event = new(DefiSwapAndLiquify)
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
func (it *DefiSwapAndLiquifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiSwapAndLiquifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiSwapAndLiquify represents a SwapAndLiquify event raised by the Defi contract.
type DefiSwapAndLiquify struct {
	TokensSwapped      *big.Int
	EthReceived        *big.Int
	TokensIntoLiqudity *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapAndLiquify is a free log retrieval operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Defi *DefiFilterer) FilterSwapAndLiquify(opts *bind.FilterOpts) (*DefiSwapAndLiquifyIterator, error) {

	logs, sub, err := _Defi.contract.FilterLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return &DefiSwapAndLiquifyIterator{contract: _Defi.contract, event: "SwapAndLiquify", logs: logs, sub: sub}, nil
}

// WatchSwapAndLiquify is a free log subscription operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Defi *DefiFilterer) WatchSwapAndLiquify(opts *bind.WatchOpts, sink chan<- *DefiSwapAndLiquify) (event.Subscription, error) {

	logs, sub, err := _Defi.contract.WatchLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiSwapAndLiquify)
				if err := _Defi.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
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

// ParseSwapAndLiquify is a log parse operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Defi *DefiFilterer) ParseSwapAndLiquify(log types.Log) (*DefiSwapAndLiquify, error) {
	event := new(DefiSwapAndLiquify)
	if err := _Defi.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiSwapAndLiquifyEnabledUpdatedIterator is returned from FilterSwapAndLiquifyEnabledUpdated and is used to iterate over the raw logs and unpacked data for SwapAndLiquifyEnabledUpdated events raised by the Defi contract.
type DefiSwapAndLiquifyEnabledUpdatedIterator struct {
	Event *DefiSwapAndLiquifyEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *DefiSwapAndLiquifyEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiSwapAndLiquifyEnabledUpdated)
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
		it.Event = new(DefiSwapAndLiquifyEnabledUpdated)
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
func (it *DefiSwapAndLiquifyEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiSwapAndLiquifyEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiSwapAndLiquifyEnabledUpdated represents a SwapAndLiquifyEnabledUpdated event raised by the Defi contract.
type DefiSwapAndLiquifyEnabledUpdated struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSwapAndLiquifyEnabledUpdated is a free log retrieval operation binding the contract event 0x53726dfcaf90650aa7eb35524f4d3220f07413c8d6cb404cc8c18bf5591bc159.
//
// Solidity: event SwapAndLiquifyEnabledUpdated(bool enabled)
func (_Defi *DefiFilterer) FilterSwapAndLiquifyEnabledUpdated(opts *bind.FilterOpts) (*DefiSwapAndLiquifyEnabledUpdatedIterator, error) {

	logs, sub, err := _Defi.contract.FilterLogs(opts, "SwapAndLiquifyEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &DefiSwapAndLiquifyEnabledUpdatedIterator{contract: _Defi.contract, event: "SwapAndLiquifyEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchSwapAndLiquifyEnabledUpdated is a free log subscription operation binding the contract event 0x53726dfcaf90650aa7eb35524f4d3220f07413c8d6cb404cc8c18bf5591bc159.
//
// Solidity: event SwapAndLiquifyEnabledUpdated(bool enabled)
func (_Defi *DefiFilterer) WatchSwapAndLiquifyEnabledUpdated(opts *bind.WatchOpts, sink chan<- *DefiSwapAndLiquifyEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _Defi.contract.WatchLogs(opts, "SwapAndLiquifyEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiSwapAndLiquifyEnabledUpdated)
				if err := _Defi.contract.UnpackLog(event, "SwapAndLiquifyEnabledUpdated", log); err != nil {
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

// ParseSwapAndLiquifyEnabledUpdated is a log parse operation binding the contract event 0x53726dfcaf90650aa7eb35524f4d3220f07413c8d6cb404cc8c18bf5591bc159.
//
// Solidity: event SwapAndLiquifyEnabledUpdated(bool enabled)
func (_Defi *DefiFilterer) ParseSwapAndLiquifyEnabledUpdated(log types.Log) (*DefiSwapAndLiquifyEnabledUpdated, error) {
	event := new(DefiSwapAndLiquifyEnabledUpdated)
	if err := _Defi.contract.UnpackLog(event, "SwapAndLiquifyEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Defi contract.
type DefiTransferIterator struct {
	Event *DefiTransfer // Event containing the contract specifics and raw log

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
func (it *DefiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiTransfer)
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
		it.Event = new(DefiTransfer)
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
func (it *DefiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiTransfer represents a Transfer event raised by the Defi contract.
type DefiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Defi *DefiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DefiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Defi.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DefiTransferIterator{contract: _Defi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Defi *DefiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DefiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Defi.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiTransfer)
				if err := _Defi.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Defi *DefiFilterer) ParseTransfer(log types.Log) (*DefiTransfer, error) {
	event := new(DefiTransfer)
	if err := _Defi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

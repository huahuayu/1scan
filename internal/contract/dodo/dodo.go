// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dodo

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

// DodoMetaData contains all meta data concerning the Dodo contract.
var DodoMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payQuote\",\"type\":\"uint256\"}],\"name\":\"BuyBaseToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maintainer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChargeMaintainerFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChargePenalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"}],\"name\":\"ClaimAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"}],\"name\":\"Donate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveQuote\",\"type\":\"uint256\"}],\"name\":\"SellBaseToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldGasPriceLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateGasPriceLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldK\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newK\",\"type\":\"uint256\"}],\"name\":\"UpdateK\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidityProviderFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidityProviderFeeRate\",\"type\":\"uint256\"}],\"name\":\"UpdateLiquidityProviderFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaintainerFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaintainerFeeRate\",\"type\":\"uint256\"}],\"name\":\"UpdateMaintainerFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_BASE_BALANCE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_CAPITAL_RECEIVE_QUOTE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_CAPITAL_TOKEN_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_TOKEN_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_CLAIMED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_CLOSED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_DEPOSIT_BASE_ALLOWED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_DEPOSIT_QUOTE_ALLOWED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_GAS_PRICE_LIMIT_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_K_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_LP_FEE_RATE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MAINTAINER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MT_FEE_RATE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_NEW_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_ORACLE_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_BALANCE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_CAPITAL_RECEIVE_BASE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_CAPITAL_TOKEN_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_TOKEN_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_R_STATUS_\",\"outputs\":[{\"internalType\":\"enumTypes.RStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_SUPERVISOR_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_TARGET_BASE_TOKEN_AMOUNT_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_TARGET_QUOTE_TOKEN_AMOUNT_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_TRADE_ALLOWED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPayQuote\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"buyBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositBaseTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositQuoteTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableBaseDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableQuoteDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"donateBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"donateQuoteToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableBaseDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableQuoteDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getBaseCapitalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedTarget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseTarget\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteTarget\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getLpBaseBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lpBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getLpQuoteBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lpBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"midPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOraclePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getQuoteCapitalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBaseCapital\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalQuoteCapital\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getWithdrawBasePenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getWithdrawQuotePenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"supervisor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maintainer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"queryBuyBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"querySellBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"retrieve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReceiveQuote\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sellBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"setGasPriceLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newK\",\"type\":\"uint256\"}],\"name\":\"setK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidityPorviderFeeRate\",\"type\":\"uint256\"}],\"name\":\"setLiquidityProviderFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMaintainer\",\"type\":\"address\"}],\"name\":\"setMaintainer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaintainerFeeRate\",\"type\":\"uint256\"}],\"name\":\"setMaintainerFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSupervisor\",\"type\":\"address\"}],\"name\":\"setSupervisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawAllBaseTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawAllQuoteTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBaseTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawQuoteTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DodoABI is the input ABI used to generate the binding from.
// Deprecated: Use DodoMetaData.ABI instead.
var DodoABI = DodoMetaData.ABI

// Dodo is an auto generated Go binding around an Ethereum contract.
type Dodo struct {
	DodoCaller     // Read-only binding to the contract
	DodoTransactor // Write-only binding to the contract
	DodoFilterer   // Log filterer for contract events
}

// DodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DodoSession struct {
	Contract     *Dodo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DodoCallerSession struct {
	Contract *DodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DodoTransactorSession struct {
	Contract     *DodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DodoRaw struct {
	Contract *Dodo // Generic contract binding to access the raw methods on
}

// DodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DodoCallerRaw struct {
	Contract *DodoCaller // Generic read-only contract binding to access the raw methods on
}

// DodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DodoTransactorRaw struct {
	Contract *DodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDodo creates a new instance of Dodo, bound to a specific deployed contract.
func NewDodo(address common.Address, backend bind.ContractBackend) (*Dodo, error) {
	contract, err := bindDodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dodo{DodoCaller: DodoCaller{contract: contract}, DodoTransactor: DodoTransactor{contract: contract}, DodoFilterer: DodoFilterer{contract: contract}}, nil
}

// NewDodoCaller creates a new read-only instance of Dodo, bound to a specific deployed contract.
func NewDodoCaller(address common.Address, caller bind.ContractCaller) (*DodoCaller, error) {
	contract, err := bindDodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DodoCaller{contract: contract}, nil
}

// NewDodoTransactor creates a new write-only instance of Dodo, bound to a specific deployed contract.
func NewDodoTransactor(address common.Address, transactor bind.ContractTransactor) (*DodoTransactor, error) {
	contract, err := bindDodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DodoTransactor{contract: contract}, nil
}

// NewDodoFilterer creates a new log filterer instance of Dodo, bound to a specific deployed contract.
func NewDodoFilterer(address common.Address, filterer bind.ContractFilterer) (*DodoFilterer, error) {
	contract, err := bindDodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DodoFilterer{contract: contract}, nil
}

// bindDodo binds a generic wrapper to an already deployed contract.
func bindDodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DodoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dodo *DodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dodo.Contract.DodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dodo *DodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.Contract.DodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dodo *DodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dodo.Contract.DodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dodo *DodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dodo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dodo *DodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dodo *DodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dodo.Contract.contract.Transact(opts, method, params...)
}

// BASEBALANCE is a free data retrieval call binding the contract method 0xeab5d20e.
//
// Solidity: function _BASE_BALANCE_() view returns(uint256)
func (_Dodo *DodoCaller) BASEBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_BASE_BALANCE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEBALANCE is a free data retrieval call binding the contract method 0xeab5d20e.
//
// Solidity: function _BASE_BALANCE_() view returns(uint256)
func (_Dodo *DodoSession) BASEBALANCE() (*big.Int, error) {
	return _Dodo.Contract.BASEBALANCE(&_Dodo.CallOpts)
}

// BASEBALANCE is a free data retrieval call binding the contract method 0xeab5d20e.
//
// Solidity: function _BASE_BALANCE_() view returns(uint256)
func (_Dodo *DodoCallerSession) BASEBALANCE() (*big.Int, error) {
	return _Dodo.Contract.BASEBALANCE(&_Dodo.CallOpts)
}

// BASECAPITALRECEIVEQUOTE is a free data retrieval call binding the contract method 0xc6b73cf9.
//
// Solidity: function _BASE_CAPITAL_RECEIVE_QUOTE_() view returns(uint256)
func (_Dodo *DodoCaller) BASECAPITALRECEIVEQUOTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_BASE_CAPITAL_RECEIVE_QUOTE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASECAPITALRECEIVEQUOTE is a free data retrieval call binding the contract method 0xc6b73cf9.
//
// Solidity: function _BASE_CAPITAL_RECEIVE_QUOTE_() view returns(uint256)
func (_Dodo *DodoSession) BASECAPITALRECEIVEQUOTE() (*big.Int, error) {
	return _Dodo.Contract.BASECAPITALRECEIVEQUOTE(&_Dodo.CallOpts)
}

// BASECAPITALRECEIVEQUOTE is a free data retrieval call binding the contract method 0xc6b73cf9.
//
// Solidity: function _BASE_CAPITAL_RECEIVE_QUOTE_() view returns(uint256)
func (_Dodo *DodoCallerSession) BASECAPITALRECEIVEQUOTE() (*big.Int, error) {
	return _Dodo.Contract.BASECAPITALRECEIVEQUOTE(&_Dodo.CallOpts)
}

// BASECAPITALTOKEN is a free data retrieval call binding the contract method 0xd689107c.
//
// Solidity: function _BASE_CAPITAL_TOKEN_() view returns(address)
func (_Dodo *DodoCaller) BASECAPITALTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_BASE_CAPITAL_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASECAPITALTOKEN is a free data retrieval call binding the contract method 0xd689107c.
//
// Solidity: function _BASE_CAPITAL_TOKEN_() view returns(address)
func (_Dodo *DodoSession) BASECAPITALTOKEN() (common.Address, error) {
	return _Dodo.Contract.BASECAPITALTOKEN(&_Dodo.CallOpts)
}

// BASECAPITALTOKEN is a free data retrieval call binding the contract method 0xd689107c.
//
// Solidity: function _BASE_CAPITAL_TOKEN_() view returns(address)
func (_Dodo *DodoCallerSession) BASECAPITALTOKEN() (common.Address, error) {
	return _Dodo.Contract.BASECAPITALTOKEN(&_Dodo.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_Dodo *DodoCaller) BASETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_BASE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_Dodo *DodoSession) BASETOKEN() (common.Address, error) {
	return _Dodo.Contract.BASETOKEN(&_Dodo.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_Dodo *DodoCallerSession) BASETOKEN() (common.Address, error) {
	return _Dodo.Contract.BASETOKEN(&_Dodo.CallOpts)
}

// CLAIMED is a free data retrieval call binding the contract method 0x68be20ad.
//
// Solidity: function _CLAIMED_(address ) view returns(bool)
func (_Dodo *DodoCaller) CLAIMED(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_CLAIMED_", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CLAIMED is a free data retrieval call binding the contract method 0x68be20ad.
//
// Solidity: function _CLAIMED_(address ) view returns(bool)
func (_Dodo *DodoSession) CLAIMED(arg0 common.Address) (bool, error) {
	return _Dodo.Contract.CLAIMED(&_Dodo.CallOpts, arg0)
}

// CLAIMED is a free data retrieval call binding the contract method 0x68be20ad.
//
// Solidity: function _CLAIMED_(address ) view returns(bool)
func (_Dodo *DodoCallerSession) CLAIMED(arg0 common.Address) (bool, error) {
	return _Dodo.Contract.CLAIMED(&_Dodo.CallOpts, arg0)
}

// CLOSED is a free data retrieval call binding the contract method 0x6ec6a58d.
//
// Solidity: function _CLOSED_() view returns(bool)
func (_Dodo *DodoCaller) CLOSED(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_CLOSED_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CLOSED is a free data retrieval call binding the contract method 0x6ec6a58d.
//
// Solidity: function _CLOSED_() view returns(bool)
func (_Dodo *DodoSession) CLOSED() (bool, error) {
	return _Dodo.Contract.CLOSED(&_Dodo.CallOpts)
}

// CLOSED is a free data retrieval call binding the contract method 0x6ec6a58d.
//
// Solidity: function _CLOSED_() view returns(bool)
func (_Dodo *DodoCallerSession) CLOSED() (bool, error) {
	return _Dodo.Contract.CLOSED(&_Dodo.CallOpts)
}

// DEPOSITBASEALLOWED is a free data retrieval call binding the contract method 0xa598aca7.
//
// Solidity: function _DEPOSIT_BASE_ALLOWED_() view returns(bool)
func (_Dodo *DodoCaller) DEPOSITBASEALLOWED(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_DEPOSIT_BASE_ALLOWED_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DEPOSITBASEALLOWED is a free data retrieval call binding the contract method 0xa598aca7.
//
// Solidity: function _DEPOSIT_BASE_ALLOWED_() view returns(bool)
func (_Dodo *DodoSession) DEPOSITBASEALLOWED() (bool, error) {
	return _Dodo.Contract.DEPOSITBASEALLOWED(&_Dodo.CallOpts)
}

// DEPOSITBASEALLOWED is a free data retrieval call binding the contract method 0xa598aca7.
//
// Solidity: function _DEPOSIT_BASE_ALLOWED_() view returns(bool)
func (_Dodo *DodoCallerSession) DEPOSITBASEALLOWED() (bool, error) {
	return _Dodo.Contract.DEPOSITBASEALLOWED(&_Dodo.CallOpts)
}

// DEPOSITQUOTEALLOWED is a free data retrieval call binding the contract method 0xc5bbffe8.
//
// Solidity: function _DEPOSIT_QUOTE_ALLOWED_() view returns(bool)
func (_Dodo *DodoCaller) DEPOSITQUOTEALLOWED(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_DEPOSIT_QUOTE_ALLOWED_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DEPOSITQUOTEALLOWED is a free data retrieval call binding the contract method 0xc5bbffe8.
//
// Solidity: function _DEPOSIT_QUOTE_ALLOWED_() view returns(bool)
func (_Dodo *DodoSession) DEPOSITQUOTEALLOWED() (bool, error) {
	return _Dodo.Contract.DEPOSITQUOTEALLOWED(&_Dodo.CallOpts)
}

// DEPOSITQUOTEALLOWED is a free data retrieval call binding the contract method 0xc5bbffe8.
//
// Solidity: function _DEPOSIT_QUOTE_ALLOWED_() view returns(bool)
func (_Dodo *DodoCallerSession) DEPOSITQUOTEALLOWED() (bool, error) {
	return _Dodo.Contract.DEPOSITQUOTEALLOWED(&_Dodo.CallOpts)
}

// GASPRICELIMIT is a free data retrieval call binding the contract method 0x4de4527e.
//
// Solidity: function _GAS_PRICE_LIMIT_() view returns(uint256)
func (_Dodo *DodoCaller) GASPRICELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_GAS_PRICE_LIMIT_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASPRICELIMIT is a free data retrieval call binding the contract method 0x4de4527e.
//
// Solidity: function _GAS_PRICE_LIMIT_() view returns(uint256)
func (_Dodo *DodoSession) GASPRICELIMIT() (*big.Int, error) {
	return _Dodo.Contract.GASPRICELIMIT(&_Dodo.CallOpts)
}

// GASPRICELIMIT is a free data retrieval call binding the contract method 0x4de4527e.
//
// Solidity: function _GAS_PRICE_LIMIT_() view returns(uint256)
func (_Dodo *DodoCallerSession) GASPRICELIMIT() (*big.Int, error) {
	return _Dodo.Contract.GASPRICELIMIT(&_Dodo.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint256)
func (_Dodo *DodoCaller) K(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_K_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint256)
func (_Dodo *DodoSession) K() (*big.Int, error) {
	return _Dodo.Contract.K(&_Dodo.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint256)
func (_Dodo *DodoCallerSession) K() (*big.Int, error) {
	return _Dodo.Contract.K(&_Dodo.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint256)
func (_Dodo *DodoCaller) LPFEERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_LP_FEE_RATE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint256)
func (_Dodo *DodoSession) LPFEERATE() (*big.Int, error) {
	return _Dodo.Contract.LPFEERATE(&_Dodo.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint256)
func (_Dodo *DodoCallerSession) LPFEERATE() (*big.Int, error) {
	return _Dodo.Contract.LPFEERATE(&_Dodo.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_Dodo *DodoCaller) MAINTAINER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_MAINTAINER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_Dodo *DodoSession) MAINTAINER() (common.Address, error) {
	return _Dodo.Contract.MAINTAINER(&_Dodo.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_Dodo *DodoCallerSession) MAINTAINER() (common.Address, error) {
	return _Dodo.Contract.MAINTAINER(&_Dodo.CallOpts)
}

// MTFEERATE is a free data retrieval call binding the contract method 0xc0ffa178.
//
// Solidity: function _MT_FEE_RATE_() view returns(uint256)
func (_Dodo *DodoCaller) MTFEERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_MT_FEE_RATE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MTFEERATE is a free data retrieval call binding the contract method 0xc0ffa178.
//
// Solidity: function _MT_FEE_RATE_() view returns(uint256)
func (_Dodo *DodoSession) MTFEERATE() (*big.Int, error) {
	return _Dodo.Contract.MTFEERATE(&_Dodo.CallOpts)
}

// MTFEERATE is a free data retrieval call binding the contract method 0xc0ffa178.
//
// Solidity: function _MT_FEE_RATE_() view returns(uint256)
func (_Dodo *DodoCallerSession) MTFEERATE() (*big.Int, error) {
	return _Dodo.Contract.MTFEERATE(&_Dodo.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_Dodo *DodoCaller) NEWOWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_NEW_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_Dodo *DodoSession) NEWOWNER() (common.Address, error) {
	return _Dodo.Contract.NEWOWNER(&_Dodo.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_Dodo *DodoCallerSession) NEWOWNER() (common.Address, error) {
	return _Dodo.Contract.NEWOWNER(&_Dodo.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x73a2ab7c.
//
// Solidity: function _ORACLE_() view returns(address)
func (_Dodo *DodoCaller) ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_ORACLE_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLE is a free data retrieval call binding the contract method 0x73a2ab7c.
//
// Solidity: function _ORACLE_() view returns(address)
func (_Dodo *DodoSession) ORACLE() (common.Address, error) {
	return _Dodo.Contract.ORACLE(&_Dodo.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x73a2ab7c.
//
// Solidity: function _ORACLE_() view returns(address)
func (_Dodo *DodoCallerSession) ORACLE() (common.Address, error) {
	return _Dodo.Contract.ORACLE(&_Dodo.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_Dodo *DodoCaller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_Dodo *DodoSession) OWNER() (common.Address, error) {
	return _Dodo.Contract.OWNER(&_Dodo.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_Dodo *DodoCallerSession) OWNER() (common.Address, error) {
	return _Dodo.Contract.OWNER(&_Dodo.CallOpts)
}

// QUOTEBALANCE is a free data retrieval call binding the contract method 0x7c9b8e89.
//
// Solidity: function _QUOTE_BALANCE_() view returns(uint256)
func (_Dodo *DodoCaller) QUOTEBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_QUOTE_BALANCE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTEBALANCE is a free data retrieval call binding the contract method 0x7c9b8e89.
//
// Solidity: function _QUOTE_BALANCE_() view returns(uint256)
func (_Dodo *DodoSession) QUOTEBALANCE() (*big.Int, error) {
	return _Dodo.Contract.QUOTEBALANCE(&_Dodo.CallOpts)
}

// QUOTEBALANCE is a free data retrieval call binding the contract method 0x7c9b8e89.
//
// Solidity: function _QUOTE_BALANCE_() view returns(uint256)
func (_Dodo *DodoCallerSession) QUOTEBALANCE() (*big.Int, error) {
	return _Dodo.Contract.QUOTEBALANCE(&_Dodo.CallOpts)
}

// QUOTECAPITALRECEIVEBASE is a free data retrieval call binding the contract method 0x0e6518e9.
//
// Solidity: function _QUOTE_CAPITAL_RECEIVE_BASE_() view returns(uint256)
func (_Dodo *DodoCaller) QUOTECAPITALRECEIVEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_QUOTE_CAPITAL_RECEIVE_BASE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTECAPITALRECEIVEBASE is a free data retrieval call binding the contract method 0x0e6518e9.
//
// Solidity: function _QUOTE_CAPITAL_RECEIVE_BASE_() view returns(uint256)
func (_Dodo *DodoSession) QUOTECAPITALRECEIVEBASE() (*big.Int, error) {
	return _Dodo.Contract.QUOTECAPITALRECEIVEBASE(&_Dodo.CallOpts)
}

// QUOTECAPITALRECEIVEBASE is a free data retrieval call binding the contract method 0x0e6518e9.
//
// Solidity: function _QUOTE_CAPITAL_RECEIVE_BASE_() view returns(uint256)
func (_Dodo *DodoCallerSession) QUOTECAPITALRECEIVEBASE() (*big.Int, error) {
	return _Dodo.Contract.QUOTECAPITALRECEIVEBASE(&_Dodo.CallOpts)
}

// QUOTECAPITALTOKEN is a free data retrieval call binding the contract method 0xac1fbc98.
//
// Solidity: function _QUOTE_CAPITAL_TOKEN_() view returns(address)
func (_Dodo *DodoCaller) QUOTECAPITALTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_QUOTE_CAPITAL_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTECAPITALTOKEN is a free data retrieval call binding the contract method 0xac1fbc98.
//
// Solidity: function _QUOTE_CAPITAL_TOKEN_() view returns(address)
func (_Dodo *DodoSession) QUOTECAPITALTOKEN() (common.Address, error) {
	return _Dodo.Contract.QUOTECAPITALTOKEN(&_Dodo.CallOpts)
}

// QUOTECAPITALTOKEN is a free data retrieval call binding the contract method 0xac1fbc98.
//
// Solidity: function _QUOTE_CAPITAL_TOKEN_() view returns(address)
func (_Dodo *DodoCallerSession) QUOTECAPITALTOKEN() (common.Address, error) {
	return _Dodo.Contract.QUOTECAPITALTOKEN(&_Dodo.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_Dodo *DodoCaller) QUOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_QUOTE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_Dodo *DodoSession) QUOTETOKEN() (common.Address, error) {
	return _Dodo.Contract.QUOTETOKEN(&_Dodo.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_Dodo *DodoCallerSession) QUOTETOKEN() (common.Address, error) {
	return _Dodo.Contract.QUOTETOKEN(&_Dodo.CallOpts)
}

// RSTATUS is a free data retrieval call binding the contract method 0x17be952e.
//
// Solidity: function _R_STATUS_() view returns(uint8)
func (_Dodo *DodoCaller) RSTATUS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_R_STATUS_")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RSTATUS is a free data retrieval call binding the contract method 0x17be952e.
//
// Solidity: function _R_STATUS_() view returns(uint8)
func (_Dodo *DodoSession) RSTATUS() (uint8, error) {
	return _Dodo.Contract.RSTATUS(&_Dodo.CallOpts)
}

// RSTATUS is a free data retrieval call binding the contract method 0x17be952e.
//
// Solidity: function _R_STATUS_() view returns(uint8)
func (_Dodo *DodoCallerSession) RSTATUS() (uint8, error) {
	return _Dodo.Contract.RSTATUS(&_Dodo.CallOpts)
}

// SUPERVISOR is a free data retrieval call binding the contract method 0x3960f142.
//
// Solidity: function _SUPERVISOR_() view returns(address)
func (_Dodo *DodoCaller) SUPERVISOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_SUPERVISOR_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SUPERVISOR is a free data retrieval call binding the contract method 0x3960f142.
//
// Solidity: function _SUPERVISOR_() view returns(address)
func (_Dodo *DodoSession) SUPERVISOR() (common.Address, error) {
	return _Dodo.Contract.SUPERVISOR(&_Dodo.CallOpts)
}

// SUPERVISOR is a free data retrieval call binding the contract method 0x3960f142.
//
// Solidity: function _SUPERVISOR_() view returns(address)
func (_Dodo *DodoCallerSession) SUPERVISOR() (common.Address, error) {
	return _Dodo.Contract.SUPERVISOR(&_Dodo.CallOpts)
}

// TARGETBASETOKENAMOUNT is a free data retrieval call binding the contract method 0xb2094fd3.
//
// Solidity: function _TARGET_BASE_TOKEN_AMOUNT_() view returns(uint256)
func (_Dodo *DodoCaller) TARGETBASETOKENAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_TARGET_BASE_TOKEN_AMOUNT_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TARGETBASETOKENAMOUNT is a free data retrieval call binding the contract method 0xb2094fd3.
//
// Solidity: function _TARGET_BASE_TOKEN_AMOUNT_() view returns(uint256)
func (_Dodo *DodoSession) TARGETBASETOKENAMOUNT() (*big.Int, error) {
	return _Dodo.Contract.TARGETBASETOKENAMOUNT(&_Dodo.CallOpts)
}

// TARGETBASETOKENAMOUNT is a free data retrieval call binding the contract method 0xb2094fd3.
//
// Solidity: function _TARGET_BASE_TOKEN_AMOUNT_() view returns(uint256)
func (_Dodo *DodoCallerSession) TARGETBASETOKENAMOUNT() (*big.Int, error) {
	return _Dodo.Contract.TARGETBASETOKENAMOUNT(&_Dodo.CallOpts)
}

// TARGETQUOTETOKENAMOUNT is a free data retrieval call binding the contract method 0x245c9685.
//
// Solidity: function _TARGET_QUOTE_TOKEN_AMOUNT_() view returns(uint256)
func (_Dodo *DodoCaller) TARGETQUOTETOKENAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_TARGET_QUOTE_TOKEN_AMOUNT_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TARGETQUOTETOKENAMOUNT is a free data retrieval call binding the contract method 0x245c9685.
//
// Solidity: function _TARGET_QUOTE_TOKEN_AMOUNT_() view returns(uint256)
func (_Dodo *DodoSession) TARGETQUOTETOKENAMOUNT() (*big.Int, error) {
	return _Dodo.Contract.TARGETQUOTETOKENAMOUNT(&_Dodo.CallOpts)
}

// TARGETQUOTETOKENAMOUNT is a free data retrieval call binding the contract method 0x245c9685.
//
// Solidity: function _TARGET_QUOTE_TOKEN_AMOUNT_() view returns(uint256)
func (_Dodo *DodoCallerSession) TARGETQUOTETOKENAMOUNT() (*big.Int, error) {
	return _Dodo.Contract.TARGETQUOTETOKENAMOUNT(&_Dodo.CallOpts)
}

// TRADEALLOWED is a free data retrieval call binding the contract method 0xdd58b41c.
//
// Solidity: function _TRADE_ALLOWED_() view returns(bool)
func (_Dodo *DodoCaller) TRADEALLOWED(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "_TRADE_ALLOWED_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TRADEALLOWED is a free data retrieval call binding the contract method 0xdd58b41c.
//
// Solidity: function _TRADE_ALLOWED_() view returns(bool)
func (_Dodo *DodoSession) TRADEALLOWED() (bool, error) {
	return _Dodo.Contract.TRADEALLOWED(&_Dodo.CallOpts)
}

// TRADEALLOWED is a free data retrieval call binding the contract method 0xdd58b41c.
//
// Solidity: function _TRADE_ALLOWED_() view returns(bool)
func (_Dodo *DodoCallerSession) TRADEALLOWED() (bool, error) {
	return _Dodo.Contract.TRADEALLOWED(&_Dodo.CallOpts)
}

// GetBaseCapitalBalanceOf is a free data retrieval call binding the contract method 0x7aed942d.
//
// Solidity: function getBaseCapitalBalanceOf(address lp) view returns(uint256)
func (_Dodo *DodoCaller) GetBaseCapitalBalanceOf(opts *bind.CallOpts, lp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getBaseCapitalBalanceOf", lp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseCapitalBalanceOf is a free data retrieval call binding the contract method 0x7aed942d.
//
// Solidity: function getBaseCapitalBalanceOf(address lp) view returns(uint256)
func (_Dodo *DodoSession) GetBaseCapitalBalanceOf(lp common.Address) (*big.Int, error) {
	return _Dodo.Contract.GetBaseCapitalBalanceOf(&_Dodo.CallOpts, lp)
}

// GetBaseCapitalBalanceOf is a free data retrieval call binding the contract method 0x7aed942d.
//
// Solidity: function getBaseCapitalBalanceOf(address lp) view returns(uint256)
func (_Dodo *DodoCallerSession) GetBaseCapitalBalanceOf(lp common.Address) (*big.Int, error) {
	return _Dodo.Contract.GetBaseCapitalBalanceOf(&_Dodo.CallOpts, lp)
}

// GetExpectedTarget is a free data retrieval call binding the contract method 0xffa64225.
//
// Solidity: function getExpectedTarget() view returns(uint256 baseTarget, uint256 quoteTarget)
func (_Dodo *DodoCaller) GetExpectedTarget(opts *bind.CallOpts) (struct {
	BaseTarget  *big.Int
	QuoteTarget *big.Int
}, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getExpectedTarget")

	outstruct := new(struct {
		BaseTarget  *big.Int
		QuoteTarget *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseTarget = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QuoteTarget = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetExpectedTarget is a free data retrieval call binding the contract method 0xffa64225.
//
// Solidity: function getExpectedTarget() view returns(uint256 baseTarget, uint256 quoteTarget)
func (_Dodo *DodoSession) GetExpectedTarget() (struct {
	BaseTarget  *big.Int
	QuoteTarget *big.Int
}, error) {
	return _Dodo.Contract.GetExpectedTarget(&_Dodo.CallOpts)
}

// GetExpectedTarget is a free data retrieval call binding the contract method 0xffa64225.
//
// Solidity: function getExpectedTarget() view returns(uint256 baseTarget, uint256 quoteTarget)
func (_Dodo *DodoCallerSession) GetExpectedTarget() (struct {
	BaseTarget  *big.Int
	QuoteTarget *big.Int
}, error) {
	return _Dodo.Contract.GetExpectedTarget(&_Dodo.CallOpts)
}

// GetLpBaseBalance is a free data retrieval call binding the contract method 0x95faa5f6.
//
// Solidity: function getLpBaseBalance(address lp) view returns(uint256 lpBalance)
func (_Dodo *DodoCaller) GetLpBaseBalance(opts *bind.CallOpts, lp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getLpBaseBalance", lp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLpBaseBalance is a free data retrieval call binding the contract method 0x95faa5f6.
//
// Solidity: function getLpBaseBalance(address lp) view returns(uint256 lpBalance)
func (_Dodo *DodoSession) GetLpBaseBalance(lp common.Address) (*big.Int, error) {
	return _Dodo.Contract.GetLpBaseBalance(&_Dodo.CallOpts, lp)
}

// GetLpBaseBalance is a free data retrieval call binding the contract method 0x95faa5f6.
//
// Solidity: function getLpBaseBalance(address lp) view returns(uint256 lpBalance)
func (_Dodo *DodoCallerSession) GetLpBaseBalance(lp common.Address) (*big.Int, error) {
	return _Dodo.Contract.GetLpBaseBalance(&_Dodo.CallOpts, lp)
}

// GetLpQuoteBalance is a free data retrieval call binding the contract method 0x36a53bbb.
//
// Solidity: function getLpQuoteBalance(address lp) view returns(uint256 lpBalance)
func (_Dodo *DodoCaller) GetLpQuoteBalance(opts *bind.CallOpts, lp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getLpQuoteBalance", lp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLpQuoteBalance is a free data retrieval call binding the contract method 0x36a53bbb.
//
// Solidity: function getLpQuoteBalance(address lp) view returns(uint256 lpBalance)
func (_Dodo *DodoSession) GetLpQuoteBalance(lp common.Address) (*big.Int, error) {
	return _Dodo.Contract.GetLpQuoteBalance(&_Dodo.CallOpts, lp)
}

// GetLpQuoteBalance is a free data retrieval call binding the contract method 0x36a53bbb.
//
// Solidity: function getLpQuoteBalance(address lp) view returns(uint256 lpBalance)
func (_Dodo *DodoCallerSession) GetLpQuoteBalance(lp common.Address) (*big.Int, error) {
	return _Dodo.Contract.GetLpQuoteBalance(&_Dodo.CallOpts, lp)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_Dodo *DodoCaller) GetMidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getMidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_Dodo *DodoSession) GetMidPrice() (*big.Int, error) {
	return _Dodo.Contract.GetMidPrice(&_Dodo.CallOpts)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_Dodo *DodoCallerSession) GetMidPrice() (*big.Int, error) {
	return _Dodo.Contract.GetMidPrice(&_Dodo.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(uint256)
func (_Dodo *DodoCaller) GetOraclePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getOraclePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(uint256)
func (_Dodo *DodoSession) GetOraclePrice() (*big.Int, error) {
	return _Dodo.Contract.GetOraclePrice(&_Dodo.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(uint256)
func (_Dodo *DodoCallerSession) GetOraclePrice() (*big.Int, error) {
	return _Dodo.Contract.GetOraclePrice(&_Dodo.CallOpts)
}

// GetQuoteCapitalBalanceOf is a free data retrieval call binding the contract method 0xf67ed448.
//
// Solidity: function getQuoteCapitalBalanceOf(address lp) view returns(uint256)
func (_Dodo *DodoCaller) GetQuoteCapitalBalanceOf(opts *bind.CallOpts, lp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getQuoteCapitalBalanceOf", lp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteCapitalBalanceOf is a free data retrieval call binding the contract method 0xf67ed448.
//
// Solidity: function getQuoteCapitalBalanceOf(address lp) view returns(uint256)
func (_Dodo *DodoSession) GetQuoteCapitalBalanceOf(lp common.Address) (*big.Int, error) {
	return _Dodo.Contract.GetQuoteCapitalBalanceOf(&_Dodo.CallOpts, lp)
}

// GetQuoteCapitalBalanceOf is a free data retrieval call binding the contract method 0xf67ed448.
//
// Solidity: function getQuoteCapitalBalanceOf(address lp) view returns(uint256)
func (_Dodo *DodoCallerSession) GetQuoteCapitalBalanceOf(lp common.Address) (*big.Int, error) {
	return _Dodo.Contract.GetQuoteCapitalBalanceOf(&_Dodo.CallOpts, lp)
}

// GetTotalBaseCapital is a free data retrieval call binding the contract method 0x0cd1667d.
//
// Solidity: function getTotalBaseCapital() view returns(uint256)
func (_Dodo *DodoCaller) GetTotalBaseCapital(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getTotalBaseCapital")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBaseCapital is a free data retrieval call binding the contract method 0x0cd1667d.
//
// Solidity: function getTotalBaseCapital() view returns(uint256)
func (_Dodo *DodoSession) GetTotalBaseCapital() (*big.Int, error) {
	return _Dodo.Contract.GetTotalBaseCapital(&_Dodo.CallOpts)
}

// GetTotalBaseCapital is a free data retrieval call binding the contract method 0x0cd1667d.
//
// Solidity: function getTotalBaseCapital() view returns(uint256)
func (_Dodo *DodoCallerSession) GetTotalBaseCapital() (*big.Int, error) {
	return _Dodo.Contract.GetTotalBaseCapital(&_Dodo.CallOpts)
}

// GetTotalQuoteCapital is a free data retrieval call binding the contract method 0x2aa82c65.
//
// Solidity: function getTotalQuoteCapital() view returns(uint256)
func (_Dodo *DodoCaller) GetTotalQuoteCapital(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getTotalQuoteCapital")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalQuoteCapital is a free data retrieval call binding the contract method 0x2aa82c65.
//
// Solidity: function getTotalQuoteCapital() view returns(uint256)
func (_Dodo *DodoSession) GetTotalQuoteCapital() (*big.Int, error) {
	return _Dodo.Contract.GetTotalQuoteCapital(&_Dodo.CallOpts)
}

// GetTotalQuoteCapital is a free data retrieval call binding the contract method 0x2aa82c65.
//
// Solidity: function getTotalQuoteCapital() view returns(uint256)
func (_Dodo *DodoCallerSession) GetTotalQuoteCapital() (*big.Int, error) {
	return _Dodo.Contract.GetTotalQuoteCapital(&_Dodo.CallOpts)
}

// GetWithdrawBasePenalty is a free data retrieval call binding the contract method 0xee5150b3.
//
// Solidity: function getWithdrawBasePenalty(uint256 amount) view returns(uint256 penalty)
func (_Dodo *DodoCaller) GetWithdrawBasePenalty(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getWithdrawBasePenalty", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawBasePenalty is a free data retrieval call binding the contract method 0xee5150b3.
//
// Solidity: function getWithdrawBasePenalty(uint256 amount) view returns(uint256 penalty)
func (_Dodo *DodoSession) GetWithdrawBasePenalty(amount *big.Int) (*big.Int, error) {
	return _Dodo.Contract.GetWithdrawBasePenalty(&_Dodo.CallOpts, amount)
}

// GetWithdrawBasePenalty is a free data retrieval call binding the contract method 0xee5150b3.
//
// Solidity: function getWithdrawBasePenalty(uint256 amount) view returns(uint256 penalty)
func (_Dodo *DodoCallerSession) GetWithdrawBasePenalty(amount *big.Int) (*big.Int, error) {
	return _Dodo.Contract.GetWithdrawBasePenalty(&_Dodo.CallOpts, amount)
}

// GetWithdrawQuotePenalty is a free data retrieval call binding the contract method 0x0c9f7bd0.
//
// Solidity: function getWithdrawQuotePenalty(uint256 amount) view returns(uint256 penalty)
func (_Dodo *DodoCaller) GetWithdrawQuotePenalty(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "getWithdrawQuotePenalty", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawQuotePenalty is a free data retrieval call binding the contract method 0x0c9f7bd0.
//
// Solidity: function getWithdrawQuotePenalty(uint256 amount) view returns(uint256 penalty)
func (_Dodo *DodoSession) GetWithdrawQuotePenalty(amount *big.Int) (*big.Int, error) {
	return _Dodo.Contract.GetWithdrawQuotePenalty(&_Dodo.CallOpts, amount)
}

// GetWithdrawQuotePenalty is a free data retrieval call binding the contract method 0x0c9f7bd0.
//
// Solidity: function getWithdrawQuotePenalty(uint256 amount) view returns(uint256 penalty)
func (_Dodo *DodoCallerSession) GetWithdrawQuotePenalty(amount *big.Int) (*big.Int, error) {
	return _Dodo.Contract.GetWithdrawQuotePenalty(&_Dodo.CallOpts, amount)
}

// QueryBuyBaseToken is a free data retrieval call binding the contract method 0x18c0bbe4.
//
// Solidity: function queryBuyBaseToken(uint256 amount) view returns(uint256 payQuote)
func (_Dodo *DodoCaller) QueryBuyBaseToken(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "queryBuyBaseToken", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryBuyBaseToken is a free data retrieval call binding the contract method 0x18c0bbe4.
//
// Solidity: function queryBuyBaseToken(uint256 amount) view returns(uint256 payQuote)
func (_Dodo *DodoSession) QueryBuyBaseToken(amount *big.Int) (*big.Int, error) {
	return _Dodo.Contract.QueryBuyBaseToken(&_Dodo.CallOpts, amount)
}

// QueryBuyBaseToken is a free data retrieval call binding the contract method 0x18c0bbe4.
//
// Solidity: function queryBuyBaseToken(uint256 amount) view returns(uint256 payQuote)
func (_Dodo *DodoCallerSession) QueryBuyBaseToken(amount *big.Int) (*big.Int, error) {
	return _Dodo.Contract.QueryBuyBaseToken(&_Dodo.CallOpts, amount)
}

// QuerySellBaseToken is a free data retrieval call binding the contract method 0xa2801e16.
//
// Solidity: function querySellBaseToken(uint256 amount) view returns(uint256 receiveQuote)
func (_Dodo *DodoCaller) QuerySellBaseToken(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "querySellBaseToken", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellBaseToken is a free data retrieval call binding the contract method 0xa2801e16.
//
// Solidity: function querySellBaseToken(uint256 amount) view returns(uint256 receiveQuote)
func (_Dodo *DodoSession) QuerySellBaseToken(amount *big.Int) (*big.Int, error) {
	return _Dodo.Contract.QuerySellBaseToken(&_Dodo.CallOpts, amount)
}

// QuerySellBaseToken is a free data retrieval call binding the contract method 0xa2801e16.
//
// Solidity: function querySellBaseToken(uint256 amount) view returns(uint256 receiveQuote)
func (_Dodo *DodoCallerSession) QuerySellBaseToken(amount *big.Int) (*big.Int, error) {
	return _Dodo.Contract.QuerySellBaseToken(&_Dodo.CallOpts, amount)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Dodo *DodoCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dodo.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Dodo *DodoSession) Version() (*big.Int, error) {
	return _Dodo.Contract.Version(&_Dodo.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Dodo *DodoCallerSession) Version() (*big.Int, error) {
	return _Dodo.Contract.Version(&_Dodo.CallOpts)
}

// BuyBaseToken is a paid mutator transaction binding the contract method 0xe67ce706.
//
// Solidity: function buyBaseToken(uint256 amount, uint256 maxPayQuote, bytes data) returns(uint256)
func (_Dodo *DodoTransactor) BuyBaseToken(opts *bind.TransactOpts, amount *big.Int, maxPayQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "buyBaseToken", amount, maxPayQuote, data)
}

// BuyBaseToken is a paid mutator transaction binding the contract method 0xe67ce706.
//
// Solidity: function buyBaseToken(uint256 amount, uint256 maxPayQuote, bytes data) returns(uint256)
func (_Dodo *DodoSession) BuyBaseToken(amount *big.Int, maxPayQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Dodo.Contract.BuyBaseToken(&_Dodo.TransactOpts, amount, maxPayQuote, data)
}

// BuyBaseToken is a paid mutator transaction binding the contract method 0xe67ce706.
//
// Solidity: function buyBaseToken(uint256 amount, uint256 maxPayQuote, bytes data) returns(uint256)
func (_Dodo *DodoTransactorSession) BuyBaseToken(amount *big.Int, maxPayQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Dodo.Contract.BuyBaseToken(&_Dodo.TransactOpts, amount, maxPayQuote, data)
}

// ClaimAssets is a paid mutator transaction binding the contract method 0x1f3c156e.
//
// Solidity: function claimAssets() returns()
func (_Dodo *DodoTransactor) ClaimAssets(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "claimAssets")
}

// ClaimAssets is a paid mutator transaction binding the contract method 0x1f3c156e.
//
// Solidity: function claimAssets() returns()
func (_Dodo *DodoSession) ClaimAssets() (*types.Transaction, error) {
	return _Dodo.Contract.ClaimAssets(&_Dodo.TransactOpts)
}

// ClaimAssets is a paid mutator transaction binding the contract method 0x1f3c156e.
//
// Solidity: function claimAssets() returns()
func (_Dodo *DodoTransactorSession) ClaimAssets() (*types.Transaction, error) {
	return _Dodo.Contract.ClaimAssets(&_Dodo.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Dodo *DodoTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Dodo *DodoSession) ClaimOwnership() (*types.Transaction, error) {
	return _Dodo.Contract.ClaimOwnership(&_Dodo.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Dodo *DodoTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Dodo.Contract.ClaimOwnership(&_Dodo.TransactOpts)
}

// DepositBase is a paid mutator transaction binding the contract method 0x27bed8ee.
//
// Solidity: function depositBase(uint256 amount) returns(uint256)
func (_Dodo *DodoTransactor) DepositBase(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "depositBase", amount)
}

// DepositBase is a paid mutator transaction binding the contract method 0x27bed8ee.
//
// Solidity: function depositBase(uint256 amount) returns(uint256)
func (_Dodo *DodoSession) DepositBase(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DepositBase(&_Dodo.TransactOpts, amount)
}

// DepositBase is a paid mutator transaction binding the contract method 0x27bed8ee.
//
// Solidity: function depositBase(uint256 amount) returns(uint256)
func (_Dodo *DodoTransactorSession) DepositBase(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DepositBase(&_Dodo.TransactOpts, amount)
}

// DepositBaseTo is a paid mutator transaction binding the contract method 0xaa06ce9b.
//
// Solidity: function depositBaseTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoTransactor) DepositBaseTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "depositBaseTo", to, amount)
}

// DepositBaseTo is a paid mutator transaction binding the contract method 0xaa06ce9b.
//
// Solidity: function depositBaseTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoSession) DepositBaseTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DepositBaseTo(&_Dodo.TransactOpts, to, amount)
}

// DepositBaseTo is a paid mutator transaction binding the contract method 0xaa06ce9b.
//
// Solidity: function depositBaseTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoTransactorSession) DepositBaseTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DepositBaseTo(&_Dodo.TransactOpts, to, amount)
}

// DepositQuote is a paid mutator transaction binding the contract method 0xf3ae6c5f.
//
// Solidity: function depositQuote(uint256 amount) returns(uint256)
func (_Dodo *DodoTransactor) DepositQuote(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "depositQuote", amount)
}

// DepositQuote is a paid mutator transaction binding the contract method 0xf3ae6c5f.
//
// Solidity: function depositQuote(uint256 amount) returns(uint256)
func (_Dodo *DodoSession) DepositQuote(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DepositQuote(&_Dodo.TransactOpts, amount)
}

// DepositQuote is a paid mutator transaction binding the contract method 0xf3ae6c5f.
//
// Solidity: function depositQuote(uint256 amount) returns(uint256)
func (_Dodo *DodoTransactorSession) DepositQuote(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DepositQuote(&_Dodo.TransactOpts, amount)
}

// DepositQuoteTo is a paid mutator transaction binding the contract method 0x5f179f64.
//
// Solidity: function depositQuoteTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoTransactor) DepositQuoteTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "depositQuoteTo", to, amount)
}

// DepositQuoteTo is a paid mutator transaction binding the contract method 0x5f179f64.
//
// Solidity: function depositQuoteTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoSession) DepositQuoteTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DepositQuoteTo(&_Dodo.TransactOpts, to, amount)
}

// DepositQuoteTo is a paid mutator transaction binding the contract method 0x5f179f64.
//
// Solidity: function depositQuoteTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoTransactorSession) DepositQuoteTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DepositQuoteTo(&_Dodo.TransactOpts, to, amount)
}

// DisableBaseDeposit is a paid mutator transaction binding the contract method 0x13c57624.
//
// Solidity: function disableBaseDeposit() returns()
func (_Dodo *DodoTransactor) DisableBaseDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "disableBaseDeposit")
}

// DisableBaseDeposit is a paid mutator transaction binding the contract method 0x13c57624.
//
// Solidity: function disableBaseDeposit() returns()
func (_Dodo *DodoSession) DisableBaseDeposit() (*types.Transaction, error) {
	return _Dodo.Contract.DisableBaseDeposit(&_Dodo.TransactOpts)
}

// DisableBaseDeposit is a paid mutator transaction binding the contract method 0x13c57624.
//
// Solidity: function disableBaseDeposit() returns()
func (_Dodo *DodoTransactorSession) DisableBaseDeposit() (*types.Transaction, error) {
	return _Dodo.Contract.DisableBaseDeposit(&_Dodo.TransactOpts)
}

// DisableQuoteDeposit is a paid mutator transaction binding the contract method 0xbc7d679d.
//
// Solidity: function disableQuoteDeposit() returns()
func (_Dodo *DodoTransactor) DisableQuoteDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "disableQuoteDeposit")
}

// DisableQuoteDeposit is a paid mutator transaction binding the contract method 0xbc7d679d.
//
// Solidity: function disableQuoteDeposit() returns()
func (_Dodo *DodoSession) DisableQuoteDeposit() (*types.Transaction, error) {
	return _Dodo.Contract.DisableQuoteDeposit(&_Dodo.TransactOpts)
}

// DisableQuoteDeposit is a paid mutator transaction binding the contract method 0xbc7d679d.
//
// Solidity: function disableQuoteDeposit() returns()
func (_Dodo *DodoTransactorSession) DisableQuoteDeposit() (*types.Transaction, error) {
	return _Dodo.Contract.DisableQuoteDeposit(&_Dodo.TransactOpts)
}

// DisableTrading is a paid mutator transaction binding the contract method 0x17700f01.
//
// Solidity: function disableTrading() returns()
func (_Dodo *DodoTransactor) DisableTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "disableTrading")
}

// DisableTrading is a paid mutator transaction binding the contract method 0x17700f01.
//
// Solidity: function disableTrading() returns()
func (_Dodo *DodoSession) DisableTrading() (*types.Transaction, error) {
	return _Dodo.Contract.DisableTrading(&_Dodo.TransactOpts)
}

// DisableTrading is a paid mutator transaction binding the contract method 0x17700f01.
//
// Solidity: function disableTrading() returns()
func (_Dodo *DodoTransactorSession) DisableTrading() (*types.Transaction, error) {
	return _Dodo.Contract.DisableTrading(&_Dodo.TransactOpts)
}

// DonateBaseToken is a paid mutator transaction binding the contract method 0xed0aa428.
//
// Solidity: function donateBaseToken(uint256 amount) returns()
func (_Dodo *DodoTransactor) DonateBaseToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "donateBaseToken", amount)
}

// DonateBaseToken is a paid mutator transaction binding the contract method 0xed0aa428.
//
// Solidity: function donateBaseToken(uint256 amount) returns()
func (_Dodo *DodoSession) DonateBaseToken(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DonateBaseToken(&_Dodo.TransactOpts, amount)
}

// DonateBaseToken is a paid mutator transaction binding the contract method 0xed0aa428.
//
// Solidity: function donateBaseToken(uint256 amount) returns()
func (_Dodo *DodoTransactorSession) DonateBaseToken(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DonateBaseToken(&_Dodo.TransactOpts, amount)
}

// DonateQuoteToken is a paid mutator transaction binding the contract method 0x387b0c11.
//
// Solidity: function donateQuoteToken(uint256 amount) returns()
func (_Dodo *DodoTransactor) DonateQuoteToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "donateQuoteToken", amount)
}

// DonateQuoteToken is a paid mutator transaction binding the contract method 0x387b0c11.
//
// Solidity: function donateQuoteToken(uint256 amount) returns()
func (_Dodo *DodoSession) DonateQuoteToken(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DonateQuoteToken(&_Dodo.TransactOpts, amount)
}

// DonateQuoteToken is a paid mutator transaction binding the contract method 0x387b0c11.
//
// Solidity: function donateQuoteToken(uint256 amount) returns()
func (_Dodo *DodoTransactorSession) DonateQuoteToken(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.DonateQuoteToken(&_Dodo.TransactOpts, amount)
}

// EnableBaseDeposit is a paid mutator transaction binding the contract method 0x1184d8be.
//
// Solidity: function enableBaseDeposit() returns()
func (_Dodo *DodoTransactor) EnableBaseDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "enableBaseDeposit")
}

// EnableBaseDeposit is a paid mutator transaction binding the contract method 0x1184d8be.
//
// Solidity: function enableBaseDeposit() returns()
func (_Dodo *DodoSession) EnableBaseDeposit() (*types.Transaction, error) {
	return _Dodo.Contract.EnableBaseDeposit(&_Dodo.TransactOpts)
}

// EnableBaseDeposit is a paid mutator transaction binding the contract method 0x1184d8be.
//
// Solidity: function enableBaseDeposit() returns()
func (_Dodo *DodoTransactorSession) EnableBaseDeposit() (*types.Transaction, error) {
	return _Dodo.Contract.EnableBaseDeposit(&_Dodo.TransactOpts)
}

// EnableQuoteDeposit is a paid mutator transaction binding the contract method 0x36ac41a8.
//
// Solidity: function enableQuoteDeposit() returns()
func (_Dodo *DodoTransactor) EnableQuoteDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "enableQuoteDeposit")
}

// EnableQuoteDeposit is a paid mutator transaction binding the contract method 0x36ac41a8.
//
// Solidity: function enableQuoteDeposit() returns()
func (_Dodo *DodoSession) EnableQuoteDeposit() (*types.Transaction, error) {
	return _Dodo.Contract.EnableQuoteDeposit(&_Dodo.TransactOpts)
}

// EnableQuoteDeposit is a paid mutator transaction binding the contract method 0x36ac41a8.
//
// Solidity: function enableQuoteDeposit() returns()
func (_Dodo *DodoTransactorSession) EnableQuoteDeposit() (*types.Transaction, error) {
	return _Dodo.Contract.EnableQuoteDeposit(&_Dodo.TransactOpts)
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Dodo *DodoTransactor) EnableTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "enableTrading")
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Dodo *DodoSession) EnableTrading() (*types.Transaction, error) {
	return _Dodo.Contract.EnableTrading(&_Dodo.TransactOpts)
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Dodo *DodoTransactorSession) EnableTrading() (*types.Transaction, error) {
	return _Dodo.Contract.EnableTrading(&_Dodo.TransactOpts)
}

// FinalSettlement is a paid mutator transaction binding the contract method 0x648a4fac.
//
// Solidity: function finalSettlement() returns()
func (_Dodo *DodoTransactor) FinalSettlement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "finalSettlement")
}

// FinalSettlement is a paid mutator transaction binding the contract method 0x648a4fac.
//
// Solidity: function finalSettlement() returns()
func (_Dodo *DodoSession) FinalSettlement() (*types.Transaction, error) {
	return _Dodo.Contract.FinalSettlement(&_Dodo.TransactOpts)
}

// FinalSettlement is a paid mutator transaction binding the contract method 0x648a4fac.
//
// Solidity: function finalSettlement() returns()
func (_Dodo *DodoTransactorSession) FinalSettlement() (*types.Transaction, error) {
	return _Dodo.Contract.FinalSettlement(&_Dodo.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xf00f9a09.
//
// Solidity: function init(address owner, address supervisor, address maintainer, address baseToken, address quoteToken, address oracle, uint256 lpFeeRate, uint256 mtFeeRate, uint256 k, uint256 gasPriceLimit) returns()
func (_Dodo *DodoTransactor) Init(opts *bind.TransactOpts, owner common.Address, supervisor common.Address, maintainer common.Address, baseToken common.Address, quoteToken common.Address, oracle common.Address, lpFeeRate *big.Int, mtFeeRate *big.Int, k *big.Int, gasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "init", owner, supervisor, maintainer, baseToken, quoteToken, oracle, lpFeeRate, mtFeeRate, k, gasPriceLimit)
}

// Init is a paid mutator transaction binding the contract method 0xf00f9a09.
//
// Solidity: function init(address owner, address supervisor, address maintainer, address baseToken, address quoteToken, address oracle, uint256 lpFeeRate, uint256 mtFeeRate, uint256 k, uint256 gasPriceLimit) returns()
func (_Dodo *DodoSession) Init(owner common.Address, supervisor common.Address, maintainer common.Address, baseToken common.Address, quoteToken common.Address, oracle common.Address, lpFeeRate *big.Int, mtFeeRate *big.Int, k *big.Int, gasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.Init(&_Dodo.TransactOpts, owner, supervisor, maintainer, baseToken, quoteToken, oracle, lpFeeRate, mtFeeRate, k, gasPriceLimit)
}

// Init is a paid mutator transaction binding the contract method 0xf00f9a09.
//
// Solidity: function init(address owner, address supervisor, address maintainer, address baseToken, address quoteToken, address oracle, uint256 lpFeeRate, uint256 mtFeeRate, uint256 k, uint256 gasPriceLimit) returns()
func (_Dodo *DodoTransactorSession) Init(owner common.Address, supervisor common.Address, maintainer common.Address, baseToken common.Address, quoteToken common.Address, oracle common.Address, lpFeeRate *big.Int, mtFeeRate *big.Int, k *big.Int, gasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.Init(&_Dodo.TransactOpts, owner, supervisor, maintainer, baseToken, quoteToken, oracle, lpFeeRate, mtFeeRate, k, gasPriceLimit)
}

// Retrieve is a paid mutator transaction binding the contract method 0xc3a2a665.
//
// Solidity: function retrieve(address token, uint256 amount) returns()
func (_Dodo *DodoTransactor) Retrieve(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "retrieve", token, amount)
}

// Retrieve is a paid mutator transaction binding the contract method 0xc3a2a665.
//
// Solidity: function retrieve(address token, uint256 amount) returns()
func (_Dodo *DodoSession) Retrieve(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.Retrieve(&_Dodo.TransactOpts, token, amount)
}

// Retrieve is a paid mutator transaction binding the contract method 0xc3a2a665.
//
// Solidity: function retrieve(address token, uint256 amount) returns()
func (_Dodo *DodoTransactorSession) Retrieve(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.Retrieve(&_Dodo.TransactOpts, token, amount)
}

// SellBaseToken is a paid mutator transaction binding the contract method 0x8dae7333.
//
// Solidity: function sellBaseToken(uint256 amount, uint256 minReceiveQuote, bytes data) returns(uint256)
func (_Dodo *DodoTransactor) SellBaseToken(opts *bind.TransactOpts, amount *big.Int, minReceiveQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "sellBaseToken", amount, minReceiveQuote, data)
}

// SellBaseToken is a paid mutator transaction binding the contract method 0x8dae7333.
//
// Solidity: function sellBaseToken(uint256 amount, uint256 minReceiveQuote, bytes data) returns(uint256)
func (_Dodo *DodoSession) SellBaseToken(amount *big.Int, minReceiveQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Dodo.Contract.SellBaseToken(&_Dodo.TransactOpts, amount, minReceiveQuote, data)
}

// SellBaseToken is a paid mutator transaction binding the contract method 0x8dae7333.
//
// Solidity: function sellBaseToken(uint256 amount, uint256 minReceiveQuote, bytes data) returns(uint256)
func (_Dodo *DodoTransactorSession) SellBaseToken(amount *big.Int, minReceiveQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Dodo.Contract.SellBaseToken(&_Dodo.TransactOpts, amount, minReceiveQuote, data)
}

// SetGasPriceLimit is a paid mutator transaction binding the contract method 0x09231602.
//
// Solidity: function setGasPriceLimit(uint256 newGasPriceLimit) returns()
func (_Dodo *DodoTransactor) SetGasPriceLimit(opts *bind.TransactOpts, newGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "setGasPriceLimit", newGasPriceLimit)
}

// SetGasPriceLimit is a paid mutator transaction binding the contract method 0x09231602.
//
// Solidity: function setGasPriceLimit(uint256 newGasPriceLimit) returns()
func (_Dodo *DodoSession) SetGasPriceLimit(newGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.SetGasPriceLimit(&_Dodo.TransactOpts, newGasPriceLimit)
}

// SetGasPriceLimit is a paid mutator transaction binding the contract method 0x09231602.
//
// Solidity: function setGasPriceLimit(uint256 newGasPriceLimit) returns()
func (_Dodo *DodoTransactorSession) SetGasPriceLimit(newGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.SetGasPriceLimit(&_Dodo.TransactOpts, newGasPriceLimit)
}

// SetK is a paid mutator transaction binding the contract method 0x67de8be9.
//
// Solidity: function setK(uint256 newK) returns()
func (_Dodo *DodoTransactor) SetK(opts *bind.TransactOpts, newK *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "setK", newK)
}

// SetK is a paid mutator transaction binding the contract method 0x67de8be9.
//
// Solidity: function setK(uint256 newK) returns()
func (_Dodo *DodoSession) SetK(newK *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.SetK(&_Dodo.TransactOpts, newK)
}

// SetK is a paid mutator transaction binding the contract method 0x67de8be9.
//
// Solidity: function setK(uint256 newK) returns()
func (_Dodo *DodoTransactorSession) SetK(newK *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.SetK(&_Dodo.TransactOpts, newK)
}

// SetLiquidityProviderFeeRate is a paid mutator transaction binding the contract method 0x5bb7552a.
//
// Solidity: function setLiquidityProviderFeeRate(uint256 newLiquidityPorviderFeeRate) returns()
func (_Dodo *DodoTransactor) SetLiquidityProviderFeeRate(opts *bind.TransactOpts, newLiquidityPorviderFeeRate *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "setLiquidityProviderFeeRate", newLiquidityPorviderFeeRate)
}

// SetLiquidityProviderFeeRate is a paid mutator transaction binding the contract method 0x5bb7552a.
//
// Solidity: function setLiquidityProviderFeeRate(uint256 newLiquidityPorviderFeeRate) returns()
func (_Dodo *DodoSession) SetLiquidityProviderFeeRate(newLiquidityPorviderFeeRate *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.SetLiquidityProviderFeeRate(&_Dodo.TransactOpts, newLiquidityPorviderFeeRate)
}

// SetLiquidityProviderFeeRate is a paid mutator transaction binding the contract method 0x5bb7552a.
//
// Solidity: function setLiquidityProviderFeeRate(uint256 newLiquidityPorviderFeeRate) returns()
func (_Dodo *DodoTransactorSession) SetLiquidityProviderFeeRate(newLiquidityPorviderFeeRate *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.SetLiquidityProviderFeeRate(&_Dodo.TransactOpts, newLiquidityPorviderFeeRate)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address newMaintainer) returns()
func (_Dodo *DodoTransactor) SetMaintainer(opts *bind.TransactOpts, newMaintainer common.Address) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "setMaintainer", newMaintainer)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address newMaintainer) returns()
func (_Dodo *DodoSession) SetMaintainer(newMaintainer common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.SetMaintainer(&_Dodo.TransactOpts, newMaintainer)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address newMaintainer) returns()
func (_Dodo *DodoTransactorSession) SetMaintainer(newMaintainer common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.SetMaintainer(&_Dodo.TransactOpts, newMaintainer)
}

// SetMaintainerFeeRate is a paid mutator transaction binding the contract method 0xf2220416.
//
// Solidity: function setMaintainerFeeRate(uint256 newMaintainerFeeRate) returns()
func (_Dodo *DodoTransactor) SetMaintainerFeeRate(opts *bind.TransactOpts, newMaintainerFeeRate *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "setMaintainerFeeRate", newMaintainerFeeRate)
}

// SetMaintainerFeeRate is a paid mutator transaction binding the contract method 0xf2220416.
//
// Solidity: function setMaintainerFeeRate(uint256 newMaintainerFeeRate) returns()
func (_Dodo *DodoSession) SetMaintainerFeeRate(newMaintainerFeeRate *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.SetMaintainerFeeRate(&_Dodo.TransactOpts, newMaintainerFeeRate)
}

// SetMaintainerFeeRate is a paid mutator transaction binding the contract method 0xf2220416.
//
// Solidity: function setMaintainerFeeRate(uint256 newMaintainerFeeRate) returns()
func (_Dodo *DodoTransactorSession) SetMaintainerFeeRate(newMaintainerFeeRate *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.SetMaintainerFeeRate(&_Dodo.TransactOpts, newMaintainerFeeRate)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address newOracle) returns()
func (_Dodo *DodoTransactor) SetOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "setOracle", newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address newOracle) returns()
func (_Dodo *DodoSession) SetOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.SetOracle(&_Dodo.TransactOpts, newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address newOracle) returns()
func (_Dodo *DodoTransactorSession) SetOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.SetOracle(&_Dodo.TransactOpts, newOracle)
}

// SetSupervisor is a paid mutator transaction binding the contract method 0x9299eb30.
//
// Solidity: function setSupervisor(address newSupervisor) returns()
func (_Dodo *DodoTransactor) SetSupervisor(opts *bind.TransactOpts, newSupervisor common.Address) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "setSupervisor", newSupervisor)
}

// SetSupervisor is a paid mutator transaction binding the contract method 0x9299eb30.
//
// Solidity: function setSupervisor(address newSupervisor) returns()
func (_Dodo *DodoSession) SetSupervisor(newSupervisor common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.SetSupervisor(&_Dodo.TransactOpts, newSupervisor)
}

// SetSupervisor is a paid mutator transaction binding the contract method 0x9299eb30.
//
// Solidity: function setSupervisor(address newSupervisor) returns()
func (_Dodo *DodoTransactorSession) SetSupervisor(newSupervisor common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.SetSupervisor(&_Dodo.TransactOpts, newSupervisor)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dodo *DodoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dodo *DodoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.TransferOwnership(&_Dodo.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dodo *DodoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.TransferOwnership(&_Dodo.TransactOpts, newOwner)
}

// WithdrawAllBase is a paid mutator transaction binding the contract method 0xd47eaa37.
//
// Solidity: function withdrawAllBase() returns(uint256)
func (_Dodo *DodoTransactor) WithdrawAllBase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "withdrawAllBase")
}

// WithdrawAllBase is a paid mutator transaction binding the contract method 0xd47eaa37.
//
// Solidity: function withdrawAllBase() returns(uint256)
func (_Dodo *DodoSession) WithdrawAllBase() (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawAllBase(&_Dodo.TransactOpts)
}

// WithdrawAllBase is a paid mutator transaction binding the contract method 0xd47eaa37.
//
// Solidity: function withdrawAllBase() returns(uint256)
func (_Dodo *DodoTransactorSession) WithdrawAllBase() (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawAllBase(&_Dodo.TransactOpts)
}

// WithdrawAllBaseTo is a paid mutator transaction binding the contract method 0x1e34b9cc.
//
// Solidity: function withdrawAllBaseTo(address to) returns(uint256)
func (_Dodo *DodoTransactor) WithdrawAllBaseTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "withdrawAllBaseTo", to)
}

// WithdrawAllBaseTo is a paid mutator transaction binding the contract method 0x1e34b9cc.
//
// Solidity: function withdrawAllBaseTo(address to) returns(uint256)
func (_Dodo *DodoSession) WithdrawAllBaseTo(to common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawAllBaseTo(&_Dodo.TransactOpts, to)
}

// WithdrawAllBaseTo is a paid mutator transaction binding the contract method 0x1e34b9cc.
//
// Solidity: function withdrawAllBaseTo(address to) returns(uint256)
func (_Dodo *DodoTransactorSession) WithdrawAllBaseTo(to common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawAllBaseTo(&_Dodo.TransactOpts, to)
}

// WithdrawAllQuote is a paid mutator transaction binding the contract method 0xc59203af.
//
// Solidity: function withdrawAllQuote() returns(uint256)
func (_Dodo *DodoTransactor) WithdrawAllQuote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "withdrawAllQuote")
}

// WithdrawAllQuote is a paid mutator transaction binding the contract method 0xc59203af.
//
// Solidity: function withdrawAllQuote() returns(uint256)
func (_Dodo *DodoSession) WithdrawAllQuote() (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawAllQuote(&_Dodo.TransactOpts)
}

// WithdrawAllQuote is a paid mutator transaction binding the contract method 0xc59203af.
//
// Solidity: function withdrawAllQuote() returns(uint256)
func (_Dodo *DodoTransactorSession) WithdrawAllQuote() (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawAllQuote(&_Dodo.TransactOpts)
}

// WithdrawAllQuoteTo is a paid mutator transaction binding the contract method 0x04512dc4.
//
// Solidity: function withdrawAllQuoteTo(address to) returns(uint256)
func (_Dodo *DodoTransactor) WithdrawAllQuoteTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "withdrawAllQuoteTo", to)
}

// WithdrawAllQuoteTo is a paid mutator transaction binding the contract method 0x04512dc4.
//
// Solidity: function withdrawAllQuoteTo(address to) returns(uint256)
func (_Dodo *DodoSession) WithdrawAllQuoteTo(to common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawAllQuoteTo(&_Dodo.TransactOpts, to)
}

// WithdrawAllQuoteTo is a paid mutator transaction binding the contract method 0x04512dc4.
//
// Solidity: function withdrawAllQuoteTo(address to) returns(uint256)
func (_Dodo *DodoTransactorSession) WithdrawAllQuoteTo(to common.Address) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawAllQuoteTo(&_Dodo.TransactOpts, to)
}

// WithdrawBase is a paid mutator transaction binding the contract method 0xf98bea15.
//
// Solidity: function withdrawBase(uint256 amount) returns(uint256)
func (_Dodo *DodoTransactor) WithdrawBase(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "withdrawBase", amount)
}

// WithdrawBase is a paid mutator transaction binding the contract method 0xf98bea15.
//
// Solidity: function withdrawBase(uint256 amount) returns(uint256)
func (_Dodo *DodoSession) WithdrawBase(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawBase(&_Dodo.TransactOpts, amount)
}

// WithdrawBase is a paid mutator transaction binding the contract method 0xf98bea15.
//
// Solidity: function withdrawBase(uint256 amount) returns(uint256)
func (_Dodo *DodoTransactorSession) WithdrawBase(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawBase(&_Dodo.TransactOpts, amount)
}

// WithdrawBaseTo is a paid mutator transaction binding the contract method 0x947cf92b.
//
// Solidity: function withdrawBaseTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoTransactor) WithdrawBaseTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "withdrawBaseTo", to, amount)
}

// WithdrawBaseTo is a paid mutator transaction binding the contract method 0x947cf92b.
//
// Solidity: function withdrawBaseTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoSession) WithdrawBaseTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawBaseTo(&_Dodo.TransactOpts, to, amount)
}

// WithdrawBaseTo is a paid mutator transaction binding the contract method 0x947cf92b.
//
// Solidity: function withdrawBaseTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoTransactorSession) WithdrawBaseTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawBaseTo(&_Dodo.TransactOpts, to, amount)
}

// WithdrawQuote is a paid mutator transaction binding the contract method 0xc0a5f6ff.
//
// Solidity: function withdrawQuote(uint256 amount) returns(uint256)
func (_Dodo *DodoTransactor) WithdrawQuote(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "withdrawQuote", amount)
}

// WithdrawQuote is a paid mutator transaction binding the contract method 0xc0a5f6ff.
//
// Solidity: function withdrawQuote(uint256 amount) returns(uint256)
func (_Dodo *DodoSession) WithdrawQuote(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawQuote(&_Dodo.TransactOpts, amount)
}

// WithdrawQuote is a paid mutator transaction binding the contract method 0xc0a5f6ff.
//
// Solidity: function withdrawQuote(uint256 amount) returns(uint256)
func (_Dodo *DodoTransactorSession) WithdrawQuote(amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawQuote(&_Dodo.TransactOpts, amount)
}

// WithdrawQuoteTo is a paid mutator transaction binding the contract method 0x108db744.
//
// Solidity: function withdrawQuoteTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoTransactor) WithdrawQuoteTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.contract.Transact(opts, "withdrawQuoteTo", to, amount)
}

// WithdrawQuoteTo is a paid mutator transaction binding the contract method 0x108db744.
//
// Solidity: function withdrawQuoteTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoSession) WithdrawQuoteTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawQuoteTo(&_Dodo.TransactOpts, to, amount)
}

// WithdrawQuoteTo is a paid mutator transaction binding the contract method 0x108db744.
//
// Solidity: function withdrawQuoteTo(address to, uint256 amount) returns(uint256)
func (_Dodo *DodoTransactorSession) WithdrawQuoteTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dodo.Contract.WithdrawQuoteTo(&_Dodo.TransactOpts, to, amount)
}

// DodoBuyBaseTokenIterator is returned from FilterBuyBaseToken and is used to iterate over the raw logs and unpacked data for BuyBaseToken events raised by the Dodo contract.
type DodoBuyBaseTokenIterator struct {
	Event *DodoBuyBaseToken // Event containing the contract specifics and raw log

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
func (it *DodoBuyBaseTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoBuyBaseToken)
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
		it.Event = new(DodoBuyBaseToken)
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
func (it *DodoBuyBaseTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoBuyBaseTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoBuyBaseToken represents a BuyBaseToken event raised by the Dodo contract.
type DodoBuyBaseToken struct {
	Buyer       common.Address
	ReceiveBase *big.Int
	PayQuote    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBuyBaseToken is a free log retrieval operation binding the contract event 0xe93ad76094f247c0dafc1c61adc2187de1ac2738f7a3b49cb20b2263420251a3.
//
// Solidity: event BuyBaseToken(address indexed buyer, uint256 receiveBase, uint256 payQuote)
func (_Dodo *DodoFilterer) FilterBuyBaseToken(opts *bind.FilterOpts, buyer []common.Address) (*DodoBuyBaseTokenIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "BuyBaseToken", buyerRule)
	if err != nil {
		return nil, err
	}
	return &DodoBuyBaseTokenIterator{contract: _Dodo.contract, event: "BuyBaseToken", logs: logs, sub: sub}, nil
}

// WatchBuyBaseToken is a free log subscription operation binding the contract event 0xe93ad76094f247c0dafc1c61adc2187de1ac2738f7a3b49cb20b2263420251a3.
//
// Solidity: event BuyBaseToken(address indexed buyer, uint256 receiveBase, uint256 payQuote)
func (_Dodo *DodoFilterer) WatchBuyBaseToken(opts *bind.WatchOpts, sink chan<- *DodoBuyBaseToken, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "BuyBaseToken", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoBuyBaseToken)
				if err := _Dodo.contract.UnpackLog(event, "BuyBaseToken", log); err != nil {
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

// ParseBuyBaseToken is a log parse operation binding the contract event 0xe93ad76094f247c0dafc1c61adc2187de1ac2738f7a3b49cb20b2263420251a3.
//
// Solidity: event BuyBaseToken(address indexed buyer, uint256 receiveBase, uint256 payQuote)
func (_Dodo *DodoFilterer) ParseBuyBaseToken(log types.Log) (*DodoBuyBaseToken, error) {
	event := new(DodoBuyBaseToken)
	if err := _Dodo.contract.UnpackLog(event, "BuyBaseToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoChargeMaintainerFeeIterator is returned from FilterChargeMaintainerFee and is used to iterate over the raw logs and unpacked data for ChargeMaintainerFee events raised by the Dodo contract.
type DodoChargeMaintainerFeeIterator struct {
	Event *DodoChargeMaintainerFee // Event containing the contract specifics and raw log

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
func (it *DodoChargeMaintainerFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoChargeMaintainerFee)
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
		it.Event = new(DodoChargeMaintainerFee)
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
func (it *DodoChargeMaintainerFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoChargeMaintainerFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoChargeMaintainerFee represents a ChargeMaintainerFee event raised by the Dodo contract.
type DodoChargeMaintainerFee struct {
	Maintainer  common.Address
	IsBaseToken bool
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChargeMaintainerFee is a free log retrieval operation binding the contract event 0xe4fed5362e2669c70e5da5a18942d1e617d8917f6adc0164d9668bd3a6d0cebe.
//
// Solidity: event ChargeMaintainerFee(address indexed maintainer, bool isBaseToken, uint256 amount)
func (_Dodo *DodoFilterer) FilterChargeMaintainerFee(opts *bind.FilterOpts, maintainer []common.Address) (*DodoChargeMaintainerFeeIterator, error) {

	var maintainerRule []interface{}
	for _, maintainerItem := range maintainer {
		maintainerRule = append(maintainerRule, maintainerItem)
	}

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "ChargeMaintainerFee", maintainerRule)
	if err != nil {
		return nil, err
	}
	return &DodoChargeMaintainerFeeIterator{contract: _Dodo.contract, event: "ChargeMaintainerFee", logs: logs, sub: sub}, nil
}

// WatchChargeMaintainerFee is a free log subscription operation binding the contract event 0xe4fed5362e2669c70e5da5a18942d1e617d8917f6adc0164d9668bd3a6d0cebe.
//
// Solidity: event ChargeMaintainerFee(address indexed maintainer, bool isBaseToken, uint256 amount)
func (_Dodo *DodoFilterer) WatchChargeMaintainerFee(opts *bind.WatchOpts, sink chan<- *DodoChargeMaintainerFee, maintainer []common.Address) (event.Subscription, error) {

	var maintainerRule []interface{}
	for _, maintainerItem := range maintainer {
		maintainerRule = append(maintainerRule, maintainerItem)
	}

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "ChargeMaintainerFee", maintainerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoChargeMaintainerFee)
				if err := _Dodo.contract.UnpackLog(event, "ChargeMaintainerFee", log); err != nil {
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

// ParseChargeMaintainerFee is a log parse operation binding the contract event 0xe4fed5362e2669c70e5da5a18942d1e617d8917f6adc0164d9668bd3a6d0cebe.
//
// Solidity: event ChargeMaintainerFee(address indexed maintainer, bool isBaseToken, uint256 amount)
func (_Dodo *DodoFilterer) ParseChargeMaintainerFee(log types.Log) (*DodoChargeMaintainerFee, error) {
	event := new(DodoChargeMaintainerFee)
	if err := _Dodo.contract.UnpackLog(event, "ChargeMaintainerFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoChargePenaltyIterator is returned from FilterChargePenalty and is used to iterate over the raw logs and unpacked data for ChargePenalty events raised by the Dodo contract.
type DodoChargePenaltyIterator struct {
	Event *DodoChargePenalty // Event containing the contract specifics and raw log

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
func (it *DodoChargePenaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoChargePenalty)
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
		it.Event = new(DodoChargePenalty)
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
func (it *DodoChargePenaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoChargePenaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoChargePenalty represents a ChargePenalty event raised by the Dodo contract.
type DodoChargePenalty struct {
	Payer       common.Address
	IsBaseToken bool
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChargePenalty is a free log retrieval operation binding the contract event 0x581f351e2bdb9fa9021bb2a24def989f06ac236f8a92aac14bcbc618ddf3826a.
//
// Solidity: event ChargePenalty(address indexed payer, bool isBaseToken, uint256 amount)
func (_Dodo *DodoFilterer) FilterChargePenalty(opts *bind.FilterOpts, payer []common.Address) (*DodoChargePenaltyIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "ChargePenalty", payerRule)
	if err != nil {
		return nil, err
	}
	return &DodoChargePenaltyIterator{contract: _Dodo.contract, event: "ChargePenalty", logs: logs, sub: sub}, nil
}

// WatchChargePenalty is a free log subscription operation binding the contract event 0x581f351e2bdb9fa9021bb2a24def989f06ac236f8a92aac14bcbc618ddf3826a.
//
// Solidity: event ChargePenalty(address indexed payer, bool isBaseToken, uint256 amount)
func (_Dodo *DodoFilterer) WatchChargePenalty(opts *bind.WatchOpts, sink chan<- *DodoChargePenalty, payer []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "ChargePenalty", payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoChargePenalty)
				if err := _Dodo.contract.UnpackLog(event, "ChargePenalty", log); err != nil {
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

// ParseChargePenalty is a log parse operation binding the contract event 0x581f351e2bdb9fa9021bb2a24def989f06ac236f8a92aac14bcbc618ddf3826a.
//
// Solidity: event ChargePenalty(address indexed payer, bool isBaseToken, uint256 amount)
func (_Dodo *DodoFilterer) ParseChargePenalty(log types.Log) (*DodoChargePenalty, error) {
	event := new(DodoChargePenalty)
	if err := _Dodo.contract.UnpackLog(event, "ChargePenalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoClaimAssetsIterator is returned from FilterClaimAssets and is used to iterate over the raw logs and unpacked data for ClaimAssets events raised by the Dodo contract.
type DodoClaimAssetsIterator struct {
	Event *DodoClaimAssets // Event containing the contract specifics and raw log

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
func (it *DodoClaimAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoClaimAssets)
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
		it.Event = new(DodoClaimAssets)
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
func (it *DodoClaimAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoClaimAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoClaimAssets represents a ClaimAssets event raised by the Dodo contract.
type DodoClaimAssets struct {
	User             common.Address
	BaseTokenAmount  *big.Int
	QuoteTokenAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimAssets is a free log retrieval operation binding the contract event 0xbe5f7fe66d16c6a87bb5b8b08a96634fe4f1c2bac9e5e413efe41a782d4d0c43.
//
// Solidity: event ClaimAssets(address indexed user, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Dodo *DodoFilterer) FilterClaimAssets(opts *bind.FilterOpts, user []common.Address) (*DodoClaimAssetsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "ClaimAssets", userRule)
	if err != nil {
		return nil, err
	}
	return &DodoClaimAssetsIterator{contract: _Dodo.contract, event: "ClaimAssets", logs: logs, sub: sub}, nil
}

// WatchClaimAssets is a free log subscription operation binding the contract event 0xbe5f7fe66d16c6a87bb5b8b08a96634fe4f1c2bac9e5e413efe41a782d4d0c43.
//
// Solidity: event ClaimAssets(address indexed user, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Dodo *DodoFilterer) WatchClaimAssets(opts *bind.WatchOpts, sink chan<- *DodoClaimAssets, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "ClaimAssets", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoClaimAssets)
				if err := _Dodo.contract.UnpackLog(event, "ClaimAssets", log); err != nil {
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

// ParseClaimAssets is a log parse operation binding the contract event 0xbe5f7fe66d16c6a87bb5b8b08a96634fe4f1c2bac9e5e413efe41a782d4d0c43.
//
// Solidity: event ClaimAssets(address indexed user, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Dodo *DodoFilterer) ParseClaimAssets(log types.Log) (*DodoClaimAssets, error) {
	event := new(DodoClaimAssets)
	if err := _Dodo.contract.UnpackLog(event, "ClaimAssets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Dodo contract.
type DodoDepositIterator struct {
	Event *DodoDeposit // Event containing the contract specifics and raw log

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
func (it *DodoDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoDeposit)
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
		it.Event = new(DodoDeposit)
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
func (it *DodoDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoDeposit represents a Deposit event raised by the Dodo contract.
type DodoDeposit struct {
	Payer         common.Address
	Receiver      common.Address
	IsBaseToken   bool
	Amount        *big.Int
	LpTokenAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x18081cde2fa64894914e1080b98cca17bb6d1acf633e57f6e26ebdb945ad830b.
//
// Solidity: event Deposit(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Dodo *DodoFilterer) FilterDeposit(opts *bind.FilterOpts, payer []common.Address, receiver []common.Address) (*DodoDepositIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "Deposit", payerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &DodoDepositIterator{contract: _Dodo.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x18081cde2fa64894914e1080b98cca17bb6d1acf633e57f6e26ebdb945ad830b.
//
// Solidity: event Deposit(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Dodo *DodoFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *DodoDeposit, payer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "Deposit", payerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoDeposit)
				if err := _Dodo.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x18081cde2fa64894914e1080b98cca17bb6d1acf633e57f6e26ebdb945ad830b.
//
// Solidity: event Deposit(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Dodo *DodoFilterer) ParseDeposit(log types.Log) (*DodoDeposit, error) {
	event := new(DodoDeposit)
	if err := _Dodo.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoDonateIterator is returned from FilterDonate and is used to iterate over the raw logs and unpacked data for Donate events raised by the Dodo contract.
type DodoDonateIterator struct {
	Event *DodoDonate // Event containing the contract specifics and raw log

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
func (it *DodoDonateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoDonate)
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
		it.Event = new(DodoDonate)
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
func (it *DodoDonateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoDonateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoDonate represents a Donate event raised by the Dodo contract.
type DodoDonate struct {
	Amount      *big.Int
	IsBaseToken bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDonate is a free log retrieval operation binding the contract event 0xa259c93818139b6bc90fb80e8feb75122b42edaae49560f81392cf4e1946726e.
//
// Solidity: event Donate(uint256 amount, bool isBaseToken)
func (_Dodo *DodoFilterer) FilterDonate(opts *bind.FilterOpts) (*DodoDonateIterator, error) {

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "Donate")
	if err != nil {
		return nil, err
	}
	return &DodoDonateIterator{contract: _Dodo.contract, event: "Donate", logs: logs, sub: sub}, nil
}

// WatchDonate is a free log subscription operation binding the contract event 0xa259c93818139b6bc90fb80e8feb75122b42edaae49560f81392cf4e1946726e.
//
// Solidity: event Donate(uint256 amount, bool isBaseToken)
func (_Dodo *DodoFilterer) WatchDonate(opts *bind.WatchOpts, sink chan<- *DodoDonate) (event.Subscription, error) {

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "Donate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoDonate)
				if err := _Dodo.contract.UnpackLog(event, "Donate", log); err != nil {
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

// ParseDonate is a log parse operation binding the contract event 0xa259c93818139b6bc90fb80e8feb75122b42edaae49560f81392cf4e1946726e.
//
// Solidity: event Donate(uint256 amount, bool isBaseToken)
func (_Dodo *DodoFilterer) ParseDonate(log types.Log) (*DodoDonate, error) {
	event := new(DodoDonate)
	if err := _Dodo.contract.UnpackLog(event, "Donate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoOwnershipTransferPreparedIterator is returned from FilterOwnershipTransferPrepared and is used to iterate over the raw logs and unpacked data for OwnershipTransferPrepared events raised by the Dodo contract.
type DodoOwnershipTransferPreparedIterator struct {
	Event *DodoOwnershipTransferPrepared // Event containing the contract specifics and raw log

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
func (it *DodoOwnershipTransferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoOwnershipTransferPrepared)
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
		it.Event = new(DodoOwnershipTransferPrepared)
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
func (it *DodoOwnershipTransferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoOwnershipTransferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoOwnershipTransferPrepared represents a OwnershipTransferPrepared event raised by the Dodo contract.
type DodoOwnershipTransferPrepared struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferPrepared is a free log retrieval operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_Dodo *DodoFilterer) FilterOwnershipTransferPrepared(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DodoOwnershipTransferPreparedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DodoOwnershipTransferPreparedIterator{contract: _Dodo.contract, event: "OwnershipTransferPrepared", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferPrepared is a free log subscription operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_Dodo *DodoFilterer) WatchOwnershipTransferPrepared(opts *bind.WatchOpts, sink chan<- *DodoOwnershipTransferPrepared, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoOwnershipTransferPrepared)
				if err := _Dodo.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
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

// ParseOwnershipTransferPrepared is a log parse operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_Dodo *DodoFilterer) ParseOwnershipTransferPrepared(log types.Log) (*DodoOwnershipTransferPrepared, error) {
	event := new(DodoOwnershipTransferPrepared)
	if err := _Dodo.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dodo contract.
type DodoOwnershipTransferredIterator struct {
	Event *DodoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DodoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoOwnershipTransferred)
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
		it.Event = new(DodoOwnershipTransferred)
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
func (it *DodoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoOwnershipTransferred represents a OwnershipTransferred event raised by the Dodo contract.
type DodoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dodo *DodoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DodoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DodoOwnershipTransferredIterator{contract: _Dodo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dodo *DodoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DodoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoOwnershipTransferred)
				if err := _Dodo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dodo *DodoFilterer) ParseOwnershipTransferred(log types.Log) (*DodoOwnershipTransferred, error) {
	event := new(DodoOwnershipTransferred)
	if err := _Dodo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoSellBaseTokenIterator is returned from FilterSellBaseToken and is used to iterate over the raw logs and unpacked data for SellBaseToken events raised by the Dodo contract.
type DodoSellBaseTokenIterator struct {
	Event *DodoSellBaseToken // Event containing the contract specifics and raw log

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
func (it *DodoSellBaseTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoSellBaseToken)
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
		it.Event = new(DodoSellBaseToken)
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
func (it *DodoSellBaseTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoSellBaseTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoSellBaseToken represents a SellBaseToken event raised by the Dodo contract.
type DodoSellBaseToken struct {
	Seller       common.Address
	PayBase      *big.Int
	ReceiveQuote *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSellBaseToken is a free log retrieval operation binding the contract event 0xd8648b6ac54162763c86fd54bf2005af8ecd2f9cb273a5775921fd7f91e17b2d.
//
// Solidity: event SellBaseToken(address indexed seller, uint256 payBase, uint256 receiveQuote)
func (_Dodo *DodoFilterer) FilterSellBaseToken(opts *bind.FilterOpts, seller []common.Address) (*DodoSellBaseTokenIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "SellBaseToken", sellerRule)
	if err != nil {
		return nil, err
	}
	return &DodoSellBaseTokenIterator{contract: _Dodo.contract, event: "SellBaseToken", logs: logs, sub: sub}, nil
}

// WatchSellBaseToken is a free log subscription operation binding the contract event 0xd8648b6ac54162763c86fd54bf2005af8ecd2f9cb273a5775921fd7f91e17b2d.
//
// Solidity: event SellBaseToken(address indexed seller, uint256 payBase, uint256 receiveQuote)
func (_Dodo *DodoFilterer) WatchSellBaseToken(opts *bind.WatchOpts, sink chan<- *DodoSellBaseToken, seller []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "SellBaseToken", sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoSellBaseToken)
				if err := _Dodo.contract.UnpackLog(event, "SellBaseToken", log); err != nil {
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

// ParseSellBaseToken is a log parse operation binding the contract event 0xd8648b6ac54162763c86fd54bf2005af8ecd2f9cb273a5775921fd7f91e17b2d.
//
// Solidity: event SellBaseToken(address indexed seller, uint256 payBase, uint256 receiveQuote)
func (_Dodo *DodoFilterer) ParseSellBaseToken(log types.Log) (*DodoSellBaseToken, error) {
	event := new(DodoSellBaseToken)
	if err := _Dodo.contract.UnpackLog(event, "SellBaseToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoUpdateGasPriceLimitIterator is returned from FilterUpdateGasPriceLimit and is used to iterate over the raw logs and unpacked data for UpdateGasPriceLimit events raised by the Dodo contract.
type DodoUpdateGasPriceLimitIterator struct {
	Event *DodoUpdateGasPriceLimit // Event containing the contract specifics and raw log

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
func (it *DodoUpdateGasPriceLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoUpdateGasPriceLimit)
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
		it.Event = new(DodoUpdateGasPriceLimit)
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
func (it *DodoUpdateGasPriceLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoUpdateGasPriceLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoUpdateGasPriceLimit represents a UpdateGasPriceLimit event raised by the Dodo contract.
type DodoUpdateGasPriceLimit struct {
	OldGasPriceLimit *big.Int
	NewGasPriceLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateGasPriceLimit is a free log retrieval operation binding the contract event 0x808f99cfd15f1be6019f93dc76c81d5bd82e0b3e0b3d23a54f5a2e647a6cc3cc.
//
// Solidity: event UpdateGasPriceLimit(uint256 oldGasPriceLimit, uint256 newGasPriceLimit)
func (_Dodo *DodoFilterer) FilterUpdateGasPriceLimit(opts *bind.FilterOpts) (*DodoUpdateGasPriceLimitIterator, error) {

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "UpdateGasPriceLimit")
	if err != nil {
		return nil, err
	}
	return &DodoUpdateGasPriceLimitIterator{contract: _Dodo.contract, event: "UpdateGasPriceLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateGasPriceLimit is a free log subscription operation binding the contract event 0x808f99cfd15f1be6019f93dc76c81d5bd82e0b3e0b3d23a54f5a2e647a6cc3cc.
//
// Solidity: event UpdateGasPriceLimit(uint256 oldGasPriceLimit, uint256 newGasPriceLimit)
func (_Dodo *DodoFilterer) WatchUpdateGasPriceLimit(opts *bind.WatchOpts, sink chan<- *DodoUpdateGasPriceLimit) (event.Subscription, error) {

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "UpdateGasPriceLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoUpdateGasPriceLimit)
				if err := _Dodo.contract.UnpackLog(event, "UpdateGasPriceLimit", log); err != nil {
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

// ParseUpdateGasPriceLimit is a log parse operation binding the contract event 0x808f99cfd15f1be6019f93dc76c81d5bd82e0b3e0b3d23a54f5a2e647a6cc3cc.
//
// Solidity: event UpdateGasPriceLimit(uint256 oldGasPriceLimit, uint256 newGasPriceLimit)
func (_Dodo *DodoFilterer) ParseUpdateGasPriceLimit(log types.Log) (*DodoUpdateGasPriceLimit, error) {
	event := new(DodoUpdateGasPriceLimit)
	if err := _Dodo.contract.UnpackLog(event, "UpdateGasPriceLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoUpdateKIterator is returned from FilterUpdateK and is used to iterate over the raw logs and unpacked data for UpdateK events raised by the Dodo contract.
type DodoUpdateKIterator struct {
	Event *DodoUpdateK // Event containing the contract specifics and raw log

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
func (it *DodoUpdateKIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoUpdateK)
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
		it.Event = new(DodoUpdateK)
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
func (it *DodoUpdateKIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoUpdateKIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoUpdateK represents a UpdateK event raised by the Dodo contract.
type DodoUpdateK struct {
	OldK *big.Int
	NewK *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateK is a free log retrieval operation binding the contract event 0x023a40bebf7ac113f81c3d628073246cf9e0bc49980a9d6a9531498ce9e3dd1c.
//
// Solidity: event UpdateK(uint256 oldK, uint256 newK)
func (_Dodo *DodoFilterer) FilterUpdateK(opts *bind.FilterOpts) (*DodoUpdateKIterator, error) {

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "UpdateK")
	if err != nil {
		return nil, err
	}
	return &DodoUpdateKIterator{contract: _Dodo.contract, event: "UpdateK", logs: logs, sub: sub}, nil
}

// WatchUpdateK is a free log subscription operation binding the contract event 0x023a40bebf7ac113f81c3d628073246cf9e0bc49980a9d6a9531498ce9e3dd1c.
//
// Solidity: event UpdateK(uint256 oldK, uint256 newK)
func (_Dodo *DodoFilterer) WatchUpdateK(opts *bind.WatchOpts, sink chan<- *DodoUpdateK) (event.Subscription, error) {

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "UpdateK")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoUpdateK)
				if err := _Dodo.contract.UnpackLog(event, "UpdateK", log); err != nil {
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

// ParseUpdateK is a log parse operation binding the contract event 0x023a40bebf7ac113f81c3d628073246cf9e0bc49980a9d6a9531498ce9e3dd1c.
//
// Solidity: event UpdateK(uint256 oldK, uint256 newK)
func (_Dodo *DodoFilterer) ParseUpdateK(log types.Log) (*DodoUpdateK, error) {
	event := new(DodoUpdateK)
	if err := _Dodo.contract.UnpackLog(event, "UpdateK", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoUpdateLiquidityProviderFeeRateIterator is returned from FilterUpdateLiquidityProviderFeeRate and is used to iterate over the raw logs and unpacked data for UpdateLiquidityProviderFeeRate events raised by the Dodo contract.
type DodoUpdateLiquidityProviderFeeRateIterator struct {
	Event *DodoUpdateLiquidityProviderFeeRate // Event containing the contract specifics and raw log

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
func (it *DodoUpdateLiquidityProviderFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoUpdateLiquidityProviderFeeRate)
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
		it.Event = new(DodoUpdateLiquidityProviderFeeRate)
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
func (it *DodoUpdateLiquidityProviderFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoUpdateLiquidityProviderFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoUpdateLiquidityProviderFeeRate represents a UpdateLiquidityProviderFeeRate event raised by the Dodo contract.
type DodoUpdateLiquidityProviderFeeRate struct {
	OldLiquidityProviderFeeRate *big.Int
	NewLiquidityProviderFeeRate *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterUpdateLiquidityProviderFeeRate is a free log retrieval operation binding the contract event 0x3ce6ea91adda496b7a0546fa6558e5b52c3a509de6015820efb00ca4020e0a07.
//
// Solidity: event UpdateLiquidityProviderFeeRate(uint256 oldLiquidityProviderFeeRate, uint256 newLiquidityProviderFeeRate)
func (_Dodo *DodoFilterer) FilterUpdateLiquidityProviderFeeRate(opts *bind.FilterOpts) (*DodoUpdateLiquidityProviderFeeRateIterator, error) {

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "UpdateLiquidityProviderFeeRate")
	if err != nil {
		return nil, err
	}
	return &DodoUpdateLiquidityProviderFeeRateIterator{contract: _Dodo.contract, event: "UpdateLiquidityProviderFeeRate", logs: logs, sub: sub}, nil
}

// WatchUpdateLiquidityProviderFeeRate is a free log subscription operation binding the contract event 0x3ce6ea91adda496b7a0546fa6558e5b52c3a509de6015820efb00ca4020e0a07.
//
// Solidity: event UpdateLiquidityProviderFeeRate(uint256 oldLiquidityProviderFeeRate, uint256 newLiquidityProviderFeeRate)
func (_Dodo *DodoFilterer) WatchUpdateLiquidityProviderFeeRate(opts *bind.WatchOpts, sink chan<- *DodoUpdateLiquidityProviderFeeRate) (event.Subscription, error) {

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "UpdateLiquidityProviderFeeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoUpdateLiquidityProviderFeeRate)
				if err := _Dodo.contract.UnpackLog(event, "UpdateLiquidityProviderFeeRate", log); err != nil {
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

// ParseUpdateLiquidityProviderFeeRate is a log parse operation binding the contract event 0x3ce6ea91adda496b7a0546fa6558e5b52c3a509de6015820efb00ca4020e0a07.
//
// Solidity: event UpdateLiquidityProviderFeeRate(uint256 oldLiquidityProviderFeeRate, uint256 newLiquidityProviderFeeRate)
func (_Dodo *DodoFilterer) ParseUpdateLiquidityProviderFeeRate(log types.Log) (*DodoUpdateLiquidityProviderFeeRate, error) {
	event := new(DodoUpdateLiquidityProviderFeeRate)
	if err := _Dodo.contract.UnpackLog(event, "UpdateLiquidityProviderFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoUpdateMaintainerFeeRateIterator is returned from FilterUpdateMaintainerFeeRate and is used to iterate over the raw logs and unpacked data for UpdateMaintainerFeeRate events raised by the Dodo contract.
type DodoUpdateMaintainerFeeRateIterator struct {
	Event *DodoUpdateMaintainerFeeRate // Event containing the contract specifics and raw log

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
func (it *DodoUpdateMaintainerFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoUpdateMaintainerFeeRate)
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
		it.Event = new(DodoUpdateMaintainerFeeRate)
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
func (it *DodoUpdateMaintainerFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoUpdateMaintainerFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoUpdateMaintainerFeeRate represents a UpdateMaintainerFeeRate event raised by the Dodo contract.
type DodoUpdateMaintainerFeeRate struct {
	OldMaintainerFeeRate *big.Int
	NewMaintainerFeeRate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaintainerFeeRate is a free log retrieval operation binding the contract event 0x6b04da3d58e4b37d99652babb3ea2bc25ce94379bfff3059f03d61b26c59e553.
//
// Solidity: event UpdateMaintainerFeeRate(uint256 oldMaintainerFeeRate, uint256 newMaintainerFeeRate)
func (_Dodo *DodoFilterer) FilterUpdateMaintainerFeeRate(opts *bind.FilterOpts) (*DodoUpdateMaintainerFeeRateIterator, error) {

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "UpdateMaintainerFeeRate")
	if err != nil {
		return nil, err
	}
	return &DodoUpdateMaintainerFeeRateIterator{contract: _Dodo.contract, event: "UpdateMaintainerFeeRate", logs: logs, sub: sub}, nil
}

// WatchUpdateMaintainerFeeRate is a free log subscription operation binding the contract event 0x6b04da3d58e4b37d99652babb3ea2bc25ce94379bfff3059f03d61b26c59e553.
//
// Solidity: event UpdateMaintainerFeeRate(uint256 oldMaintainerFeeRate, uint256 newMaintainerFeeRate)
func (_Dodo *DodoFilterer) WatchUpdateMaintainerFeeRate(opts *bind.WatchOpts, sink chan<- *DodoUpdateMaintainerFeeRate) (event.Subscription, error) {

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "UpdateMaintainerFeeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoUpdateMaintainerFeeRate)
				if err := _Dodo.contract.UnpackLog(event, "UpdateMaintainerFeeRate", log); err != nil {
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

// ParseUpdateMaintainerFeeRate is a log parse operation binding the contract event 0x6b04da3d58e4b37d99652babb3ea2bc25ce94379bfff3059f03d61b26c59e553.
//
// Solidity: event UpdateMaintainerFeeRate(uint256 oldMaintainerFeeRate, uint256 newMaintainerFeeRate)
func (_Dodo *DodoFilterer) ParseUpdateMaintainerFeeRate(log types.Log) (*DodoUpdateMaintainerFeeRate, error) {
	event := new(DodoUpdateMaintainerFeeRate)
	if err := _Dodo.contract.UnpackLog(event, "UpdateMaintainerFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DodoWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Dodo contract.
type DodoWithdrawIterator struct {
	Event *DodoWithdraw // Event containing the contract specifics and raw log

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
func (it *DodoWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DodoWithdraw)
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
		it.Event = new(DodoWithdraw)
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
func (it *DodoWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DodoWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DodoWithdraw represents a Withdraw event raised by the Dodo contract.
type DodoWithdraw struct {
	Payer         common.Address
	Receiver      common.Address
	IsBaseToken   bool
	Amount        *big.Int
	LpTokenAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe89c586bd81ee35a18f7eac22a732b56e589a2821497cce12a0208828540a36d.
//
// Solidity: event Withdraw(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Dodo *DodoFilterer) FilterWithdraw(opts *bind.FilterOpts, payer []common.Address, receiver []common.Address) (*DodoWithdrawIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dodo.contract.FilterLogs(opts, "Withdraw", payerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &DodoWithdrawIterator{contract: _Dodo.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe89c586bd81ee35a18f7eac22a732b56e589a2821497cce12a0208828540a36d.
//
// Solidity: event Withdraw(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Dodo *DodoFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *DodoWithdraw, payer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dodo.contract.WatchLogs(opts, "Withdraw", payerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DodoWithdraw)
				if err := _Dodo.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xe89c586bd81ee35a18f7eac22a732b56e589a2821497cce12a0208828540a36d.
//
// Solidity: event Withdraw(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Dodo *DodoFilterer) ParseWithdraw(log types.Log) (*DodoWithdraw, error) {
	event := new(DodoWithdraw)
	if err := _Dodo.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

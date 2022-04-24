// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vote

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

// VoteABI is the input ABI used to generate the binding from.
const VoteABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"clearVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getListNameFromID\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"linkName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"linkVoter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"countVoting\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmnt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasProfile\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasVote\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listNames\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endRegisterTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endVotingTime\",\"type\":\"uint256\"}],\"name\":\"registerVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"submitVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endRegisterTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endVotingTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"activeID\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// GetListNameFromID is a free data retrieval call binding the contract method 0x88b96c09.
//
// Solidity: function getListNameFromID(uint256 _id) view returns(bytes32[])
func (_Vote *VoteCaller) GetListNameFromID(opts *bind.CallOpts, _id *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getListNameFromID", _id)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetListNameFromID is a free data retrieval call binding the contract method 0x88b96c09.
//
// Solidity: function getListNameFromID(uint256 _id) view returns(bytes32[])
func (_Vote *VoteSession) GetListNameFromID(_id *big.Int) ([][32]byte, error) {
	return _Vote.Contract.GetListNameFromID(&_Vote.CallOpts, _id)
}

// GetListNameFromID is a free data retrieval call binding the contract method 0x88b96c09.
//
// Solidity: function getListNameFromID(uint256 _id) view returns(bytes32[])
func (_Vote *VoteCallerSession) GetListNameFromID(_id *big.Int) ([][32]byte, error) {
	return _Vote.Contract.GetListNameFromID(&_Vote.CallOpts, _id)
}

// LinkName is a free data retrieval call binding the contract method 0x1019d432.
//
// Solidity: function linkName(uint256 , bytes32 ) view returns(address)
func (_Vote *VoteCaller) LinkName(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "linkName", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkName is a free data retrieval call binding the contract method 0x1019d432.
//
// Solidity: function linkName(uint256 , bytes32 ) view returns(address)
func (_Vote *VoteSession) LinkName(arg0 *big.Int, arg1 [32]byte) (common.Address, error) {
	return _Vote.Contract.LinkName(&_Vote.CallOpts, arg0, arg1)
}

// LinkName is a free data retrieval call binding the contract method 0x1019d432.
//
// Solidity: function linkName(uint256 , bytes32 ) view returns(address)
func (_Vote *VoteCallerSession) LinkName(arg0 *big.Int, arg1 [32]byte) (common.Address, error) {
	return _Vote.Contract.LinkName(&_Vote.CallOpts, arg0, arg1)
}

// LinkVoter is a free data retrieval call binding the contract method 0x5449518f.
//
// Solidity: function linkVoter(uint256 , address ) view returns(bytes32 name, uint256 countVoting, uint256 depositAmnt, bool hasProfile, bool hasVote)
func (_Vote *VoteCaller) LinkVoter(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Name        [32]byte
	CountVoting *big.Int
	DepositAmnt *big.Int
	HasProfile  bool
	HasVote     bool
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "linkVoter", arg0, arg1)

	outstruct := new(struct {
		Name        [32]byte
		CountVoting *big.Int
		DepositAmnt *big.Int
		HasProfile  bool
		HasVote     bool
	})

	outstruct.Name = out[0].([32]byte)
	outstruct.CountVoting = out[1].(*big.Int)
	outstruct.DepositAmnt = out[2].(*big.Int)
	outstruct.HasProfile = out[3].(bool)
	outstruct.HasVote = out[4].(bool)

	return *outstruct, err

}

// LinkVoter is a free data retrieval call binding the contract method 0x5449518f.
//
// Solidity: function linkVoter(uint256 , address ) view returns(bytes32 name, uint256 countVoting, uint256 depositAmnt, bool hasProfile, bool hasVote)
func (_Vote *VoteSession) LinkVoter(arg0 *big.Int, arg1 common.Address) (struct {
	Name        [32]byte
	CountVoting *big.Int
	DepositAmnt *big.Int
	HasProfile  bool
	HasVote     bool
}, error) {
	return _Vote.Contract.LinkVoter(&_Vote.CallOpts, arg0, arg1)
}

// LinkVoter is a free data retrieval call binding the contract method 0x5449518f.
//
// Solidity: function linkVoter(uint256 , address ) view returns(bytes32 name, uint256 countVoting, uint256 depositAmnt, bool hasProfile, bool hasVote)
func (_Vote *VoteCallerSession) LinkVoter(arg0 *big.Int, arg1 common.Address) (struct {
	Name        [32]byte
	CountVoting *big.Int
	DepositAmnt *big.Int
	HasProfile  bool
	HasVote     bool
}, error) {
	return _Vote.Contract.LinkVoter(&_Vote.CallOpts, arg0, arg1)
}

// ListNames is a free data retrieval call binding the contract method 0xa3c4be58.
//
// Solidity: function listNames(uint256 , uint256 ) view returns(bytes32)
func (_Vote *VoteCaller) ListNames(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "listNames", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ListNames is a free data retrieval call binding the contract method 0xa3c4be58.
//
// Solidity: function listNames(uint256 , uint256 ) view returns(bytes32)
func (_Vote *VoteSession) ListNames(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _Vote.Contract.ListNames(&_Vote.CallOpts, arg0, arg1)
}

// ListNames is a free data retrieval call binding the contract method 0xa3c4be58.
//
// Solidity: function listNames(uint256 , uint256 ) view returns(bytes32)
func (_Vote *VoteCallerSession) ListNames(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _Vote.Contract.ListNames(&_Vote.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCallerSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// Voting is a free data retrieval call binding the contract method 0xfd4a77f1.
//
// Solidity: function voting(uint256 ) view returns(uint256 id, uint256 minDeposit, uint256 totalBalance, uint256 totalVoters, uint256 totalVotes, uint256 endRegisterTime, uint256 endVotingTime, bool activeID)
func (_Vote *VoteCaller) Voting(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              *big.Int
	MinDeposit      *big.Int
	TotalBalance    *big.Int
	TotalVoters     *big.Int
	TotalVotes      *big.Int
	EndRegisterTime *big.Int
	EndVotingTime   *big.Int
	ActiveID        bool
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "voting", arg0)

	outstruct := new(struct {
		Id              *big.Int
		MinDeposit      *big.Int
		TotalBalance    *big.Int
		TotalVoters     *big.Int
		TotalVotes      *big.Int
		EndRegisterTime *big.Int
		EndVotingTime   *big.Int
		ActiveID        bool
	})

	outstruct.Id = out[0].(*big.Int)
	outstruct.MinDeposit = out[1].(*big.Int)
	outstruct.TotalBalance = out[2].(*big.Int)
	outstruct.TotalVoters = out[3].(*big.Int)
	outstruct.TotalVotes = out[4].(*big.Int)
	outstruct.EndRegisterTime = out[5].(*big.Int)
	outstruct.EndVotingTime = out[6].(*big.Int)
	outstruct.ActiveID = out[7].(bool)

	return *outstruct, err

}

// Voting is a free data retrieval call binding the contract method 0xfd4a77f1.
//
// Solidity: function voting(uint256 ) view returns(uint256 id, uint256 minDeposit, uint256 totalBalance, uint256 totalVoters, uint256 totalVotes, uint256 endRegisterTime, uint256 endVotingTime, bool activeID)
func (_Vote *VoteSession) Voting(arg0 *big.Int) (struct {
	Id              *big.Int
	MinDeposit      *big.Int
	TotalBalance    *big.Int
	TotalVoters     *big.Int
	TotalVotes      *big.Int
	EndRegisterTime *big.Int
	EndVotingTime   *big.Int
	ActiveID        bool
}, error) {
	return _Vote.Contract.Voting(&_Vote.CallOpts, arg0)
}

// Voting is a free data retrieval call binding the contract method 0xfd4a77f1.
//
// Solidity: function voting(uint256 ) view returns(uint256 id, uint256 minDeposit, uint256 totalBalance, uint256 totalVoters, uint256 totalVotes, uint256 endRegisterTime, uint256 endVotingTime, bool activeID)
func (_Vote *VoteCallerSession) Voting(arg0 *big.Int) (struct {
	Id              *big.Int
	MinDeposit      *big.Int
	TotalBalance    *big.Int
	TotalVoters     *big.Int
	TotalVotes      *big.Int
	EndRegisterTime *big.Int
	EndVotingTime   *big.Int
	ActiveID        bool
}, error) {
	return _Vote.Contract.Voting(&_Vote.CallOpts, arg0)
}

// ClearVote is a paid mutator transaction binding the contract method 0x241c6d82.
//
// Solidity: function clearVote(uint256 _id, bytes32 name) returns()
func (_Vote *VoteTransactor) ClearVote(opts *bind.TransactOpts, _id *big.Int, name [32]byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "clearVote", _id, name)
}

// ClearVote is a paid mutator transaction binding the contract method 0x241c6d82.
//
// Solidity: function clearVote(uint256 _id, bytes32 name) returns()
func (_Vote *VoteSession) ClearVote(_id *big.Int, name [32]byte) (*types.Transaction, error) {
	return _Vote.Contract.ClearVote(&_Vote.TransactOpts, _id, name)
}

// ClearVote is a paid mutator transaction binding the contract method 0x241c6d82.
//
// Solidity: function clearVote(uint256 _id, bytes32 name) returns()
func (_Vote *VoteTransactorSession) ClearVote(_id *big.Int, name [32]byte) (*types.Transaction, error) {
	return _Vote.Contract.ClearVote(&_Vote.TransactOpts, _id, name)
}

// Deposit is a paid mutator transaction binding the contract method 0xc9630cb0.
//
// Solidity: function deposit(uint256 _id, bytes32 _name) payable returns()
func (_Vote *VoteTransactor) Deposit(opts *bind.TransactOpts, _id *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "deposit", _id, _name)
}

// Deposit is a paid mutator transaction binding the contract method 0xc9630cb0.
//
// Solidity: function deposit(uint256 _id, bytes32 _name) payable returns()
func (_Vote *VoteSession) Deposit(_id *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _Vote.Contract.Deposit(&_Vote.TransactOpts, _id, _name)
}

// Deposit is a paid mutator transaction binding the contract method 0xc9630cb0.
//
// Solidity: function deposit(uint256 _id, bytes32 _name) payable returns()
func (_Vote *VoteTransactorSession) Deposit(_id *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _Vote.Contract.Deposit(&_Vote.TransactOpts, _id, _name)
}

// RegisterVoting is a paid mutator transaction binding the contract method 0x0eaadeb7.
//
// Solidity: function registerVoting(uint256 _id, uint256 minDeposit, uint256 endRegisterTime, uint256 endVotingTime) returns()
func (_Vote *VoteTransactor) RegisterVoting(opts *bind.TransactOpts, _id *big.Int, minDeposit *big.Int, endRegisterTime *big.Int, endVotingTime *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "registerVoting", _id, minDeposit, endRegisterTime, endVotingTime)
}

// RegisterVoting is a paid mutator transaction binding the contract method 0x0eaadeb7.
//
// Solidity: function registerVoting(uint256 _id, uint256 minDeposit, uint256 endRegisterTime, uint256 endVotingTime) returns()
func (_Vote *VoteSession) RegisterVoting(_id *big.Int, minDeposit *big.Int, endRegisterTime *big.Int, endVotingTime *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.RegisterVoting(&_Vote.TransactOpts, _id, minDeposit, endRegisterTime, endVotingTime)
}

// RegisterVoting is a paid mutator transaction binding the contract method 0x0eaadeb7.
//
// Solidity: function registerVoting(uint256 _id, uint256 minDeposit, uint256 endRegisterTime, uint256 endVotingTime) returns()
func (_Vote *VoteTransactorSession) RegisterVoting(_id *big.Int, minDeposit *big.Int, endRegisterTime *big.Int, endVotingTime *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.RegisterVoting(&_Vote.TransactOpts, _id, minDeposit, endRegisterTime, endVotingTime)
}

// SubmitVote is a paid mutator transaction binding the contract method 0xac8f38c8.
//
// Solidity: function submitVote(uint256 _id, bytes32 name) returns()
func (_Vote *VoteTransactor) SubmitVote(opts *bind.TransactOpts, _id *big.Int, name [32]byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "submitVote", _id, name)
}

// SubmitVote is a paid mutator transaction binding the contract method 0xac8f38c8.
//
// Solidity: function submitVote(uint256 _id, bytes32 name) returns()
func (_Vote *VoteSession) SubmitVote(_id *big.Int, name [32]byte) (*types.Transaction, error) {
	return _Vote.Contract.SubmitVote(&_Vote.TransactOpts, _id, name)
}

// SubmitVote is a paid mutator transaction binding the contract method 0xac8f38c8.
//
// Solidity: function submitVote(uint256 _id, bytes32 name) returns()
func (_Vote *VoteTransactorSession) SubmitVote(_id *big.Int, name [32]byte) (*types.Transaction, error) {
	return _Vote.Contract.SubmitVote(&_Vote.TransactOpts, _id, name)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _id) payable returns()
func (_Vote *VoteTransactor) Withdraw(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "withdraw", _id)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _id) payable returns()
func (_Vote *VoteSession) Withdraw(_id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Withdraw(&_Vote.TransactOpts, _id)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _id) payable returns()
func (_Vote *VoteTransactorSession) Withdraw(_id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Withdraw(&_Vote.TransactOpts, _id)
}

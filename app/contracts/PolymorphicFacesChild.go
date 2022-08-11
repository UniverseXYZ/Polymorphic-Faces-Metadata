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

const PolymorphicFacesChildABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_daoAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_maticWETHAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_baseGenomeChangePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_randomizeGenomePrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_arweaveAssetsJSON\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"arweaveAssetsJSON\",\"type\":\"string\"}],\"name\":\"ArweaveAssetsJSONChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGenomeChange\",\"type\":\"uint256\"}],\"name\":\"BaseGenomeChangePriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ConsumerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRandomizeGenomePriceChange\",\"type\":\"uint256\"}],\"name\":\"RandomizeGenomePriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumPolymorphicFaces.FacesEventType\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"TokenMorphed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arweaveAssetsJSON\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseGenomeChangePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGenomeChangePrice\",\"type\":\"uint256\"}],\"name\":\"changeBaseGenomeChangePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consumer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"changeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRandomizeGenomePrice\",\"type\":\"uint256\"}],\"name\":\"changeRandomizeGenomePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"consumerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"geneOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"genomeChanges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"genomeChnages\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isNotVirgin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maticWETH\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"}],\"name\":\"mintFaceWithInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genePosition\",\"type\":\"uint256\"}],\"name\":\"morphGene\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"priceForGenomeChange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"randomizeGenome\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomizeGenomePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_arweaveAssetsJSON\",\"type\":\"string\"}],\"name\":\"setArweaveAssetsJSON\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maticWETHAddress\",\"type\":\"address\"}],\"name\":\"setMaticWETHContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"whitelistBridgeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistTunnelAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVirgin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"genomeChangesCount\",\"type\":\"uint256\"}],\"name\":\"wormholeUpdateGene\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"


// PolymorphicFacesChild is an auto generated Go binding around an Ethereum contract.
type PolymorphicFacesChild struct {
	PolymorphicFacesChildCaller     // Read-only binding to the contract
	PolymorphicFacesChildTransactor // Write-only binding to the contract
	PolymorphicFacesChildFilterer   // Log filterer for contract events
}

// PolymorphicFacesChildCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolymorphicFacesChildCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphicFacesChildTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolymorphicFacesChildTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphicFacesChildFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolymorphicFacesChildFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphicFacesChildSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolymorphicFacesChildSession struct {
	Contract     *PolymorphicFacesChild // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PolymorphicFacesChildCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolymorphicFacesChildCallerSession struct {
	Contract *PolymorphicFacesChildCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PolymorphicFacesChildTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolymorphicFacesChildTransactorSession struct {
	Contract     *PolymorphicFacesChildTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PolymorphicFacesChildRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolymorphicFacesChildRaw struct {
	Contract *PolymorphicFacesChild // Generic contract binding to access the raw methods on
}

// PolymorphicFacesChildCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolymorphicFacesChildCallerRaw struct {
	Contract *PolymorphicFacesChildCaller // Generic read-only contract binding to access the raw methods on
}

// PolymorphicFacesChildTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolymorphicFacesChildTransactorRaw struct {
	Contract *PolymorphicFacesChildTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolymorphicFacesChild creates a new instance of PolymorphicFacesChild, bound to a specific deployed contract.
func NewPolymorphicFacesChild(address common.Address, backend bind.ContractBackend) (*PolymorphicFacesChild, error) {
	contract, err := bindPolymorphicFacesChild(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChild{PolymorphicFacesChildCaller: PolymorphicFacesChildCaller{contract: contract}, PolymorphicFacesChildTransactor: PolymorphicFacesChildTransactor{contract: contract}, PolymorphicFacesChildFilterer: PolymorphicFacesChildFilterer{contract: contract}}, nil
}

// NewPolymorphicFacesChildCaller creates a new read-only instance of PolymorphicFacesChild, bound to a specific deployed contract.
func NewPolymorphicFacesChildCaller(address common.Address, caller bind.ContractCaller) (*PolymorphicFacesChildCaller, error) {
	contract, err := bindPolymorphicFacesChild(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildCaller{contract: contract}, nil
}

// NewPolymorphicFacesChildTransactor creates a new write-only instance of PolymorphicFacesChild, bound to a specific deployed contract.
func NewPolymorphicFacesChildTransactor(address common.Address, transactor bind.ContractTransactor) (*PolymorphicFacesChildTransactor, error) {
	contract, err := bindPolymorphicFacesChild(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildTransactor{contract: contract}, nil
}

// NewPolymorphicFacesChildFilterer creates a new log filterer instance of PolymorphicFacesChild, bound to a specific deployed contract.
func NewPolymorphicFacesChildFilterer(address common.Address, filterer bind.ContractFilterer) (*PolymorphicFacesChildFilterer, error) {
	contract, err := bindPolymorphicFacesChild(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildFilterer{contract: contract}, nil
}

// bindPolymorphicFacesChild binds a generic wrapper to an already deployed contract.
func bindPolymorphicFacesChild(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolymorphicFacesChildABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolymorphicFacesChild *PolymorphicFacesChildRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolymorphicFacesChild.Contract.PolymorphicFacesChildCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolymorphicFacesChild *PolymorphicFacesChildRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.PolymorphicFacesChildTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolymorphicFacesChild *PolymorphicFacesChildRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.PolymorphicFacesChildTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolymorphicFacesChild.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PolymorphicFacesChild.Contract.DEFAULTADMINROLE(&_PolymorphicFacesChild.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PolymorphicFacesChild.Contract.DEFAULTADMINROLE(&_PolymorphicFacesChild.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) MINTERROLE() ([32]byte, error) {
	return _PolymorphicFacesChild.Contract.MINTERROLE(&_PolymorphicFacesChild.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) MINTERROLE() ([32]byte, error) {
	return _PolymorphicFacesChild.Contract.MINTERROLE(&_PolymorphicFacesChild.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) PAUSERROLE() ([32]byte, error) {
	return _PolymorphicFacesChild.Contract.PAUSERROLE(&_PolymorphicFacesChild.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) PAUSERROLE() ([32]byte, error) {
	return _PolymorphicFacesChild.Contract.PAUSERROLE(&_PolymorphicFacesChild.CallOpts)
}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) ArweaveAssetsJSON(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "arweaveAssetsJSON")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) ArweaveAssetsJSON() (string, error) {
	return _PolymorphicFacesChild.Contract.ArweaveAssetsJSON(&_PolymorphicFacesChild.CallOpts)
}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) ArweaveAssetsJSON() (string, error) {
	return _PolymorphicFacesChild.Contract.ArweaveAssetsJSON(&_PolymorphicFacesChild.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.BalanceOf(&_PolymorphicFacesChild.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.BalanceOf(&_PolymorphicFacesChild.CallOpts, owner)
}

// BaseGenomeChangePrice is a free data retrieval call binding the contract method 0xce14617d.
//
// Solidity: function baseGenomeChangePrice() view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) BaseGenomeChangePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "baseGenomeChangePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseGenomeChangePrice is a free data retrieval call binding the contract method 0xce14617d.
//
// Solidity: function baseGenomeChangePrice() view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) BaseGenomeChangePrice() (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.BaseGenomeChangePrice(&_PolymorphicFacesChild.CallOpts)
}

// BaseGenomeChangePrice is a free data retrieval call binding the contract method 0xce14617d.
//
// Solidity: function baseGenomeChangePrice() view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) BaseGenomeChangePrice() (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.BaseGenomeChangePrice(&_PolymorphicFacesChild.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) BaseURI() (string, error) {
	return _PolymorphicFacesChild.Contract.BaseURI(&_PolymorphicFacesChild.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) BaseURI() (string, error) {
	return _PolymorphicFacesChild.Contract.BaseURI(&_PolymorphicFacesChild.CallOpts)
}

// ConsumerOf is a free data retrieval call binding the contract method 0xe5892331.
//
// Solidity: function consumerOf(uint256 _tokenId) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) ConsumerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "consumerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConsumerOf is a free data retrieval call binding the contract method 0xe5892331.
//
// Solidity: function consumerOf(uint256 _tokenId) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) ConsumerOf(_tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesChild.Contract.ConsumerOf(&_PolymorphicFacesChild.CallOpts, _tokenId)
}

// ConsumerOf is a free data retrieval call binding the contract method 0xe5892331.
//
// Solidity: function consumerOf(uint256 _tokenId) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) ConsumerOf(_tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesChild.Contract.ConsumerOf(&_PolymorphicFacesChild.CallOpts, _tokenId)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "daoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) DaoAddress() (common.Address, error) {
	return _PolymorphicFacesChild.Contract.DaoAddress(&_PolymorphicFacesChild.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) DaoAddress() (common.Address, error) {
	return _PolymorphicFacesChild.Contract.DaoAddress(&_PolymorphicFacesChild.CallOpts)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) GeneOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "geneOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.GeneOf(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.GeneOf(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// GenomeChanges is a free data retrieval call binding the contract method 0x23c8d07a.
//
// Solidity: function genomeChanges(uint256 tokenId) view returns(uint256 genomeChnages)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) GenomeChanges(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "genomeChanges", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenomeChanges is a free data retrieval call binding the contract method 0x23c8d07a.
//
// Solidity: function genomeChanges(uint256 tokenId) view returns(uint256 genomeChnages)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) GenomeChanges(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.GenomeChanges(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// GenomeChanges is a free data retrieval call binding the contract method 0x23c8d07a.
//
// Solidity: function genomeChanges(uint256 tokenId) view returns(uint256 genomeChnages)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) GenomeChanges(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.GenomeChanges(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesChild.Contract.GetApproved(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesChild.Contract.GetApproved(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PolymorphicFacesChild.Contract.GetRoleAdmin(&_PolymorphicFacesChild.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PolymorphicFacesChild.Contract.GetRoleAdmin(&_PolymorphicFacesChild.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _PolymorphicFacesChild.Contract.GetRoleMember(&_PolymorphicFacesChild.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _PolymorphicFacesChild.Contract.GetRoleMember(&_PolymorphicFacesChild.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.GetRoleMemberCount(&_PolymorphicFacesChild.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.GetRoleMemberCount(&_PolymorphicFacesChild.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PolymorphicFacesChild.Contract.HasRole(&_PolymorphicFacesChild.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PolymorphicFacesChild.Contract.HasRole(&_PolymorphicFacesChild.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PolymorphicFacesChild.Contract.IsApprovedForAll(&_PolymorphicFacesChild.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PolymorphicFacesChild.Contract.IsApprovedForAll(&_PolymorphicFacesChild.CallOpts, owner, operator)
}

// IsNotVirgin is a free data retrieval call binding the contract method 0x4df77416.
//
// Solidity: function isNotVirgin(uint256 ) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) IsNotVirgin(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "isNotVirgin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNotVirgin is a free data retrieval call binding the contract method 0x4df77416.
//
// Solidity: function isNotVirgin(uint256 ) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) IsNotVirgin(arg0 *big.Int) (bool, error) {
	return _PolymorphicFacesChild.Contract.IsNotVirgin(&_PolymorphicFacesChild.CallOpts, arg0)
}

// IsNotVirgin is a free data retrieval call binding the contract method 0x4df77416.
//
// Solidity: function isNotVirgin(uint256 ) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) IsNotVirgin(arg0 *big.Int) (bool, error) {
	return _PolymorphicFacesChild.Contract.IsNotVirgin(&_PolymorphicFacesChild.CallOpts, arg0)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) LastTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "lastTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) LastTokenId() (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.LastTokenId(&_PolymorphicFacesChild.CallOpts)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) LastTokenId() (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.LastTokenId(&_PolymorphicFacesChild.CallOpts)
}

// MaticWETH is a free data retrieval call binding the contract method 0x24ae3fb5.
//
// Solidity: function maticWETH() view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) MaticWETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "maticWETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MaticWETH is a free data retrieval call binding the contract method 0x24ae3fb5.
//
// Solidity: function maticWETH() view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) MaticWETH() (common.Address, error) {
	return _PolymorphicFacesChild.Contract.MaticWETH(&_PolymorphicFacesChild.CallOpts)
}

// MaticWETH is a free data retrieval call binding the contract method 0x24ae3fb5.
//
// Solidity: function maticWETH() view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) MaticWETH() (common.Address, error) {
	return _PolymorphicFacesChild.Contract.MaticWETH(&_PolymorphicFacesChild.CallOpts)
}

// Mint is a free data retrieval call binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) pure returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) Mint(opts *bind.CallOpts, to common.Address) error {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "mint", to)

	if err != nil {
		return err
	}

	return err

}

// Mint is a free data retrieval call binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) pure returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) Mint(to common.Address) error {
	return _PolymorphicFacesChild.Contract.Mint(&_PolymorphicFacesChild.CallOpts, to)
}

// Mint is a free data retrieval call binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) pure returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) Mint(to common.Address) error {
	return _PolymorphicFacesChild.Contract.Mint(&_PolymorphicFacesChild.CallOpts, to)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) Name() (string, error) {
	return _PolymorphicFacesChild.Contract.Name(&_PolymorphicFacesChild.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) Name() (string, error) {
	return _PolymorphicFacesChild.Contract.Name(&_PolymorphicFacesChild.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) Owner() (common.Address, error) {
	return _PolymorphicFacesChild.Contract.Owner(&_PolymorphicFacesChild.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) Owner() (common.Address, error) {
	return _PolymorphicFacesChild.Contract.Owner(&_PolymorphicFacesChild.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesChild.Contract.OwnerOf(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesChild.Contract.OwnerOf(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) Paused() (bool, error) {
	return _PolymorphicFacesChild.Contract.Paused(&_PolymorphicFacesChild.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) Paused() (bool, error) {
	return _PolymorphicFacesChild.Contract.Paused(&_PolymorphicFacesChild.CallOpts)
}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) PriceForGenomeChange(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "priceForGenomeChange", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) PriceForGenomeChange(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.PriceForGenomeChange(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) PriceForGenomeChange(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.PriceForGenomeChange(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// RandomizeGenomePrice is a free data retrieval call binding the contract method 0xec9c074c.
//
// Solidity: function randomizeGenomePrice() view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) RandomizeGenomePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "randomizeGenomePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomizeGenomePrice is a free data retrieval call binding the contract method 0xec9c074c.
//
// Solidity: function randomizeGenomePrice() view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) RandomizeGenomePrice() (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.RandomizeGenomePrice(&_PolymorphicFacesChild.CallOpts)
}

// RandomizeGenomePrice is a free data retrieval call binding the contract method 0xec9c074c.
//
// Solidity: function randomizeGenomePrice() view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) RandomizeGenomePrice() (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.RandomizeGenomePrice(&_PolymorphicFacesChild.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _PolymorphicFacesChild.Contract.RoyaltyInfo(&_PolymorphicFacesChild.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _PolymorphicFacesChild.Contract.RoyaltyInfo(&_PolymorphicFacesChild.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PolymorphicFacesChild.Contract.SupportsInterface(&_PolymorphicFacesChild.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PolymorphicFacesChild.Contract.SupportsInterface(&_PolymorphicFacesChild.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) Symbol() (string, error) {
	return _PolymorphicFacesChild.Contract.Symbol(&_PolymorphicFacesChild.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) Symbol() (string, error) {
	return _PolymorphicFacesChild.Contract.Symbol(&_PolymorphicFacesChild.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.TokenByIndex(&_PolymorphicFacesChild.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.TokenByIndex(&_PolymorphicFacesChild.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.TokenOfOwnerByIndex(&_PolymorphicFacesChild.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.TokenOfOwnerByIndex(&_PolymorphicFacesChild.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PolymorphicFacesChild.Contract.TokenURI(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PolymorphicFacesChild.Contract.TokenURI(&_PolymorphicFacesChild.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) TotalSupply() (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.TotalSupply(&_PolymorphicFacesChild.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) TotalSupply() (*big.Int, error) {
	return _PolymorphicFacesChild.Contract.TotalSupply(&_PolymorphicFacesChild.CallOpts)
}

// WhitelistTunnelAddresses is a free data retrieval call binding the contract method 0xcccb6d0d.
//
// Solidity: function whitelistTunnelAddresses(address ) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCaller) WhitelistTunnelAddresses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesChild.contract.Call(opts, &out, "whitelistTunnelAddresses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistTunnelAddresses is a free data retrieval call binding the contract method 0xcccb6d0d.
//
// Solidity: function whitelistTunnelAddresses(address ) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) WhitelistTunnelAddresses(arg0 common.Address) (bool, error) {
	return _PolymorphicFacesChild.Contract.WhitelistTunnelAddresses(&_PolymorphicFacesChild.CallOpts, arg0)
}

// WhitelistTunnelAddresses is a free data retrieval call binding the contract method 0xcccb6d0d.
//
// Solidity: function whitelistTunnelAddresses(address ) view returns(bool)
func (_PolymorphicFacesChild *PolymorphicFacesChildCallerSession) WhitelistTunnelAddresses(arg0 common.Address) (bool, error) {
	return _PolymorphicFacesChild.Contract.WhitelistTunnelAddresses(&_PolymorphicFacesChild.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.Approve(&_PolymorphicFacesChild.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.Approve(&_PolymorphicFacesChild.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.Burn(&_PolymorphicFacesChild.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.Burn(&_PolymorphicFacesChild.TransactOpts, tokenId)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) ChangeBaseGenomeChangePrice(opts *bind.TransactOpts, newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "changeBaseGenomeChangePrice", newGenomeChangePrice)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) ChangeBaseGenomeChangePrice(newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.ChangeBaseGenomeChangePrice(&_PolymorphicFacesChild.TransactOpts, newGenomeChangePrice)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) ChangeBaseGenomeChangePrice(newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.ChangeBaseGenomeChangePrice(&_PolymorphicFacesChild.TransactOpts, newGenomeChangePrice)
}

// ChangeConsumer is a paid mutator transaction binding the contract method 0x70b5aecb.
//
// Solidity: function changeConsumer(address _consumer, uint256 _tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) ChangeConsumer(opts *bind.TransactOpts, _consumer common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "changeConsumer", _consumer, _tokenId)
}

// ChangeConsumer is a paid mutator transaction binding the contract method 0x70b5aecb.
//
// Solidity: function changeConsumer(address _consumer, uint256 _tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) ChangeConsumer(_consumer common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.ChangeConsumer(&_PolymorphicFacesChild.TransactOpts, _consumer, _tokenId)
}

// ChangeConsumer is a paid mutator transaction binding the contract method 0x70b5aecb.
//
// Solidity: function changeConsumer(address _consumer, uint256 _tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) ChangeConsumer(_consumer common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.ChangeConsumer(&_PolymorphicFacesChild.TransactOpts, _consumer, _tokenId)
}

// ChangeRandomizeGenomePrice is a paid mutator transaction binding the contract method 0x98c5c078.
//
// Solidity: function changeRandomizeGenomePrice(uint256 newRandomizeGenomePrice) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) ChangeRandomizeGenomePrice(opts *bind.TransactOpts, newRandomizeGenomePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "changeRandomizeGenomePrice", newRandomizeGenomePrice)
}

// ChangeRandomizeGenomePrice is a paid mutator transaction binding the contract method 0x98c5c078.
//
// Solidity: function changeRandomizeGenomePrice(uint256 newRandomizeGenomePrice) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) ChangeRandomizeGenomePrice(newRandomizeGenomePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.ChangeRandomizeGenomePrice(&_PolymorphicFacesChild.TransactOpts, newRandomizeGenomePrice)
}

// ChangeRandomizeGenomePrice is a paid mutator transaction binding the contract method 0x98c5c078.
//
// Solidity: function changeRandomizeGenomePrice(uint256 newRandomizeGenomePrice) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) ChangeRandomizeGenomePrice(newRandomizeGenomePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.ChangeRandomizeGenomePrice(&_PolymorphicFacesChild.TransactOpts, newRandomizeGenomePrice)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.GrantRole(&_PolymorphicFacesChild.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.GrantRole(&_PolymorphicFacesChild.TransactOpts, role, account)
}

// MintFaceWithInfo is a paid mutator transaction binding the contract method 0x5a5fc8e5.
//
// Solidity: function mintFaceWithInfo(uint256 tokenId, address ownerAddress, uint256 gene) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) MintFaceWithInfo(opts *bind.TransactOpts, tokenId *big.Int, ownerAddress common.Address, gene *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "mintFaceWithInfo", tokenId, ownerAddress, gene)
}

// MintFaceWithInfo is a paid mutator transaction binding the contract method 0x5a5fc8e5.
//
// Solidity: function mintFaceWithInfo(uint256 tokenId, address ownerAddress, uint256 gene) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) MintFaceWithInfo(tokenId *big.Int, ownerAddress common.Address, gene *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.MintFaceWithInfo(&_PolymorphicFacesChild.TransactOpts, tokenId, ownerAddress, gene)
}

// MintFaceWithInfo is a paid mutator transaction binding the contract method 0x5a5fc8e5.
//
// Solidity: function mintFaceWithInfo(uint256 tokenId, address ownerAddress, uint256 gene) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) MintFaceWithInfo(tokenId *big.Int, ownerAddress common.Address, gene *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.MintFaceWithInfo(&_PolymorphicFacesChild.TransactOpts, tokenId, ownerAddress, gene)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) MorphGene(opts *bind.TransactOpts, tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "morphGene", tokenId, genePosition)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) MorphGene(tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.MorphGene(&_PolymorphicFacesChild.TransactOpts, tokenId, genePosition)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) MorphGene(tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.MorphGene(&_PolymorphicFacesChild.TransactOpts, tokenId, genePosition)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) Pause() (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.Pause(&_PolymorphicFacesChild.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) Pause() (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.Pause(&_PolymorphicFacesChild.TransactOpts)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) RandomizeGenome(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "randomizeGenome", tokenId)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) RandomizeGenome(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.RandomizeGenome(&_PolymorphicFacesChild.TransactOpts, tokenId)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) RandomizeGenome(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.RandomizeGenome(&_PolymorphicFacesChild.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) RenounceOwnership() (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.RenounceOwnership(&_PolymorphicFacesChild.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.RenounceOwnership(&_PolymorphicFacesChild.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.RenounceRole(&_PolymorphicFacesChild.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.RenounceRole(&_PolymorphicFacesChild.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.RevokeRole(&_PolymorphicFacesChild.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.RevokeRole(&_PolymorphicFacesChild.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SafeTransferFrom(&_PolymorphicFacesChild.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SafeTransferFrom(&_PolymorphicFacesChild.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SafeTransferFrom0(&_PolymorphicFacesChild.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SafeTransferFrom0(&_PolymorphicFacesChild.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SetApprovalForAll(&_PolymorphicFacesChild.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SetApprovalForAll(&_PolymorphicFacesChild.TransactOpts, operator, approved)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) SetArweaveAssetsJSON(opts *bind.TransactOpts, _arweaveAssetsJSON string) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "setArweaveAssetsJSON", _arweaveAssetsJSON)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) SetArweaveAssetsJSON(_arweaveAssetsJSON string) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SetArweaveAssetsJSON(&_PolymorphicFacesChild.TransactOpts, _arweaveAssetsJSON)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) SetArweaveAssetsJSON(_arweaveAssetsJSON string) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SetArweaveAssetsJSON(&_PolymorphicFacesChild.TransactOpts, _arweaveAssetsJSON)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SetBaseURI(&_PolymorphicFacesChild.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SetBaseURI(&_PolymorphicFacesChild.TransactOpts, _baseURI)
}

// SetMaticWETHContract is a paid mutator transaction binding the contract method 0x87087e61.
//
// Solidity: function setMaticWETHContract(address _maticWETHAddress) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) SetMaticWETHContract(opts *bind.TransactOpts, _maticWETHAddress common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "setMaticWETHContract", _maticWETHAddress)
}

// SetMaticWETHContract is a paid mutator transaction binding the contract method 0x87087e61.
//
// Solidity: function setMaticWETHContract(address _maticWETHAddress) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) SetMaticWETHContract(_maticWETHAddress common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SetMaticWETHContract(&_PolymorphicFacesChild.TransactOpts, _maticWETHAddress)
}

// SetMaticWETHContract is a paid mutator transaction binding the contract method 0x87087e61.
//
// Solidity: function setMaticWETHContract(address _maticWETHAddress) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) SetMaticWETHContract(_maticWETHAddress common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.SetMaticWETHContract(&_PolymorphicFacesChild.TransactOpts, _maticWETHAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.TransferFrom(&_PolymorphicFacesChild.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.TransferFrom(&_PolymorphicFacesChild.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.TransferOwnership(&_PolymorphicFacesChild.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.TransferOwnership(&_PolymorphicFacesChild.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) Unpause() (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.Unpause(&_PolymorphicFacesChild.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) Unpause() (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.Unpause(&_PolymorphicFacesChild.TransactOpts)
}

// WhitelistBridgeAddress is a paid mutator transaction binding the contract method 0xab39a3c8.
//
// Solidity: function whitelistBridgeAddress(address bridgeAddress, bool status) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) WhitelistBridgeAddress(opts *bind.TransactOpts, bridgeAddress common.Address, status bool) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "whitelistBridgeAddress", bridgeAddress, status)
}

// WhitelistBridgeAddress is a paid mutator transaction binding the contract method 0xab39a3c8.
//
// Solidity: function whitelistBridgeAddress(address bridgeAddress, bool status) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) WhitelistBridgeAddress(bridgeAddress common.Address, status bool) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.WhitelistBridgeAddress(&_PolymorphicFacesChild.TransactOpts, bridgeAddress, status)
}

// WhitelistBridgeAddress is a paid mutator transaction binding the contract method 0xab39a3c8.
//
// Solidity: function whitelistBridgeAddress(address bridgeAddress, bool status) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) WhitelistBridgeAddress(bridgeAddress common.Address, status bool) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.WhitelistBridgeAddress(&_PolymorphicFacesChild.TransactOpts, bridgeAddress, status)
}

// WormholeUpdateGene is a paid mutator transaction binding the contract method 0x6a1c03dc.
//
// Solidity: function wormholeUpdateGene(uint256 tokenId, uint256 gene, bool isVirgin, uint256 genomeChangesCount) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactor) WormholeUpdateGene(opts *bind.TransactOpts, tokenId *big.Int, gene *big.Int, isVirgin bool, genomeChangesCount *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.contract.Transact(opts, "wormholeUpdateGene", tokenId, gene, isVirgin, genomeChangesCount)
}

// WormholeUpdateGene is a paid mutator transaction binding the contract method 0x6a1c03dc.
//
// Solidity: function wormholeUpdateGene(uint256 tokenId, uint256 gene, bool isVirgin, uint256 genomeChangesCount) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildSession) WormholeUpdateGene(tokenId *big.Int, gene *big.Int, isVirgin bool, genomeChangesCount *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.WormholeUpdateGene(&_PolymorphicFacesChild.TransactOpts, tokenId, gene, isVirgin, genomeChangesCount)
}

// WormholeUpdateGene is a paid mutator transaction binding the contract method 0x6a1c03dc.
//
// Solidity: function wormholeUpdateGene(uint256 tokenId, uint256 gene, bool isVirgin, uint256 genomeChangesCount) returns()
func (_PolymorphicFacesChild *PolymorphicFacesChildTransactorSession) WormholeUpdateGene(tokenId *big.Int, gene *big.Int, isVirgin bool, genomeChangesCount *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesChild.Contract.WormholeUpdateGene(&_PolymorphicFacesChild.TransactOpts, tokenId, gene, isVirgin, genomeChangesCount)
}

// PolymorphicFacesChildApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildApprovalIterator struct {
	Event *PolymorphicFacesChildApproval // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildApproval)
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
		it.Event = new(PolymorphicFacesChildApproval)
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
func (it *PolymorphicFacesChildApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildApproval represents a Approval event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PolymorphicFacesChildApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildApprovalIterator{contract: _PolymorphicFacesChild.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildApproval)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseApproval(log types.Log) (*PolymorphicFacesChildApproval, error) {
	event := new(PolymorphicFacesChildApproval)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildApprovalForAllIterator struct {
	Event *PolymorphicFacesChildApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildApprovalForAll)
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
		it.Event = new(PolymorphicFacesChildApprovalForAll)
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
func (it *PolymorphicFacesChildApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildApprovalForAll represents a ApprovalForAll event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PolymorphicFacesChildApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildApprovalForAllIterator{contract: _PolymorphicFacesChild.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildApprovalForAll)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseApprovalForAll(log types.Log) (*PolymorphicFacesChildApprovalForAll, error) {
	event := new(PolymorphicFacesChildApprovalForAll)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildArweaveAssetsJSONChangedIterator is returned from FilterArweaveAssetsJSONChanged and is used to iterate over the raw logs and unpacked data for ArweaveAssetsJSONChanged events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildArweaveAssetsJSONChangedIterator struct {
	Event *PolymorphicFacesChildArweaveAssetsJSONChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildArweaveAssetsJSONChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildArweaveAssetsJSONChanged)
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
		it.Event = new(PolymorphicFacesChildArweaveAssetsJSONChanged)
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
func (it *PolymorphicFacesChildArweaveAssetsJSONChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildArweaveAssetsJSONChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildArweaveAssetsJSONChanged represents a ArweaveAssetsJSONChanged event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildArweaveAssetsJSONChanged struct {
	ArweaveAssetsJSON string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterArweaveAssetsJSONChanged is a free log retrieval operation binding the contract event 0x4a826ca029d05af64e411551e15f7ee1e70af0b9bc43a31154ace86a863397b4.
//
// Solidity: event ArweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterArweaveAssetsJSONChanged(opts *bind.FilterOpts) (*PolymorphicFacesChildArweaveAssetsJSONChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "ArweaveAssetsJSONChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildArweaveAssetsJSONChangedIterator{contract: _PolymorphicFacesChild.contract, event: "ArweaveAssetsJSONChanged", logs: logs, sub: sub}, nil
}

// WatchArweaveAssetsJSONChanged is a free log subscription operation binding the contract event 0x4a826ca029d05af64e411551e15f7ee1e70af0b9bc43a31154ace86a863397b4.
//
// Solidity: event ArweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchArweaveAssetsJSONChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildArweaveAssetsJSONChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "ArweaveAssetsJSONChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildArweaveAssetsJSONChanged)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "ArweaveAssetsJSONChanged", log); err != nil {
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

// ParseArweaveAssetsJSONChanged is a log parse operation binding the contract event 0x4a826ca029d05af64e411551e15f7ee1e70af0b9bc43a31154ace86a863397b4.
//
// Solidity: event ArweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseArweaveAssetsJSONChanged(log types.Log) (*PolymorphicFacesChildArweaveAssetsJSONChanged, error) {
	event := new(PolymorphicFacesChildArweaveAssetsJSONChanged)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "ArweaveAssetsJSONChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildBaseGenomeChangePriceChangedIterator is returned from FilterBaseGenomeChangePriceChanged and is used to iterate over the raw logs and unpacked data for BaseGenomeChangePriceChanged events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildBaseGenomeChangePriceChangedIterator struct {
	Event *PolymorphicFacesChildBaseGenomeChangePriceChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildBaseGenomeChangePriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildBaseGenomeChangePriceChanged)
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
		it.Event = new(PolymorphicFacesChildBaseGenomeChangePriceChanged)
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
func (it *PolymorphicFacesChildBaseGenomeChangePriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildBaseGenomeChangePriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildBaseGenomeChangePriceChanged represents a BaseGenomeChangePriceChanged event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildBaseGenomeChangePriceChanged struct {
	NewGenomeChange *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBaseGenomeChangePriceChanged is a free log retrieval operation binding the contract event 0xb1d78271daba9a366098d40b64d642a1399cabaa22c5234bacc87e92cef82ae6.
//
// Solidity: event BaseGenomeChangePriceChanged(uint256 newGenomeChange)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterBaseGenomeChangePriceChanged(opts *bind.FilterOpts) (*PolymorphicFacesChildBaseGenomeChangePriceChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "BaseGenomeChangePriceChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildBaseGenomeChangePriceChangedIterator{contract: _PolymorphicFacesChild.contract, event: "BaseGenomeChangePriceChanged", logs: logs, sub: sub}, nil
}

// WatchBaseGenomeChangePriceChanged is a free log subscription operation binding the contract event 0xb1d78271daba9a366098d40b64d642a1399cabaa22c5234bacc87e92cef82ae6.
//
// Solidity: event BaseGenomeChangePriceChanged(uint256 newGenomeChange)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchBaseGenomeChangePriceChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildBaseGenomeChangePriceChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "BaseGenomeChangePriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildBaseGenomeChangePriceChanged)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "BaseGenomeChangePriceChanged", log); err != nil {
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

// ParseBaseGenomeChangePriceChanged is a log parse operation binding the contract event 0xb1d78271daba9a366098d40b64d642a1399cabaa22c5234bacc87e92cef82ae6.
//
// Solidity: event BaseGenomeChangePriceChanged(uint256 newGenomeChange)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseBaseGenomeChangePriceChanged(log types.Log) (*PolymorphicFacesChildBaseGenomeChangePriceChanged, error) {
	event := new(PolymorphicFacesChildBaseGenomeChangePriceChanged)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "BaseGenomeChangePriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildBaseURIChangedIterator is returned from FilterBaseURIChanged and is used to iterate over the raw logs and unpacked data for BaseURIChanged events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildBaseURIChangedIterator struct {
	Event *PolymorphicFacesChildBaseURIChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildBaseURIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildBaseURIChanged)
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
		it.Event = new(PolymorphicFacesChildBaseURIChanged)
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
func (it *PolymorphicFacesChildBaseURIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildBaseURIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildBaseURIChanged represents a BaseURIChanged event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildBaseURIChanged struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseURIChanged is a free log retrieval operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterBaseURIChanged(opts *bind.FilterOpts) (*PolymorphicFacesChildBaseURIChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildBaseURIChangedIterator{contract: _PolymorphicFacesChild.contract, event: "BaseURIChanged", logs: logs, sub: sub}, nil
}

// WatchBaseURIChanged is a free log subscription operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchBaseURIChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildBaseURIChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildBaseURIChanged)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
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

// ParseBaseURIChanged is a log parse operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseBaseURIChanged(log types.Log) (*PolymorphicFacesChildBaseURIChanged, error) {
	event := new(PolymorphicFacesChildBaseURIChanged)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildConsumerChangedIterator is returned from FilterConsumerChanged and is used to iterate over the raw logs and unpacked data for ConsumerChanged events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildConsumerChangedIterator struct {
	Event *PolymorphicFacesChildConsumerChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildConsumerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildConsumerChanged)
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
		it.Event = new(PolymorphicFacesChildConsumerChanged)
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
func (it *PolymorphicFacesChildConsumerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildConsumerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildConsumerChanged represents a ConsumerChanged event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildConsumerChanged struct {
	Owner    common.Address
	Consumer common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsumerChanged is a free log retrieval operation binding the contract event 0x42ef856c2602f37ce625d252830bed486c5c8e9a4de8aa36cc3d15f304eb662b.
//
// Solidity: event ConsumerChanged(address indexed owner, address indexed consumer, uint256 indexed tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterConsumerChanged(opts *bind.FilterOpts, owner []common.Address, consumer []common.Address, tokenId []*big.Int) (*PolymorphicFacesChildConsumerChangedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "ConsumerChanged", ownerRule, consumerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildConsumerChangedIterator{contract: _PolymorphicFacesChild.contract, event: "ConsumerChanged", logs: logs, sub: sub}, nil
}

// WatchConsumerChanged is a free log subscription operation binding the contract event 0x42ef856c2602f37ce625d252830bed486c5c8e9a4de8aa36cc3d15f304eb662b.
//
// Solidity: event ConsumerChanged(address indexed owner, address indexed consumer, uint256 indexed tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchConsumerChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildConsumerChanged, owner []common.Address, consumer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "ConsumerChanged", ownerRule, consumerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildConsumerChanged)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "ConsumerChanged", log); err != nil {
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

// ParseConsumerChanged is a log parse operation binding the contract event 0x42ef856c2602f37ce625d252830bed486c5c8e9a4de8aa36cc3d15f304eb662b.
//
// Solidity: event ConsumerChanged(address indexed owner, address indexed consumer, uint256 indexed tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseConsumerChanged(log types.Log) (*PolymorphicFacesChildConsumerChanged, error) {
	event := new(PolymorphicFacesChildConsumerChanged)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "ConsumerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildOwnershipTransferredIterator struct {
	Event *PolymorphicFacesChildOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildOwnershipTransferred)
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
		it.Event = new(PolymorphicFacesChildOwnershipTransferred)
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
func (it *PolymorphicFacesChildOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildOwnershipTransferred represents a OwnershipTransferred event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolymorphicFacesChildOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildOwnershipTransferredIterator{contract: _PolymorphicFacesChild.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildOwnershipTransferred)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseOwnershipTransferred(log types.Log) (*PolymorphicFacesChildOwnershipTransferred, error) {
	event := new(PolymorphicFacesChildOwnershipTransferred)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildPausedIterator struct {
	Event *PolymorphicFacesChildPaused // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildPaused)
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
		it.Event = new(PolymorphicFacesChildPaused)
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
func (it *PolymorphicFacesChildPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildPaused represents a Paused event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterPaused(opts *bind.FilterOpts) (*PolymorphicFacesChildPausedIterator, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildPausedIterator{contract: _PolymorphicFacesChild.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildPaused) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildPaused)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParsePaused(log types.Log) (*PolymorphicFacesChildPaused, error) {
	event := new(PolymorphicFacesChildPaused)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildRandomizeGenomePriceChangedIterator is returned from FilterRandomizeGenomePriceChanged and is used to iterate over the raw logs and unpacked data for RandomizeGenomePriceChanged events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildRandomizeGenomePriceChangedIterator struct {
	Event *PolymorphicFacesChildRandomizeGenomePriceChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildRandomizeGenomePriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildRandomizeGenomePriceChanged)
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
		it.Event = new(PolymorphicFacesChildRandomizeGenomePriceChanged)
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
func (it *PolymorphicFacesChildRandomizeGenomePriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildRandomizeGenomePriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildRandomizeGenomePriceChanged represents a RandomizeGenomePriceChanged event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildRandomizeGenomePriceChanged struct {
	NewRandomizeGenomePriceChange *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterRandomizeGenomePriceChanged is a free log retrieval operation binding the contract event 0xff4da8d01e7184cc8c9d6c57d64b336b1de6d676b6215408967bd071c8da7e3d.
//
// Solidity: event RandomizeGenomePriceChanged(uint256 newRandomizeGenomePriceChange)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterRandomizeGenomePriceChanged(opts *bind.FilterOpts) (*PolymorphicFacesChildRandomizeGenomePriceChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "RandomizeGenomePriceChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildRandomizeGenomePriceChangedIterator{contract: _PolymorphicFacesChild.contract, event: "RandomizeGenomePriceChanged", logs: logs, sub: sub}, nil
}

// WatchRandomizeGenomePriceChanged is a free log subscription operation binding the contract event 0xff4da8d01e7184cc8c9d6c57d64b336b1de6d676b6215408967bd071c8da7e3d.
//
// Solidity: event RandomizeGenomePriceChanged(uint256 newRandomizeGenomePriceChange)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchRandomizeGenomePriceChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildRandomizeGenomePriceChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "RandomizeGenomePriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildRandomizeGenomePriceChanged)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "RandomizeGenomePriceChanged", log); err != nil {
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

// ParseRandomizeGenomePriceChanged is a log parse operation binding the contract event 0xff4da8d01e7184cc8c9d6c57d64b336b1de6d676b6215408967bd071c8da7e3d.
//
// Solidity: event RandomizeGenomePriceChanged(uint256 newRandomizeGenomePriceChange)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseRandomizeGenomePriceChanged(log types.Log) (*PolymorphicFacesChildRandomizeGenomePriceChanged, error) {
	event := new(PolymorphicFacesChildRandomizeGenomePriceChanged)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "RandomizeGenomePriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildRoleAdminChangedIterator struct {
	Event *PolymorphicFacesChildRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildRoleAdminChanged)
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
		it.Event = new(PolymorphicFacesChildRoleAdminChanged)
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
func (it *PolymorphicFacesChildRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildRoleAdminChanged represents a RoleAdminChanged event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolymorphicFacesChildRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildRoleAdminChangedIterator{contract: _PolymorphicFacesChild.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildRoleAdminChanged)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseRoleAdminChanged(log types.Log) (*PolymorphicFacesChildRoleAdminChanged, error) {
	event := new(PolymorphicFacesChildRoleAdminChanged)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildRoleGrantedIterator struct {
	Event *PolymorphicFacesChildRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildRoleGranted)
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
		it.Event = new(PolymorphicFacesChildRoleGranted)
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
func (it *PolymorphicFacesChildRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildRoleGranted represents a RoleGranted event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolymorphicFacesChildRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildRoleGrantedIterator{contract: _PolymorphicFacesChild.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildRoleGranted)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseRoleGranted(log types.Log) (*PolymorphicFacesChildRoleGranted, error) {
	event := new(PolymorphicFacesChildRoleGranted)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildRoleRevokedIterator struct {
	Event *PolymorphicFacesChildRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildRoleRevoked)
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
		it.Event = new(PolymorphicFacesChildRoleRevoked)
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
func (it *PolymorphicFacesChildRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildRoleRevoked represents a RoleRevoked event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolymorphicFacesChildRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildRoleRevokedIterator{contract: _PolymorphicFacesChild.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildRoleRevoked)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseRoleRevoked(log types.Log) (*PolymorphicFacesChildRoleRevoked, error) {
	event := new(PolymorphicFacesChildRoleRevoked)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildTokenMintedIterator struct {
	Event *PolymorphicFacesChildTokenMinted // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildTokenMinted)
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
		it.Event = new(PolymorphicFacesChildTokenMinted)
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
func (it *PolymorphicFacesChildTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildTokenMinted represents a TokenMinted event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildTokenMinted struct {
	TokenId *big.Int
	NewGene *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterTokenMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*PolymorphicFacesChildTokenMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildTokenMintedIterator{contract: _PolymorphicFacesChild.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildTokenMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildTokenMinted)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseTokenMinted(log types.Log) (*PolymorphicFacesChildTokenMinted, error) {
	event := new(PolymorphicFacesChildTokenMinted)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildTokenMorphedIterator is returned from FilterTokenMorphed and is used to iterate over the raw logs and unpacked data for TokenMorphed events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildTokenMorphedIterator struct {
	Event *PolymorphicFacesChildTokenMorphed // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildTokenMorphedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildTokenMorphed)
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
		it.Event = new(PolymorphicFacesChildTokenMorphed)
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
func (it *PolymorphicFacesChildTokenMorphedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildTokenMorphedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildTokenMorphed represents a TokenMorphed event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildTokenMorphed struct {
	TokenId   *big.Int
	OldGene   *big.Int
	NewGene   *big.Int
	Price     *big.Int
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenMorphed is a free log retrieval operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterTokenMorphed(opts *bind.FilterOpts, tokenId []*big.Int) (*PolymorphicFacesChildTokenMorphedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildTokenMorphedIterator{contract: _PolymorphicFacesChild.contract, event: "TokenMorphed", logs: logs, sub: sub}, nil
}

// WatchTokenMorphed is a free log subscription operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchTokenMorphed(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildTokenMorphed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildTokenMorphed)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
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

// ParseTokenMorphed is a log parse operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseTokenMorphed(log types.Log) (*PolymorphicFacesChildTokenMorphed, error) {
	event := new(PolymorphicFacesChildTokenMorphed)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildTransferIterator struct {
	Event *PolymorphicFacesChildTransfer // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildTransfer)
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
		it.Event = new(PolymorphicFacesChildTransfer)
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
func (it *PolymorphicFacesChildTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildTransfer represents a Transfer event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PolymorphicFacesChildTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildTransferIterator{contract: _PolymorphicFacesChild.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildTransfer)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseTransfer(log types.Log) (*PolymorphicFacesChildTransfer, error) {
	event := new(PolymorphicFacesChildTransfer)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesChildUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildUnpausedIterator struct {
	Event *PolymorphicFacesChildUnpaused // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesChildUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesChildUnpaused)
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
		it.Event = new(PolymorphicFacesChildUnpaused)
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
func (it *PolymorphicFacesChildUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesChildUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesChildUnpaused represents a Unpaused event raised by the PolymorphicFacesChild contract.
type PolymorphicFacesChildUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PolymorphicFacesChildUnpausedIterator, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesChildUnpausedIterator{contract: _PolymorphicFacesChild.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesChildUnpaused) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesChild.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesChildUnpaused)
				if err := _PolymorphicFacesChild.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PolymorphicFacesChild *PolymorphicFacesChildFilterer) ParseUnpaused(log types.Log) (*PolymorphicFacesChildUnpaused, error) {
	event := new(PolymorphicFacesChildUnpaused)
	if err := _PolymorphicFacesChild.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

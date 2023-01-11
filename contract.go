package w3utils

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Contract is a utility struct that provides just-enough contract-specific information
// for small Go backend services.
type Contract struct {
	Name           string         `json:"name"`    // Arbitrary contract name
	Address        common.Address `json:"address"` // Deployment address
	ContractABI    abi.ABI        `json:"-"`       // Contract's parsed abi.ABI
	ContractEvents []abi.Event    `json:"events"`  // Contract's arbitrary interesting events
}

// NewContract constructs a new Contract based on arguments.
// If eventKeys length is 0, the Contract returned will have all
// of its events from |jsonABI|.
func NewContract(
	name string, // Arbitrary contract name
	jsonABI string, // An Ethereum ABI JSON string
	addrString string, // Hexadecimal Ethereum address with 0x prepended
	eventKeys ...string, // Arbitrary event key strings
) Contract {
	contractABI, err := abi.JSON(strings.NewReader(jsonABI))
	if err != nil {
		panic("error parsing contract ABI: " + err.Error())
	}

	return Contract{
		Name:           name,
		Address:        common.HexToAddress(addrString),
		ContractABI:    contractABI,
		ContractEvents: AccrueEvents(contractABI, eventKeys...),
	}
}

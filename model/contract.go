package model

import (
	"encoding/json"
	"fmt"
)

type Contract struct {
	ContractName string `json:"contractName"`
	AbiArr       []Abi  `json:"abi"`
}

func ContractFromBytes(b []byte) *Contract {
	var contract *Contract
	json.Unmarshal(b, &contract)
	fmt.Println(contract)
	return contract
}

func (c *Contract) GetName() string {
	return c.ContractName
}

package model

import (
	"encoding/json"
	"fmt"
)

type Contract struct {
	ContractName string `json:"contractName"`
	Abi          []struct {
		Constant bool `json:"constant,omitempty"`
		Inputs   []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"inputs"`
		Name    string `json:"name,omitempty"`
		Outputs []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"outputs,omitempty"`
		Payable         bool   `json:"payable"`
		StateMutability string `json:"stateMutability"`
		Type            string `json:"type"`
	} `json:"abi"`
}

func ContractFromBytes(b []byte) *Contract {
	var contract *Contract
	json.Unmarshal(b, &contract)
	fmt.Println(contract)
	return contract
}

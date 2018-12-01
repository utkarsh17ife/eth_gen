package model

type Abi struct {
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
}

func (abi *Abi) GetName() string {
	return abi.Name
}

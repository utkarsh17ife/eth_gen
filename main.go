package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/utkarsh17ife/eth_gen/model"
)

func main() {

	cfg := model.LoanConfig()

	files, err := ioutil.ReadDir(cfg.AbiPath)
	if err != nil {
		log.Fatal(err)
	}

	var contractsCount uint8

	for _, f := range files {
		if f.Name() == "Migrations.json" {
			continue
		}

		b, err := ioutil.ReadFile(cfg.AbiPath + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		contract := model.ContractFromBytes(b)
		
		
	}

	if contractsCount > 0 {
		fmt.Printf("%d contract detected \n", contractsCount)
	}

}

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/utkarsh17ife/eth_gen/appCodeGen"
	"github.com/utkarsh17ife/eth_gen/appFileGen"
	"github.com/utkarsh17ife/eth_gen/model"
)

type Application struct {
	cfg *model.Config
}

func main() {

	App := &Application{}
	App.cfg = model.LoanConfig()

	files, err := ioutil.ReadDir(App.cfg.AbiPath)
	if err != nil {
		log.Fatal(err)
	}

	var contractsCount uint8

	//create dir structure
	appFileGen.SetupNodeProjectDir()

	for _, f := range files {
		if f.Name() == "Migrations.json" {
			continue
		}

		b, err := ioutil.ReadFile(App.cfg.AbiPath + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		contract := model.ContractFromBytes(b)

		afh := appFileGen.NewAppFileHandler(contract.GetName())
		apiCodeGen := appCodeGen.NewApiCodeGen(contract.GetName())

		//common code
		apiCodeGen.GenApiHeaderCode()

		afh.PushRouteCode(apiCodeGen.GetRouteCodeBytes())
		afh.PushControllerCode(apiCodeGen.GetControllerCodeBytes())
		afh.PushModelCode(apiCodeGen.GetModelCodeBytes())
		// for _, abi := range contract.AbiArr {

		// 	afh.PushRouteCode(codeGen.RouteCode(abi))

		// 	nodegen.GenRouteCode(abi)
		// 	nodegen.GenControllerCode(abi)
		// 	nodegen.GenModelCode(abi)
		// }

		afh.CloseFiles()
	}

	if contractsCount > 0 {
		fmt.Printf("%d contract detected \n", contractsCount)
	}

}

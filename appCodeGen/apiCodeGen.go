package appCodeGen

// import "github.com/utkarsh17ife/eth_gen/model"

type ApiCodeGen struct {
	MainName       string
	RouteCode      string
	ControllerCode string
	ModelCode      string
}

func NewApiCodeGen(name string) *ApiCodeGen {
	acg := &ApiCodeGen{}
	acg.MainName = name
	return acg
}

func (acg *ApiCodeGen) GenApiHeaderCode() {
	acg.CreateRouterHeaderCode()
	acg.CreateControllerHeaderCode()
	acg.CreateModelHeaderCode()
}

func (acg *ApiCodeGen) CreateRouterHeaderCode() {
	acg.RouteCode = `
const model = require('../controller.js')
`
}
func (acg *ApiCodeGen) CreateControllerHeaderCode() {
	acg.ControllerCode = `
const controller = require('../controller.js')
`
}
func (acg *ApiCodeGen) CreateModelHeaderCode() {
	acg.ModelCode = `
const mongoose = require('mongoose')
`
}

func (acg *ApiCodeGen) GetRouteCodeBytes() []byte {
	return []byte(acg.RouteCode)
}
func (acg *ApiCodeGen) GetControllerCodeBytes() []byte {
	return []byte(acg.ControllerCode)
}
func (acg *ApiCodeGen) GetModelCodeBytes() []byte {
	return []byte(acg.ModelCode)
}

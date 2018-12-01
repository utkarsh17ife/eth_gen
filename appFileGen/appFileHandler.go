package appFileGen

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	routeFileSuffix      = ".route.js"
	controllerFileSuffix = ".controller.js"
	modelFileSuffix      = ".model.js"
)

type AppFileHandler struct {
	Name           string
	RouteFile      *os.File
	ControllerFile *os.File
	ModelFile      *os.File
}

func NewAppFileHandler(name string) *AppFileHandler {
	afh := &AppFileHandler{}
	afh.RouteFile = CreateFile(routePathName(name))
	afh.ControllerFile = CreateFile(controllerPathName(name))
	afh.ModelFile = CreateFile(modelPathName(name))
	return afh
}

func (afh *AppFileHandler) PushRouteCode(code []byte) bool {
	afh.RouteFile.Write(code)
	return true
}
func (afh *AppFileHandler) PushControllerCode(code []byte) bool {
	afh.ControllerFile.Write(code)
	return true
}
func (afh *AppFileHandler) PushModelCode(code []byte) bool {
	afh.ModelFile.Write(code)
	return true
}

func (afh *AppFileHandler) CloseFiles() bool {
	afh.RouteFile.Close()
	afh.ControllerFile.Close()
	afh.ModelFile.Close()
	return true
}

func routePathName(name string) string {
	fName := fmt.Sprintf("%v%v", name, routeFileSuffix)
	return filepath.Join(RoutesPath, fName)
}
func controllerPathName(name string) string {
	fName := fmt.Sprintf("%v%v", name, controllerFileSuffix)
	return filepath.Join(ControllersPath, fName)
}
func modelPathName(name string) string {
	fName := fmt.Sprintf("%v%v", name, modelFileSuffix)
	return filepath.Join(ModelsPath, fName)
}

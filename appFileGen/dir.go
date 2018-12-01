package appFileGen

import (
	"os"
	"path/filepath"
)

var (
	ProjName        = "node_gen_project"
	ProjPath        = filepath.Join("./", ProjName)
	ModelsPath      = filepath.Join(ProjName, "/models")
	RoutesPath      = filepath.Join(ProjName, "/routes")
	ControllersPath = filepath.Join(ProjName, "/controllers")
)

var nodePojectDirArray = []string{ProjPath, ModelsPath, RoutesPath, ControllersPath}

func SetupNodeProjectDir() {

	for i := 0; i < len(nodePojectDirArray); i++ {
		os.Mkdir(nodePojectDirArray[i], os.ModePerm)
	}

}

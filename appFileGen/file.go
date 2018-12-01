package appFileGen

import (
	"log"
	"os"
)

func CreateFile(fileName string) *os.File {
	if f, err := os.Create(fileName); err != nil {
		log.Fatal("unable to create File %v", err.Error())
		return nil
	} else {
		return f
	}
}

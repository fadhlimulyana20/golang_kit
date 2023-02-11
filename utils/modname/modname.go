package modname

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/mod/modfile"
)

func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("./go.mod")
	if err != nil {
		log.Fatal(err.Error())
	}

	modName := modfile.ModulePath(goModBytes)
	fmt.Fprintf(os.Stdout, "modName=%+v\n", modName)

	return modName
}

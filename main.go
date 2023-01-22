package main

import (
	"template/cmd"
	_ "template/utils/env"
	_ "template/utils/log"
)

func main() {
	cmd.Execute()
}

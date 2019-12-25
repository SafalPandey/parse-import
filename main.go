package main

import (
	"encoding/json"
	"fmt"
	"os"

	"./core"
	"./utils"
)

func main() {
	var names []string
	importMap := make(map[string]interface{})

	names = append(names, os.Args[1:]...)
	names = utils.GetAbs(names)

	_, err := core.ParseImport(names, importMap)
	utils.CheckError(err)

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(importMap)

	fmt.Println("Done")
}

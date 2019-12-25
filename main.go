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

	str, err := json.MarshalIndent(importMap, "", "  ")
	utils.CheckError(err)

	f, err := os.Create("imports.json")
	utils.CheckError(err)

	defer f.Close()

	_, err = f.Write(str)
	utils.CheckError(err)

	fmt.Println("Done")
}

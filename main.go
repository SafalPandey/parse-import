package main

import (
	"encoding/json"
	"log"
	"os"

	"./core"
	"./utils"
)

var outputFile = "./imports.json"

func main() {
	var names []string
	importMap := make(map[string]interface{})

	names = append(names, os.Args[1:]...)
	names = utils.GetAbs(names)

	log.Printf("Parsing imports for: %s", names)
	_, err := core.ParseImport(names, importMap)
	utils.CheckError(err)

	str, err := json.MarshalIndent(importMap, "", "  ")
	utils.CheckError(err)

	log.Printf("Writing output to: %s", outputFile)
	f, err := os.Create(outputFile)
	utils.CheckError(err)

	defer f.Close()

	_, err = f.Write(str)
	utils.CheckError(err)

	log.Printf("Imports detected: %d", len(importMap))
	log.Printf("Done")
}

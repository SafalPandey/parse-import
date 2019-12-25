package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"./core"
	"./utils"
)

var defaultOutputFile = "./imports.json"

func main() {
	var names []string

	var filename string
	var tsconfig string
	var outputFile string

	importMap := make(map[string]interface{})

	flag.StringVar(&filename, "f", "", "File to parse")
	flag.StringVar(&tsconfig, "tsconfig", "", "Path to tsconfig")
	flag.StringVar(&outputFile, "o", defaultOutputFile, "Output file path")
	flag.Parse()

	names = append(names, filename)
	names = utils.GetAbs(names)

	if tsconfig != "" {
		log.Printf("Parsing using tsconfig file: %s", tsconfig)
		core.FindLocalDirs(tsconfig)
	}

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

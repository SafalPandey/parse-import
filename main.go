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

	var showHelp bool
	var filename string
	var tsconfig string
	var outputFile string
	var entryPoint string

	flag.BoolVar(&showHelp, "h", false, "Show usage")
	flag.StringVar(&filename, "f", "", "File to parse")
	flag.StringVar(&core.Language, "l", "ts", "Language to parse")
	flag.StringVar(&entryPoint, "entry-point", "", "Path to main.py")
	flag.StringVar(&tsconfig, "tsconfig", "", "Path to tsconfig file")
	flag.StringVar(&outputFile, "o", defaultOutputFile, "Output file path")
	flag.Parse()

	if showHelp {
		flag.Usage()

		return
	}

	core.ComputeConstants()

	names = append(names, filename)
	names = utils.GetAbs(names)

	if tsconfig != "" && core.Language == "ts" {
		log.Printf("Parsing using tsconfig file: %s", tsconfig)
		core.FindLocalDirs(tsconfig)
	}

	if entryPoint != "" {
		log.Printf("Parsing using entry point file: %s", entryPoint)
		core.FindLocalDirs(entryPoint)
	}

	core.ValidateEntrypoints(names)
	entrypointMap := core.CreateEntrypointMap(names)

	log.Printf("Parsing imports for: %s", names)
	importMap := core.ParseImport(names)
	log.Printf("Imports detected: %d", len(importMap))

	str, err := json.MarshalIndent(importMap, "", "  ")
	utils.CheckError(err)

	log.Printf("Writing output to: %s", outputFile)
	f, err := os.Create(outputFile)
	utils.CheckError(err)

	defer f.Close()

	_, err = f.Write(str)
	utils.CheckError(err)

	log.Printf("Done")
}

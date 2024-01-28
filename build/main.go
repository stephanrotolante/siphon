package main

import (
	"log"
	"os"
	"text/template"

	"github.com/stephanrotolante/siphon/cli"
	"github.com/stephanrotolante/siphon/config"
)

func main() {
	cliFlags := cli.ParseBuildFlags()

	siphonConfig := config.ReadConfig(&cliFlags)

	file, err := os.ReadFile("/home/stephan/Projects/siphon/build/handler.template.txt")

	if err != nil {
		log.Fatalf("Failed to read file %v", err)
	}

	templateFile, err := template.New("handler").Funcs(CustomFns).Parse(string(file))

	if err != nil {
		log.Fatalf("Failed to create template %v", err)
	}

	// Create a new file to write the output
	outputFile, err := os.Create("/home/stephan/Projects/siphon/handler.go")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	data := struct {
		AllowedVerbs []string
	}{
		AllowedVerbs: siphonConfig.AllowedVerbs,
	}

	templateFile.Execute(outputFile, data)
}

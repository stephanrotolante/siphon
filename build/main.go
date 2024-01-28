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

	handlerTemplateFile, err := os.ReadFile("templates/handler.template.txt")

	if err != nil {
		log.Fatalf("Failed to read handler file %v", err)
	}

	parsedHandlerTemplateFile, err := template.New("handler").Funcs(CustomFns).Parse(string(handlerTemplateFile))

	if err != nil {
		log.Fatalf("Failed to create handler template %v", err)
	}

	handlerOutputFile, err := os.Create("../handler.go")
	if err != nil {
		panic(err)
	}
	defer handlerOutputFile.Close()

	parsedHandlerTemplateFile.Execute(handlerOutputFile, siphonConfig)

	// ****************************************************************

	mainTemplateFile, err := os.ReadFile("templates/main.template.txt")

	if err != nil {
		log.Fatalf("Failed to read handler file %v", err)
	}

	parsedMainTemplateFile, err := template.New("main").Funcs(CustomFns).Parse(string(mainTemplateFile))

	if err != nil {
		log.Fatalf("Failed to create handler template %v", err)
	}

	mainOutputFile, err := os.Create("../main.go")
	if err != nil {
		panic(err)
	}
	defer mainOutputFile.Close()

	parsedMainTemplateFile.Execute(mainOutputFile, siphonConfig)
}

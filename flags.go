package main

import (
	"flag"
	"fmt"
)

type CliFlags struct {
	configPath string
	verbose    bool
}

func ParseFlags() CliFlags {

	flags := CliFlags{}

	configPath := flag.String("config", "", "Path to your siphon config")

	verbose := flag.Bool("verbose", false, "Toggle verbose mode")

	flag.Parse()

	if *configPath == "" {
		*configPath = "siphon.config.yml"
	}

	flags.configPath = *configPath
	flags.verbose = *verbose

	if flags.verbose {
		fmt.Println(fmt.Sprintf("Config location %s", flags.configPath))

	}

	return flags

}

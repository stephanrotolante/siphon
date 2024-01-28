package cli

import (
	"flag"
	"fmt"
)

type CliFlags struct {
	ConfigPath string
	Verbose    bool
}

func ParseBuildFlags() CliFlags {

	flags := CliFlags{}

	configPath := flag.String("config", "", "Path to your siphon config")

	flag.Parse()

	if *configPath == "" {
		*configPath = "siphon.config.yml"
	}

	flags.ConfigPath = *configPath

	if flags.Verbose {
		fmt.Println(fmt.Sprintf("Config location %s", flags.ConfigPath))

	}

	return flags
}

func ParseRunFlags() CliFlags {

	flags := CliFlags{}

	configPath := flag.String("config", "", "Path to your siphon config")

	verbose := flag.Bool("verbose", false, "Toggle verbose mode")

	flag.Parse()

	if *configPath == "" {
		*configPath = "siphon.config.yml"
	}

	flags.ConfigPath = *configPath
	flags.Verbose = *verbose

	if flags.Verbose {
		fmt.Println(fmt.Sprintf("Config location %s", flags.ConfigPath))

	}

	return flags

}

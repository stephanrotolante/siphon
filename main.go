package main

import (
	"fmt"
	"net/http"

	"github.com/stephanrotolante/siphon/cli"
	"github.com/stephanrotolante/siphon/config"
)

func main() {

	cliFlags := cli.ParseRunFlags()

	siphonConfig := config.ReadConfig(&cliFlags)

	fmt.Println(fmt.Sprintf("Starting siphon on port %d...", siphonConfig.Proxy.Port))

	if siphonConfig.Https.Enabled && siphonConfig.Https.CertPath != "" && siphonConfig.Https.KeyPath != "" {
		err := http.ListenAndServeTLS(
			fmt.Sprintf("%s:%d", siphonConfig.Proxy.Host, siphonConfig.Proxy.Port),
			siphonConfig.Https.CertPath,
			siphonConfig.Https.KeyPath,
			Handler(&siphonConfig))

		if err != nil {
			fmt.Println("Error starting the tls server:", err)
		}
	} else {
		err := http.ListenAndServe(
			fmt.Sprintf("%s:%d", siphonConfig.Proxy.Host, siphonConfig.Proxy.Port),
			Handler(&siphonConfig))

		if err != nil {
			fmt.Println("Error starting the server:", err)
		}
	}

}

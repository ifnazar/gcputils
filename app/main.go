package main

import (
	"gcputils/commands"
	"gcputils/utils/logger"
	"os"
)

var log = logger.New(map[string]interface{}{
	"class": "main",
})

func main() {
	log.Debug("Starting application")

	if err := commands.RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

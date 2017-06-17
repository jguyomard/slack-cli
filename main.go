package main

import (
	"fmt"

	"github.com/jguyomard/slack-cli/src/flags"
	"github.com/jguyomard/slack-cli/src/config"
	"github.com/jguyomard/slack-cli/src/commands"
)

func main() {

	// Read command options
	args, opts := flags.Parse()

	// Set config filepath
	config.SetFilePath(opts.ConfigFile)

	// Find command name
	command := ""
	if len(args) > 0 {
		command = args[0]
	}

	// Run Command
	switch command {

		case "purge-files":
			commands.PurgeFiles(opts);

		default:
			fmt.Println("Available commands:")
			fmt.Println(" - purge-files [--config-file=...] [--dry-run]")

	}
}

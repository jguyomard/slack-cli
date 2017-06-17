package flags

import 	"github.com/jessevdk/go-flags"

// Options read from command line
type Options struct {
	AssumeYes bool `short:"y" long:"assume-yes" description:"When a yes/no prompt would be presented, assume that the user entered 'yes'"`
	ConfigFile string `long:"config-file" default:"/etc/slack-cli/config.yaml" description:"Config filepath"`
	DryRun bool `long:"dry-run" description:"Enable dry-run mode"`
}

// Parse command line options
func Parse() ([]string, Options) {
	var opts Options
	args, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}
	return args, opts
}

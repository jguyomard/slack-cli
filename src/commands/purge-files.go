package commands

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/nlopes/slack"

	"github.com/jguyomard/slack-cli/src/config"
	"github.com/jguyomard/slack-cli/src/flags"
)

const (
	filesMinDuration = 30 * 24 * time.Hour
	filesMaxCount = 100
)

// PurgeFiles delete files older than `filesMinDuration` (limited to `filesMaxCount` files)
func PurgeFiles(opts flags.Options) {
	conf := config.Get()

	api := slack.New(conf.SlackToken)
	api.SetDebug(conf.DebugMode)

	// Get old Files
	filesParams := slack.NewGetFilesParameters();
	filesParams.TimestampTo = slack.JSONTime(time.Now().Add(-filesMinDuration).Unix())
	filesParams.Count = filesMaxCount
	files, _, err := api.GetFiles(filesParams)
	if err != nil {
		panic(err)
	}

	// Delete files?
	for _, file := range files {
		fmt.Printf("Delete \"%s\" (created=%s, size=%s, url=%s)? ", file.Title, file.Created, humanize.Bytes(uint64(file.Size)), file.URLPrivate)
		if opts.DryRun {
			fmt.Println("No (dry-run).")
			continue
		}
		if !opts.AssumeYes && !readYesNo() {
			fmt.Println("No!")
			continue
		}

		// Yes, delete this file!
		fmt.Print("Yes. Deleting... ")
		err := api.DeleteFile(file.ID)
		if err == nil {
			fmt.Println("Deleted.")
		} else {
			fmt.Printf("ERROR: %v\n", err);
		}
	}
}

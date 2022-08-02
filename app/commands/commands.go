package commands

import (
	"gcputils/commands/bucket"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "gcputils",
}

func init() {
	RootCmd.PersistentFlags().String("project", "", "Google project name")
	RootCmd.MarkPersistentFlagRequired("project")
	RootCmd.AddCommand(bucket.BucketCmd)
}

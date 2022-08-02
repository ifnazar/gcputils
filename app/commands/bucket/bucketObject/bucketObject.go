package bucketObject

import (
	"gcputils/commands/bucket/bucketObject/size"

	"github.com/spf13/cobra"
)

var BucketObjectCmd = &cobra.Command{
	Use: "object",
}

func init() {
	BucketObjectCmd.AddCommand(size.BucketObjectSizeCmd)
}

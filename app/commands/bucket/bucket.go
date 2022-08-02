package bucket

import (
	"gcputils/commands/bucket/bucketObject"
	"gcputils/commands/bucket/size"

	"github.com/spf13/cobra"
)

var BucketCmd = &cobra.Command{
	Use: "bucket",
}

func init() {
	BucketCmd.AddCommand(size.BucketSizeCmd)
	BucketCmd.AddCommand(bucketObject.BucketObjectCmd)

}

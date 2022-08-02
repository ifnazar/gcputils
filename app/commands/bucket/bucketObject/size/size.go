package size

import (
	"gcputils/utils/bucketUtils"
	"gcputils/utils/logger"
	"time"

	"context"

	"github.com/spf13/cobra"

	"cloud.google.com/go/storage"
)

var log = logger.New(map[string]interface{}{})

var BucketObjectSizeCmd = &cobra.Command{
	Use:   "size",
	Short: `Object size`,
	Long:  ``,
	Run:   run,
}

func init() {
	var defaultDate = time.Now().Format("20060102")

	BucketObjectSizeCmd.Flags().String("created-before", defaultDate, "created before Ex:20211031")
}

func getParamCreatedBefore(cmd *cobra.Command) time.Time {
	createdBeforeString, err := cmd.Flags().GetString("created-before")
	if err != nil {
		log.Panic(err)
	}
	createdBefore, err := time.Parse("20060102", createdBeforeString)
	if err != nil {
		log.Error("Invalid parameter value for created-before")
		log.Panic(err)
	}
	return createdBefore
}

func getParamProject(cmd *cobra.Command) string {
	var project, err = cmd.Flags().GetString("project")
	if err != nil {
		log.Panic(err)
	}
	return project
}

func run(cmd *cobra.Command, args []string) {
	var project = getParamProject(cmd)
	var createdBefore = getParamCreatedBefore(cmd)

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Panic("Failed to create client: %v", err)
	}
	defer client.Close()

	bucketUtils.ProcessBuckets(ctx, client, project, createdBefore, false, true)
}

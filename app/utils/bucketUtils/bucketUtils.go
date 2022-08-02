package bucketUtils

import (
	"context"
	"fmt"
	"gcputils/utils/logger"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gammazero/workerpool"
	"google.golang.org/api/iterator"
)

var log = logger.New(map[string]interface{}{})

func getBuckets(ctx context.Context, client *storage.Client, project string) []string {
	var result []string

	var buckets = client.Buckets(ctx, project)

	for {
		attrs, err := buckets.Next()
		if err == iterator.Done {
			break
		}
		result = append(result, attrs.Name)
	}

	return result
}

func ProcessBuckets(ctx context.Context, client *storage.Client, project string, createdFrom time.Time, printBucket bool, printObject bool) {

	if printBucket {
		log.Info(fmt.Sprintf("%s;%s;%s", "Bucket", "Count", "Size"))
	}

	if printObject {
		log.Info(fmt.Sprintf("%s;%s;%s", "Name", "Size", "Updated"))
	}

	var buckets = getBuckets(ctx, client, project)

	buckets = append(buckets, "liferaycloud-development-backup-cvfstfnkaiaiwjrtwu")

	wp := workerpool.New(4)
	for _, bucket := range buckets {
		bucket := bucket
		wp.Submit(func() {
			processObjects(ctx, project, client, bucket, createdFrom, printBucket, printObject)
		})
	}

	wp.StopWait()
}

func processObjects(ctx context.Context, project string, client *storage.Client, bucketName string, createdFrom time.Time, printBucket bool, printObject bool) {
	bucket := client.Bucket(bucketName).UserProject(project)

	var sum int64 = 0
	var count int = 0

	it := bucket.Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		if createdFrom.After(attrs.Updated) {
			if printObject {
				fmt.Printf("%s/%s;%d;%s\n", attrs.Bucket, attrs.Name, attrs.Size, attrs.Updated)
			}

			sum += attrs.Size
			count++
		}
	}

	if printBucket {
		fmt.Printf("%s;%d;%d\n", bucketName, count, sum)
	}
}

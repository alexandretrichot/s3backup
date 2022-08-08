package mongo

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/alexandretrichot/s3backup/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func backup(cmd *cobra.Command) error {
	s3Bucket, _ := cmd.Flags().GetString("s3Bucket")

	mongoUri, _ := cmd.Flags().GetString("mongoUri")
	mongodumpPath, _ := cmd.Flags().GetString("mongodumpPath")

	wg, ctx := errgroup.WithContext(context.Background())

	session, err := common.GetS3Session(cmd)
	if err != nil {
		return err
	}

	mongodump := exec.Command(mongodumpPath, "-vvvv", "--archive", "--uri", mongoUri)
	mongodump.Stderr = os.Stderr

	mongodumpStdout, err := mongodump.StdoutPipe()
	if err != nil {
		return fmt.Errorf("backup: %w", err)
	}

	uploader := s3manager.NewUploader(session)

	wg.Go(mongodump.Run)

	wg.Go(func() error {
		info, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
			Bucket: aws.String(s3Bucket),
			Key:    aws.String("dump.bson"),
			Body:   &common.Reader{R: mongodumpStdout},
		})
		if err != nil {
			return err
		}

		fmt.Println("Dump must be found at:", info.Location)

		return err
	})

	return wg.Wait()
}

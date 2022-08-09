package mongo

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"os/exec"
	"strconv"
	"time"

	"github.com/alexandretrichot/s3backup/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func backup(cmd *cobra.Command) {

	// Get Flags
	name, _ := cmd.Flags().GetString("name")
	s3Bucket, _ := cmd.Flags().GetString("s3Bucket")
	mongoUri, _ := cmd.Flags().GetString("mongoUri")
	mongodumpPath, _ := cmd.Flags().GetString("mongodumpPath")

	startTime := time.Now()

	backupName := fmt.Sprintf("%v_%v.bson", name, startTime.UTC().Format("2006-01-02T15:04:05"))

	common.AppLog.Println("Starting MongoDB Backup")
	common.AppLog.Println("")
	common.AppLog.Println("Backup Name:       ", backupName)
	common.AppLog.Println("Start time:        ", startTime.UTC())
	common.AppLog.Println("")
	common.AppLog.Println("mongodump location:", mongodumpPath)
	common.AppLog.Println("URI:               ", mongoUri)
	common.AppLog.Println("S3 Bucket:         ", s3Bucket)
	common.AppLog.Println("")

	wg, ctx := errgroup.WithContext(context.Background())

	session := common.GetS3Session(cmd)
	uploader := s3manager.NewUploader(session)

	mongodump := exec.Command(mongodumpPath, "-vvvv", "--archive", "--uri", mongoUri)
	mongodumpStdout, err := mongodump.StdoutPipe()
	if err != nil {
		common.AppErrLog.Fatal("error creating StdoutPipe for mongodump: %w", err)
	}
	defer mongodumpStdout.Close()

	mongodumpStderr, err := mongodump.StderrPipe()
	if err != nil {
		common.AppErrLog.Fatal("error creating StderrPipe for mongodump: %w", err)
	}
	defer mongodumpStderr.Close()
	stderrScanner := bufio.NewScanner(mongodumpStderr)
	go func() {
		for stderrScanner.Scan() {
			mongodumpLog.Println(stderrScanner.Text())
		}
	}()

	wg.Go(mongodump.Run)

	wg.Go(func() error {
		reader := &common.Reader{R: mongodumpStdout}
		info, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
			Bucket: aws.String(s3Bucket),
			Key:    aws.String(backupName),
			Body:   reader,
		})
		if err != nil {
			return err
		}

		backupDuration := time.Since(startTime)

		common.AppLog.Println("Success !")
		common.AppLog.Printf("Backup finished in %vh%v:%v\n", color.BlueString(strconv.Itoa(int(math.Trunc(backupDuration.Hours())))), color.BlueString(strconv.Itoa(int(math.Trunc(backupDuration.Minutes())))), color.BlueString(strconv.Itoa(int(math.Trunc(backupDuration.Seconds())))))
		common.AppLog.Printf("Backup size: %.1fMb\n", float64(reader.ReadedBytes)/1024)

		common.AppLog.Println("Dump must be found at:", info.Location)

		return nil
	})

	// Wait for group and check for error
	if err := wg.Wait(); err != nil {
		log.Fatal(err)
	}
}

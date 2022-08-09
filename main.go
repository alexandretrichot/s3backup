package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexandretrichot/s3backup/config"
	"github.com/alexandretrichot/s3backup/mongo"
	"github.com/caarlos0/env/v6"
	"github.com/spf13/cobra"
)

var version = "development"

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	rootCmd := &cobra.Command{
		Use:     "s3backup",
		Short:   "Back up your database data to a safe place in an S3 bucket. 🪣",
		Version: version,
	}
	rootCmd.PersistentFlags().String("s3Endpoint", cfg.S3Endpoint, "The s3 endpoint URL (env: S3_ENDPOINT)")
	rootCmd.PersistentFlags().String("s3Region", cfg.S3Region, "The region to use for the backup (env: S3_REGION)")
	rootCmd.PersistentFlags().String("s3AccessKeyId", cfg.S3AccessKeyId, "The s3 access key id (env: S3_ACCESS_KEY_ID)")
	rootCmd.PersistentFlags().String("s3SecretKey", cfg.S3SecretKey, "The s3 secret key (env: S3_SECRET_KEY)")
	rootCmd.PersistentFlags().String("s3Bucket", cfg.S3Bucket, "The name of the bucket (env: S3_BUCKET)")

	rootCmd.AddCommand(mongo.BuildRootCmd(cfg))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

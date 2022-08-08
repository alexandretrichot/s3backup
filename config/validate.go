package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ValidateS3Flags(cmd *cobra.Command) error {
	s3Region, _ := cmd.Flags().GetString("s3Region")
	s3AccessKeyId, _ := cmd.Flags().GetString("s3AccessKeyId")
	s3SecretKey, _ := cmd.Flags().GetString("s3SecretKey")
	s3Bucket, _ := cmd.Flags().GetString("s3Bucket")

	if s3Region == "" {
		return fmt.Errorf("required '--s3Region' flag or 'S3_REGION' env var")
	}
	if s3AccessKeyId == "" {
		return fmt.Errorf("required '--s3AccessKeyId' flag or 'S3_ACCESS_KEY_ID' env var")
	}
	if s3SecretKey == "" {
		return fmt.Errorf("required '--s3SecretKey' flag or 'S3_SECRET_KEY' env var")
	}
	if s3Bucket == "" {
		return fmt.Errorf("required '--s3Bucket' flag or 'S3_BUCKET' env var")
	}

	return nil
}

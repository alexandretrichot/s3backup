package config

import (
	"github.com/alexandretrichot/s3backup/common"
	"github.com/spf13/cobra"
)

func ValidateS3Flags(cmd *cobra.Command) {
	name, _ := cmd.Flags().GetString("name")
	s3Region, _ := cmd.Flags().GetString("s3Region")
	s3AccessKeyId, _ := cmd.Flags().GetString("s3AccessKeyId")
	s3SecretKey, _ := cmd.Flags().GetString("s3SecretKey")
	s3Bucket, _ := cmd.Flags().GetString("s3Bucket")

	if name == "" {
		common.AppErrLog.Fatal("required '--name' flag or 'BACKUP_NAME' env var")
	}
	if s3Region == "" {
		common.AppErrLog.Fatal("required '--s3Region' flag or 'S3_REGION' env var")
	}
	if s3AccessKeyId == "" {
		common.AppErrLog.Fatal("required '--s3AccessKeyId' flag or 'S3_ACCESS_KEY_ID' env var")
	}
	if s3SecretKey == "" {
		common.AppErrLog.Fatal("required '--s3SecretKey' flag or 'S3_SECRET_KEY' env var")
	}
	if s3Bucket == "" {
		common.AppErrLog.Fatal("required '--s3Bucket' flag or 'S3_BUCKET' env var")
	}
}

package common

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/cobra"
)

func GetS3Session(cmd *cobra.Command) (*session.Session, error) {
	s3Endpoint, _ := cmd.Flags().GetString("s3Endpoint")
	s3Region, _ := cmd.Flags().GetString("s3Region")
	s3AccessKeyId, _ := cmd.Flags().GetString("s3AccessKeyId")
	s3SecretKey, _ := cmd.Flags().GetString("s3SecretKey")

	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3Region),
		Endpoint:    aws.String(s3Endpoint),
		Credentials: credentials.NewStaticCredentials(s3AccessKeyId, s3SecretKey, ""),
	})
	if err != nil {
		return session, fmt.Errorf("GetSession: %w", err)
	}

	return session, nil
}

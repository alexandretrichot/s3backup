package mongo

import (
	"github.com/alexandretrichot/s3backup/config"
	"github.com/spf13/cobra"
)

func BuildRootCmd(cfg config.Config) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mongo",
		Short: "🍃 Back up and restore your MongoDB to S3",
	}

	rootCmd.PersistentFlags().StringP("mongoUri", "u", cfg.MongoURI, "The Database URI string (env: MONGO_URI)")
	rootCmd.PersistentFlags().String("mongodumpPath", cfg.MongodumpPath, "The path of 'mongodump' binary (env: MONGODUMP_PATH)")

	var backupCmd = &cobra.Command{
		Use:   "backup",
		Short: "MongoDB Backup",
		Run: func(cmd *cobra.Command, args []string) {
			config.ValidateS3Flags(cmd)
			validateMongoFlags(cmd)

			backup(cmd)
		},
	}

	rootCmd.AddCommand(backupCmd)

	return rootCmd
}

package mongo

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

func validateMongoFlags(cmd *cobra.Command) {
	if cmd.Flag("mongoUri").Value.String() == "" {
		log.Fatal(errors.New("required '--mongoUri' flag or 'MONGO_URI' env var"))
	}
}

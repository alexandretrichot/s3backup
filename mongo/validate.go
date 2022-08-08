package mongo

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ValidateMongoFlags(cmd *cobra.Command) error {
	if cmd.Flag("mongoUri").Value.String() == "" {
		return fmt.Errorf("required '--mongoUri' flag or 'MONGO_URI' env var")
	}

	return nil
}

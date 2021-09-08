package cmd

import (
	"github.com/spf13/cobra"
)

type Config struct {
	DB_Host string
	DB_Port int
	DB_User string
}

func NewMigrateCmd() *cobra.Command {
	var config Config

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "postgres migrations",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(config)
		},
	}

	cmd.Flags().StringVarP(&config.DB_Host, "host", "h", "", "db host")
	cmd.Flags().IntVarP(&config.DB_Port, "port", "p", 8080, "db port")
	cmd.Flags().StringVarP(&config.DB_User, "user", "u", "", "db user")

	cmd.AddCommand(newMigrateDownCmd(config))
	cmd.AddCommand(newMigrateUpCmd(config))

	return cmd
}

//break - roll down
//execute from file

func preRun(config Config) error {
	return nil
}

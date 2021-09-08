package cmd

import "github.com/spf13/cobra"

func newMigrateUpCmd(config Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "up postgres migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}

package cmd

import "github.com/spf13/cobra"

func newMigrateDownCmd(config Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down",
		Short: "down postgres migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}

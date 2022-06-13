package cmd

import (
	"github.com/cold-hometown/hey/config"
	"github.com/spf13/cobra"
)

func newConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "config hey",
		Long:  "Set configuration to hey",
	}

	cmd.AddCommand(nil)

	return cmd
}

func setConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [property] [value]",
		Short: "Set config",
		Args:  cobra.ExactValidArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			key, value := args[0], args[1]

			cfg := &config.Config{}

			cfg.Load()

			
		},
	}

	return cmd
}

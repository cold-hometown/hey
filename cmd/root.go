package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version,
	}
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(newConfigCommand())
	return cmd
}

func Execute() error {
	cmd := NewRootCmd()

	return cmd.Execute()
}

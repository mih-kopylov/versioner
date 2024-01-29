package cmd

import (
	"github.com/mih-kopylov/versioner/internal/app"
	"github.com/mih-kopylov/versioner/internal/fileops"
	"github.com/mih-kopylov/versioner/internal/versionops"
	"github.com/spf13/cobra"
)

func NewSetCmd(parent *cobra.Command, config *app.Config) *cobra.Command {
	var command = &cobra.Command{
		Use:   "set",
		Short: "Set version to the specified value",
		Long: `Set version to the specified value.

versioner set 1.2.3-SNAPSHOT -- changes the current version to 1.2.3-SNAPSHOT`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			nextVersion := args[0]
			err := versionops.Verify(nextVersion)
			if err != nil {
				return err
			}

			err = fileops.SetVersion(config, nextVersion)
			if err != nil {
				return err
			}

			return nil
		},
	}

	parent.AddCommand(command)
	return command
}

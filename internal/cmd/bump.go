package cmd

import (
	"github.com/mih-kopylov/versioner/internal/app"
	"github.com/mih-kopylov/versioner/internal/fileops"
	"github.com/mih-kopylov/versioner/internal/versionops"
	"github.com/spf13/cobra"
)

func NewBumpCommand(parent *cobra.Command, config *app.Config) *cobra.Command {
	var command = &cobra.Command{
		Use:   "bump",
		Short: "Bump the specified version component",
		Long: `Bump the specified version component.

For the current version 1.2.3-SNAPSHOT:
- versioner bump major -- changes version to 2.0.0-SNAPSHOT
- versioner bump minor -- changes version to 1.3.0-SNAPSHOT
- versioner bump patch -- changes version to 1.2.4-SNAPSHOT
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			currentVersion, err := fileops.GetVersion(config)
			if err != nil {
				return err
			}

			nextVersion, err := versionops.Increase(currentVersion, args[0])
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

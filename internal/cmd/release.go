package cmd

import (
	"github.com/mih-kopylov/versioner/internal/app"
	"github.com/mih-kopylov/versioner/internal/fileops"
	"github.com/mih-kopylov/versioner/internal/versionops"
	"github.com/spf13/cobra"
)

func NewReleaseCommand(parent *cobra.Command, config *app.Config) *cobra.Command {
	var command = &cobra.Command{
		Use:   "release",
		Short: "Removes suffix from the version, if any",
		Long: `Removes suffix from the version, if any.

For the snapshot version 1.2.3-SNAPSTHOT:
- versioner release -- changes version to 1.2.3
For the release version 1.2.3:
- versioner release -- doesn't change anything, just return 1.2.3
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			currentVersion, err := fileops.GetVersion(config)
			if err != nil {
				return err
			}

			nextVersion, err := versionops.Release(currentVersion)
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

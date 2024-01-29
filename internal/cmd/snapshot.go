package cmd

import (
	"github.com/mih-kopylov/versioner/internal/app"
	"github.com/mih-kopylov/versioner/internal/fileops"
	"github.com/mih-kopylov/versioner/internal/versionops"
	"github.com/spf13/cobra"
)

func NewSnapshotCommand(parent *cobra.Command, config *app.Config) *cobra.Command {
	var command = &cobra.Command{
		Use:   "snapshot",
		Short: "Adds suffix to the version",
		Long: `Adds suffix to the version.
Opposite to "release" command, puts the suffix back to the version.

For the release version 1.2.3:
- versioner snapshot -- changes version to 1.2.3-SNAPSHOT

For the snapshot version 1.2.3-SNAPSHOT:
- versioner snapshot -- doesn't change anything, just returns 1.2.3-SNAPSHOT
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			currentVersion, err := fileops.GetVersion(config)
			if err != nil {
				return err
			}

			nextVersion, err := versionops.Snapshot(currentVersion)
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

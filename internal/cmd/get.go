package cmd

import (
	"fmt"

	"github.com/mih-kopylov/versioner/internal/app"
	"github.com/mih-kopylov/versioner/internal/fileops"
	"github.com/mih-kopylov/versioner/internal/versionops"
	"github.com/spf13/cobra"
)

func NewGetCmd(parent *cobra.Command, config *app.Config) *cobra.Command {
	var releaseOnly bool
	var command = &cobra.Command{
		Use:   "get",
		Short: "Print current version",
		Long: `Print current version.

For the current version 1.2.3-SNAPSHOT:
- versioner get major           -- prints 1-SNAPSHOT
- versioner get major --release -- prints 1
- versioner get minor           -- prints 1.2-SNAPSHOT
- versioner get minor --release -- prints 1.2
- versioner get patch           -- prints 1.2.3-SNAPSHOT
- versioner get patch --release -- prints 1.2.3
- versioner get                 -- prints 1.2.3-SNAPSHOT -- same for the "patch" mode
- versioner get --release       -- prints 1.2.3 -- same for the "patch" mode

If used with 'major' argument then '1.2.3' becomes just '1'.
If used with 'minor' argument then '1.2.3' becomes just '1.2'.
Optional suffix is kept in both cases.
`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			currentVersion, err := fileops.GetVersion(config)
			if err != nil {
				return err
			}

			result := currentVersion
			if releaseOnly {
				result, err = versionops.Release(result)
				if err != nil {
					return err
				}
			}

			modeString := "patch"
			if len(args) > 0 {
				modeString = args[0]
			}
			result, err = versionops.HandleMode(
				result, modeString, versionops.RemoveMinor,
				versionops.RemovePatch, func(s string) (string, error) {
					return s, nil
				},
			)
			if err != nil {
				return err
			}

			fmt.Println(result)

			return nil
		},
	}
	command.Flags().BoolVar(&releaseOnly, "release", false, "If used, version suffix is trimmed")
	parent.AddCommand(command)
	return command
}

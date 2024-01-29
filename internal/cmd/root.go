package cmd

import (
	"github.com/mih-kopylov/versioner/internal/app"
	"github.com/spf13/cobra"
)

func NewRootCommand(info *app.Info) *cobra.Command {
	var command = &cobra.Command{
		Use:   "versioner",
		Short: "A version management tool",
		Long: `Application version is usually stored in project files, such as pom.xml, package.json etc.
Sometimes there are more then 1 file that store the same version.
This tool helps to manage version and cut a release in such cases. 
`,
	}
	command.SetVersionTemplate(`{{printf "%s" .Version}}`)

	command.Version = info.Version

	return command
}

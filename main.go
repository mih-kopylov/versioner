package main

import (
	_ "embed"
	"fmt"

	"github.com/mih-kopylov/versioner/internal/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/mih-kopylov/versioner/internal/cmd"
)

//go:embed app.yaml
var appJsonContent []byte

func main() {
	command, err := createCommand()
	if err != nil {
		panic(fmt.Errorf("%+v", err))
	}
	err = command.Execute()
	if err != nil {
		panic(fmt.Errorf("%+v", err))
	}
}

func createCommand() (*cobra.Command, error) {
	info, err := app.NewInfo(appJsonContent)
	if err != nil {
		return nil, err
	}

	config, err := app.ReadConfig()
	if err != nil {
		return nil, err
	}

	if config.Debug {
		logrus.Debug("Debug logging enabled")
	}

	rootCommand := cmd.NewRootCommand(info)

	_ = cmd.NewBumpCommand(rootCommand, config)
	_ = cmd.NewReleaseCommand(rootCommand, config)
	_ = cmd.NewSnapshotCommand(rootCommand, config)

	_ = cmd.NewGetCmd(rootCommand, config)
	_ = cmd.NewSetCmd(rootCommand, config)

	return rootCommand, nil
}

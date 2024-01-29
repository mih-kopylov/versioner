package fileops

import (
	"os"
	"path/filepath"

	"github.com/joomcode/errorx"
	"github.com/mih-kopylov/versioner/internal/app"
	"github.com/mih-kopylov/versioner/internal/versionmanager"
	"github.com/sirupsen/logrus"
)

var (
	ErrFailedGlobFiles        = errorx.NewType(errorx.CommonErrors, "ErrFailedGlobFiles")
	ErrFailedReadFileInfo     = errorx.NewType(errorx.CommonErrors, "ErrFailedReadFileInfo")
	ErrFailedWriteFileContent = errorx.NewType(errorx.CommonErrors, "ErrFailedWriteFileContent")
)

// GetVersion gets project version from the first file in config
func GetVersion(config *app.Config) (string, error) {
	configFiles, err := GetConfigFiles(config)
	if err != nil {
		return "", err
	}

	fileConfig := configFiles[0]
	versionManager, err := versionmanager.NewVersionManager(fileConfig)
	if err != nil {
		return "", err
	}

	result, err := versionManager.Read()
	if err != nil {
		return "", err
	}

	logrus.Debugf("Current version '%v'", result)

	return result, nil
}

// SetVersion writes a version in all config files
func SetVersion(config *app.Config, version string) error {
	logrus.Debugf("Set version to '%v'", version)

	configFiles, err := GetConfigFiles(config)
	if err != nil {
		return err
	}

	for _, fileConfig := range configFiles {
		versionManager, err := versionmanager.NewVersionManager(fileConfig)
		if err != nil {
			return err
		}

		newFileContent, err := versionManager.Write(version)
		if err != nil {
			return err
		}

		err = writeFileContent(fileConfig, newFileContent)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetConfigFiles resolves Glob pattern in Name, returning a slice of app.File with resoved file names
func GetConfigFiles(config *app.Config) ([]app.File, error) {
	var result []app.File

	for _, fileConfig := range config.Files {
		matches, err := filepath.Glob(fileConfig.Name)
		if err != nil {
			return nil, ErrFailedGlobFiles.WrapWithNoMessage(err)
		}
		if matches == nil {
			continue
		}
		for _, match := range matches {
			result = append(result, app.File{Name: match, Path: fileConfig.Path, Regexp: fileConfig.Regexp})
		}
	}

	if len(result) == 0 {
		return nil, errorx.AssertionFailed.New("No files found for the config")
	}

	logrus.Debugf("Files found: '%v'", result)

	return result, nil
}

func writeFileContent(fileConfig app.File, content string) error {
	stat, err := os.Stat(fileConfig.Name)
	if err != nil {
		return ErrFailedReadFileInfo.WrapWithNoMessage(err)
	}
	perm := stat.Mode().Perm()

	err = os.WriteFile(fileConfig.Name, []byte(content), perm)
	if err != nil {
		return ErrFailedWriteFileContent.WrapWithNoMessage(err)
	}
	return err
}

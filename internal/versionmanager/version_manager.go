package versionmanager

import (
	"bytes"
	"os"
	"strings"

	"github.com/joomcode/errorx"
	"github.com/mih-kopylov/versioner/internal/app"
	"github.com/mih-kopylov/versioner/internal/pathfinder"
)

var (
	ErrFailedReadFileContent = errorx.NewType(errorx.CommonErrors, "ErrFailedReadFileContent")
)

// VersionManager manages version in string content
type VersionManager interface {
	Read() (string, error)
	Write(version string) (string, error)
}

func NewVersionManager(fileConfig app.File) (VersionManager, error) {
	if len(fileConfig.Regexp) == 0 && len(fileConfig.Path) == 0 {
		return nil, errorx.AssertionFailed.New(
			"Either `regexp` or `path` field is expected in file config '%s'", fileConfig.Name,
		)
	}

	content, err := readFileContent(fileConfig)
	if err != nil {
		return nil, err
	}

	if len(fileConfig.Regexp) > 0 {
		return &RegexpVersionManager{content, fileConfig.Name, fileConfig.Regexp}, nil
	}

	pathFinder, err := pathfinder.NewPathFinder(fileConfig.Name)
	if err != nil {
		return nil, err
	}

	return &PathVersionManager{content, pathFinder, fileConfig.Path, fileConfig.Name}, nil
}

func readFileContent(fileConfig app.File) (string, error) {
	if len(fileConfig.Name) == 0 {
		return "", errorx.AssertionFailed.New("No file name provided, check config structure: %#v", fileConfig)
	}

	contentBytes, err := os.ReadFile(fileConfig.Name)
	if err != nil {
		return "", ErrFailedReadFileContent.WrapWithNoMessage(err)
	}

	//Make sure the file has \n line endings
	contentBytes = forceLineEndings(contentBytes)

	result := string(contentBytes)
	//Make sure files have new line, so that PathFinders parse content correctly.
	if !strings.HasSuffix(result, "\n") {
		result = result + "\n"
	}

	return result, nil
}

func forceLineEndings(content []byte) []byte {
	// replace CR LF \r\n (windows) with LF \n (unix)
	result := bytes.ReplaceAll(content, []byte{13, 10}, []byte{10})

	// replace CF \r (mac) with LF \n (unix)
	result = bytes.ReplaceAll(result, []byte{13}, []byte{10})

	return result
}

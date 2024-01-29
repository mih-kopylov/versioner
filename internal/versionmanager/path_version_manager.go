package versionmanager

import (
	"strings"

	"github.com/joomcode/errorx"
	"github.com/mih-kopylov/versioner/internal/pathfinder"
	"github.com/sirupsen/logrus"
)

type PathVersionManager struct {
	content    string
	pathFinder pathfinder.PathFinder
	path       string
	fileName   string
}

func (r PathVersionManager) Read() (string, error) {
	logrus.Debugf("Read version with path '%v' from file '%v'", r.path, r.fileName)

	parts, err := getPathParts(r.path, r.fileName)
	if err != nil {
		return "", err
	}

	result, err := r.pathFinder.Find(r.content, parts)
	if err != nil {
		return "", err
	}

	return result.Value, nil
}

func (r PathVersionManager) Write(version string) (string, error) {
	logrus.Debugf("Write version with path '%v' to file '%v'", r.path, r.fileName)

	parts, err := getPathParts(r.path, r.fileName)
	if err != nil {
		return "", err
	}

	result, err := r.pathFinder.Find(r.content, parts)
	if err != nil {
		return "", err
	}

	return r.content[:result.Start] + version + r.content[result.End:], nil
}

func getPathParts(path string, fileName string) ([]string, error) {
	if path == "" {
		return nil, errorx.AssertionFailed.New(
			"Either `regexp` or `path` field is expected in file config '%s'", fileName,
		)
	}

	parts := strings.Split(path, ".")
	if len(parts) <= 1 {
		return nil, errorx.AssertionFailed.New(
			"Path is expected to have more than 1 element in file config '%s'", fileName,
		)
	}

	if parts[0] != "$" {
		return nil, errorx.AssertionFailed.New("Path is expected to start from `$` in file config '%s'", fileName)
	}

	return parts[1:], nil
}

package versionops

import (
	"regexp"

	"github.com/coreos/go-semver/semver"
	"github.com/joomcode/errorx"
	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidVersionFormat = errorx.NewType(errorx.CommonErrors, "ErrInvalidVersionFormat")
	ErrUnknownMode          = errorx.NewType(errorx.CommonErrors, "ErrUnknownMode")
)

type Mode string

const (
	MajorMode = "major"
	MinorMode = "minor"
	PatchMode = "patch"
)

func Verify(version string) error {
	_, err := semver.NewVersion(version)
	if err != nil {
		return ErrInvalidVersionFormat.WrapWithNoMessage(err)
	}

	return nil
}

func IncreaseMajor(version string) (string, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return "", ErrInvalidVersionFormat.WrapWithNoMessage(err)
	}

	v.Major = v.Major + 1
	v.Minor = 0
	v.Patch = 0
	result := v.String()
	logrus.Debugf("Increase major version from '%v' to '%v'", version, result)
	return result, nil
}

func IncreaseMinor(version string) (string, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return "", ErrInvalidVersionFormat.WrapWithNoMessage(err)
	}

	v.Minor = v.Minor + 1
	v.Patch = 0
	result := v.String()
	logrus.Debugf("Increase minor version from '%v' to '%v'", version, result)
	return result, nil
}

func IncreasePatch(version string) (string, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return "", ErrInvalidVersionFormat.WrapWithNoMessage(err)
	}

	v.Patch = v.Patch + 1
	result := v.String()
	logrus.Debugf("Increase patch version from '%v' to '%v'", version, result)
	return result, nil
}

func Release(version string) (string, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return "", ErrInvalidVersionFormat.WrapWithNoMessage(err)
	}

	v.PreRelease = ""
	result := v.String()
	logrus.Debugf("Release version from '%v' to '%v'", version, result)
	return result, nil
}

func Snapshot(version string) (string, error) {
	return SuffixSnapshot(version, "")
}

func SuffixSnapshot(version string, suffix string) (string, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return "", ErrInvalidVersionFormat.WrapWithNoMessage(err)
	}

	suffixToAdd := "SNAPSHOT"
	if len(suffix) > 0 {
		suffixToAdd = suffix + "-" + suffixToAdd
	}
	v.PreRelease = semver.PreRelease(suffixToAdd)
	result := v.String()
	logrus.Debugf("Add suffix to version from '%v' to '%v'", version, result)
	return result, nil
}

func RemoveMinor(version string) (string, error) {
	reg, err := regexp.Compile(`(\d+)(\.\d+\.\d+)(-.+)?`)
	if err != nil {
		return "", err
	}
	submatch := reg.FindStringSubmatch(version)
	result := submatch[1] + submatch[3]
	return result, nil
}

func RemovePatch(version string) (string, error) {
	reg, err := regexp.Compile(`(\d+\.\d+)(\.\d+)(-.+)?`)
	if err != nil {
		return "", err
	}
	submatch := reg.FindStringSubmatch(version)
	result := submatch[1] + submatch[3]
	return result, nil
}

func Increase(version string, modeString string) (string, error) {
	return HandleMode(version, modeString, IncreaseMajor, IncreaseMinor, IncreasePatch)
}

type VersionHandler func(string) (string, error)

func HandleMode(
	version string, modeString string, majorHandler VersionHandler, minorHandler VersionHandler,
	patchHandler VersionHandler,
) (string, error) {
	mode := Mode(modeString)
	switch mode {
	case MajorMode:
		return majorHandler(version)
	case MinorMode:
		return minorHandler(version)
	case PatchMode:
		return patchHandler(version)
	default:
		return "", ErrUnknownMode.New(modeString)
	}
}

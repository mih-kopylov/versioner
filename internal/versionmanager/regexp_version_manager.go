package versionmanager

import (
	"regexp"

	"github.com/joomcode/errorx"
	"github.com/sirupsen/logrus"
)

var (
	ErrWrongRegexp = errorx.NewType(errorx.CommonErrors, "ErrWrongRegexp")
)

type RegexpVersionManager struct {
	content  string
	fileName string
	regexp   string
}

func (r RegexpVersionManager) Read() (string, error) {
	logrus.Debugf("Read version with regexp '%v' from file '%v'", r.regexp, r.fileName)

	reg, err := regexp.Compile(r.regexp)
	if err != nil {
		return "", ErrWrongRegexp.Wrap(err, r.regexp)
	}

	if reg.NumSubexp() != 1 {
		return "", errorx.AssertionFailed.New("Regexp '%s' should have a single group with version", r.regexp)
	}

	if !reg.MatchString(r.content) {
		return "", errorx.AssertionFailed.New("File '%s' content doesn't match regexp '%s'", r.fileName, r.regexp)
	}

	submatches := reg.FindStringSubmatch(r.content)
	return submatches[1], nil
}

func (r RegexpVersionManager) Write(version string) (string, error) {
	logrus.Debugf("Write version with regexp '%v' to file '%v'", r.regexp, r.fileName)

	reg, err := regexp.Compile(r.regexp)
	if err != nil {
		return "", ErrWrongRegexp.Wrap(err, r.regexp)
	}

	if reg.NumSubexp() != 1 {
		return "", errorx.AssertionFailed.New("Regexp '%s' should have a single group with version", r.regexp)
	}

	if !reg.MatchString(r.content) {
		return "", errorx.AssertionFailed.New("File '%s' content doesn't match regexp '%s'", r.fileName, r.regexp)
	}

	result := reg.ReplaceAllStringFunc(
		r.content, func(s string) string {
			//start-end positions of the first group
			firstGroupIndex := reg.FindStringSubmatchIndex(s)[2:4]
			return s[:firstGroupIndex[0]] + version + s[firstGroupIndex[1]:]
		},
	)

	return result, nil
}

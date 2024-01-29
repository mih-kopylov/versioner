package app

import (
	"github.com/joomcode/errorx"
	"gopkg.in/yaml.v3"
)

type Info struct {
	Version string `json:"version"`
}

var (
	ErrFailedUnmarshallInfo = errorx.NewType(errorx.CommonErrors, "ErrFailedUnmarshallInfo")
)

func NewInfo(content []byte) (*Info, error) {
	info := &Info{}
	err := yaml.Unmarshal(content, info)
	if err != nil {
		return nil, ErrFailedUnmarshallInfo.WrapWithNoMessage(err)
	}

	return info, nil
}

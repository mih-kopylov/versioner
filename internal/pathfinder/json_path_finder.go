package pathfinder

import (
	"github.com/buger/jsonparser"
	"github.com/joomcode/errorx"
)

var (
	ErrFailedParseJson = errorx.NewType(errorx.CommonErrors, "ErrFailedParseJson")
)

type JsonPathFinder struct {
}

func (f JsonPathFinder) Find(content string, parts []string) (*Result, error) {
	value, dataType, offset, err := jsonparser.Get([]byte(content), parts...)
	if err != nil {
		return nil, ErrFailedParseJson.WrapWithNoMessage(err)
	}
	if dataType == jsonparser.String {
		//subtract 1 additionally because string is quoted
		return &Result{offset - len(value) - 1, offset - 1, string(value)}, nil
	}
	if dataType == jsonparser.Number {
		return &Result{offset - len(value), offset, string(value)}, nil
	}
	return nil, errorx.AssertionFailed.New(
		"It's expected that element '%v' is a string or number, but '%v' found", parts, dataType,
	)
}

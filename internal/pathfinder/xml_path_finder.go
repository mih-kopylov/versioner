package pathfinder

import (
	"encoding/xml"

	"strings"

	"github.com/joomcode/errorx"
)

var (
	ErrFailedParseXml = errorx.NewType(errorx.CommonErrors, "ErrFailedParseXml")
)

type XmlPathFinder struct {
}

func (f XmlPathFinder) Find(content string, parts []string) (*Result, error) {
	reader := strings.NewReader(content)
	decoder := xml.NewDecoder(reader)

	for _, part := range parts {
		for {
			token, err := decoder.Token()
			if err != nil {
				return nil, ErrFailedParseXml.WrapWithNoMessage(err)
			}
			if startElement, ok := token.(xml.StartElement); ok {
				if startElement.Name.Local == part {
					//current level element found
					break
				} else {
					//it's another element found, skip it
					err := decoder.Skip()
					if err != nil {
						return nil, ErrFailedParseXml.WrapWithNoMessage(err)
					}
				}
			} else {
				//it's not an element, some other token. Need to seek for another one
				continue
			}
		}
	}
	//attempt to take inner text element
	pathToken, err := decoder.Token()
	if err != nil {
		return nil, ErrFailedParseXml.WrapWithNoMessage(err)
	}
	if charData, ok := pathToken.(xml.CharData); ok {
		end := int(decoder.InputOffset())
		start := end - len(charData)
		return &Result{start, end, string(charData)}, nil
	} else {
		return nil, errorx.AssertionFailed.New("Element in path '%v' is not a text element", parts)
	}
}

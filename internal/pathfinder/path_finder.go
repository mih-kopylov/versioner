package pathfinder

import (
	"path/filepath"

	"github.com/joomcode/errorx"
)

var (
	ErrUnsupportedFileExtension = errorx.NewType(errorx.CommonErrors, "ErrUnsupportedFileExtension")
)

// PathFinder finds a position of the version in string content using path array.
// Path array is an array of node names.
// It's agnostic to content type, and it's up to implementation to interpret them
type PathFinder interface {
	Find(content string, parts []string) (*Result, error)
}

// Result stores result of the version that was found and its position in string content
type Result struct {
	Start int
	End   int
	Value string
}

func NewPathFinder(fileName string) (PathFinder, error) {
	extension := filepath.Ext(fileName)
	switch extension {
	case ".json":
		return JsonPathFinder{}, nil
	case ".yaml":
		fallthrough
	case ".yml":
		return YamlPathFinder{}, nil
	case ".xml":
		return XmlPathFinder{}, nil
	default:
		return nil, ErrUnsupportedFileExtension.New(fileName)
	}
}

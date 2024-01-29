package pathfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleYaml(t *testing.T) {
	content := `
project:
  version: 123
`
	finder := YamlPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "123", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

func TestYamlSnapshotVersion(t *testing.T) {
	content := `
project:
  version: 1.2.3-SNAPSHOT
`
	finder := YamlPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "1.2.3-SNAPSHOT", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

func TestYamlReleaseVersion(t *testing.T) {
	content := `
project:
  version: 1.2.3
`
	finder := YamlPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "1.2.3", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

func TestYamlWithMultipleNodes(t *testing.T) {
	content := `
init:
  version: 111
project:
  version: 123
another:
  version: 222
`
	finder := YamlPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "123", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

func TestYamlWithArrayNodes(t *testing.T) {
	content := `
init:
  versions:
  - first: 111
  - second: 222
project:
  version: 123
`
	finder := YamlPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "123", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

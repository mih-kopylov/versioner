package pathfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleJson(t *testing.T) {
	content := `
{
  "project": {
    "version": "123"
  }
}
`
	finder := JsonPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "123", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

func TestJsonNotFound(t *testing.T) {
	content := `
{
  "project": {
    "version": "123"
  }
}
`
	finder := JsonPathFinder{}
	_, err := finder.Find(content, []string{"project", "other"})
	assert.Error(t, err)
}

func TestMalformedJson(t *testing.T) {
	content := `
{
  "project":
    "version": "123"
  }
}
`
	finder := JsonPathFinder{}
	_, err := finder.Find(content, []string{"project", "version"})
	assert.Error(t, err)
}

func TestJsonObject(t *testing.T) {
	content := `
{
  "project": {
    "version": {
      "value": "123"
    }
  }
}
`
	finder := JsonPathFinder{}
	_, err := finder.Find(content, []string{"project", "version"})
	assert.Error(t, err)
}

func TestJsonNumber(t *testing.T) {
	content := `
{
  "project": {
    "version": 123
  }
}
`
	finder := JsonPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "123", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

package pathfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleXml(t *testing.T) {
	content := `
<project>
  <version>123</version>
</project>
`
	finder := XmlPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "123", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

func TestXmlWithAttributes(t *testing.T) {
	content := `
<project attr="val">
  <version attr="val">123</version>
</project>
`
	finder := XmlPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "123", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

func TestXmlWithSeveralVersions(t *testing.T) {
	content := `
<project>
  <parent>
    <version>ParentVersion</version>
  </parent>
  <version>123</version>
  <repository>
    <version>RepositoryVersion</version>
  </repository>
</project>
`
	finder := XmlPathFinder{}
	result, err := finder.Find(content, []string{"project", "version"})
	if assert.NoError(t, err) {
		assert.Equal(t, "123", result.Value)
		assert.Equal(t, result.Value, content[result.Start:result.End])
	}
}

func TestMalformedXml(t *testing.T) {
	content := `
project
  <version>123</version>
</project>
`
	finder := XmlPathFinder{}
	_, err := finder.Find(content, []string{"project", "version"})
	assert.Error(t, err)
}

func TestXmlPathNotFound(t *testing.T) {
	content := `
<project>
  <version>123</version>
</project>
`
	finder := XmlPathFinder{}
	_, err := finder.Find(content, []string{"project", "another"})
	assert.Error(t, err)
}

package versionmanager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexpVersionManager(t *testing.T) {
	content := `
project.version=123
`
	expectedContent := `
project.version=321
`
	regexp := "project\\.version=(.+)"
	versionManager := RegexpVersionManager{content, "123.txt", regexp}
	version, err := versionManager.Read()
	if assert.NoError(t, err) {
		assert.Equal(t, "123", version)
	}

	result, err := versionManager.Write("321")
	if assert.NoError(t, err) {
		assert.Equal(t, expectedContent, result)
	}

}

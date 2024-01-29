package versionops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreaseMajorSnapshot(t *testing.T) {
	actual, err := IncreaseMajor("1.2.3-SNAPSHOT")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "2.0.0-SNAPSHOT", actual)
}

func TestIncreaseMajorRelease(t *testing.T) {
	actual, err := IncreaseMajor("1.2.3")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "2.0.0", actual)
}

func TestIncreaseMinorSnapshot(t *testing.T) {
	actual, err := IncreaseMinor("1.2.3-SNAPSHOT")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.3.0-SNAPSHOT", actual)
}

func TestIncreaseMinorRelease(t *testing.T) {
	actual, err := IncreaseMinor("1.2.3")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.3.0", actual)
}

func TestIncreasePatchSnapshot(t *testing.T) {
	actual, err := IncreasePatch("1.2.3-SNAPSHOT")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.2.4-SNAPSHOT", actual)
}

func TestIncreasePatchRelease(t *testing.T) {
	actual, err := IncreasePatch("1.2.3")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.2.4", actual)
}

func TestReleaseSnapshot(t *testing.T) {
	actual, err := Release("1.2.3-SNAPSHOT")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.2.3", actual)
}

func TestReleaseRelease(t *testing.T) {
	actual, err := Release("1.2.3")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.2.3", actual)
}

func TestSnapshotSnapshot(t *testing.T) {
	actual, err := Snapshot("1.2.3-SNAPSHOT")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.2.3-SNAPSHOT", actual)
}

func TestSnapshotRelease(t *testing.T) {
	actual, err := Snapshot("1.2.3")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.2.3-SNAPSHOT", actual)
}

func TestSuffixSnapshotSnapshot(t *testing.T) {
	actual, err := SuffixSnapshot("1.2.3-SNAPSHOT", "AAA")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.2.3-AAA-SNAPSHOT", actual)
}

func TestSuffixSnapshotSnapshotOverride(t *testing.T) {
	actual, err := SuffixSnapshot("1.2.3-BBB-SNAPSHOT", "AAA")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.2.3-AAA-SNAPSHOT", actual)
}

func TestSuffixSnapshotRelease(t *testing.T) {
	actual, err := SuffixSnapshot("1.2.3", "AAA")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "1.2.3-AAA-SNAPSHOT", actual)
}

func TestVerifyValid(t *testing.T) {
	tests := []string{
		"1.2.3",
		"1.2.3-SNAPSHOT",
		"1.2.3-SNAPSHOTUS",
		"1.2.3-ANOTHER-SNAPSHOT",
	}
	for _, test := range tests {
		t.Run(
			test, func(t *testing.T) {
				assert.NoError(t, Verify(test))
			},
		)
	}
}

func TestVerifyInvalid(t *testing.T) {
	tests := []string{
		"a.2.3",
		"1,2,3",
		"1.2.3-ANOTHER_SNAPSHOT",
		"1.2.3_SNAPSHOT",
	}
	for _, test := range tests {
		t.Run(
			test, func(t *testing.T) {
				assert.Error(t, Verify(test))
			},
		)
	}
}

func TestRemoveMinor(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected string
	}{
		{name: "Release", version: "1.2.3", expected: "1"},
		{name: "Snapshot", version: "1.2.3-SNAPSHOT", expected: "1-SNAPSHOT"},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				actual, err := RemoveMinor(test.version)
				assert.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			},
		)
	}
}

func TestRemovePatch(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected string
	}{
		{name: "Release", version: "1.2.3", expected: "1.2"},
		{name: "Snapshot", version: "1.2.3-SNAPSHOT", expected: "1.2-SNAPSHOT"},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				actual, err := RemovePatch(test.version)
				assert.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			},
		)
	}
}

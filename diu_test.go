package diu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImageManifest(t *testing.T) {
	manifest, err := GetImageManifest("hello-world")

	assert.Nil(t, err)
	if assert.NotNil(t, manifest) {
		// Test layer count
		assert.Equal(t, 1, len(manifest.SchemaV2Manifest.Layers))
	}
}

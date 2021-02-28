package diu

import (
	"github.com/compose-generator/docker-inspect-utils/model"
	"github.com/compose-generator/docker-inspect-utils/parser"
)

// GetImageManifest returns the manifest of a remote Docker image
func GetImageManifest(image string) (model.DockerManifest, error) {
	return parser.ParseDockerManifest(image)
}

// GetExistingVolumes returns details about the existing volumes of the docker instance
func GetExistingVolumes() ([]model.DockerVolume, error) {
	return parser.ParseDockerVolumes()
}
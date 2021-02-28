package parser

import (
	"docker-inspect-utils/model"
	"docker-inspect-utils/utils"
	"encoding/json"
	"errors"
	"strings"
)

// ParseDockerManifest retrieves the manifest of a remote Docker image and bundles it to an object
func ParseDockerManifest(imageName string) (manifest model.DockerManifest, err error) {
	manifestJSON := utils.ExecuteAndWaitWithOutput("docker", "manifest", "inspect", "-v", imageName)
	if strings.HasPrefix(manifestJSON, "[") {
		var manifestArray []model.DockerManifest
		json.Unmarshal([]byte(manifestJSON), &manifestArray)
		manifest = manifestArray[0]
		return
	} else if strings.HasPrefix(manifestJSON, "{") {
		json.Unmarshal([]byte(manifestJSON), &manifest)
		return
	}
	err = errors.New("Could not parse manifest")
	return
}

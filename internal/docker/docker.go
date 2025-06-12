package docker

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// "github.com/docker/docker/api/types/build"
// "github.com/docker/docker/client"
func BuildImage(path string) error {
	app := "docker"
	// Print the output
	log.Info(path)
	cmd := exec.Command(app, "build", "-f", "Dockerfile", ".")
	stdout, err := cmd.Output()

	if err != nil {
		log.Error(err.Error())
		return err
	}
	// Print the output
	log.Info(string(stdout))
	return nil
}

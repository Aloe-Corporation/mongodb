package docker

import (
	"fmt"
	"os/exec"
)

var _ Manager = (*Compose)(nil)

// Compose is a docker manager for up and down docker-compose file.
type Compose struct {
	PathFile string
}

// Up is use for up docker-compose file.
func (c *Compose) Up() error {
	// #nosec
	cmd := exec.Command("docker-compose", "-f", c.PathFile, "up", "-d", "--build", "--force-recreate")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command \"%s\" fail.\nOutput : %s .\nError : %w", cmd.String(), string(out), err)
	}
	return nil
}

// Down is use for down docker-compose file.
func (c *Compose) Down() error {
	cmd := exec.Command("docker-compose", "-f", c.PathFile, "down", "-v")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command \"%s\" fail.\nOutput : %s .\nError : %w", cmd.String(), string(out), err)
	}
	return nil
}

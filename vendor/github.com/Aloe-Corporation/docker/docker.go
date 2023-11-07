package docker

import (
	"fmt"
	"os/exec"
)

// Manager is an interface for create a manager of docker container.
type Manager interface {
	// Up the container
	Up() error
	// Down the container
	Down() error
}

// Run a new container with the given image and the given name, you can precise more parameters to the docker run command.
func Run(image, name string, args ...string) error {
	return RunWithFlag(image, name, nil, args...)
}

// RunWithFlag a new container with the given image and the given name, you can pass flag to the image as a []string, you also can precise more parameters to the docker run command.
func RunWithFlag(image, name string, flags []string, args ...string) error {
	arg := []string{"run", "--rm", "--name", name}
	arg = append(arg, args...)
	arg = append(arg, image)
	if len(flags) > 0 {
		arg = append(arg, flags...)
	}

	cmd := exec.Command("docker", arg...) /* #nosec  G204 */

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command \"%s\" fail.\nOutput : %s .\nError : %w", cmd.String(), string(out), err)
	}
	return nil
}

// Down container with the given name. This function should NOT be called if you launched the container with the other functions of this package.
func Down(name string) error {
	// #nosec
	cmd := exec.Command("docker", "stop", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command \"%s\" fail.\nOutput : %s .\nError : %w", cmd.String(), string(out), err)
	}
	return nil
}

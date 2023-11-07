package test

import "os"

var (
	// WorkDir where the root of the project.
	WorkDir       = os.Getenv("GOPATH") + "/src/github.com/Aloe-Corporation/mongodb"
	DockerCompose = WorkDir + "/test/deployment/docker-compose.yaml"
)

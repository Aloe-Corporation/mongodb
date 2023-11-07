# docker

This project is a module that allows to manage docker container

## Usage

### Interface

`docker.Manager` is an interface to create your docker manager

```go
// Manager is an interface for create a manager of docker container
type Manager interface {
	// Up the container
	Up() error
	// Down the container
	Down() error
}
```

### Compose

`docker.Compose` is a struct for manage docker-compose. It implement `docker.Manager` interface

```go
dc := docker.Compose{PathFile: "path/docker-compose.yaml"}

err := dc.Up()
if err != nil {
	return fmt.Errorf("fail to up docker-compose: %w", err)
}

err := dc.Down()
if err != nil {
	return fmt.Errorf("fail to down docker-compose: %w", err)
}
```

### Basique

```go
// func Run(image, name string, args ...string) error
err := docker.Run("nginx", "docker_run_nginx", "-d")
if err := nil {
    return fmt.Errorf("Fail to run: %w ", err)
}
```

```go
// func RunWithFlag(image, name string, flags []string, args ...string) error
flags := []string{"-v"}

err := RunWithFlag("nginx", "docker_run_nginx_flag", flags, "-d")
if err := nil {
    return fmt.Errorf("Fail to run: %w ", err)
}

```

```go
// func Down(name string) error
err = Down("tdocker_run_nginx")
if err := nil {
    return fmt.Errorf("Fail to down: %w ", err)
}
```

## Test

- `make test`

`test` folder contain docker-compose file

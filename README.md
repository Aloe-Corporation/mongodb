# mongodb

This project is a module for MongoDB Connector.

## Example to use
### Configuration
The `mongodb.Conf` use a YAML tags, it's easy to load PostgreSQL config with configuration file in your project
```go
type Conf struct {
	DB       string `yaml:"db"`
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"`
}
```

### Create new connector
For create new MongoDB Connector use this function with as configuration the structure `mongodb.FactoryConnector(c mongodb.Conf) (*mongodb.Connector, error)` and try connection with `mongodb.Connector.TryConnection() err`
```go
var config := mongodb.Conf{
	DB:       "my_database",
	Addr:     "mongodb://localhost",
	Port:     27017,
	Username: "user",
	Password: "pass",
	Timeout:  10,
}

md, err = mongodb.FactoryConnector(config)
if err != nil {
	return fmt.Errorf("fail to init MongoDB connector: %w", err)
}

err = md.TryConnection()
if err != nil {
	return fmt.Errorf("fail to ping MongoDB: %w", err)
}

```

## Test
- `make test`

## Documentation
- `make godoc`
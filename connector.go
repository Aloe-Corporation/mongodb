package mongodb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Conf contains all information to connect to a MongoDB server.
type Conf struct {
	DB         string            `mapstructure:"db" yaml:"db"`                         // Name of the database.
	Host       string            `mapstructure:"host" yaml:"host"`                     // URL to reach the mongoDB server.
	Port       int               `mapstructure:"port,omitempty" yaml:"port,omitempty"` // Optionnal port, if set to 0 it won't be processed.
	Username   string            `mapstructure:"username" yaml:"username"`             // Credential to authenticate to the db.
	Password   string            `mapstructure:"password" yaml:"password"`             // Credential to authenticate to the db.
	AuthSource string            `mapstructure:"auth_source" yaml:"auth_source"`       // Database to check authentication
	Timeout    int               `mapstructure:"timeout" yaml:"timeout"`               // Connection timeout in seconds
	Options    map[string]string `mapstructure:"options" yaml:"options"`               // List of connection options
}

// Connector is the connector used to communicate with MongoDB database server.
// It embeds a native mongo.Client so it can be used as is and is supercharged with
// additionnal methods.
type Connector struct {
	*mongo.Client
	DB          string
	Collections map[string]*mongo.Collection
}

// Collection returns the  *mongo.Collection identified its name.
// If the specified collections doesn't exists on con.Collections map
// then add it.
func (con *Connector) Collection(collectionName string) *mongo.Collection {
	if con.Collections == nil {
		con.Collections = make(map[string]*mongo.Collection)
	}

	if _, ok := con.Collections[collectionName]; !ok {
		con.Collections[collectionName] = con.Database(con.DB).Collection(collectionName)
	}

	return con.Collections[collectionName]
}

// TryConnection tests ping, it end if the ping is a success or timeout.
func (con *Connector) TryConnection() error {
	if err := con.Ping(context.Background(), nil); err != nil {
		return fmt.Errorf("fail to ping mongo: %w", err)
	}

	return nil
}

// FactoryConnector instanciates a new *Connector with the given params.
func FactoryConnector(c Conf) (*Connector, error) {
	connectionURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority%s", c.Username, c.Password, c.Host, c.AuthSource, buildOptions(c.Options))
	if c.Port != 0 {
		connectionURI = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?retryWrites=true&w=majority", c.Username, c.Password, c.Host, c.Port, c.AuthSource)

	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().ApplyURI(connectionURI).SetServerAPIOptions(serverAPI)
	err := clientOptions.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid client options: %w", err)
	}

	timeout := time.Second * time.Duration(c.Timeout)

	clientOptions.ServerSelectionTimeout = &timeout

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("fail to create mongo client: %w", err)
	}

	con := &Connector{
		DB:          c.DB,
		Client:      client,
		Collections: make(map[string]*mongo.Collection),
	}
	return con, nil
}

func buildOptions(options map[string]string) string {
	if len(options) == 0 {
		return ""
	}

	mergedOptions := []string{}
	for name, value := range options {
		mergedOptions = append(mergedOptions, name+"="+value)
	}

	return "&" + strings.Join(mergedOptions, "&")
}

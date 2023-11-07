package mongodb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Conf structure to open the connection to the database.
type Conf struct {
	DB         string `yaml:"db"`          // name of the database.
	Addr       string `yaml:"addr"`        // Address of the mongoDB.
	Username   string `yaml:"username"`    // credential to authenticate to the db.
	Password   string `yaml:"password"`    // credential to authenticate to the db.
	AuthSource string `yaml:"auth_source"` // Database to check authentication
	Port       int    `yaml:"port"`        // Port of the mongoDB
	Timeout    int    `yaml:"timeout"`
}

// Connector is the connector used to communicate with MongoDB database server.
type Connector struct {
	*mongo.Client
	DB string
}

func (con *Connector) Collection(collectionName string) *mongo.Collection {
	return con.Database(con.DB).Collection(collectionName)
}

// TryConnection test ping, it end if the ping is a success or timeout.
func (con *Connector) TryConnection() error {
	if err := con.Ping(context.Background(), nil); err != nil {
		return fmt.Errorf("fail to ping mongo: %w", err)
	}

	return nil
}

// FactoryConnector instanciates a new *Connector with the given params.
func FactoryConnector(c Conf) (*Connector, error) {
	credential := new(options.Credential)
	credential.Username = c.Username
	credential.Password = c.Password
	credential.AuthSource = c.AuthSource

	// Set client options
	clientOptions := options.Client().ApplyURI(c.Addr + ":" + strconv.Itoa(c.Port) + "/" + c.DB).SetAuth(*credential)
	timeout := time.Second * time.Duration(c.Timeout)

	clientOptions.ServerSelectionTimeout = &timeout

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("fail to create mongo client: %w", err)
	}

	con := &Connector{
		DB:     c.DB,
		Client: client,
	}
	return con, nil
}

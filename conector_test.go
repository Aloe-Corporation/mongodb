package mongodb

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/Aloe-Corporation/docker"
	"github.com/Aloe-Corporation/mongodb/test"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Up docker-compose
	dc := docker.Compose{PathFile: test.DockerCompose}
	err := dc.Up()
	if err != nil {
		_ = dc.Down()
		panic(fmt.Errorf("fail to up docker-compose: %w", err))
	}

	if err != nil {
		_ = dc.Down()
		panic(err)
	}

	r := m.Run()

	// Down docker-compose
	err = dc.Down()
	if err != nil {
		panic(fmt.Errorf("fail to down docker-compose: %w", err))
	}

	os.Exit(r)
}

type connectorCollectionTestData struct {
	Config         Conf
	CollectionName string
}

var connectorCollectionTestCases = [...]connectorCollectionTestData{
	{
		Config: Conf{
			DB:         "mongo",
			Addr:       "mongodb://localhost",
			Port:       49998,
			Username:   "user",
			Password:   "pass",
			AuthSource: "admin",
			Timeout:    10,
		},
		CollectionName: "user",
	},
}

func TestConnectorCollection(t *testing.T) {
	for i, testCase := range connectorCollectionTestCases {
		t.Run("Case "+strconv.Itoa(i), func(t *testing.T) {
			connector, err := FactoryConnector(testCase.Config)
			assert.NoError(t, err)

			collection := connector.Collection(testCase.CollectionName)
			assert.NotNil(t, collection)
		})
	}
}

type connectorTryConnectionTestData struct {
	Config     Conf
	ShouldFail bool
}

var connectorTryConnectionTestCases = [...]connectorTryConnectionTestData{
	{
		Config: Conf{
			DB:         "mongo",
			Addr:       "mongodb://localhost",
			Port:       49998,
			Username:   "user",
			Password:   "pass",
			AuthSource: "admin",
			Timeout:    10,
		},
		ShouldFail: false,
	},
	{ // Fail case with wrong auth source
		Config: Conf{
			DB:         "mongo",
			Addr:       "mongodb://localhost",
			Port:       49998,
			Username:   "user",
			Password:   "pass",
			AuthSource: "wrong",
			Timeout:    10,
		},
		ShouldFail: true,
	},
	{
		Config: Conf{
			DB:         "bad",
			Addr:       "mongodb://badaddr",
			Port:       38017,
			Username:   "user",
			Password:   "pass",
			AuthSource: "admin",
			Timeout:    10,
		},
		ShouldFail: true,
	},
}

func TestConnectorTryConnection(t *testing.T) {
	for i, testCase := range connectorTryConnectionTestCases {
		t.Run("Case "+strconv.Itoa(i), func(t *testing.T) {
			connector, err := FactoryConnector(testCase.Config)
			assert.NoError(t, err)
			err = connector.TryConnection()
			if testCase.ShouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, connector)
			}
		})
	}
}

type factoryConnectorTestData struct {
	Config     Conf
	ShouldFail bool
}

var factoryConnectorTestCases = []factoryConnectorTestData{
	{
		Config: Conf{
			DB:         "mongo",
			Addr:       "mongodb://localhost",
			Port:       27017,
			Username:   "user",
			Password:   "pass",
			AuthSource: "admin",
			Timeout:    5,
		},
		ShouldFail: false,
	},
	{
		Config: Conf{
			DB:         "mongo",
			Addr:       "mongodb://localhost::::::",
			Port:       27017,
			Username:   "user",
			Password:   "pass",
			AuthSource: "admin",
			Timeout:    5,
		},
		ShouldFail: true,
	},
}

func TestFactoryConnector(t *testing.T) {
	for i, testCase := range factoryConnectorTestCases {
		t.Run("Case "+strconv.Itoa(i), func(t *testing.T) {
			connector, err := FactoryConnector(testCase.Config)
			if testCase.ShouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, connector)
			}
		})
	}
}

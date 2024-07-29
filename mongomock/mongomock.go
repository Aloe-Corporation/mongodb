package mongomock

import (
	"github.com/Aloe-Corporation/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var _ mongodb.IConnector = (*MockConnector)(nil)

type MockConnector struct {
	MockClient *mtest.T
}

func (mc *MockConnector) TryConnection() error {
	return nil
}

func (mc *MockConnector) Collection(name string) *mongo.Collection {
	return mc.MockClient.Coll
}

func (mc *MockConnector) GetClient() *mongo.Client {
	return mc.MockClient.Client
}

func New(t *mtest.T) *MockConnector {
	return &MockConnector{
		MockClient: t,
	}
}

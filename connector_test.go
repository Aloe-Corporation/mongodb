package mongodb

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	defaultConfig = Conf{
		DB:         os.Getenv("MONGO_DB_NAME"),
		Host:       os.Getenv("MONGO_DB_HOST"),
		Username:   os.Getenv("MONGO_DB_USERNAME"),
		Password:   os.Getenv("MONGO_DB_PASSWORD"),
		AuthSource: os.Getenv("MONGO_DB_AUTH_SOURCE"),
		Port:       0,
		Timeout:    10,
	}
	defaultClient *Connector
)

func TestMain(m *testing.M) {
	var err error
	// instanciate a client, please verify that the env is set before running this tests.
	defaultClient, err = FactoryConnector(defaultConfig)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = defaultClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := m.Run()

	os.Exit(r)
}

func TestConnector_Collection(t *testing.T) {
	type fields struct {
		Client      *mongo.Client
		DB          string
		Collections map[string]*mongo.Collection
	}
	type args struct {
		collectionName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Success case, with not nil Collection",
			fields: fields{
				Client:      defaultClient.Client,
				DB:          "testing",
				Collections: make(map[string]*mongo.Collection),
			},
			args: args{
				collectionName: "test",
			},
		},
		{
			name: "Success case, with nil Collection",
			fields: fields{
				Client:      defaultClient.Client,
				DB:          "testing",
				Collections: nil,
			},
			args: args{
				collectionName: "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := &Connector{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			got := con.Collection(tt.args.collectionName)
			if got == nil {
				t.Errorf("Connector.Collection() = %v, should not be nil", got)
			}
		})
	}
}

func TestConnector_TryConnection(t *testing.T) {
	type fields struct {
		Client *mongo.Client
		DB     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Succes case",
			fields: fields{
				Client: defaultClient.Client,
				DB:     "testing",
			},
			wantErr: false,
		},
		{
			name: "Fail case: invalid client",
			fields: fields{
				Client: &mongo.Client{},
				DB:     "testing",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := &Connector{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			if err := con.TryConnection(); (err != nil) != tt.wantErr {
				t.Errorf("Connector.TryConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFactoryConnector(t *testing.T) {
	type args struct {
		c Conf
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success case",
			args: args{
				c: defaultConfig,
			},
			wantErr: false,
		},
		{
			name: "Success case with port",
			args: args{
				c: Conf{
					DB:         defaultClient.DB,
					Host:       defaultConfig.Host,
					Port:       27017,
					Username:   defaultConfig.Username,
					Password:   defaultConfig.Password,
					AuthSource: defaultConfig.AuthSource,
					Timeout:    defaultConfig.Timeout,
				},
			},
			wantErr: false,
		},
		{
			name: "Fail case: wrong Host",
			args: args{
				c: Conf{
					DB:         defaultClient.DB,
					Host:       "unknown",
					Username:   defaultConfig.Username,
					Password:   defaultConfig.Password,
					AuthSource: defaultConfig.AuthSource,
					Timeout:    defaultConfig.Timeout,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FactoryConnector(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("FactoryConnector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && got == nil {
				t.Errorf("FactoryConnector() error = %v, client should not be nil", err)
				return
			}

		})
	}
}

func Test_buildOptions(t *testing.T) {

	tests := []struct {
		name    string
		options map[string]string
		want    string
	}{
		{
			name: "Multiple options provided",
			options: map[string]string{
				"opt1": "va1",
				"opt2": "va2",
			},
			want: "&opt1=va1&opt2=va2",
		},
		{
			name: "One option provided",
			options: map[string]string{
				"opt1": "va1",
			},
			want: "&opt1=va1",
		},
		{
			name:    "Empty",
			options: map[string]string{},
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildOptions(tt.options); got != tt.want {
				t.Errorf("buildOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

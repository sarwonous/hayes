package elastic

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/olivere/elastic.v5"
)

var elastics = make(map[string]*Client, 0)

// Client elastic
type Client struct {
	Client *elastic.Client
	DSN    string
}

// Connect retrieve elastic client connection
func Connect(id string) (*Client, error) {
	if e, ok := elastics[id]; ok {
		return e, nil
	}

	return newConnection(id)
}

// Shutdown close all established elastic connections
func Shutdown() {
	for _, c := range elastics {
		c.Client.Stop()
	}
}

func newConnection(id string) (*Client, error) {
	elasticConfig := viper.GetStringMap("elastic")
	if _, ok := elasticConfig[id]; !ok {
		return nil, fmt.Errorf("elastic configuration for [%s] does not exists", id)
	}
	config := elasticConfig[id].(map[string]string)
	e, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config["dsn"]),
	)

	if err != nil {
		return nil, err
	}

	elastics[id] = &Client{
		Client: e,
		DSN:    config["dsn"],
	}

	return elastics[id], nil
}

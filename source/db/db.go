package db

import (
	"database/sql"
	// using mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type database struct {
	Node      string
	Status    int
	master    *sql.DB
	slave     *sql.DB
	masterURL string
	slaveURL  string
	maxOpen   int
}

var connectedDB = make(map[string]*database, 0)
var activeNode string

func Init() {
	config := viper.GetStringMap("database")
	nodes := config["nodes"].(map[string]interface{})
	for key, node := range nodes {
		dbnode := node.(map[string]interface{})
		connectedDB[key] = &database{
			Node:      key,
			Status:    cast.ToInt(dbnode["status"]),
			masterURL: cast.ToString(dbnode["master"]),
			slaveURL:  cast.ToString(dbnode["slave"]),
			maxOpen:   cast.ToInt(dbnode["max_open_connection"]),
		}
		if cast.ToInt(dbnode["status"]) == 1 {
			activeNode = key
		}
	}
}

// Master get master Db
func Master() (*sql.DB, error) {
	if connectedDB[activeNode] == nil {
		master, err := Open(connectedDB[activeNode].masterURL)
		if err != nil {
			return nil, err
		}
		connectedDB[activeNode].master = master
		return connectedDB[activeNode].master, nil
	}
	return connectedDB[activeNode].master, nil
}

func Slave() (*sql.DB, error) {
	if connectedDB[activeNode] == nil {
		slave, err := Open(connectedDB[activeNode].slaveURL)
		if err != nil {
			return nil, err
		}
		connectedDB[activeNode].slave = slave
		return connectedDB[activeNode].slave, nil
	}
	return connectedDB[activeNode].slave, nil
}

// Open db
func Open(config string) (*sql.DB, error) {
	return nil, nil
}

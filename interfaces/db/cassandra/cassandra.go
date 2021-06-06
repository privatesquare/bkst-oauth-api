package cassandra

import (
	"github.com/gocql/gocql"
	"io/ioutil"
	"log"
)

var (
	cluster *gocql.ClusterConfig
	session *gocql.Session
)

type Cfg struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBKeySpace string `mapstructure:"DB_KEYSPACE"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

func NewCluster(cfg Cfg) error {
	cluster = gocql.NewCluster(cfg.DBHost)
	cluster.Keyspace = cfg.DBKeySpace
	cluster.Consistency = gocql.Quorum
	cluster.Logger = log.New(ioutil.Discard, "", 0)
	return nil
}

func OpenSession() error {
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		return err
	}
	return nil
}

func CloseSession() {
	session.Close()
}


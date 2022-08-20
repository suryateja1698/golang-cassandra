package db

import (
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
)

var Tables = []string{
	`CREATE TABLE IF NOT EXISTS player_details 
	(id text PRIMARY KEY, first_name text, last_name text, age int, country text, position text, joined_year int, joined_month int,
	 created_at TIMESTAMP, updated_at TIMESTAMP );`,
}

var Session *gocql.Session

func Init() *gocql.Session {
	if Session != nil {
		return Session
	}

	var err error
	keyspace := os.Getenv("CASSANDRA_KEYSPACE")
	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_CLUSTER"))

	cluster.Keyspace = keyspace
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Printf("cannot create session, :%w", err)
	} else {
		err = MigrateTables()
		if err != nil {
			log.Println(err.Error())
		}
	}
	return Session
}

func MigrateTables() error {
	for _, createTable := range Tables {
		query := Session.Query(createTable)

		err := query.Exec()

		if err != nil {
			return fmt.Errorf("cannot create tables :%w", err)
		}
	}
	return nil
}

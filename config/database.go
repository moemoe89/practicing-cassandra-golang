//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package config

import (
	"fmt"

	"github.com/gocql/gocql"
)

// InitDB will create a variable that represent the gocql.Session
func InitDB() (*gocql.Session, error) {
	cluster := gocql.NewCluster(Configuration.Cassandra.Addr)
	cluster.Port = Configuration.Cassandra.Port
	cluster.Keyspace = Configuration.Cassandra.Keyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: Configuration.Cassandra.Username,
		Password: Configuration.Cassandra.Password,
	}
	sess, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping connection to cassandra: %s", err.Error())
	}

	return sess, nil
}

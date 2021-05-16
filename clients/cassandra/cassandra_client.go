package cassandra

import (
	// "fmt"
	"github.com/gocql/gocql"
	// "github.com/abhishekbhr/bookstore-oauth-api/utils/errors"
)

var (
	session *gocql.Session
)


func init() {
	//connect to cassandra cluster
	cluster := gocql.NewCluster("172.20.0.4")
	cluster.Keyspace 	= "oauth"
	cluster.Consistency = gocql.Quorum
	var err error

	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}


func GetSession() *gocql.Session{	
	return session 
}
package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/grafana/metrictank/util"
	log "github.com/sirupsen/logrus"
)

// EnsureTableExists checks if the specified table exists or not. If it does not exist and the
// create-keyspace flag is true, then it will create it, if it doesn't exist and the create-keyspace
// flag is false, then it will return an error. If the table exists then it just returns nil
// session:     cassandra session
// schemaFile:  file containing table definition
// entryName:   identifier of the schema within the file
// tableName:   name of the table in cassandra
func EnsureTableExists(session *gocql.Session, createKeyspace bool, keyspace, schemaFile, entryName, tableName string) error {
	var err error
	tableSchema := util.ReadEntry(schemaFile, entryName).(string)

	if createKeyspace {
		log.Infof("cassandra-idx: ensuring that table %s exists.", tableName)
		err = session.Query(fmt.Sprintf(tableSchema, keyspace, tableName)).Exec()
		if err != nil {
			return fmt.Errorf("cassandra-idx: failed to initialize cassandra table: %s", err)
		}
	} else {
		var keyspaceMetadata *gocql.KeyspaceMetadata
		keyspaceMetadata, err = session.KeyspaceMetadata(keyspace)
		if err != nil {
			return fmt.Errorf("cassandra-idx: failed to read cassandra tables: %s", err)
		}
		if _, ok := keyspaceMetadata.Tables[tableName]; !ok {
			return fmt.Errorf("cassandra-idx: table %s does not exist", tableName)
		}
	}
	return nil
}

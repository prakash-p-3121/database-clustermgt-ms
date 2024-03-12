package database_shard_service

import (
	"database/sql"
	serviceImpl "github.com/prakash-p-3121/directory-database-lib/service/database_shard_service/impl"
)

func NewDatabaseShardService(databaseConnection *sql.DB) DatabaseShardService {
	return &serviceImpl.DatabaseShardServiceImpl{DatabaseConnection: databaseConnection}
}

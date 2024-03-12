package database_shard_repository

import (
	"database/sql"
	impl "github.com/prakash-p-3121/directory-database-lib/repository/database_shard_repository/impl"
)

func NewDatabaseShardRepository(databaseConnection *sql.DB) DatabaseShardRepository {
	return &impl.DatabaseShardRepositoryImpl{DatabaseConnection: databaseConnection}
}

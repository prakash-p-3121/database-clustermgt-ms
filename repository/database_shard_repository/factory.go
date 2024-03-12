package database_shard_repository

import (
	"database/sql"
	impl "github.com/prakash-p-3121/database-clustermgt-ms/repository/database_shard_repository/impl"
)

func NewDatabaseShardRepository(databaseConnection *sql.DB) DatabaseShardRepository {
	return &impl.DatabaseShardRepositoryImpl{DatabaseConnection: databaseConnection}
}

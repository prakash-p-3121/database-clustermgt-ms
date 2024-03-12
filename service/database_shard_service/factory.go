package database_shard_service

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/repository/database_shard_repository"
	serviceImpl "github.com/prakash-p-3121/directory-database-lib/service/database_shard_service/impl"
)

func NewDatabaseShardService(databaseConnection *sql.DB) DatabaseShardService {
	shardRepo := database_shard_repository.NewDatabaseShardRepository(databaseConnection)
	return &serviceImpl.DatabaseShardServiceImpl{DatabaseShardRepository: shardRepo}
}

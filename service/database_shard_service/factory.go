package database_shard_service

import (
	"database/sql"
	"github.com/prakash-p-3121/database-clustermgt-ms/repository/database_shard_repository"
	serviceImpl "github.com/prakash-p-3121/database-clustermgt-ms/service/database_shard_service/impl"
)

func NewDatabaseShardService(databaseConnection *sql.DB) DatabaseShardService {
	shardRepo := database_shard_repository.NewDatabaseShardRepository(databaseConnection)
	return &serviceImpl.DatabaseShardServiceImpl{DatabaseShardRepository: shardRepo}
}

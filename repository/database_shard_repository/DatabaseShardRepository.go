package database_shard_repository

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseShardRepository interface {
	CreateShard(shard *model.DatabaseShard) (int64, errorlib.AppError)
	FindShardByID(shardID int64) (*model.DatabaseShard, errorlib.AppError)
	FindShardsByClusterID(clusterID int64) ([]*model.DatabaseShard, errorlib.AppError)
}

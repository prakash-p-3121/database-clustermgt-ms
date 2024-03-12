package database_shard_service

import (
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseShardService interface {
	CreateShard(shard *model.DatabaseShard) (int64, errorlib.AppError)
	FindShardByID(id int64) (*model.DatabaseShard, errorlib.AppError)
}

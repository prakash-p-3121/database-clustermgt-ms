package database_shard_service

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseShardService interface {
	CreateShard(shard *model.DatabaseShardCreateReq) (int64, errorlib.AppError)
	FindShardByID(id int64) (*model.DatabaseShard, errorlib.AppError)
}

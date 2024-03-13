package impl

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/database-clustermgt-ms/repository/database_shard_repository"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseShardServiceImpl struct {
	DatabaseShardRepository database_shard_repository.DatabaseShardRepository
}

func (service *DatabaseShardServiceImpl) CreateShard(shard *model.DatabaseShardCreateReq) (int64, errorlib.AppError) {
	return service.DatabaseShardRepository.CreateShard(shard)
}

func (service *DatabaseShardServiceImpl) FindShardByID(id int64) (*model.DatabaseShard, errorlib.AppError) {
	return service.DatabaseShardRepository.FindShardByID(id)
}

package impl

import (
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/directory-database-lib/repository/database_shard_repository"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseShardServiceImpl struct {
	DatabaseShardRepository database_shard_repository.DatabaseShardRepository
}

func (service *DatabaseShardServiceImpl) CreateShard(shard *model.DatabaseShard) (int64, errorlib.AppError) {
	return service.DatabaseShardRepository.CreateShard(shard)
}

func (service *DatabaseShardServiceImpl) FindShardByID(id int64) (*model.DatabaseShard, errorlib.AppError) {
	return service.DatabaseShardRepository.FindShardByID(id)
}

package impl

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/model"
	repository2 "github.com/prakash-p-3121/directory-database-lib/repository"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseShardServiceImpl struct {
	DatabaseConnection *sql.DB
}

func (service *DatabaseShardServiceImpl) CreateShard(ipAddr string) (*model.DatabaseShard, errorlib.AppError) {
	repository := repository2.NewDatabaseShardRepository(service.DatabaseConnection)
	return repository.CreateShard(ipAddr)
}

func (service *DatabaseShardServiceImpl) ReadShardByID(id int64) (*model.DatabaseShard, errorlib.AppError) {
	repository := repository2.NewDatabaseShardRepository(service.DatabaseConnection)
	return repository.ReadShardByID(id)
}

func (service *DatabaseShardServiceImpl) UpdateShard(id int64, ipAddr string) (int64, errorlib.AppError) {
	repository := repository2.NewDatabaseShardRepository(service.DatabaseConnection)
	return repository.UpdateShard(id, ipAddr)
}

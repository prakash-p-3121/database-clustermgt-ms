package impl

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/directory-database-lib/repository"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseClusterServiceImpl struct {
	DatabaseConnection *sql.DB
}

func (service *DatabaseClusterServiceImpl) CreateCluster(tableName string,
	shardList []*model.DatabaseShard) (*model.DatabaseCluster, errorlib.AppError) {
	clusterRepo := repository.NewDatabaseClusterRepository(service.DatabaseConnection)
	return clusterRepo.CreateCluster(tableName, shardList)
}

func (service *DatabaseClusterServiceImpl) ReadClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError) {
	clusterRepo := repository.NewDatabaseClusterRepository(service.DatabaseConnection)
	return clusterRepo.ReadClusterByID(id)
}

func (service *DatabaseClusterServiceImpl) ReadCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError) {
	clusterRepo := repository.NewDatabaseClusterRepository(service.DatabaseConnection)
	return clusterRepo.ReadCurrentWriteClusterByTableName(tableName)
}

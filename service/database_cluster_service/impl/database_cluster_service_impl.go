package impl

import (
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/directory-database-lib/repository/database_cluster_repository"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseClusterServiceImpl struct {
	DatabaseClusterRepository database_cluster_repository.DatabaseClusterRepository
}

func (service *DatabaseClusterServiceImpl) CreateCluster(tableName string,
	shardList []*model.DatabaseShard) (*model.DatabaseCluster, errorlib.AppError) {
	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.CreateCluster(tableName, shardList)
}

func (service *DatabaseClusterServiceImpl) ReadClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError) {
	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.ReadClusterByID(id)
}

func (service *DatabaseClusterServiceImpl) FindCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError) {
	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.FindCurrentWriteClusterByTableName(tableName)
}

func (service *DatabaseClusterServiceImpl) FindCurrentWriteShardByTableName(tableName string,
	id string) (*model.DatabaseShard, errorlib.AppError) {
	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.FindCurrentWriteShardByTableName(tableName, id)
}

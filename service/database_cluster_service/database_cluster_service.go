package database_cluster_service

import (
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseClusterService interface {
	CreateCluster(tableName string, shardList []*model.DatabaseShard) (*model.DatabaseCluster, errorlib.AppError)
	ReadClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError)
	FindCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError)
	FindCurrentWriteShardByTableName(tableName string, id string) (*model.DatabaseShard, errorlib.AppError)
}

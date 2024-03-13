package database_cluster_repository

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseClusterRepository interface {
	CreateCluster(tableName string, shardList []int64) (*model.DatabaseCluster, errorlib.AppError)
	FindClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError)
	FindCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError)
	FindCurrentWriteShardByTableName(tableName, id string) (*model.DatabaseShard, errorlib.AppError)
}

package database_cluster_repository

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseClusterRepository interface {
	CreateCluster(tableName string, shardingType uint8, shardList []int64) (*model.DatabaseCluster, errorlib.AppError)
	FindClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError)
	FindClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError)
	FindShardByNumber(tableName, id string) (*model.DatabaseShard, errorlib.AppError)
	FindShardByChar(tableName string, id rune) (*model.DatabaseShard, errorlib.AppError)
	FindAllShardsByTable(tableName string) ([]*model.DatabaseShard, errorlib.AppError)
}

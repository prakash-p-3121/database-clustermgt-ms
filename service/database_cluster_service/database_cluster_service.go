package database_cluster_service

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseClusterService interface {
	CreateCluster(req *model.DatabaseClusterCreateReq) (*model.DatabaseCluster, errorlib.AppError)
	ReadClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError)
	FindClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError)
	FindShard(tableName string, id string) (*model.DatabaseShard, errorlib.AppError)
	FindAllShardsByTable(tableName string) ([]*model.DatabaseShard, errorlib.AppError)
}

package repository

import (
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseClusterRepository interface {
	CreateCluster(tableName string, shardList []*model.DatabaseShard) (*model.DatabaseCluster, errorlib.AppError)
	ReadClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError)
	ReadCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError)
}

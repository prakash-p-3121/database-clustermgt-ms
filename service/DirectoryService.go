package service

import (
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/errorlib"
)

type UserDirectoryService interface {
	LookUpByEmailID(emailID string) (*model.EmailIDLookUpResp, errorlib.AppError)
	LookUpByUserID(userID string) (*model.UserIDLookUpResp, errorlib.AppError)
}

type DatabaseShardService interface {
	CreateShard(ipAddr string) (*model.DatabaseShard, errorlib.AppError)
	ReadShardByID(id int64) (*model.DatabaseShard, errorlib.AppError)
	UpdateShard(id int64, ipAddr string) (int64, errorlib.AppError) /* returns affectedRows, AppError  */
}

type DatabaseClusterService interface {
	CreateCluster(tableName string, shardList []*model.DatabaseShard) (*model.DatabaseCluster, errorlib.AppError)
	ReadClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError)
	ReadCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError)
}

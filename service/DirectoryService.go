package service

import (
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/errorlib"
)

type UserDirectoryService interface {
	LookUpByEmailID(emailID string) (*model.EmailIDLookUpResp, errorlib.AppError)
	LookUpByUserID(userID string) (*model.UserIDLookUpResp, errorlib.AppError)
}

type ShardServerService interface {
	CreateShard(ipAddr string) (*model.DatabaseShard, errorlib.AppError)
	ReadShardByID(id uint64) (*model.DatabaseShard, errorlib.AppError)
	UpdateShard(id uint64, ipAddr string) (uint, errorlib.AppError) /* returns affectedRows, AppError  */
}

type ClusterService interface {
	CreateCluster(tableName string, shardList []*model.DatabaseShard) (*model.DatabaseShard, errorlib.AppError)
	//UpdateCurrentWriteCluster(tableName string, clusterID uint64) (uint, errorlib.AppError)
	ReadClusterByID(id uint64) (*model.DatabaseCluster, errorlib.AppError)
	ReadCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError)
}

package repository

import (
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseShardRepository interface {
	CreateShard(ipAddr string) (*model.DatabaseShard, errorlib.AppError)
	ReadShardByID(id int64) (*model.DatabaseShard, errorlib.AppError)
	UpdateShard(id int64, ipAddr string) (int64, errorlib.AppError)
}

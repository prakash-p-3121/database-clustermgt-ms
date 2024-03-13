package database_shard_controller

import (
	"github.com/prakash-p-3121/restlib"
)

type DatabaseShardController interface {
	CreateShard(restCtx restlib.RestContext)
	FindShardByID(restCtx restlib.RestContext)
}

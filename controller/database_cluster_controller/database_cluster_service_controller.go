package database_cluster_controller

import (
	"github.com/prakash-p-3121/restlib"
)

type DatabaseClusterController interface {
	FindCurrentWriteShardByTableName(restCtx restlib.RestContext)
}

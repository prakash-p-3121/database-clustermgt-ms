package database_cluster_controller

import (
	"github.com/prakash-p-3121/restlib"
)

type DatabaseClusterController interface {
	CreateCluster(restCtx restlib.RestContext)
	FindCurrentWriteShard(restCtx restlib.RestContext)
}

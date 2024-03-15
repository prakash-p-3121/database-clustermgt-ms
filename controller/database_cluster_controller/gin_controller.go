package database_cluster_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/restlib"
)

func FindShard(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewDatabaseClusterController()
	controller.FindShard(ginRestCtx)
}

func CreateCluster(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewDatabaseClusterController()
	controller.CreateCluster(ginRestCtx)
}

func FindAllShardsByTable(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewDatabaseClusterController()
	controller.FindAllShardsByTable(ginRestCtx)
}

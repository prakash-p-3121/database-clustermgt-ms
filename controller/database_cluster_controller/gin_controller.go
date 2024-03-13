package database_cluster_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/restlib"
)

func FindCurrentWriteShard(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewDatabaseClusterController()
	controller.FindCurrentWriteShard(ginRestCtx)
}

func CreateCluster(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewDatabaseClusterController()
	controller.CreateCluster(ginRestCtx)
}

func FindAllShardsByTableName(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewDatabaseClusterController()
	controller.FindAllShardsByTableName(ginRestCtx)
}

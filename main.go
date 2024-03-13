package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/database-clustermgt-ms/controller/database_cluster_controller"
	"github.com/prakash-p-3121/database-clustermgt-ms/controller/database_shard_controller"
	"github.com/prakash-p-3121/database-clustermgt-ms/database"
	"github.com/prakash-p-3121/mysqllib"
)

func main() {

	databaseInst, err := mysqllib.CreateDatabaseConnectionWithRetry("conf/database.toml")
	if err != nil {
		panic(err)
	}
	database.SetDatabaseInstance(databaseInst)

	router := gin.Default()
	routerGroup := router.Group("/database/clustermgt")

	routerGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routerGroup.POST("/database-shard", database_shard_controller.CreateShard)
	routerGroup.POST("/database-cluster", database_cluster_controller.CreateCluster)
	routerGroup.GET("/find-current/write-shard", database_cluster_controller.FindCurrentWriteShard)

	err = router.Run("127.0.0.1:3000")
	if err != nil {
		panic("Error Starting database-clustermgt-ms")
	}
}

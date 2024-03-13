package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/database-clustermgt-ms/controller/database_cluster_controller"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routerGroup := router.Group("/database/clustermgt")
	routerGroup.GET("/find-current/write-shard", database_cluster_controller.FindCurrentWriteShard)

	err := router.Run("127.0.0.1:3000")
	if err != nil {
		panic("Error Starting UserMgtMs")
	}
}

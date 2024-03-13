package database_shard_controller

import (
	"github.com/prakash-p-3121/database-clustermgt-ms/controller/database_shard_controller/impl"
	"github.com/prakash-p-3121/database-clustermgt-ms/database"
	"github.com/prakash-p-3121/database-clustermgt-ms/service/database_shard_service"
)

func NewDatabaseShardController() DatabaseShardController {
	service := database_shard_service.NewDatabaseShardService(database.GetDatabaseInstance())
	return &impl.DatabaseShardControllerImpl{DatabaseShardService: service}
}

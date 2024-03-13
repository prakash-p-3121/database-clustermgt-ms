package database_cluster_controller

import (
	"github.com/prakash-p-3121/database-clustermgt-ms/controller/database_cluster_controller/impl"
	"github.com/prakash-p-3121/database-clustermgt-ms/database"
	"github.com/prakash-p-3121/database-clustermgt-ms/service/database_cluster_service"
)

func NewDatabaseClusterController() DatabaseClusterController {
	service := database_cluster_service.NewDatabaseClusterService(database.GetDatabaseInstance())
	return &impl.DatabaseClusterControllerImpl{DatabaseClusterService: service}
}

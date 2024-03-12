package database_cluster_service

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/repository/database_cluster_repository"
	serviceImpl "github.com/prakash-p-3121/directory-database-lib/service/database_cluster_service/impl"
)

func NewDatabaseClusterService(databaseConnection *sql.DB) DatabaseClusterService {
	clusterRepo := database_cluster_repository.NewDatabaseClusterRepository(databaseConnection)
	return &serviceImpl.DatabaseClusterServiceImpl{DatabaseClusterRepository: clusterRepo}
}

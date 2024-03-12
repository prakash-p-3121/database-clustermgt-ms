package database_cluster_service

import (
	"database/sql"
	serviceImpl "github.com/prakash-p-3121/directory-database-lib/service/database_cluster_service/impl"
)

func NewDatabaseClusterService(databaseConnection *sql.DB) DatabaseClusterService {
	return &serviceImpl.DatabaseClusterServiceImpl{DatabaseConnection: databaseConnection}
}

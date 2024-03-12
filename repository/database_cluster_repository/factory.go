package database_cluster_repository

import (
	"database/sql"
	impl "github.com/prakash-p-3121/directory-database-lib/repository/database_cluster_repository/impl"
)

func NewDatabaseClusterRepository(databaseConnection *sql.DB) DatabaseClusterRepository {
	return &impl.DatabaseClusterRepositoryImpl{DatabaseConnection: databaseConnection}
}

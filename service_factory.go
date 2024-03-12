package directory_database_lib

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/service"
	service_impl "github.com/prakash-p-3121/directory-database-lib/service/impl"
)

func NewUserDirectoryService(databaseConnection *sql.DB) service.UserDirectoryService {
	return &service_impl.UserDirectoryServiceImpl{DatabaseConnection: databaseConnection}
}

func NewDatabaseShardService(databaseConnection *sql.DB) service.DatabaseShardService {
	return &service_impl.DatabaseShardServiceImpl{DatabaseConnection: databaseConnection}
}

func NewDatabaseClusterService(databaseConnection *sql.DB) service.DatabaseClusterService {
	return &service_impl.DatabaseClusterServiceImpl{DatabaseConnection: databaseConnection}
}

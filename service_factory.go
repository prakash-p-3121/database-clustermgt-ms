package directory_database_lib

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/service"
	service_impl "github.com/prakash-p-3121/directory-database-lib/service/impl"
)

func NewUserDirectoryService(databaseConnection *sql.DB) service.UserDirectoryService {
	return &service_impl.UserDirectoryServiceImpl{DatabaseConnection: databaseConnection}
}

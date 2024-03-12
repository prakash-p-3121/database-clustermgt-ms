package user_directory_service

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/service/user_directory_service/impl"
)

func NewUserDirectoryService(databaseConnection *sql.DB) UserDirectoryService {
	return &impl.UserDirectoryServiceImpl{DatabaseConnection: databaseConnection}
}

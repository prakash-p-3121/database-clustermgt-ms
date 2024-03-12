package user_directory_repository

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/repository/user_directory_repository/impl"
)

func NewUserDirectoryRepository(databaseConnection *sql.DB) UserDirectoryRepository {
	return &impl.UserDirectoryRepositoryImpl{DatabaseConnection: databaseConnection}
}

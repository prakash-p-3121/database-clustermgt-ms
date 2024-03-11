package repository

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/repository/impl"
)

func NewUserDirectoryRepository(databaseConnection *sql.DB) UserDirectoryRepository {
	return &impl.UserDirectoryRepositoryImpl{DatabaseConnection: databaseConnection}
}

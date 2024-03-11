package impl

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/directory-database-lib/repository"
	"github.com/prakash-p-3121/errorlib"
)

type UserDirectoryServiceImpl struct {
	DatabaseConnection *sql.DB
}

func (service *UserDirectoryServiceImpl) LookUpByEmailID(emailID string) (*model.EmailIDLookUpResp,
	errorlib.AppError) {
	repo := repository.NewUserDirectoryRepository(service.DatabaseConnection)
	return repo.LookUpByEmailID(emailID)
}

func (service *UserDirectoryServiceImpl) LookUpByUserID(userID string) (*model.UserIDLookUpResp, errorlib.AppError) {
	repo := repository.NewUserDirectoryRepository(service.DatabaseConnection)
	return repo.LookUpByUserID(userID)
}

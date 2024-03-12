package impl

import (
	"database/sql"
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/directory-database-lib/repository/user_directory_repository"
	databaseClusterService "github.com/prakash-p-3121/directory-database-lib/service/database_cluster_service"
	"github.com/prakash-p-3121/errorlib"
)

const (
	userTable = "users"
)

type UserDirectoryServiceImpl struct {
	DatabaseConnection *sql.DB
}

func (userService *UserDirectoryServiceImpl) LookUpByEmailID(emailID string) (*model.EmailIDLookUpResp,
	errorlib.AppError) {
	repo := user_directory_repository.NewUserDirectoryRepository(userService.DatabaseConnection)
	return repo.LookUpByEmailID(emailID)
}

func (userService *UserDirectoryServiceImpl) LookUpByUserID(userID string) (*model.UserIDLookUpResp, errorlib.AppError) {
	repo := user_directory_repository.NewUserDirectoryRepository(userService.DatabaseConnection)
	return repo.LookUpByUserID(userID)
}

func (userService *UserDirectoryServiceImpl) LookUpCurrentWriteShard(userID string) (*model.DatabaseShard, errorlib.AppError) {
	clusterService := databaseClusterService.NewDatabaseClusterService(userService.DatabaseConnection)
	shard, err := clusterService.FindCurrentWriteShardByTableName(userTable, userID)
	if err != nil {
		return nil, err
	}
	return shard, nil
}

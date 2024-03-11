package impl

import (
	"database/sql"
	"errors"
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/errorlib"
)

type UserDirectoryRepositoryImpl struct {
	DatabaseConnection *sql.DB
}

func (repository *UserDirectoryRepositoryImpl) LookUpByEmailID(emailID string) (*model.EmailIDLookUpResp,
	errorlib.AppError) {
	db := repository.DatabaseConnection
	query := `SELECT user_id, cluster_id FROM users_table_directory WHERE email_id=?;`
	row := db.QueryRow(query, emailID)
	var resp model.EmailIDLookUpResp
	err := row.Scan(&resp.UserID, &resp.ClusterID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("emailID=" + emailID + "not-found")
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &resp, nil
}

func (repository *UserDirectoryRepositoryImpl) LookUpByUserID(userID string) (*model.UserIDLookUpResp,
	errorlib.AppError) {
	db := repository.DatabaseConnection
	query := `SELECT email_id, cluster_id FROM users_table_directory WHERE user_id=?;`
	row := db.QueryRow(query, userID)
	var resp model.UserIDLookUpResp
	err := row.Scan(&resp.EmailID, &resp.ClusterID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("userID=" + userID + "not-found")
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &resp, nil
}

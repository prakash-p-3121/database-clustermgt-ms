package impl

import (
	"database/sql"
	"errors"
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/errorlib"
	"strconv"
)

type DatabaseShardRepositoryImpl struct {
	DatabaseConnection *sql.DB
}

func (repository *DatabaseShardRepositoryImpl) CreateShard(ipAddr string) (*model.DatabaseShard, errorlib.AppError) {
	db := repository.DatabaseConnection
	qry := `INSERT INTO database_shards  (ip_address) VALUES (?); `
	result, err := db.Exec(qry, ipAddr)
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	ipAddress := new(string)
	*ipAddress = ipAddr
	return &model.DatabaseShard{ID: &id, IPAddress: ipAddress}, nil
}

func (repository *DatabaseShardRepositoryImpl) ReadShardByID(id int64) (*model.DatabaseShard, errorlib.AppError) {
	db := repository.DatabaseConnection
	qry := `SELECT ip_address from database_shards WHERE id=?;`
	row := db.QueryRow(qry, id)
	var ipAddr string
	err := row.Scan(&ipAddr)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("shard-info-not-found-for" + strconv.FormatInt(id, 10))
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &model.DatabaseShard{ID: &id, IPAddress: &ipAddr}, nil
}

func (repository *DatabaseShardRepositoryImpl) UpdateShard(id int64, ipAddr string) (int64, errorlib.AppError) {
	db := repository.DatabaseConnection
	qry := `UPDATE database_shards SET ip_address =? WHERE id=?;`
	result, err := db.Exec(qry, ipAddr, id)
	if err != nil {
		return 0, errorlib.NewInternalServerError(err.Error())
	}
	affRows, err := result.RowsAffected()
	if err != nil {
		return 0, errorlib.NewInternalServerError(err.Error())
	}
	return affRows, nil
}

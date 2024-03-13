package impl

import (
	"database/sql"
	"errors"
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
	"strconv"
)

type DatabaseShardRepositoryImpl struct {
	DatabaseConnection *sql.DB
}

func (repository *DatabaseShardRepositoryImpl) CreateShard(shardPtr *model.DatabaseShard) (int64, errorlib.AppError) {
	db := repository.DatabaseConnection
	qry := `INSERT INTO database_shards  (ip_address, cluster_id, port, user_name, password, database_name) VALUES (?,?,?,?,?,?); `
	result, err := db.Exec(qry, *shardPtr.IPAddress,
		*shardPtr.ClusterID,
		*shardPtr.Port,
		*shardPtr.UserName,
		*shardPtr.Password,
		*shardPtr.DatabaseName)
	if err != nil {
		return 0, errorlib.NewInternalServerError(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, errorlib.NewInternalServerError(err.Error())
	}
	return id, nil
}

func (repository *DatabaseShardRepositoryImpl) FindShardByID(id int64) (*model.DatabaseShard, errorlib.AppError) {
	db := repository.DatabaseConnection
	qry := `SELECT ip_address, cluster_id, port, user_name, password, database_name from database_shards WHERE id=?;`
	row := db.QueryRow(qry, id)
	var shard model.DatabaseShard
	err := row.Scan(&shard.IPAddress,
		&shard.ClusterID,
		&shard.Port,
		&shard.UserName,
		&shard.Password,
		&shard.DatabaseName)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("shard-info-not-found-for" + strconv.FormatInt(id, 10))
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &shard, nil
}

func (repository *DatabaseShardRepositoryImpl) FindShardsByClusterID(clusterID int64) ([]*model.DatabaseShard, errorlib.AppError) {
	db := repository.DatabaseConnection
	qry := `SELECT id, 
       		ip_address, 
    		cluster_id, 
    		port, 
    		user_name, 
    		password, 
    		database_name, 
    		created_at,
    		updated_at FROM database_shards A WHERE A.cluster_id = ? ORDER BY A.id ASC;`
	rows, err := db.Query(qry, clusterID)
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	shardList := make([]*model.DatabaseShard, 0)
	for rows.Next() {
		var shard model.DatabaseShard
		err := rows.Scan(&shard.ID,
			&shard.IPAddress,
			&shard.ClusterID,
			&shard.Port,
			&shard.UserName,
			&shard.Password,
			&shard.DatabaseName,
			&shard.CreatedAt,
			&shard.UpdatedAt)
		if err != nil {
			return nil, errorlib.NewInternalServerError(err.Error())
		}
		shardList = append(shardList, &shard)
	}
	return shardList, nil
}

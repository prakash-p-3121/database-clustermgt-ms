package impl

import (
	"context"
	"crypto/md5"
	"database/sql"
	"errors"
	"github.com/prakash-p-3121/directory-database-lib/model"
	shardRepo "github.com/prakash-p-3121/directory-database-lib/repository/database_shard_repository"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/mysqllib"
	"strconv"
)

type DatabaseClusterRepositoryImpl struct {
	DatabaseConnection *sql.DB
}

func (repository *DatabaseClusterRepositoryImpl) CreateCluster(tableName string,
	shardList []*model.DatabaseShard) (*model.DatabaseCluster, errorlib.AppError) {
	db := repository.DatabaseConnection
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	cluster, err := repository.createCluster(tx, tableName, shardList)
	if err != nil {
		return nil, errorlib.NewInternalServerError(mysqllib.RollbackTx(tx, err).Error())
	}
	err = tx.Commit()
	if err != nil {
		return nil, errorlib.NewInternalServerError(mysqllib.RollbackTx(tx, err).Error())
	}
	return cluster, nil
}

func (repository *DatabaseClusterRepositoryImpl) createCluster(tx *sql.Tx, tableName string,
	shardList []*model.DatabaseShard) (*model.DatabaseCluster, error) {
	shardSize := int64(len(shardList))
	qry := `INSERT INTO database_clusters (table_name, shard_size) VALUES (?, ?);`
	result, err := tx.Exec(qry, tableName, shardSize)
	if err != nil {
		return nil, err
	}
	clusterID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	clusterPtr := &model.DatabaseCluster{ID: &clusterID, TableName: &tableName, ShardSize: &shardSize}
	err = repository.createClusterToShardRelationship(tx, clusterPtr, shardList)
	if err != nil {
		return nil, err
	}

	qry = `INSERT INTO current_write_clusters (table_name, cluster_id) VALUES (?,?) ON DUPLICATE KEY UPDATE 
    			table_name=VALUES(table_name) AND cluster_id=VALUES(cluster_id);`
	_, err = tx.Exec(qry, tableName, clusterID)
	if err != nil {
		return nil, err
	}
	return clusterPtr, nil
}

func (repository *DatabaseClusterRepositoryImpl) createClusterToShardRelationship(tx *sql.Tx,
	clusterPtr *model.DatabaseCluster,
	shardPtrList []*model.DatabaseShard) error {

	for _, shardPtr := range shardPtrList {
		qry := `UPDATE shards SET cluster_id=? WHEE id=?;`
		_, err := tx.Exec(qry, clusterPtr.ID, shardPtr.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repository *DatabaseClusterRepositoryImpl) ReadClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError) {
	db := repository.DatabaseConnection
	qry := `SELECT id, table_name, shard_size FROM database_clusters where id=?;`
	row := db.QueryRow(qry, id)
	var cluster model.DatabaseCluster
	err := row.Scan(&cluster.ID, &cluster.TableName, &cluster.ShardSize)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("cluster-id-not-found")
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &cluster, nil
}

func (repository *DatabaseClusterRepositoryImpl) FindCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError) {
	db := repository.DatabaseConnection
	qry := `SELECT id, table_name, shard_size FROM database_clusters A INNER JOIN current_write_clusters B ON B.table_name = ? AND B.cluster_id = A.id ;`
	row := db.QueryRow(qry, tableName)
	var cluster model.DatabaseCluster
	err := row.Scan(&cluster.ID, &cluster.TableName, &cluster.ShardSize)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("write-cluster-not-found-for-tableName=" + tableName)
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &cluster, nil
}

func (repository *DatabaseClusterRepositoryImpl) FindCurrentWriteShardByTableName(tableName string,
	id string) (*model.DatabaseShard, errorlib.AppError) {
	cluster, err := repository.FindCurrentWriteClusterByTableName(tableName)
	if err != nil {
		return nil, err
	}
	shardRepoInst := shardRepo.NewDatabaseShardRepository(repository.DatabaseConnection)
	shardList, err := shardRepoInst.FindShardsByClusterID(*cluster.ID)
	if err != nil {
		return nil, err
	}
	if len(shardList) == 0 {
		return nil, errorlib.NewNotFoundError("shards-not-found-for-cluster_id=" + strconv.FormatInt(*cluster.ID, 10))
	}
	shard, err := repository.FindWriteShard(shardList, id)
	if err != nil {
		return nil, err
	}
	return shard, nil
}

func (repository *DatabaseClusterRepositoryImpl) FindWriteShard(shardList []*model.DatabaseShard, id string) (*model.DatabaseShard, errorlib.AppError) {
	md5HashInt64 := repository.computeMD5Hash(id)
	shardNo := int(md5HashInt64 % int64(len(shardList)))
	return shardList[shardNo], nil
}

func (repository *DatabaseClusterRepositoryImpl) computeMD5Hash(id string) int64 {
	hash := md5.Sum([]byte(id))

	var hashInt64 int64
	for i := range hash {
		hashInt64 = (hashInt64 << 8) | int64(hash[i])
	}

	if hashInt64 < 0 {
		hashInt64 = -hashInt64
	}
	return hashInt64
}

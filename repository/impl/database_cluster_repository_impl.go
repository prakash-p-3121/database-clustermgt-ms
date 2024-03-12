package impl

import (
	"context"
	"database/sql"
	"errors"
	"github.com/prakash-p-3121/directory-database-lib/model"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/mysqllib"
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
		qry := `INSERT INTO cluster_to_shard_relationships (cluster_id, shard_id) VALUES (cluster_id, shard_id)
				ON DUPLICATE KEY UPDATE cluster_id=VALUES(cluster_id), shard_id=VALUES(shard_id);`
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

func (repository *DatabaseClusterRepositoryImpl) ReadCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError) {
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

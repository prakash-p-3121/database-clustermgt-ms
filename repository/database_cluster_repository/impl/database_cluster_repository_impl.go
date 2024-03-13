package impl

import (
	"context"
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	model "github.com/prakash-p-3121/database-clustermgt-model"
	shardRepo "github.com/prakash-p-3121/database-clustermgt-ms/repository/database_shard_repository"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/mysqllib"
	"strconv"
	"time"
)

type DatabaseClusterRepositoryImpl struct {
	DatabaseConnection *sql.DB
}

func (repository *DatabaseClusterRepositoryImpl) CreateCluster(tableName string,
	shardIDList []int64) (*model.DatabaseCluster, errorlib.AppError) {
	db := repository.DatabaseConnection
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	cluster, err := repository.createCluster(tx, tableName, shardIDList)
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
	shardIDList []int64) (*model.DatabaseCluster, error) {
	shardSize := int64(len(shardIDList))
	createdAt := time.Now().UTC()
	qry := `INSERT INTO database_clusters (table_name, shard_size, created_at, updated_at) VALUES (?, ?, ?, ?);`
	result, err := tx.Exec(qry, tableName, shardSize, createdAt, createdAt)
	if err != nil {
		return nil, err
	}
	clusterID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	clusterPtr := &model.DatabaseCluster{ID: &clusterID, TableName: &tableName, ShardSize: &shardSize, CreatedAt: createdAt, UpdatedAt: createdAt}
	err = repository.createClusterToShardRelationship(tx, clusterPtr, shardIDList)
	if err != nil {
		return nil, err
	}
	fmt.Println("ASASDASD")
	qry = `INSERT INTO current_write_database_clusters (table_name, cluster_id) VALUES (?,?) ON DUPLICATE KEY UPDATE 
    	   cluster_id=VALUES(cluster_id);`
	_, err = tx.Exec(qry, tableName, clusterID)
	if err != nil {
		return nil, err
	}
	fmt.Println("RFASDASD")

	return clusterPtr, nil
}

func (repository *DatabaseClusterRepositoryImpl) createClusterToShardRelationship(tx *sql.Tx,
	clusterPtr *model.DatabaseCluster,
	shardIDList []int64) error {

	for _, shardID := range shardIDList {
		qry := `UPDATE database_shards SET cluster_id=? WHERE id=?;`
		_, err := tx.Exec(qry, clusterPtr.ID, shardID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repository *DatabaseClusterRepositoryImpl) FindClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError) {
	db := repository.DatabaseConnection
	qry := `SELECT id, table_name, shard_size FROM database_clusters where id=?;`
	row := db.QueryRow(qry, id)
	var cluster model.DatabaseCluster
	fmt.Println("find cluster by id")
	err := row.Scan(&cluster.ID, &cluster.TableName, &cluster.ShardSize)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("cluster-id-not-found")
	}
	if err != nil {

		fmt.Println("ASDASDASDASDSA")
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	fmt.Println("find cluster by id")
	return &cluster, nil
}

func (repository *DatabaseClusterRepositoryImpl) FindCurrentWriteClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError) {
	db := repository.DatabaseConnection
	fmt.Println("find write -cluster by table name")
	qry := `SELECT A.id, A.table_name, A.shard_size  FROM database_clusters A INNER JOIN current_write_database_clusters B ON B.table_name = ? AND B.cluster_id = A.id ;`
	row := db.QueryRow(qry, tableName)
	var cluster model.DatabaseCluster
	err := row.Scan(&cluster.ID, &cluster.TableName, &cluster.ShardSize)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("write-cluster-not-found-for-tableName=" + tableName)
	}
	if err != nil {
		fmt.Println("AASDASD")
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
		fmt.Println("find-shards-by-cluster-id")
		return nil, err
	}
	if len(shardList) == 0 {
		return nil, errorlib.NewNotFoundError("shards-not-found-for-cluster_id=" + strconv.FormatInt(*cluster.ID, 10))
	}
	fmt.Println("find-current-write-shard-by-table-name")
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

func (repository *DatabaseClusterRepositoryImpl) FindAllShardsByTableName(tableName string) ([]*model.DatabaseShard, errorlib.AppError) {
	qry := `SELECT A.id, A.ip_address, A.cluster_id, A.port, A.user_name, A.password, A.database_name  FROM database_shards A 
				INNER JOIN database_clusters B ON B.table_name = ? AND A.cluster_id = B.id ORDER BY A.id desc;
           `
	db := repository.DatabaseConnection
	rows, err := db.Query(qry, tableName)
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	shardPtrList := make([]*model.DatabaseShard, 0)
	for rows.Next() {
		var shard model.DatabaseShard
		err := rows.Scan(
			&shard.ID,
			&shard.IPAddress,
			&shard.ClusterID,
			&shard.Port,
			&shard.UserName,
			&shard.Password,
			&shard.DatabaseName,
		)
		if err != nil {
			return nil, errorlib.NewInternalServerError(err.Error())
		}
		shardPtrList = append(shardPtrList, &shard)
	}
	return shardPtrList, nil
}

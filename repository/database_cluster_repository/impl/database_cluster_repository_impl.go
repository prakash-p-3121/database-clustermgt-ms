package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	model "github.com/prakash-p-3121/database-clustermgt-model"
	shardRepo "github.com/prakash-p-3121/database-clustermgt-ms/repository/database_shard_repository"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/mysqllib"
	"log"
	"math/big"
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
	createdAt := time.Now().UTC()
	qry := `INSERT INTO database_clusters (table_name, created_at, updated_at) VALUES (?, ?, ?);`
	result, err := tx.Exec(qry, tableName, createdAt, createdAt)
	if err != nil {
		return nil, err
	}
	clusterID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	clusterPtr := &model.DatabaseCluster{ID: &clusterID, TableName: &tableName, CreatedAt: createdAt, UpdatedAt: createdAt}
	err = repository.updateClusterToShardRelationship(tx, clusterPtr, shardIDList)
	if err != nil {
		return nil, err
	}
	fmt.Println("RFASDASD")
	return clusterPtr, nil
}

func (repository *DatabaseClusterRepositoryImpl) updateClusterToShardRelationship(tx *sql.Tx,
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
	qry := `SELECT id, table_name FROM database_clusters where id=?;`
	row := db.QueryRow(qry, id)
	var cluster model.DatabaseCluster
	fmt.Println("find cluster by id")
	err := row.Scan(&cluster.ID, &cluster.TableName)
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

func (repository *DatabaseClusterRepositoryImpl) FindClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError) {
	db := repository.DatabaseConnection
	fmt.Println("find write -cluster by table name")
	qry := `SELECT A.id, A.table_name  FROM database_clusters A WHERE A.table_name=?;`
	row := db.QueryRow(qry, tableName)
	var cluster model.DatabaseCluster
	err := row.Scan(&cluster.ID, &cluster.TableName)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("NotFound error")
		return nil, errorlib.NewNotFoundError("write-cluster-not-found-for-tableName=" + tableName)
	}
	if err != nil {
		fmt.Println("AASDASD")
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &cluster, nil
}

func (repository *DatabaseClusterRepositoryImpl) FindShard(tableName string,
	id string) (*model.DatabaseShard, errorlib.AppError) {
	cluster, appErr := repository.FindClusterByTableName(tableName)
	if appErr != nil {
		log.Println("Find Write Cluster By Table Name")
		return nil, appErr
	}
	shardRepoInst := shardRepo.NewDatabaseShardRepository(repository.DatabaseConnection)
	shardList, appErr := shardRepoInst.FindShardsByClusterID(*cluster.ID)
	if appErr != nil {
		fmt.Println("find-shards-by-cluster-id")
		return nil, appErr
	}
	log.Println("shard-len-check")
	if len(shardList) == 0 {
		return nil, errorlib.NewNotFoundError("shards-not-found-for-cluster_id=" + strconv.FormatInt(*cluster.ID, 10))
	}
	fmt.Println("find-current-write-shard-by-table-name")
	shard, appErr := repository.findShard(shardList, id)
	if appErr != nil {
		return nil, appErr
	}
	return shard, nil
}

func (repository *DatabaseClusterRepositoryImpl) findShard(shardList []*model.DatabaseShard, id string) (*model.DatabaseShard, errorlib.AppError) {
	decimalBase := 10
	idBigInt := new(big.Int)
	idBigInt, ok := idBigInt.SetString(id, decimalBase)
	if !ok {
		return nil, errorlib.NewInternalServerError("database-shard-not-found")
	}
	for _, shardPtr := range shardList {
		log.Println("shardID=", *shardPtr.ID)
		startRangeBigInt, ok := new(big.Int).SetString(*shardPtr.StartRange, decimalBase)
		if !ok {
			return nil, errorlib.NewInternalServerError("start-range-invalid")
		}
		var endRangeBigInt *big.Int
		if shardPtr.EndRange != nil {
			endRangeBigInt, ok = new(big.Int).SetString(*shardPtr.EndRange, decimalBase)
			if !ok {
				return nil, errorlib.NewInternalServerError("start-range-invalid")
			}
		}
		startRangeCmpRes := startRangeBigInt.Cmp(idBigInt)
		log.Println("start-range-cmp-res=", startRangeCmpRes)
		if endRangeBigInt != nil {
			endRangeCmpRes := endRangeBigInt.Cmp(idBigInt)
			log.Println("end-range-cmp-res=", endRangeCmpRes)
			if (startRangeCmpRes == -1 || startRangeCmpRes == 0) && (endRangeCmpRes == 1 || endRangeCmpRes == 0) {
				return shardPtr, nil
			}
		} else {
			if startRangeCmpRes == -1 || startRangeCmpRes == 0 {
				return shardPtr, nil
			}
		}
	}
	return nil, errorlib.NewInternalServerError("shard-not-found-for-id=" + id)
}

func (repository *DatabaseClusterRepositoryImpl) FindAllShardsByTable(tableName string) ([]*model.DatabaseShard, errorlib.AppError) {
	qry := `SELECT A.id, A.ip_address, A.cluster_id, A.port, A.user_name, A.password, A.database_name, A.start_range, A.end_range  FROM database_shards A 
				INNER JOIN database_clusters B ON B.table_name = ? AND A.cluster_id = B.id ORDER BY A.id ASC;
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
			&shard.StartRange,
			&shard.EndRange,
		)
		if err != nil {
			return nil, errorlib.NewInternalServerError(err.Error())
		}
		shardPtrList = append(shardPtrList, &shard)
	}
	return shardPtrList, nil
}

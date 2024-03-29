package impl

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/database-clustermgt-ms/repository/database_cluster_repository"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseClusterServiceImpl struct {
	DatabaseClusterRepository database_cluster_repository.DatabaseClusterRepository
}

func (service *DatabaseClusterServiceImpl) CreateCluster(req *model.DatabaseClusterCreateReq) (*model.DatabaseCluster, errorlib.AppError) {
	if req == nil {
		badReqErr := errorlib.NewBadReqError("req-null")
		return nil, badReqErr
	}
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.CreateCluster(*req.TableName, *req.ShardingType, req.ShardIDList)
}

func (service *DatabaseClusterServiceImpl) ReadClusterByID(id int64) (*model.DatabaseCluster, errorlib.AppError) {
	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.FindClusterByID(id)
}

func (service *DatabaseClusterServiceImpl) FindClusterByTableName(tableName string) (*model.DatabaseCluster, errorlib.AppError) {
	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.FindClusterByTableName(tableName)
}

func (service *DatabaseClusterServiceImpl) FindShardByNumber(tableName string,
	id string) (*model.DatabaseShard, errorlib.AppError) {
	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.FindShardByNumber(tableName, id)
}

func (service *DatabaseClusterServiceImpl) FindShardByChar(tableName string,
	id rune) (*model.DatabaseShard, errorlib.AppError) {
	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.FindShardByChar(tableName, id)
}

func (service *DatabaseClusterServiceImpl) FindAllShardsByTable(tableName string) ([]*model.DatabaseShard, errorlib.AppError) {
	clusterRepo := service.DatabaseClusterRepository
	return clusterRepo.FindAllShardsByTable(tableName)
}

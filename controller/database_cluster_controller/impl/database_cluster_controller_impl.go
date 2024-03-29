package impl

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/database-clustermgt-ms/service/database_cluster_service"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/restlib"
	"log"
	"strings"
)

type DatabaseClusterControllerImpl struct {
	DatabaseClusterService database_cluster_service.DatabaseClusterService
}

func (controller *DatabaseClusterControllerImpl) CreateCluster(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	var req model.DatabaseClusterCreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		badReqErr := errorlib.NewInternalServerError("payload-serialization-err=" + err.Error())
		badReqErr.SendRestResponse(ctx)
		return
	}

	clusterPtr, appErr := controller.DatabaseClusterService.CreateCluster(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, *clusterPtr)
}

func (controller *DatabaseClusterControllerImpl) FindShard(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	tableName := ctx.Query("table-name")
	if len(strings.TrimSpace(tableName)) == 0 {
		badReqErr := errorlib.NewBadReqError("table-name-empty")
		badReqErr.SendRestResponse(ctx)
		return
	}

	strIsByNumber := ctx.Query("is-by-number")
	if len(strings.TrimSpace(strIsByNumber)) == 0 {
		badReqErr := errorlib.NewBadReqError("is-by-number-empty")
		badReqErr.SendRestResponse(ctx)
		return
	}

	id := ctx.Query("resource-id")
	if len(strings.TrimSpace(id)) == 0 {
		badReqErr := errorlib.NewBadReqError("resource-id-empty")
		badReqErr.SendRestResponse(ctx)
		return
	}

	var shardPtr *model.DatabaseShard
	var appErr errorlib.AppError
	shardPtr, appErr = controller.DatabaseClusterService.FindShard(tableName, id)
	if appErr != nil {
		log.Println("FindCurrentWriteShardByTableName Err")
		appErr.SendRestResponse(ctx)
		return
	}

	restlib.OkResponse(ctx, *shardPtr)
}

func (controller *DatabaseClusterControllerImpl) FindAllShardsByTable(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	tableName := ctx.Query("table-name")
	if len(strings.TrimSpace(tableName)) == 0 {
		badReqErr := errorlib.NewBadReqError("table-name-empty")
		badReqErr.SendRestResponse(ctx)
		return
	}

	shardPtrList, appErr := controller.DatabaseClusterService.FindAllShardsByTable(tableName)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}

	restlib.OkResponse(ctx, shardPtrList)
}

package impl

import (
	"github.com/prakash-p-3121/directory-database-lib/service/database_cluster_service"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/restlib"
	"strings"
)

type DatabaseClusterControllerImpl struct {
	DatabaseClusterService database_cluster_service.DatabaseClusterService
}

func (controller *DatabaseClusterControllerImpl) FindCurrentWriteShard(restCtx restlib.RestContext) {
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

	id := ctx.Query("resource-id")
	if len(strings.TrimSpace(id)) == 0 {
		badReqErr := errorlib.NewBadReqError("resource-id-empty")
		badReqErr.SendRestResponse(ctx)
		return
	}

	shardPtr, appErr := controller.DatabaseClusterService.FindCurrentWriteShardByTableName(tableName, id)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}

	restlib.OkResponse(ctx, *shardPtr)
}

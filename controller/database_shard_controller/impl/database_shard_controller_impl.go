package impl

import (
	"fmt"
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/database-clustermgt-ms/service/database_shard_service"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/restlib"
	"strings"
)

type DatabaseShardControllerImpl struct {
	DatabaseShardService database_shard_service.DatabaseShardService
}

func (controller *DatabaseShardControllerImpl) CreateShard(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}
	fmt.Println("testing1")
	ctx := ginRestCtx.CtxGet()
	var req model.DatabaseShardCreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		appErr := errorlib.NewInternalServerError("payload-serialization-err=" + err.Error())
		appErr.SendRestResponse(ctx)
		return
	}
	fmt.Println("testing2")
	service := controller.DatabaseShardService
	id, appErr := service.CreateShard(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}

	resp := model.CreateShardResp{ShardID: id}
	restlib.OkResponse(ctx, resp)
}

func (controller *DatabaseShardControllerImpl) FindShardByID(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	shardIDStr := ctx.Query("shard-id")
	shardIDStr = strings.TrimSpace(shardIDStr)
	if len(shardIDStr) == 0 {
		appErr := errorlib.NewBadReqError("shard-id-empty")
		appErr.SendRestResponse(ctx)
		return
	}
	shardID, err := restlib.AsInt64(shardIDStr)
	if err != nil {
		appErr := errorlib.NewBadReqError("shard-id-not-integer")
		appErr.SendRestResponse(ctx)
		return
	}

	service := controller.DatabaseShardService
	shard, appErr := service.FindShardByID(shardID)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, shard)
}

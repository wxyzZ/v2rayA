package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/v2rayA/v2rayA/common"
	"github.com/v2rayA/v2rayA/conf"
	"github.com/v2rayA/v2rayA/db/configure"
	"github.com/v2rayA/v2rayA/pkg/util/log"
	"github.com/v2rayA/v2rayA/server/service"
)

func PostConnection(ctx *gin.Context) {
	conf.UpdatingMu.Lock()
	if updating {
		common.ResponseError(ctx, processingErr)
		conf.UpdatingMu.Unlock()
		return
	}
	updating = true
	conf.UpdatingMu.Unlock()
	defer func() {
		conf.UpdatingMu.Lock()
		updating = false
		conf.UpdatingMu.Unlock()
	}()

	var which configure.Which
	err := ctx.ShouldBindJSON(&which)
	if err != nil {
		common.ResponseError(ctx, logError("bad request"))
		return
	}
	err = service.Connect(&which)
	if err != nil {
		log.Warn("PostConnection: %v", err)
		common.ResponseError(ctx, logError(fmt.Errorf("failed to connect: %w", err)))
		return
	}
	getTouch(ctx)
}

func DeleteConnection(ctx *gin.Context) {
	conf.UpdatingMu.Lock()
	if updating {
		common.ResponseError(ctx, processingErr)
		conf.UpdatingMu.Unlock()
		return
	}
	updating = true
	conf.UpdatingMu.Unlock()
	defer func() {
		conf.UpdatingMu.Lock()
		updating = false
		conf.UpdatingMu.Unlock()
	}()

	var which configure.Which
	err := ctx.ShouldBindJSON(&which)
	if err != nil {
		common.ResponseError(ctx, logError("bad request"))
		return
	}
	err = service.Disconnect(which, false)
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	getTouch(ctx)
}

func PostV2ray(ctx *gin.Context) {
	conf.UpdatingMu.Lock()
	if updating {
		common.ResponseError(ctx, processingErr)
		conf.UpdatingMu.Unlock()
		return
	}
	updating = true
	conf.UpdatingMu.Unlock()
	defer func() {
		conf.UpdatingMu.Lock()
		updating = false
		conf.UpdatingMu.Unlock()
	}()

	err := service.StartV2ray()
	if err != nil {
		common.ResponseError(ctx, logError(fmt.Errorf("failed to start v2ray-core: %w", err)))
		return
	}
	getTouch(ctx)
}

func DeleteV2ray(ctx *gin.Context) {
	conf.UpdatingMu.Lock()
	if updating {
		common.ResponseError(ctx, processingErr)
		conf.UpdatingMu.Unlock()
		return
	}
	updating = true
	conf.UpdatingMu.Unlock()
	defer func() {
		conf.UpdatingMu.Lock()
		updating = false
		conf.UpdatingMu.Unlock()
	}()

	err := service.StopV2ray()
	if err != nil {
		common.ResponseError(ctx, logError(fmt.Errorf("failed to stop v2ray-core: %w", err)))
		return
	}
	getTouch(ctx)
}

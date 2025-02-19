package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/v2rayA/v2rayA/common"
	"github.com/v2rayA/v2rayA/conf"
	"github.com/v2rayA/v2rayA/core/touch"
	"github.com/v2rayA/v2rayA/core/v2ray"
	"github.com/v2rayA/v2rayA/db/configure"
	"github.com/v2rayA/v2rayA/server/service"
)

func GetTouch(ctx *gin.Context) {
	conf.UpdatingMu.Lock()
	if updating {
		common.ResponseError(ctx, processingErr)
		conf.UpdatingMu.Unlock()
		return
	}
	conf.UpdatingMu.Unlock()
	defer func() {
		conf.UpdatingMu.Lock()
		conf.UpdatingMu.Unlock()
	}()
	getTouch(ctx)

}
func getTouch(ctx *gin.Context) {
	running := v2ray.ProcessManager.Running()
	t := touch.GenerateTouch()
	common.ResponseSuccess(ctx, gin.H{
		"running": running,
		"touch":   t,
	})
}

func DeleteTouch(ctx *gin.Context) {
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

	var ws configure.Whiches
	err := ctx.ShouldBindJSON(&ws)
	if err != nil {
		common.ResponseError(ctx, logError("bad request"))
		return
	}
	err = service.DeleteWhich(ws.Get())
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	getTouch(ctx)
}

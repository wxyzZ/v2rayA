package controller

import (
	"github.com/v2rayA/v2rayA/conf"
	"github.com/v2rayA/v2rayA/pkg/util/log"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/v2rayA/v2rayA/common"
	"github.com/v2rayA/v2rayA/db/configure"
	"github.com/v2rayA/v2rayA/server/service"
)

func GetPingLatency(ctx *gin.Context) {
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

	var wt []*configure.Which
	err := jsoniter.Unmarshal([]byte(ctx.Query("whiches")), &wt)
	if err != nil {
		common.ResponseError(ctx, logError("bad request"))
		return
	}
	wt, err = service.Ping(wt, 1*time.Second)
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"whiches": wt,
	})
}

func GetHttpLatency(ctx *gin.Context) {
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

	var wt []*configure.Which
	err := jsoniter.Unmarshal([]byte(ctx.Query("whiches")), &wt)
	if err != nil {
		common.ResponseError(ctx, logError("bad request"))
		return
	}
	log.Debug("controller/latency.go --- testurl:%s", ctx.Query("testUrl"))

	testUrl := ctx.Query("testUrl")
	//if testUrl == "" {
	//	outbound := configure.GetOutbounds()[0]
	//	outSetting := configure.GetOutboundSetting(outbound)
	//	testUrl = outSetting.ProbeURL
	//}
	conf.UpdatingMu2.Lock()
	wt, err = service.TestHttpLatency(wt, 8*time.Second, 32, false, testUrl)
	conf.UpdatingMu2.Unlock()
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"whiches": wt,
	})
}

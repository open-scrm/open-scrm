package configcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/pkg/system/model"
	"github.com/open-scrm/open-scrm/pkg/system/service"
)

func UpdateTalentInfo(ctx *gin.Context) {
	req := UpsertTalentInfoRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		vo.SendFail(ctx, "参数错误")
		return
	}

	if err := service.NewTalentService().SetTalentInfo(ctx.Request.Context(), model.Talent{
		CorpId:                            req.CorpId,
		AgentId:                           req.AgentId,
		AddressBookSecret:                 req.AddressBookSecret,
		AppSecret:                         req.AppSecret,
		ExternalContactSecret:             req.ExternalContactSecret,
		AddressBookCallbackToken:          req.AddressBookCallbackToken,
		AddressBookCallbackAesEncodingKey: req.AddressBookCallbackAesEncodingKey,
	}); err != nil {
		log.WithContext(ctx.Request.Context()).WithError(err).Errorf("更新租户信息失败")
		vo.SendFail(ctx, "更新租户信息失败")
		return
	}

	// TODO:: 刷新token
	vo.SendOK(ctx)
}

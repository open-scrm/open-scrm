package callbaclcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/wxwork"
	service2 "github.com/open-scrm/open-scrm/pkg/addressbook/service"
	"github.com/open-scrm/open-scrm/pkg/system/service"
	"io/ioutil"
)

func AddressBookCallback(ctx *gin.Context) {
	// 获取talent信息.
	talent, err := service.NewTalentService().GetTalentInfo(ctx.Request.Context())
	if err != nil {
		log.WithContext(ctx.Request.Context()).Errorf("查询配置信息失败")
		ctx.AbortWithError(500, err)
		return
	}

	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.WithContext(ctx.Request.Context()).WithError(err).Errorf("readBody failed")
		ctx.AbortWithError(500, err)
		return
	}

	crypt := wxwork.NewWXBizMsgCrypt(talent.AddressBookCallbackToken, talent.AddressBookCallbackAesEncodingKey, talent.CorpId, wxwork.XmlType)
	msg, xerr := crypt.DecryptMsg(ctx.Query("msg_signature"), ctx.Query("timestamp"), ctx.Query("nonce"), data)
	if xerr != nil {
		log.WithContext(ctx.Request.Context()).WithError(fmt.Errorf("%v", xerr)).Errorf("企微回调接码失败")
		ctx.AbortWithError(500, err)
		return
	}
	log.WithContext(ctx).WithField("data", msg).Infof("收到企微通讯录回调")
	// 推送到不同的事件接收器.
	if err := service2.NewCallbackService().HandleMessage(ctx.Request.Context(), msg); err != nil {
		log.WithContext(ctx.Request.Context()).WithError(err).Errorf("处理回调失败")
		ctx.AbortWithError(500, err)
		return
	}

	ctx.String(200, "ok")
}

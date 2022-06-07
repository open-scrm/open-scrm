package callbaclcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/wxwork"
	"github.com/open-scrm/open-scrm/pkg/system/service"
)

type AddressCallbackRequest struct {
	ToUserName string `json:"toUserName" xml:"toUserName"`
	Encrypt    string `json:"encrypt" xml:"encrypt"`
	AgentID    string `json:"agentId" xml:"agentID"`
}

func AddressBookCallback(ctx *gin.Context) {
	req := new(AddressCallbackRequest)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithContext(ctx.Request.Context()).Errorf("回调参数绑定失败")
		return
	}
	// 获取talent信息.
	talent, err := service.NewTalentService().GetTalentInfo(ctx.Request.Context())
	if err != nil {
		log.WithContext(ctx.Request.Context()).Errorf("回调参数绑定失败")
		return
	}

	crypt := wxwork.NewWXBizMsgCrypt(talent.AddressBookCallbackToken, talent.AddressBookCallbackAesEncodingKey, talent.CorpId)
	data, err := crypt.EncryptMsg(req.Encrypt, ctx.Query("timestamp"), ctx.Query("nonce"))
	fmt.Println(string(data))
}

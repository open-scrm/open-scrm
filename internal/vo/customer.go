package vo

import "github.com/open-scrm/open-scrm/lib/utils"

type ListCustomerRequest struct {
	utils.PageUtil
	Type          int      `json:"type"`
	Name          string   `json:"name"`          // 名称
	Remark        string   `json:"remark"`        // 备注名
	CorpName      string   `json:"corpName"`      // 公司名称
	Mobile        string   `json:"mobile"`        // 手机号
	AddWay        int      `json:"addWay"`        // 好友来源
	CreateTimeGte int64    `json:"createTimeGte"` // 创建时间大于
	CreateTimeLte int64    `json:"createTimeLte"` // 创建时间小于
	Tags          []string `json:"tags"`          // 企微标签
}

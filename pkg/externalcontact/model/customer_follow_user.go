package model

import (
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerFollowUser struct {
	Id                     int64    `json:"id" bson:"_id"`
	Uid                    int64    `json:"uid" bson:"uid"`
	Userid                 string   `json:"userid" bson:"userid"`                 // 添加了此外部联系人的企业成员userid
	Remark                 string   `json:"remark" bson:"remark"`                 // 该成员对此外部联系人的备注
	Description            string   `json:"description" bson:"description"`       // 该成员对此外部联系人的描述
	Createtime             int64    `json:"createtime" bson:"createtime"`         // 该成员添加此外部联系人的时间
	TagId                  []string `json:"tagId" bson:"tagId"`                   // 该成员添加此外部联系人所打企业标签的id，用户自定义类型标签（type=2）不返回
	RemarkCorpName         string   `json:"remarkCorpName" bson:"remarkCorpName"` // 该成员对此微信客户备注的企业名称（仅微信客户有该字段）
	RemarkMobiles          []string `json:"remarkMobiles" bson:"remarkMobiles"`   // 该成员对此客户备注的手机号码，代开发自建应用需要管理员授权才可以获取，第三方不可获取，上游企业不可获取下游企业客户该字段
	OperUserid             string   `json:"operUserid" bson:"operUserid"`
	AddWay                 int      `json:"addWay" bson:"addWay"` // 该成员添加此客户的来源，具体含义详见来源定义 https://developer.work.weixin.qq.com/document/path/92114#%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89
	WechatChannelsNickname string   `json:"wechatChannelsNickname" bson:"wechatChannelsNickname"`
	WechatChannelsSource   int      `json:"wechatChannelsSource" bson:"wechatChannelsSource"`
	State                  string   `json:"state" bson:"state"`
	FriendState            int      `json:"friendState" bson:"friendState"`
	ExternalUserId         string   `json:"externalUserId" bson:"externalUserId"`
}

func GetCustomerFollowUserCollection() *mongo.Collection {
	return global.GetMongoDriver().Database(configs.Get().Mongo.Database).Collection("customer_follow_user")
}

/**
0	未知来源
1	扫描二维码
2	搜索手机号
3	名片分享
4	群聊
5	手机通讯录
6	微信联系人
8	安装第三方应用时自动添加的客服人员
9	搜索邮箱
10	视频号添加
11	通过日程参与人添加
12	通过会议参与人添加
13	添加微信好友对应的企业微信
14	通过智慧硬件专属客服添加
201	内部成员共享
202	管理员/负责人分配
*/
const (
	AddWayKnown = iota // 未知来源
	AddWayScanQRCode
	AddWaySearchMobile
	AddWayCardShare
	AddWayChatroom
	AddWayMobileContact
	AddWayThirdApp
	AddWaySearchEmail
	AddWayWechatChannel
	AddWayCalendar
	AddWayMeeting
	AddWayWechatCorp
	AddWayHardware
	AddWayInternalShare = 201
	AddWayAdminDispatch = 202
)

const (
	FriendStateIsFriend       = iota + 1 // 是好友关系
	FriendStateCustomerDelete            // 客户主动删除
	FriendStateUserDelete                // 员工主动删除
)

package wxwork

const (
	accessTokenURL = `https://qyapi.weixin.qq.com/cgi-bin/gettoken`
)

// 客户联系
const (
	// ListExternalContact 列出客户列表
	ListExternalContact     = `https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list`              // ?access_token=ACCESS_TOKEN&userid=USERID
	GetExternalContext      = `https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get`               // ?access_token=ACCESS_TOKEN&external_userid=EXTERNAL_USERID&cursor=CURSOR
	BatchGetExternalContext = `https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user` // ?access_token=ACCESS_TOKEN
)

const (
	OAuthCodeExchange = `https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo`
)

// 部门管理
const (
	departmentList = `https://qyapi.weixin.qq.com/cgi-bin/department/list`
	userListByDept = `https://qyapi.weixin.qq.com/cgi-bin/user/simplelist`
	getUser        = `https://qyapi.weixin.qq.com/cgi-bin/user/get`
)

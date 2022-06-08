package service

import "encoding/xml"

type CreateUserMessage struct {
	XMLName        xml.Name `xml:"xml" json:"-"`
	Text           string   `xml:",chardata" json:"-"`
	ToUserName     string   `xml:"ToUserName" json:"-"`
	FromUserName   string   `xml:"FromUserName" json:"-"`
	CreateTime     int64    `json:"createTime" xml:"CreateTime"`
	MsgType        string   `json:"msgType" xml:"MsgType"`
	Event          string   `json:"event" xml:"Event"`
	ChangeType     string   `json:"changeType" xml:"ChangeType"`
	UserID         string   `json:"userID" xml:"UserID"`
	Name           string   `json:"name" xml:"Name"`
	Department     uint32   `json:"department" xml:"Department"`
	MainDepartment uint32   `json:"mainDepartment" xml:"MainDepartment"`
	IsLeaderInDept string   `json:"isLeaderInDept" xml:"IsLeaderInDept"`
	DirectLeader   string   `json:"directLeader" xml:"DirectLeader"`
	Position       string   `json:"position" xml:"Position"`
	Mobile         string   `json:"mobile" xml:"Mobile"`
	Gender         string   `json:"gender" xml:"Gender"`
	Email          string   `json:"email" xml:"Email"`
	BizMail        string   `json:"bizMail" xml:"BizMail"`
	Status         int8     `json:"status" xml:"Status"`
	Avatar         string   `json:"avatar" xml:"Avatar"`
	Alias          string   `json:"alias" xml:"Alias"`
	Telephone      string   `json:"telephone" xml:"Telephone"`
	Address        string   `json:"address" xml:"Address"`
}

type UpdateUserMessage struct {
	XMLName        xml.Name `xml:"xml" json:"-"`
	Text           string   `xml:",chardata" json:"-"`
	ToUserName     string   `xml:"ToUserName" json:"-"`
	FromUserName   string   `xml:"FromUserName" json:"-"`
	CreateTime     int64    `json:"createTime" xml:"CreateTime"`
	MsgType        string   `json:"msgType" xml:"MsgType"`
	Event          string   `json:"event" xml:"Event"`
	ChangeType     string   `json:"changeType" xml:"ChangeType"`
	UserID         string   `json:"userId" xml:"UserID"`
	NewUserID      string   `json:"newUserId" xml:"NewUserID"`
	Name           string   `json:"name" xml:"Name"`
	Department     string   `json:"department" xml:"Department"`
	MainDepartment string   `json:"mainDepartment" xml:"MainDepartment"`
	IsLeaderInDept string   `json:"isLeaderInDept" xml:"IsLeaderInDept"`
	Position       string   `json:"position" xml:"Position"`
	Mobile         string   `json:"mobile" xml:"Mobile"`
	Gender         string   `json:"gender" xml:"Gender"`
	Email          string   `json:"email" xml:"Email"`
	Status         string   `json:"status" xml:"Status"`
	Avatar         string   `json:"avatar" xml:"Avatar"`
	Alias          string   `json:"alias" xml:"Alias"`
	Telephone      string   `json:"telephone" xml:"Telephone"`
	Address        string   `json:"address" xml:"Address"`
}

type DeleteUserMessage struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	Text         string   `xml:",chardata" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"-"`
	FromUserName string   `xml:"FromUserName" json:"-"`
	CreateTime   int64    `json:"createTime" xml:"CreateTime"`
	MsgType      string   `json:"msgType" xml:"MsgType"`
	Event        string   `json:"event" xml:"Event"`
	ChangeType   string   `json:"changeType" xml:"ChangeType"`
	UserID       string   `json:"userID" xml:"UserID"`
}

type CreateDeptMessage struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	Text         string   `xml:",chardata" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"-"`
	FromUserName string   `xml:"FromUserName" json:"-"`
	CreateTime   int64    `json:"createTime" xml:"CreateTime"`
	MsgType      string   `json:"msgType" xml:"MsgType"`
	Event        string   `json:"event" xml:"Event"`
	ChangeType   string   `json:"changeType" xml:"ChangeType"`
	ID           uint32   `json:"id" xml:"Id"`
	Name         string   `json:"name" xml:"Name"`
	ParentId     uint32   `json:"parentId" xml:"ParentId"`
	Order        uint32   `json:"order" xml:"Order"`
}

type UpdateteDeptMessage struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	Text         string   `xml:",chardata" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"-"`
	FromUserName string   `xml:"FromUserName" json:"-"`
	CreateTime   int64    `json:"createTime" xml:"CreateTime"`
	MsgType      string   `json:"msgType" xml:"MsgType"`
	Event        string   `json:"event" xml:"Event"`
	ChangeType   string   `json:"changeType" xml:"ChangeType"`
	ID           uint32   `json:"id" xml:"Id"`
	Name         string   `json:"name" xml:"Name"`
	ParentId     uint32   `json:"parentId" xml:"ParentId"`
}

type DeleteDeptMessage struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	Text         string   `xml:",chardata" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"-"`
	FromUserName string   `xml:"FromUserName" json:"-"`
	CreateTime   int64    `json:"createTime" xml:"CreateTime"`
	MsgType      string   `json:"msgType" xml:"MsgType"`
	Event        string   `json:"event" xml:"Event"`
	ChangeType   string   `json:"changeType" xml:"ChangeType"`
	ID           uint32   `json:"id" xml:"Id"`
}

type TagChangeMessage struct {
	XMLName       xml.Name `xml:"xml" json:"-"`
	Text          string   `xml:",chardata" json:"-"`
	ToUserName    string   `xml:"ToUserName" json:"-"`
	FromUserName  string   `xml:"FromUserName" json:"-"`
	CreateTime    int64    `json:"createTime" xml:"CreateTime"`
	MsgType       string   `json:"msgType" xml:"MsgType"`
	Event         string   `json:"event" xml:"Event"`
	ChangeType    string   `json:"changeType" xml:"ChangeType"`
	TagId         string   `json:"tagId" xml:"TagId"`
	AddUserItems  string   `json:"addUserItems" xml:"AddUserItems"`
	DelUserItems  string   `json:"delUserItems" xml:"DelUserItems"`
	AddPartyItems string   `json:"addPartyItems" xml:"AddPartyItems"`
	DelPartyItems string   `json:"delPartyItems" xml:"DelPartyItems"`
}

type BatchJobFinishedMessage struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	Text         string   `xml:",chardata" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"-"`
	FromUserName int64    `xml:"FromUserName" json:"-"`
	CreateTime   string   `xml:"CreateTime" json:"createTime"`
	MsgType      string   `xml:"MsgType" json:"msgType"`
	Event        string   `xml:"Event" json:"event"`
	BatchJob     struct {
		JobId   string `xml:"JobId" json:"jobId"`
		JobType string `xml:"JobType" json:"jobType"`
		ErrCode string `xml:"ErrCode" json:"errCode"`
		ErrMsg  string `xml:"ErrMsg" json:"errMsg"`
	} `xml:"BatchJob" json:"batchJob"`
}

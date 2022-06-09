package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	var msg = `<xml><ToUserName><![CDATA[ww48fb21eab5cc8802]]></ToUserName><FromUserName><![CDATA[sys]]></FromUserName><CreateTime>1654612061</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[change_contact]]></Event><ChangeType><![CDATA[update_user]]></ChangeType><UserID><![CDATA[HuaHuaDeSaHaLa]]></UserID><Alias><![CDATA[aliasa222222asdsd]]></Alias></xml>`
	var upd interface{}
	upd = UpdateUserMessage{}
	xml.Unmarshal([]byte(msg), &upd)
	fmt.Println(upd)
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

package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/pkg/errors"
)

type CallbackService struct{}

func NewCallbackService() *CallbackService {
	return &CallbackService{}
}

type CallbackMessage struct {
	XMLName    xml.Name `xml:"xml"`
	MsgType    string   `xml:"MsgType"`
	Event      string   `xml:"Event"`
	ChangeType string   `xml:"ChangeType"`
}

const (
	MsgTypeEvent = "event"
)

const (
	MsgEventChangeContact  = "change_contact"   // 新增|更新|删除成员事件 | 新增|更新|删除部门事件 | 标签成员变更事件
	MsgEventBatchJobResult = "batch_job_result" // 异步任务完成通知
)

const (
	changeTypeCreateUser  = "create_user"  // 新增成员
	changeTypeUpdateUser  = "update_user"  // 更新成员
	changeTypeDeleteUser  = "delete_user"  // 删除成员
	changeTypeCreateParty = "create_party" // 新增部门
	changeTypeUpdateParty = "update_party" // 更新部门
	changeTypeDeleteParty = "delete_party" // 删除部门
	changeTypeUpdateTag   = "update_tag"   // 标签成员变更事件
)

func (c CallbackService) HandleMessage(ctx context.Context, msg []byte) error {
	var callbackMessage CallbackMessage
	if err := xml.Unmarshal(msg, &callbackMessage); err != nil {
		return errors.Wrap(err, "接码xml失败")
	}

	switch callbackMessage.MsgType {
	case MsgTypeEvent:
	default:
		return fmt.Errorf("invalid msg type: %v", callbackMessage.MsgType)
	}

	switch callbackMessage.Event {
	case MsgEventChangeContact:
		var (
			err  error
			data interface{}
		)
		switch callbackMessage.ChangeType {
		case changeTypeCreateUser:
			_data := CreateUserMessage{}
			err = xml.Unmarshal(msg, &_data)
			data = _data
		case changeTypeUpdateUser:
			_data := UpdateUserMessage{}
			err = xml.Unmarshal(msg, &_data)
			data = _data
		case changeTypeDeleteUser:
			_data := DeleteUserMessage{}
			err = xml.Unmarshal(msg, &_data)
			data = _data
		case changeTypeCreateParty:
			_data := CreateDeptMessage{}
			err = xml.Unmarshal(msg, &_data)
			data = _data
		case changeTypeUpdateParty:
			_data := UpdateteDeptMessage{}
			err = xml.Unmarshal(msg, &_data)
			data = _data
		case changeTypeDeleteParty:
			_data := DeleteDeptMessage{}
			err = xml.Unmarshal(msg, &_data)
			data = _data
		case changeTypeUpdateTag:
			_data := TagChangeMessage{}
			err = xml.Unmarshal(msg, &_data)
			data = _data
		default:
			return fmt.Errorf("invalid msg change type: %v", callbackMessage.ChangeType)
		}
		if err != nil {
			return errors.Wrap(err, "接码消息失败")
		}
		// 推送到不同的消息机.
		switch callbackMessage.ChangeType {
		case changeTypeCreateUser, changeTypeUpdateUser, changeTypeDeleteUser:
			err = global.GetAddressBookUserPublisher().PublishOne(ctx, data)
		case changeTypeCreateParty, changeTypeUpdateParty, changeTypeDeleteParty:
			err = global.GetAddressBookDeptPublisher().PublishOne(ctx, data)
		case changeTypeUpdateTag:
			err = global.GetAddressBookTagPublisher().PublishOne(ctx, data)
		}
		if err != nil {
			log.WithContext(ctx).WithError(err).Errorf("推送消息失败: %v", data)
		}
	case MsgEventBatchJobResult:
		data := BatchJobFinishedMessage{}
		err := xml.Unmarshal(msg, &data)
		if err != nil {
			return errors.Wrap(err, "接码消息失败")
		}
		err = global.GetAddressBookBatchJobPublisher().PublishOne(ctx, data)
		if err != nil {
			log.WithContext(ctx).WithError(err).Errorf("推送消息失败: %v", data)
		}
	default:
		return fmt.Errorf("invalid msg event: %v", callbackMessage.Event)
	}
	return nil
}

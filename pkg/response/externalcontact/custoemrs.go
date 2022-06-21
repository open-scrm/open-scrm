package externalcontact

import (
	userModel "github.com/open-scrm/open-scrm/pkg/addressbook/model"
	"github.com/open-scrm/open-scrm/pkg/externalcontact/model"
)

type CustomerListResponse struct {
	*model.Customer
	AddTime   string                `json:"addTime"`
	OwnerInfo *userModel.SimpleUser `json:"ownerInfo"`
}

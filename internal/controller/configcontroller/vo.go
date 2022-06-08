package configcontroller

type UpsertTalentInfoRequest struct {
	CorpId                            string `json:"corpId"  binding:"required"`
	AgentId                           string `json:"agentId" binding:"required"`
	AddressBookSecret                 string `json:"addressBookSecret"  binding:"required"`
	AppSecret                         string `json:"appSecret" binding:"required"`
	ExternalContactSecret             string `json:"externalContactSecret"  binding:"required"`
	AddressBookCallbackToken          string `json:"addressBookCallbackToken" binding:"required"`
	AddressBookCallbackAesEncodingKey string `json:"addressBookCallbackAesEncodingKey" bingding:"required"`
}

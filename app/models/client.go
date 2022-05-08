package models

import "strconv"

type Client struct {
	ID
	Uuid
	AppID       int64  `json:"app_id" gorm:"not null;index;comment:客户端ID"`
	SecretKey   string `json:"-"  gorm:"not null;comment:客户端密钥"`
	Frozen      string `json:"frozen"  gorm:"default:0;comment:客户端状态"`
	NickName    string `json:"nickName" gorm:"not null;size:30;index;comment:客户端名称"`
	AuthorityId string `json:"authorityId" gorm:"default:99999;comment:客户端角色ID"`
	Phone       string `json:"phone"  gorm:"not null;comment:联系手机号"`
	Email       string `json:"email"  gorm:"not null;comment:联系邮箱"`
	Addr        string `json:"addr"  gorm:"size:100;comment:联系地址"`
	Timestamps
	SoftDeletes
}

func (client Client) GetUid() string {
	return strconv.Itoa(int(client.ID.ID))
}

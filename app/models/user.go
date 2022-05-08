package models

import (
	"strconv"
)

type User struct {
	ID
	Uuid
	Name        string `json:"name" gorm:"size:30;not null;comment:用户名称"`
	Password    string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	NickName    string `json:"nickName" gorm:"not null;size:30;index;comment:用户昵称"`
	Phone       string `json:"phone" gorm:"size:24;not null;index;comment:用户手机号"`
	Email       string `json:"email"  gorm:"comment:联系邮箱"`
	Addr        string `json:"addr"  gorm:"size:100;comment:联系地址"`
	AuthorityId string `json:"authorityId" gorm:"default:99999;comment:角色ID"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}

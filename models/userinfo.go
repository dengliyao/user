package models

import "time"

type Userinfo struct {
	Id       int       `json:"id,omitempty"`       // 用户ID
	Name     string    `json:"name,omitempty"`     // 名称
	Tel      string    `json:"tel,omitempty"`      // 联系号码
	Addr     string    `json:"addr,omitempty"`     // 联系地址
	Birthday time.Time `json:"birthday"`           // 出生日期
	Password string    `json:"password,omitempty"` // 密码
}

var UsersDb []Userinfo

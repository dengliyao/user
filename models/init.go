package models

import (
	"fmt"
	"github.com/dengliyao/user/utils"
	"time"
)

// 管理员账号
func init() {
	adminUser := Userinfo{
		Id:       0,
		Name:     "admin",
		Tel:      "00000000000",
		Addr:     "admin",
		Birthday: time.Now(),
		Password: utils.Md5string("123123"),
	}
	UsersDb = append(UsersDb, adminUser)
	if err := NewJsonByUser.EncoderUsersDb(); err != nil {
		fmt.Println(err)
		return
	}
}

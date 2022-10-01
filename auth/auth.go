package auth

import (
	"fmt"
	"github.com/dengliyao/user/models"
	"github.com/dengliyao/user/utils"
)

func Auth() bool {
	var usersdb = models.NewUserInterface
	for i := 0; i < 3; i++ {
		tmpusername := utils.Input("请输入用户名: ")
		tmpPassword := utils.Md5string(utils.Input("请输入密码: "))
		if user, err := usersdb.GetUserName(tmpusername); err == nil {
			if user.Password == tmpPassword {
				return true
			} else {
				fmt.Println("用户名密码错误")
			}
		} else {
			fmt.Println(err)
			return false
		}
	}
	return false
}

package manager

import (
	"fmt"
	"github.com/dengliyao/user/models"
	"github.com/dengliyao/user/utils"
)

var routers = map[string]func(){}

func Register(op string, callback func()) {
	if _, ok := routers[op]; ok {
		panic("指令已存在")
	}
	routers[op] = callback
}

func Runc() {
	var usersdb = models.NewUserInterface
	for {
		text := utils.Input("请输入指令: ")
		if text == "exit" {
			tusersdb, err := usersdb.QueryUserAll()
			if err != nil {
				fmt.Println(err)
				return
			}
			models.TableWriterUser(tusersdb)
		}
		if action, ok := routers[text]; ok {
			action()
		} else {
			fmt.Println("指令错误")
		}
	}
}

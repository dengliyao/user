package controllers

import (
	"fmt"
	"github.com/dengliyao/user/models"
	"github.com/dengliyao/user/utils"
	"strconv"
)

var users = models.NewUserInterface
var usersjson = models.NewJsonByUser

func AddUser() {
	user := models.Userinfo{
		Id:       models.AutoId(),
		Name:     utils.Input("用户名: "),
		Tel:      utils.Input("联系电话: "),
		Addr:     utils.Input("联系地址: "),
		Birthday: utils.StringTime(utils.Input("出生日期(格式2006-01-02): ")),
		Password: utils.Md5string(utils.Input("设置密码: ")),
	}
	if _, err := users.GetUserName(user.Name); err == nil {
		fmt.Println(err)
		return
	}
	users.AddNewUser(user)
	fmt.Println("[+]添加有户成功")
}

func DelUser() {
	id, err := strconv.Atoi(utils.Input("请输入需要删除的用户ID: "))
	if err != nil {
		fmt.Println("请输入整数数字")
		return
	}

	user, err := users.GetUserId(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	users.DeleteUserById(user)
	fmt.Println("删除成功")
}

func ModifyUser() {
	id, err := strconv.Atoi(utils.Input("请输入要修改的用户ID号: "))
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpUsers := make([]models.Userinfo, 0)
	fmt.Println("你将要修改的用户信息如下:")
	tuser, err := users.GetUserId(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpUsers = append(tmpUsers, tuser)
	models.TableWriterUser(tmpUsers)

	user := models.Userinfo{
		Name:     utils.Input("用户名: "),
		Tel:      utils.Input("联系电话: "),
		Addr:     utils.Input("联系地址: "),
		Birthday: utils.StringTime(utils.Input("出生日期(格式2006-01-02): ")),
		Password: utils.Input("密码: "),
	}

	if _, err := users.GetUserName(user.Name); err == nil {
		fmt.Println("用户已存在!!!")
		return
	}
	users.ModifyUserById(user)

}

func QueryUser() {
	id, err := strconv.Atoi(utils.Input("请输入要查询的用户ID号: "))
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := users.GetUserId(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpUsers := make([]models.Userinfo, 0)
	fmt.Println("用户信息如下:")
	tmpUsers = append(tmpUsers, user)
	models.TableWriterUser(tmpUsers)
}

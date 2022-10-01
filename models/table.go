package models

import (
	"github.com/dengliyao/user/utils"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func TableWriterUser(users []Userinfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "名字", "联系电话", "联系地址", "出生日期", "密码MD5"})
	for _, u := range users {
		table.Append([]string{strconv.Itoa(u.Id), u.Name, u.Tel, u.Addr, utils.TimeString(u.Birthday), u.Password})
	}
	table.Render()
}

package routers

import (
	"github.com/dengliyao/user/controllers"
	"github.com/dengliyao/user/manager"
)

func init() {
	manager.Register("add", controllers.AddUser)
	manager.Register("delete", controllers.DelUser)
	manager.Register("modify", controllers.ModifyUser)
	manager.Register("query", controllers.QueryUser)
}

package main

import (
	"github.com/dengliyao/user/auth"
	"github.com/dengliyao/user/manager"
	_ "github.com/dengliyao/user/routers"
)

func main() {
	if auth.Auth() {
		manager.Runc()
	}
}

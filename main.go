package main

import "github.com/uadmin/uadmin"

func main() {
	uadmin.Register()

	uadmin.Port = 8080
	uadmin.RootURL = "/uadmin/"

	uadmin.StartServer()
}

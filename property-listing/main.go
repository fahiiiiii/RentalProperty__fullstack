package main

import (
	_ "property-listing/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}


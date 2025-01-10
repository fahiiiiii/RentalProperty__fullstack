package routers

import (
    "property-listing/controllers"

    "github.com/beego/beego/v2/server/web"
)

func init() {
    web.Router("/v1/property/list", &controllers.BookingController{}, "get:ListProperties")
    web.Router("/", &controllers.BookingController{}, "get:Index")
}
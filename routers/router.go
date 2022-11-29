package routers

import (
	"shopeeadapterapi/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.RESTRouter("/health", &controllers.HealthController{})
	beego.RESTRouter("/api/v1/orders/basics", &controllers.OrderListController{})
	beego.RESTRouter("/api/v1/orders/detail", &controllers.OrderDetailController{})
	beego.RESTRouter("/api/v2/logistics/get_channel_list", &controllers.LogisticsController{})
}

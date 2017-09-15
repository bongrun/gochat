package routers

import (
	"gochat/back/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// Register routers.
	beego.Router("/", &controllers.AppController{})

	ns := beego.NewNamespace("/user",
		beego.NSInclude(
			&controllers.UserController{},
		),
	)
	beego.AddNamespace(ns)

	// User routers.
	//beego.Router("/user", &controllers.UserController{})
	// Indicate AppController.Join method to handle POST requests.
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	// Long polling.
	beego.Router("/lp", &controllers.LongPollingController{}, "get:Join")
	beego.Router("/lp/post", &controllers.LongPollingController{})
	beego.Router("/lp/fetch", &controllers.LongPollingController{}, "get:Fetch")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
}

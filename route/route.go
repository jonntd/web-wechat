package route

import "github.com/gofiber/fiber/v2"

// InitRoute 初始化路由
func InitRoute(app *fiber.App) {
	// 初始化登录相关路由
	initLoginRoute(app)

	// 初始化用户相关路由
	initUserRoute(app)
}

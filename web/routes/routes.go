package routes

import (
	"gin-sites/controller"
	"gin-sites/controller/webhook"
	"gin-sites/middleware"

	"github.com/gin-gonic/gin"
)

// GET /todo 获取列表
// GET /todo/:todoId 获取待办事项详情
// Post /todo 创建一个待办事项
// PUT /todo/:todoId 更新待办事项
// DELETE /todo/:todoId 删除待办事项

// 状态码
// 200 OK - [GET]：服务器成功返回用户请求的数据，该操作是幂等的（Idempotent）。
// 201 CREATED - [POST/PUT/PATCH]：用户新建或修改数据成功。
// 202 Accepted - [*]：表示一个请求已经进入后台排队（异步任务）
// 204 NO CONTENT - [DELETE]：用户删除数据成功。
// 400 INVALID REQUEST - [POST/PUT/PATCH]：用户发出的请求有错误，服务器没有进行新建或修改数据的操作，该操作是幂等的。
// 401 Unauthorized - [*]：表示用户没有权限（令牌、用户名、密码错误）。
// 403 Forbidden - [*] 表示用户得到授权（与401错误相对），但是访问是被禁止的。
// 404 NOT FOUND - [*]：用户发出的请求针对的是不存在的记录，服务器没有进行操作，该操作是幂等的。
// 406 Not Acceptable - [GET]：用户请求的格式不可得（比如用户请求JSON格式，但是只有XML格式）。
// 410 Gone -[GET]：用户请求的资源被永久删除，且不会再得到的。
// 422 Unprocesable entity - [POST/PUT/PATCH] 当创建一个对象时，发生一个验证错误。
// 500 INTERNAL SERVER ERROR - [*]：服务器发生错误，用户将无法判断发出的请求是否成功。

var ApiVersion = "v1"

func RoutesInit(r *gin.Engine) {

	// 设置允许跨域的 header
	r.Use(middleware.CorsMiddleware)

	public := r.Group("/" + ApiVersion)
	auth := r.Group("/" + ApiVersion)

	// 校验 JWT 是否正确
	auth.Use(middleware.JWTMiddleware)

	publicRoute(public)
	authRoute(auth)

}

// 公共接口, 不需要鉴权部分
func publicRoute(public *gin.RouterGroup) {
	// 登录
	public.POST("/login", controller.LoginHandle)
	// 注册
	public.POST("/user", controller.CreateUserHandle)

	// github webhook
	public.POST("/api/webhook", webhook.GithubWebhookHandle)

}
func authRoute(auth *gin.RouterGroup) {
	// 获取用户详情
	auth.GET("/user/:userId", controller.GetUserHandle)

	// 待办事项路由
	auth.GET("/todo", controller.TodoListHandle)              // 获取 todo 列表
	auth.POST("/todo", controller.TodoAddHandle)              // 添加 一个 todo
	auth.PUT("/todo/:todoId", controller.TodoUpdateHandle)    // 更新一个 todo
	auth.DELETE("/todo/:todoId", controller.TodoDeleteHandle) // 删除

	auth.GET("/check-token", controller.CheckTokenHandle)

}

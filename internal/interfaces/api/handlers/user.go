package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/daheige/athena/internal/application"
)

// NewIndexHandler 创建一个 IndexHandler 实例
func NewIndexHandler(userService *application.UserService) *IndexHandler {
	return &IndexHandler{
		userService: userService,
	}
}

// IndexHandler ctrl handler
type IndexHandler struct {
	baseHandler
	userService *application.UserService
}

// Home 首页
func (h *IndexHandler) Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

// Foo api request
// 访问方式：http://localhost:1337/api/foo
func (h *IndexHandler) Foo(c *gin.Context) {
	h.Success(c, "ok", EmptyObject{})
}

// UserRequest 用户信息请求结构体
type UserRequest struct {
	Id int64 `json:"id" form:"id" binding:"required,min=1"`
}

// User 获取用户信息
// http://localhost:1337/api/user?id=1
func (h *IndexHandler) User(c *gin.Context) {
	req := &UserRequest{}
	if err := c.ShouldBind(req); err != nil {
		h.Error(c, 400, "id invalid", gin.H{
			"trace_error": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	user, err := h.userService.GetUser(ctx, req.Id)
	if err != nil {
		h.Error(c, 500, "failed to get user", gin.H{
			"trace_error": err.Error(),
		})
	}

	h.Success(c, "ok", user)
}

// BatchUsersRequest 批量获取信息请求结构体
type BatchUsersRequest struct {
	Ids []int64 `json:"ids" form:"ids" binding:"required"`
}

// BatchUsers 批量获取用户信息
// 请求方式
// curl --location 'http://localhost:1337/api/users' \
// --header 'Content-Type: application/json' \
//
//	--data '{
//	   "ids":[1]
//	}'
func (h *IndexHandler) BatchUsers(c *gin.Context) {
	req := &BatchUsersRequest{}
	if err := c.ShouldBind(req); err != nil {
		h.Error(c, 400, "param invalid", gin.H{
			"trace_error": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	users, err := h.userService.BatchUsers(ctx, req.Ids)
	if err != nil {
		h.Error(c, 500, "failed to get users by ids", gin.H{
			"trace_error": err.Error(),
		})
	}

	h.Success(c, "ok", users)
}

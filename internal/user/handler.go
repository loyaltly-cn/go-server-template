package user

import (
	"net/http"
	common "server/internal/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// CreateUser
// @Summary 创建用户
// @Tags User
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "用户信息"
// @Success 200 {object} common.Response
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.Error(err.Error()))
		return
	}

	user, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.Success(user))
}

// GetUser
// @Summary 获取用户
// @Tags User
// @Param id path int true "用户ID"
// @Success 200 {object} common.Response
// @Router /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(404, common.Error("not found"))
		return
	}

	c.JSON(200, common.Success(user))
}

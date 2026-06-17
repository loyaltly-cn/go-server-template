package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(
	c *gin.Context,
) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"msg": err.Error(),
			},
		)

		return
	}

	data, err := h.service.Login(req)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"msg": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		data,
	)
}

func (h *Handler) Me(c *gin.Context) {

	userIDVal, exists := c.Get("user_id")
	fmt.Print(exists)
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	userID := userIDVal.(int64)

	user, err := h.service.GetMe(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

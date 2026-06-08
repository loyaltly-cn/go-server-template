package banner

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

// Create
func (h *Handler) Create(c *gin.Context) {
	var req CreateBannerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.Error(err.Error()))
		return
	}

	if err := h.service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, common.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.Success(nil))
}

// List
func (h *Handler) List(c *gin.Context) {
	list, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.Success(list))
}

// Get
func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	data, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.Error("not found"))
		return
	}

	c.JSON(http.StatusOK, common.Success(data))
}

// Update
func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var req UpdateBannerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.Error(err.Error()))
		return
	}

	if err := h.service.Update(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, common.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.Success(nil))
}

// Delete
func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, common.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.Success(nil))
}

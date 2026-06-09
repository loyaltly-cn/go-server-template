package banner

import (
	"net/http"
	common "server/internal/common/response"
	"server/internal/query"
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

func (h *Handler) Patch(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var req PatchBannerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, common.Error(err.Error()))
		return
	}

	if err := h.service.Patch(id, req); err != nil {
		c.JSON(500, common.Error(err.Error()))
		return
	}

	c.JSON(200, common.Success(nil))
}

func (h *Handler) Query(c *gin.Context) {

	var req query.Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			common.Error(err.Error()),
		)
		return
	}

	data, err := h.service.Query(req)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			common.Error(err.Error()),
		)
		return
	}

	c.JSON(http.StatusOK, common.Success(data))
}

package banner

type CreateBannerRequest struct {
	Image string `json:"image" binding:"required,url"`
}

type UpdateBannerRequest struct {
	Image string `json:"image" binding:"required,url"`
}

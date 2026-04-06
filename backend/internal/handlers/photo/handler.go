package photo

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime/types"
	"net/http"
	api "photo-upload-service/internal/pkg/api/photo"
	"photo-upload-service/internal/service/photo"
)

type photoService interface {
	ProcessPhoto(ctx context.Context, file types.File) (*api.EvaluateSuccessResponse, error)
}

type UploadHandler struct {
	photoService photoService
}

func NewPhotoHandler(photoService *photo.Service) *UploadHandler {
	return &UploadHandler{
		photoService: photoService,
	}
}

func (h *UploadHandler) Evaluate(c *gin.Context) {
	ctx := c.Request.Context()
	var req api.EvaluateUploadRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ErrorResponse{
			Error:  err.Error(),
			Status: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	res, err := h.photoService.ProcessPhoto(ctx, req.Image)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{
			Error:  err.Error(),
			Status: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

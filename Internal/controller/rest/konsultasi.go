package rest

import (
	"bcc-geazy/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"net/http"
)

func (p *V1) GetKonsultasi(c *gin.Context) {
	userID, err := uuid.Parse(c.Query("user_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dokterID, err := uuid.Parse(c.Query("dokter_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	konsul, err := p.usecase.KonsultasiUsecase.GetPesan(ctx, userID, dokterID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, konsul)

}

func (p *V1) CreateKonsultasi(c *gin.Context) {
	var create model.BuatPesan

	if err := c.ShouldBindJSON(&create); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	konsul, err := p.usecase.KonsultasiUsecase.KirimPesan(ctx, create)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, konsul)
}

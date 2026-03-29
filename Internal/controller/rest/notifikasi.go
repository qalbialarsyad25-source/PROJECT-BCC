package rest

import (
	"bcc-geazy/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"errors"
	"net/http"
	"strconv"
)

func (p *V1) GetNotifikasi(c *gin.Context) {
	lembar, err := strconv.Atoi(c.DefaultQuery("lembar", "1"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	pagination := model.Pagination{
		Lembar: lembar,
		Limit:  limit,
	}

	pagination.Check()

	ctx := c.Request.Context()
	notip, err := p.usecase.NotifikasiUsecase.GetNotifikasi(ctx, pagination)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := model.PaginationResponse[model.NotifikasiResponse]{
		Data:       notip,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, response)

}

func (p *V1) CreateNotifikasi(c *gin.Context) {
	var create model.BuatNotifikasi

	userID := c.MustGet("userId").(uuid.UUID)

	err := c.ShouldBindJSON(&create)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	notip, err := p.usecase.NotifikasiUsecase.CreateNotifikasi(ctx, userID, create)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	p.wsManager.SendToUser(userID.String(), map[string]interface{}{
		"type":    "notification",
		"title":   "Notifikasi Baru",
		"message": "Kamu punya notifikasi baru",
		"data":    notip,
	})

	c.JSON(http.StatusCreated, notip)
}

func (p *V1) DeleteNotifikasi(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err = p.usecase.NotifikasiUsecase.DeleteNotifikasi(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pesan": "berhasil",
	})
}

func (p *V1) EditNotifikasi(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var edit model.EditNotifikasi
	err = c.ShouldBindJSON(&edit)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	err = p.usecase.NotifikasiUsecase.EditNotifikasi(ctx, id, edit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pesan": "berhasil",
	})
}

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

func (p *V1) GetLogInformasi(c *gin.Context) {
	userId := c.MustGet("userId").(uuid.UUID)
	informasiId := c.MustGet("informasiId").(uuid.UUID)

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
	LogInformasi, err := p.usecase.LogInformasiUsecase.GetLogInformasi(ctx, userId, informasiId, pagination)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := model.PaginationResponse[model.LogInformasiResponse]{
		Data:       LogInformasi,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, response)

}

func (p *V1) DeleteLogInformasi(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err = p.usecase.LogInformasiUsecase.DeleteLogInformasi(ctx, id)
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

func (p *V1) CreateLogInformasi(c *gin.Context) {
	var create model.BuatLogInformasi

	userID := c.MustGet("userId").(uuid.UUID)

	err := c.ShouldBindJSON(&create)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	informasi, err := p.usecase.LogInformasiUsecase.CreateLogInformasi(ctx, userID, create)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, informasi)
}

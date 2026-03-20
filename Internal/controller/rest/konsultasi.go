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

func (p *V1) GetKonsultasi(c *gin.Context) {
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
	konsul, err := p.usecase.KonsultasiUsecase.GetKonsultasi(ctx, pagination)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := model.PaginationResponse[model.KonsultasiResponse]{
		Data:       konsul,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, response)

}

func (p *V1) CreateKonsultasi(c *gin.Context) {
	var create model.BuatPesan

	err := c.ShouldBindBodyWithJSON(&create)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	konsul, err := p.usecase.KonsultasiUsecase.CreateKonsultasi(ctx, create)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, konsul)
}

func (p *V1) DeleteKonsultasi(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err = p.usecase.KonsultasiUsecase.DeleteKonsultasi(ctx, id)
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

func (p *V1) EditKonsultasi(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var edit model.EditPesan
	err = c.ShouldBindBodyWithJSON(&edit)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	err = p.usecase.KonsultasiUsecase.EditKonsultasi(ctx, id, edit)
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

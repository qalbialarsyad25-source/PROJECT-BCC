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

func (p *V1) GetDokter(c *gin.Context) {
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
	dokter, err := p.usecase.DokterUsecase.GetDokter(ctx, pagination)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := model.PaginationResponse[model.DokterResponse]{
		Data:       dokter,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, response)

}

func (p *V1) CreateDokter(c *gin.Context) {
	var create model.BuatUserDokter

	err := c.ShouldBindBodyWithJSON(&create)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	dokter, err := p.usecase.DokterUsecase.CreateDataDokter(ctx, create)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, dokter)
}

func (p *V1) DeleteDokter(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err = p.usecase.DokterUsecase.DeleteDokter(ctx, id)
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

func (p *V1) EditDokter(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var edit model.EditDokter
	err = c.ShouldBindBodyWithJSON(&edit)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	err = p.usecase.DokterUsecase.EditDokter(ctx, id, edit)
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

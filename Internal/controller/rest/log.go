package rest

import (
	"bcc-geazy/internal/model"

	"github.com/google/uuid"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
	"strconv"
)

func (p *V1) GetLog(c *gin.Context) {
	idParam := c.Param("id")
	AnakID, err := uuid.Parse(idParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

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
	log, err := p.usecase.LogUsecase.GetLog(ctx, AnakID, pagination)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := model.PaginationResponse[model.LogResponse]{
		Data:       log,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, response)

}

func (p *V1) CreateLog(c *gin.Context) {
	idParam := c.Param("id")
	AnakID, err := uuid.Parse(idParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var create model.BuatLog
	if err := c.ShouldBindJSON(&create); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	err = p.usecase.LogUsecase.CreateLog(ctx, AnakID, create)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, create)
}

func (p *V1) DeleteLog(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err = p.usecase.LogUsecase.DeleteLog(ctx, id)
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

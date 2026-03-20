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

func (p *V1) GetLogMakanan(c *gin.Context) {
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
	logma, err := p.usecase.LogMakananUsecase.GetLogMakanan(ctx, pagination)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := model.PaginationResponse[model.LogMakananResponse]{
		Data:       logma,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, response)

}

func (p *V1) DeleteLogMakanan(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err = p.usecase.LogMakananUsecase.DeleteLogMakanan(ctx, id)
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

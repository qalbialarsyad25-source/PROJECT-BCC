package rest

import (
	"bcc-geazy/internal/model"

	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

func (p *V1) GetLog(c *gin.Context) {
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
	log, err := p.usecase.LogUsecase.GetLog(ctx, pagination)
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

package rest

import (
	"bcc-geazy/internal/model"
	"strings"

	"path/filepath"

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

func (p *V1) UploadFotoDokter(c *gin.Context) {
	dokterID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "id tidak valid"})
		return
	}

	userID := c.MustGet("userId").(uuid.UUID)
	role := c.MustGet("role").(string)

	ctx := c.Request.Context()

	dokter, err := p.usecase.DokterUsecase.GetDokterByID(ctx, dokterID)
	if err != nil {
		c.JSON(404, gin.H{"error": "dokter tidak ditemukan"})
		return
	}

	if role != "admin" && dokter.UserId != userID {
		c.JSON(403, gin.H{"error": "tidak punya akses"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file tidak ditemukan"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		c.JSON(400, gin.H{"error": "format harus jpg/png"})
		return
	}

	if file.Size > 2*1024*1024 {
		c.JSON(400, gin.H{"error": "maksimal 2MB"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "gagal buka file"})
		return
	}
	defer f.Close()

	filename := uuid.New().String() + ext
	url, err := p.usecase.DokterUsecase.UploadFoto(ctx, dokterID, f, filename)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "foto berhasil diupload",
		"url":     url,
	})
}

func (p *V1) GetDetailDokter(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "id tidak valid"})
		return
	}

	ctx := c.Request.Context()

	dokter, err := p.usecase.DokterUsecase.GetDokterByID(ctx, id)
	if err != nil {
		c.JSON(404, gin.H{"error": "dokter tidak ditemukan"})
		return
	}

	res := model.ToDokterResponse(dokter)

	c.JSON(200, res)
}

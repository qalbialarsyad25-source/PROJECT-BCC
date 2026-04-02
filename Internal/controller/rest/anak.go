package rest

import (
	"bcc-geazy/internal/model"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"errors"
	"net/http"
	"strconv"
)

func (p *V1) GetDataAnak(c *gin.Context) {
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
	anak, err := p.usecase.AnakUsecase.GetDataAnak(ctx, pagination)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := model.PaginationResponse[model.AnakResponse]{
		Data:       anak,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, response)

}

func (p *V1) CreateDataAnak(c *gin.Context) {
	var buat model.TambahDataAnak

	userID := c.MustGet("userId").(uuid.UUID)

	err := c.ShouldBindBodyWithJSON(&buat)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	anak, err := p.usecase.AnakUsecase.CreateDataAnak(ctx, buat, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, anak)
}

func (p *V1) DeleteDataAnak(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err = p.usecase.AnakUsecase.DeleteDataAnak(ctx, id)
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

func (p *V1) EditDataAnak(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var edit model.EditDataAnak
	err = c.ShouldBindBodyWithJSON(&edit)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	err = p.usecase.AnakUsecase.EditDataAnak(ctx, id, edit)
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

func (p *V1) GetGenderOptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": []map[string]string{
			{"label": "Laki-Laki", "value": "laki-laki"},
			{"label": "Perempuan", "value": "perempuan"},
		},
	})
}

func (p *V1) GetGolonganOption(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": []map[string]string{
			{"label": "O", "value": "O"},
			{"label": "A", "value": "A"},
			{"label": "B", "value": "B"},
			{"label": "AB", "value": "AB"},
		},
	})
}

func (p *V1) GetAnakKeOptions(c *gin.Context) {
	labels := []string{
		"Pertama",
		"Kedua",
		"Ketiga",
	}

	var pilihan []map[string]any

	for i, label := range labels {
		pilihan = append(pilihan, map[string]any{
			"label": "Anak " + label,
			"value": i + 1,
		})
	}

	c.JSON(200, gin.H{
		"data": pilihan,
	})
}

func (p *V1) GetNutrisiHarian(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	result, err := p.usecase.NutrisiUsecase.GetNutrisiHarian(ctx, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (r *V1) UploadFotoAnak(c *gin.Context) {
	userID := c.MustGet("userId").(uuid.UUID)
	anakID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid anak id"})
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

	ctx := c.Request.Context()

	url, err := r.usecase.AnakUsecase.UploadFotoAnak(ctx, userID, anakID, f, filename)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "foto berhasil diupload",
		"url":     url,
	})
}

func (r *V1) GetProfileAnak(c *gin.Context) {
	userID := c.MustGet("userId").(uuid.UUID)

	anakID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid anak id"})
		return
	}

	ctx := c.Request.Context()

	anak, err := r.usecase.AnakUsecase.GetProfileAnak(ctx, userID, anakID)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, anak)
}

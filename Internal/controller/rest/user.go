package rest

import (
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *V1) GetProfile(c *gin.Context) {
	userID := c.MustGet("userId").(uuid.UUID)

	ctx := c.Request.Context()

	user, err := r.usecase.UserUsecase.GetProfile(ctx, userID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (r *V1) UploadFotoUser(c *gin.Context) {
	userID := c.MustGet("userId").(uuid.UUID)

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

	url, err := r.usecase.UserUsecase.UploadFotoUser(ctx, userID, f, filename)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "foto berhasil diupload",
		"url":     url,
	})
}

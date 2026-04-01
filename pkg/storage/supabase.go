package storage

import (
	"mime/multipart"
	"os"

	"net/http"

	storage_go "github.com/supabase-community/storage-go"
)

func NewStorageClient() *storage_go.Client {
	storageURL := os.Getenv("SUPABASE_URL") + "/storage/v1"
	storageKey := os.Getenv("SUPABASE_SERVICE_KEY")

	storageClient := storage_go.NewClient(storageURL, storageKey, nil)
	return storageClient
}

func UploadFile(file multipart.File, filename string) (string, error) {
	client := NewStorageClient()

	bucket := os.Getenv("SUPABASE_BUCKET")

	buffer := make([]byte, 512)
	file.Read(buffer)
	file.Seek(0, 0)

	contentType := http.DetectContentType(buffer)
	upsert := true

	_, err := client.UploadFile(
		bucket,
		filename,
		file,
		storage_go.FileOptions{
			ContentType: &contentType,
			Upsert:      &upsert,
		},
	)
	if err != nil {
		return "", err
	}

	publicURL := os.Getenv("SUPABASE_URL") +
		"/storage/v1/object/public/" + bucket + "/" + filename

	return publicURL, nil
}

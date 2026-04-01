package storage

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadToSupabase(file multipart.File, filename string) (string, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	url := os.Getenv("SUPABASE_URL") + "/storage/v1/object/" + os.Getenv("SUPABASE_BUCKET") + "/" + filename

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_SERVICE_KEY"))
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("upload gagal: %s", resp.Status)
	}

	publicURL := os.Getenv("SUPABASE_URL") + "/storage/v1/object/public/" + os.Getenv("SUPABASE_BUCKET") + "/" + filename

	return publicURL, nil
}

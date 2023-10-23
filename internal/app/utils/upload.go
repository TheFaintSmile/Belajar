package utils

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	config "github.com/rumbel/belajar/internal/config"
)

func UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create Cloudinary instance
	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	resp, err := cld.Upload.Upload(ctx, fileHeader, uploader.UploadParams {
        PublicID: fileHeader.Filename,
    })

	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
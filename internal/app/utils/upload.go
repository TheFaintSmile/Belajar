package utils

import (
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	config "github.com/rumbel/belajar/internal/config"
)

func UploadFile(input interface{}) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    //create cloudinary instance
    cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
    if err != nil {
        return "", err
    }
    //upload file
    uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})

    if err != nil {
        return "", err
    }
    return uploadParam.SecureURL, nil
}
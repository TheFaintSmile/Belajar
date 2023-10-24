package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	config "github.com/rumbel/belajar/internal/config"
)

func UploadFile(fileHeader interface{}) (string, error) {
	fmt.Println("masuk kesini errornya")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}
	fmt.Println("filenya ini", fileHeader)
	fmt.Println("masuk kesini errornya 2")
	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, fileHeader, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})
	if err != nil {
		return "", err
	}

	fmt.Println("masuk kesini errornya 3")
	return uploadParam.SecureURL, nil
}

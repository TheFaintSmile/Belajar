package controller

import (
	"github.com/eugeniusms/poseidon/entity"
	"github.com/eugeniusms/poseidon/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video 
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

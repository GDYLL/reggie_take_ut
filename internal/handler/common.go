package handler

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

type CommonController struct{}

var path = "static/images/"

func (c CommonController) Download(context *gin.Context) {
	image_name := context.Query("name")
	imgUrl := filepath.Join(path, image_name)
	context.FileAttachment(imgUrl, image_name)
}

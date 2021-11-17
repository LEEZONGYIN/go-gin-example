package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/pkg/app"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/upload"
	"net/http"
)

func UploadImage(c *gin.Context) {
	appG := app.Gin{c}

	//获取上传的图片（返回提供的表单键的第一个文件）
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, e.ERROR, nil)
	}

	if image == nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullpath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullpath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusOK, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullpath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	err = c.SaveUploadedFile(image, src)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})

}

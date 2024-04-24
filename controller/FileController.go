package controller

import (
	"github.com/gin-gonic/gin"
	"path"
	"remember/common"
	"remember/config"
	"remember/utils"
	"strconv"
	"strings"
	"time"
)

func UpLoad(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	fileName := strings.Replace(file.Filename, "-", "", -1)
	name := strconv.FormatInt(time.Now().Unix(), 10) + "-" + fileName
	user := utils.GetUser(c)
	dst := path.Join(config.Configure.Remember.FilePath, user.Username)
	err = utils.MakeDir(dst)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	err = c.SaveUploadedFile(file, path.Join(dst, name))
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage("文件上传失败"))
		return
	}
	c.JSON(200, common.StatusOk().SetMessage("文件上传成功"))
}

func GetFilesList(c *gin.Context) {
	user := utils.GetUser(c)
	source := path.Join(config.Configure.Remember.FilePath, user.Username)
	err := utils.MakeDir(source)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	files := utils.ListDir(source)
	data := make([]*common.FileInfo, 0)
	for _, i := range files {
		data = append(data, common.NewFileInfo(i, user.Username))
	}
	c.JSON(200, common.StatusOk().SetMessage("获取成功").AddData("files", data))
}

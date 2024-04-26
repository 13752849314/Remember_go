package controller

import (
	"github.com/gin-gonic/gin"
	"log"
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
	var f *common.FileInfo
	for _, i := range files {
		f, err = common.NewFileInfo(i, user.Username)
		if err == nil {
			data = append(data, f)
		}
	}
	c.JSON(200, common.StatusOk().SetMessage("获取成功").AddData("files", data))
}

func Download(c *gin.Context) {
	timestamp := c.Param("flag")
	user := utils.GetUser(c)
	source := path.Join(config.Configure.Remember.FilePath, user.Username)
	files := utils.ListDir(source)
	var filename string = ""
	for _, info := range files {
		if strings.Contains(info.Name(), timestamp) {
			filename = info.Name()
			break
		}
	}
	if filename == "" {
		c.JSON(200, common.StatusErr().SetMessage("文件不存在"))
		return
	}
	// 获取下载名
	downName := strings.Split(filename, "-")[1]
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+downName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(path.Join(source, filename))
	log.Printf("用户：%s-下载%s成功\n", user.Username, filename)
}

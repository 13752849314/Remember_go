package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"remember/common"
	"remember/entity"
	"remember/service/impl"
	"remember/utils"
	"strconv"
)

var bs *impl.BillServiceImpl

func init() {
	bs = new(impl.BillServiceImpl)
}

func GetBillsByUsername(c *gin.Context) {
	user, _ := c.Get("user")
	bills := bs.GetBillsByUsername(user.(*entity.User).Username)
	c.JSON(200, common.StatusOk().SetMessage("获取成功").AddData("bills", bills))
}

func AddBill(c *gin.Context) {
	bill := new(entity.Bill)
	err := c.ShouldBindJSON(bill)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	user, _ := c.Get("user")
	bill.Username = user.(*entity.User).Username
	err = bs.AddBill(bill)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	log.Printf("账单添加：%+v\n", bill)
	c.JSON(200, common.StatusOk().SetMessage("添加成功"))
}

func DeleteBillById(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage("请求参数错误"))
		return
	}
	err = bs.DeleteBillById(id)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	log.Printf("账单删除：id=%d\n", id)
	c.JSON(200, common.StatusOk().SetMessage("删除成功"))
}

func ChangeBillInfoById(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage("请求参数错误"))
		return
	}
	bi := new(common.ChangeBillI)
	err = c.ShouldBindJSON(bi)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	mp := utils.Struct2Map(bi)
	err = bs.ChangeBillInfoById(id, mp)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	log.Println("修改信息为：", mp)
	c.JSON(200, common.StatusOk().SetMessage("信息修改成功"))
}

package job

import (
	"github.com/robfig/cron"
	"log"
	"remember/utils"
)

func AddUserJob(c *cron.Cron) *cron.Cron {
	err := c.AddFunc("* * 0/2 * * ?", utils.MaintainLogin)
	if err != nil {
		log.Println(err.Error())
	}
	return c
}

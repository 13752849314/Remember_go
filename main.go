package main

import (
	"github.com/robfig/cron"
	"log"
	"remember/config"
	"remember/job"
	"remember/router"
	"strconv"
)

func main() {
	log.Println("Hello Remember_go!")
	c := cron.New()
	c = job.AddUserJob(c)
	c.Start()
	defer c.Stop()
	r := router.Remember()
	err := r.Run(":" + strconv.Itoa(config.Configure.Service.Port))
	if err != nil {
		return
	}
}

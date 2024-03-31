package main

import (
	"fmt"
	"remember/config"
	"remember/router"
	"strconv"
)

func main() {
	fmt.Println("Hello Remember_go!")

	r := router.Remember()
	err := r.Run(":" + strconv.Itoa(config.Configure.Service.Port))
	if err != nil {
		return
	}
}

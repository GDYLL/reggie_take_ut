package main

import (
	"reggie_take_ut/config"
	"reggie_take_ut/router"
)

func main() {

	config.InitConfig()
	router.InitRouter()

}

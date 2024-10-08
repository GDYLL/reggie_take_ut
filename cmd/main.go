package main

import (
	"reggie_take_ut/config"
	"reggie_take_ut/internal/router"
)

func main() {

	config.InitConfig()
	router.InitRouter()

}

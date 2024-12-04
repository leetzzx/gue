package cmd

import (
	"awesomeProject/conf"
	"awesomeProject/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clear() {
	println("==========Clean==============")
}

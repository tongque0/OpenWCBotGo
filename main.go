package main

import (
	"openWcBotGo/config"
	"openWcBotGo/instance"
)

func main() {
	config.InitEnv()
	instance.Init()
}

package main

import (
	"github.com/coolnut12138/my-godis/conf"
	"os"
)

var startBanner = `
 __      __  ________            ________           .___.__        
/  \    /  \/  _____/           /  _____/  ____   __| _/|__| ______
\   \/\/   /   \  ___   ______ /   \  ___ /  _ \ / __ | |  |/  ___/
 \        /\    \_\  \ /_____/ \    \_\  (  <_> ) /_/ | |  |\___ \ 
  \__/\  /  \______  /          \______  /\____/\____ | |__/____  >
       \/          \/                  \/            \/         \/ 
`

func fileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	// print introduction
	print(startBanner)
	// init log

	// init config (首先从CONFIG环境变量获取配置文件路径，如果环境变量中未设置配置文件路径，
	// 则尝试读取工作路径中的redis.conf文件，若redis.conf不存在则使用自带的默认配置)
	configName := os.Getenv("SHELL")
	if configName == "" {
		if exist, err := fileExists("redis.conf"); exist && err == nil {
			conf.SetupConfig("redis.conf")
		} else {

		}
	}
	// init server

}
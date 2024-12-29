package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/daheige/athena/cmd/job/commands"
	"github.com/daheige/athena/internal/infras/config"
	"github.com/daheige/athena/internal/infras/logger"
)

func main() {
	log.Println("athena-job...")
	// 初始化配置
	cobra.OnInitialize(initConfig)

	// 命令初始化
	commands.Init()

	// 执行命令
	commands.Execute()
}

func initConfig() {
	// 读取配置文件，并初始化redis和mysql
	conf := config.InitAppConfig()

	// 初始化日志配置，如果app_debug = true 将日志输出到终端中
	logger.Default(logger.WithLogFilename("athena-web.log"), logger.WithStdout(conf.AppDebug))
}

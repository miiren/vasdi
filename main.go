package main

import (
	"flag"
	"log"
	"vasdi/adapter/job"
	"vasdi/adapter/router"
	"vasdi/infrastructure/cnf"
	log2 "vasdi/infrastructure/log"
	"vasdi/infrastructure/source"

	"github.com/miiren/mutil/mtool"
)

var appBasePath string

func init() {
	appBasePath = getAppBasePath()
	appInit(flagConfigVar())
}

func main() {
	r := (&router.GinWebRun{}).New(appBasePath)
	err := r.Run()
	if err != nil {
		log.Panicf("Server Err: %v", err)
	}
}

func getAppBasePath() string {
	curPath := mtool.GetCurrentAbPath(1)

	return curPath + "/"
}

func flagConfigVar() string {
	//获取控制台参数
	var mode string
	flag.StringVar(&mode, "m", "test", "application run mode, Default: `test`")
	flag.Parse()
	configPath := appBasePath + "resources/config-" + mode + ".yaml"
	return configPath
}

// 对项目进行初始化
func appInit(configPath string) {
	//configInit
	err := cnf.InitCnf(configPath)
	if err != nil {
		log.Panicf("init.InitCnf err: %v", err)
	}
	//初始化项目logger
	log2.InitLogger(appBasePath + "/runtimes/")
	//cache初始化
	source.InitCache()
	//数据库初始化
	source.InitDB(log2.LoggerDb)
	source.InitRedis()
	//初始化job
	job.InitJob()
}

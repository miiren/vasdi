package main

import (
	"btool/adapter/router"
	"btool/data/po"
	"btool/infrastructure/initialize"
	"btool/infrastructure/pkg/utool"
	"btool/infrastructure/source"
	"flag"
	"fmt"
	"log"
)

var appBasePath string

func init() {
	appBasePath = getAppBasePath()
	initialize.AppInit(flagConfigVar())
}

func main() {
	r := (&router.GinWebRun{}).New(appBasePath)
	err := r.Run()
	if err != nil {
		log.Panicf("Server Err: %v", err)
	}
}

func getAppBasePath() string {
	curPath := utool.GetCurrentAbPath(1)

	return curPath + "/"
}

func flagConfigVar() string {
	//获取控制台参数
	var mode string
	flag.StringVar(&mode, "m", "pro", "application run mode, Default: `pro`")
	flag.Parse()
	configPath := appBasePath + "resources/config-" + mode + ".yaml"
	return configPath
}

func initDbTable() {
	err := source.DB.AutoMigrate(
		&po.AppInfo{}, &po.ApiInfo{},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ok")
}

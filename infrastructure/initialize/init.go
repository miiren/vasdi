package initialize

import (
	"log"
	"vasdi/adapter/job"
	"vasdi/infrastructure/cnf"
	log2 "vasdi/infrastructure/log"
	"vasdi/infrastructure/source"
)

// 对项目进行初始化
func AppInit(configPath string) {
	//configInit
	err := cnf.InitCnf(configPath)
	if err != nil {
		log.Panicf("init.InitCnf err: %v", err)
	}
	//cache初始化
	source.InitCache()
	//数据库初始化
	source.InitDB(log2.LoggerDb)
	source.InitRedis()
	//初始化job
	job.InitJob()
}

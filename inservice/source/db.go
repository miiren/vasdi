package source

import (
	"fmt"
	"log"
	"time"
	"vasdi/infrastructure/cnf"

	"github.com/miiren/mbox/orm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB    *gorm.DB
	PayDB *gorm.DB
)

func InitDB(writer orm.Writer) {
	var err error
	//初始化本项目db
	DB, err = creatDbConn(writer, cnf.GetCnf().Mysql)
	if err != nil {
		log.Panic(err.Error())
	}
}

func creatDbConn(writer orm.Writer, dbCnf *cnf.MysqlConfig) (*gorm.DB, error) {
	ormConfig := orm.Config{
		MaxIdleConns: dbCnf.MaxIdleConns,
		MaxOpenConns: dbCnf.MaxOpenConns,
		UNamingStrategy: schema.NamingStrategy{
			//TablePrefix: "bam_",
			SingularTable: true,
			NoLowerCase:   false,
		},
		LoggerWriter: writer,
		LoggerConfig: logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	}
	dbConn, err := orm.New(
		fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			dbCnf.UserName,
			dbCnf.Password,
			dbCnf.Host,
			dbCnf.Port,
			dbCnf.DBName,
		),
		ormConfig,
	)
	if err != nil {
		return nil, fmt.Errorf("orm数据库链接错误：%v", err)
	}
	sqlDb, _ := dbConn.DB()
	err = sqlDb.Ping()
	if err != nil {
		return nil, fmt.Errorf("orm数据库ping错误：%v", err)
	}
	return dbConn, nil
}

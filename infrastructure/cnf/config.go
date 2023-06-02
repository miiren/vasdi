package cnf

import (
	"log"
	"time"

	set "github.com/miiren/mbox/setter"
)

// 定义全局变量
var (
	cnf    *configs
	setter *set.Setter
)

func GetCnf() *configs {
	return cnf
}

type configs struct {
	Server  *ServerConfig
	Mysql   *MysqlConfig
	GoCache *GoCacheConfig
	URedis  *URedisConfig
}

// 服务器配置
type ServerConfig struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// mysql数据库配置
type MysqlConfig struct {
	UserName     string
	Password     string
	Host         string
	Port         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// GoCache配置
type GoCacheConfig struct {
	DefaultExpiration int //默认有效期
	CleanupInterval   int //定期清理时间
}

// URedis配置
type URedisConfig struct {
	Addr     string //连接串 如 localhost:6379
	Password string //密码设置 no password set ""
	DB       int    //redis db set
}

func InitCnf(configFile string) error {
	var err error
	setter, err = set.NewSetter(configFile)
	if err != nil {
		return err
	}

	//setting
	err = setter.ReadCnf(&cnf)
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
		return err
	}

	return nil
}

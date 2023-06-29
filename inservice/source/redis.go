package source

import (
	"fmt"
	"log"
	"vasdi/inservice/cnf"

	"github.com/miiren/mbox/redis"
)

var (
	Redis *redis.URClient
)

func InitRedis() {
	var err error
	Redis, err = creatRedisConn(cnf.GetCnf().URedis)
	if err != nil {
		log.Panic(err)
	}
}

func creatRedisConn(redisCnf *cnf.URedisConfig) (*redis.URClient, error) {
	redisConn, err := redis.NewClient(redisCnf.Addr, redisCnf.Password, redisCnf.DB)
	if err != nil {
		return nil, fmt.Errorf("db engine err: %v", err)
	}

	return redisConn, nil
}

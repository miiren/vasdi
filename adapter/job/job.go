package job

import (
	"github.com/miiren/mbox/cron"
	"log"
)

func InitJob() {
	c := cron.New()
	var err error
	err = c.AddJob("*/10 * * * * *", &helloJob{})
	if err != nil {
		log.Panicf("job init err: %v", err)
	}
	c.Start()
}

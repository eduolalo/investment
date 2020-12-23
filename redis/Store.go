package redis

import (
	"log"
	"strings"

	"github.com/kalmecak/go-redis-connector"
	"github.com/mediocregopher/radix/v3"
)

// Store agrega el conteo de las asignaciones exitosas
func Store(ok bool, amount string) {

	var key strings.Builder
	if ok {
		key.WriteString("success")
	} else {
		key.WriteString("fail")
	}
	db := redis.Connect()
	if err := db.Do(radix.Cmd(nil, "INCR", key.String())); err != nil {

		log.Println("*** redis.Success.db.Do ***")
		log.Println(err.Error())
		log.Println("*** redis.Success.db.Do ***")
	}

	if amount == "0" {
		return
	}

	key.WriteString("-investment")
	if err := db.Do(radix.Cmd(nil, "INCRBY", key.String(), amount)); err != nil {

		log.Println("*** redis.Success.db.Do ***")
		log.Println(err.Error())
		log.Println("*** redis.Success.db.Do ***")
	}
}

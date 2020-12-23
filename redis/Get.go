package redis

import (
	"log"

	"github.com/kalmecak/go-redis-connector"
	"github.com/mediocregopher/radix/v3"
)

// Get obtiene los valores de un dato de redis
func Get(key string) (result string) {

	db := redis.Connect()

	if err := db.Do(radix.Cmd(&result, "GET", key)); err != nil {

		log.Println("*** redis.Get.db.Do ***")
		log.Println(err.Error())
		log.Println("*** redis.Get.db.Do ***")
	}
	return
}

package server

import (
	"kolo-assignment/redis"
)

func Init() {

	//config := config.Get()
	redis.Init()

	r := NewRouter()
	r.Run(":" + "4000")
}

package server

import (
	"github.com/gin-gonic/gin"
	comics "kolo-assignment/pkg/comics/v1"
	"kolo-assignment/redis"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	redis := redis.Client()

	pingGroup := router.Group("comics")
	//test API to test if everything works as planned
	pingGroup.GET("/", comics.NewComicModule(redis).GetController().Hello)

	//fetches all characters
	pingGroup.GET("/getCharacters", comics.NewComicModule(redis).GetController().GetComics)

	//searches characters from user query
	pingGroup.GET("/searchCharacter", comics.NewComicModule(redis).GetController().SearchCharacter)

	return router
}

package server

import (
	"github.com/gin-gonic/gin"
	comics "kolo-assignment/pkg/comics/v1"
)

func NewRouter() *gin.Engine {

	router := gin.New()

	pingGroup := router.Group("comics")
	//test API to test if everything works as planned
	pingGroup.GET("/", comics.NewComicModule().GetController().Hello)

	//fetches all characters
	pingGroup.GET("/getCharacters", comics.NewComicModule().GetController().GetComics)

	//searches characters from user query
	pingGroup.GET("/searchCharacter", comics.NewComicModule().GetController().SearchCharacter)

	return router
}

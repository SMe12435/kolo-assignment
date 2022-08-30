package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ComicController struct {
	service *ComicService
}

func NewComicController() *ComicController {
	return &ComicController{}
}

func (s *ComicController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": s.service.Hello()})
}

func (s *ComicController) GetComics(c *gin.Context) {
	c.Request.URL.Query()
	comics, err := s.service.GetComic()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"Comics": comics})
}

func (s *ComicController) SearchCharacter(c *gin.Context) {

	searchKey := c.Request.URL.Query().Get("searchKey")
	offset := c.Request.URL.Query().Get("offset")
	limit := "10"
	characters, err := s.service.SearchCharacter(searchKey, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}
	//fmt.Println(characters)
	c.JSON(http.StatusOK, gin.H{"characters": characters.Characters, "offset": characters.Offset})
}

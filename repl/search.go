package repl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type maintenance struct {
	searchQuery string
	offset      int
	size        int
}

var maint maintenance

//searchResponse struct to take care of search response
type SearchResponse struct {
	Characters []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"characters"`
	Offset int `json:"offset"`
}

//searches a marvel character, offset param is used for pagination
func searchMarvelCharacter(searchQuery string, offset string) {
	var searchData SearchResponse

	//calls API(local) to fetch marvel characters
	resp, err := http.Get("http://localhost:4000/comics/searchCharacter?searchKey=" + searchQuery + "&offset=" + offset)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {

		if err != nil {
			log.Fatal(err)
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		err = json.Unmarshal(bodyBytes, &searchData)
		if err != nil {
			panic(err)
		}
		if len(searchData.Characters) == 0 {
			log.Println("No More Data")
		} else {
			log.Println(bodyString)

			//updates maintenance params
			maint.searchQuery = searchQuery
			maint.offset = searchData.Offset
			maint.size = len(searchData.Characters)
		}
	}
}

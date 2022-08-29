package repl

import (
	"encoding/json"
	"fmt"
	"github.com/abiosoft/ishell/v2"
	"github.com/blevesearch/bleve/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//struct to maintain offset and searchQuery for next and previous functionality
type maintenance struct {
	searchQuery string
	offset      int
	size        int
}

//searchResponse struct to take care of search response
type SearchResponse struct {
	Characters []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"characters"`
	Offset int `json:"offset"`
}

var maint maintenance

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

		log.Println(bodyString)

		//updates maintenance params
		maint.searchQuery = searchQuery
		maint.offset = searchData.Offset
		maint.size = len(searchData.Characters)

	}
}

//used bleeve, an open source library to search anywhere in the datas
func search() {
	resp, err := http.Get("http://127.0.0.1:4000/ping/getComic?")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("myPath.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}

	// index some data
	index.Index("abc", data)

	// search for some text
	query := bleve.NewWildcardQuery("*ext")
	search := bleve.NewSearchRequestOptions(query, 1, 0, false)
	search.Fields = []string{"*"}
	searchResults, err := index.Search(search)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
	fmt.Println(searchResults.Hits[0].Fields)
}

func Init() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("Welcome to Marvel Universe. What do you want to do today. Try running help command for more")

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "NS",
		Help: "NS(new search) searches the Marvel Directory of Comics",
		Func: func(c *ishell.Context) {
			// disable the '>>>' for cleaner same line input.
			c.ShowPrompt(false)
			defer c.ShowPrompt(true) // yes, revert after login.

			//gets Character Search Query from User
			c.Print("Character Search Query: ")
			searchQuery := c.ReadLine()

			//searchesMarvelCharacter for the given query
			searchMarvelCharacter(searchQuery, "0")
			c.Println("\n type N for next | P for Previous | new for new search | exit to exit | help for some help")

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "N",
		Help: "Fetches the next data available",
		Func: func(c *ishell.Context) {
			// disable the '>>>' for cleaner same line input.
			c.ShowPrompt(false)
			defer c.ShowPrompt(true) // yes, revert after login.

			//fetches desired offset for next page
			nextOffset := maint.offset + maint.size

			searchMarvelCharacter(maint.searchQuery, strconv.Itoa(nextOffset))
			c.Println("\n type N for next | P for Previous | new for new search | exit to exit | help for some help")

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "P",
		Help: "Fetches the previous data available",
		Func: func(c *ishell.Context) {
			// disable the '>>>' for cleaner same line input.
			c.ShowPrompt(false)
			defer c.ShowPrompt(true) // yes, revert after login.

			//fetches desired offset for previous offset
			prevOffset := maint.offset - maint.size

			searchMarvelCharacter(maint.searchQuery, strconv.Itoa(prevOffset))
			c.Println("\n type N for next | P for Previous | new for new search | exit to exit | help for some help")

		},
	})

	// run shell
	shell.Run()
}

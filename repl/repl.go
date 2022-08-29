package repl

import (
	"github.com/abiosoft/ishell/v2"
	"strconv"
)

//struct to maintain offset and searchQuery for next and previous functionality

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

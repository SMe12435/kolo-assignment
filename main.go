package main

import (
	_ "io/ioutil"
	"kolo-assignment/repl"
	//"github.com/gin-gonic/gin"
	_ "kolo-assignment/config"
	"kolo-assignment/server"
)

func main() {
	//using routine to initialise the API server. Used this approach so that
	go server.Init()

	//initializes the REPL
	repl.Init()

}

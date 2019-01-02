package main

import (
	"TimeService/svc/server"
	"log"
)

func main()  {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	server.Run()
}



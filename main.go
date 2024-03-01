package main

import (
	"log"
	"github.com/Kumar-pai/go-restfulapi-practices/api"
)

func main()  {
	svr, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	svr.Start()
}
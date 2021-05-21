package main

import (
	"github.com/msroz/sparring-hasura/actions/route"
)

func main() {

	router := route.Init()
	router.Logger.Fatal(router.Start(":9876"))
}


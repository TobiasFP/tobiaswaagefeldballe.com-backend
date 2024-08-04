package main

import (
	routes "backend/routes"
	"flag"
)

func main() {
	flag.Parse()

	routes.StartGin()
}

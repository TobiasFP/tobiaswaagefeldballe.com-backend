package main

import (
	"backend/config"
	routes "backend/routes"
	"flag"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*environment)
	routes.StartGin()
}

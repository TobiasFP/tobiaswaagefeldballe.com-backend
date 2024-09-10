package main

import (
	routes "backend/routes"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Println("Just getting started")
	routes.StartGin()
}

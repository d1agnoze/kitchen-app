package main

import (
	"fmt"
	"log"

	"github.com/d1agnoze/kitchen/services/gateway/routers"
)

func main() {
	fmt.Println("Starting server")
	log.Fatal(routers.Start())
}

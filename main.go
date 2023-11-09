package main

import (
	"fmt"
	"log"
	"net/http"
	"hey/routers"
)

func main() {
	fmt.Println("Hello, World!")

	r := routers.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
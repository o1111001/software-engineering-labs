package main

import (
	"fmt"
	"log"
	"net/http"

	"software-engineering-labs/constants"
	"software-engineering-labs/handlers/time"
)

func main() {
	http.HandleFunc("/time", time.GetJson)

	fmt.Println("Listening on", constants.PORT)
	log.Fatal(http.ListenAndServe(constants.PORT, nil))
} 
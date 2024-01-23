package main

import (
	"fmt"
	"log"
	"net/http"

	//"html/template"
	"cyoa"
	"flag"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "Port to start the CYOA App")
	jsonPath := flag.String("path", "gopher.json", "File path for JSON Story")

	flag.Parse()

	jsonStoryData, err := os.ReadFile(*jsonPath)
	if err != nil {
		fmt.Println("Error Line 22:", err)
	}

	storyData, err := cyoa.JSONParse(jsonStoryData)
	if err != nil {
		fmt.Println("Error Line 27:", err)
	}

	//fmt.Printf("%+v\n", storyData)

	handler := cyoa.NewHandler(storyData)
	fmt.Println("Starting server on port ", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))

}

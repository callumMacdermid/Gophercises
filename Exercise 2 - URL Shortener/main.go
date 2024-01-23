package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"url-shorten/urlshort"
)

func main() {

	// var yamlPath = flag.String("yaml-path", "", "Path to YAML URL File")
	var jsonPath = flag.String("json-path", "", "Path to JSON URL File")
	flag.Parse()

	// fmt.Println(yamlPath)
	// yamlFile, err := os.ReadFile(*yamlPath)
	// if err != nil {
	// 	fmt.Print("Error reading YAML File from Path ", string(*yamlPath))
	// }

	jsonFile, err := os.ReadFile(*jsonPath)
	if err != nil {
		fmt.Print(err)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//fmt.Print(mapHandler)
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// `
	jsonHandler, err := urlshort.JSONHandler([]byte(jsonFile), mapHandler)
	if err != nil {
		panic(err)
	}

	// yamlHandler, err := urlshort.YAMLHandler([]byte(yamlFile), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(yamlHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

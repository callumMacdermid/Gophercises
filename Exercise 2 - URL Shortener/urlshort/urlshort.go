package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path            //Read incoming request path.
		dest, ok := pathsToUrls[path] //Checking a matching key/value pair is present using the path
		if ok {                       //If a key path is found, redirect to the value URL.
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r) //Fallback if no matching key is found.
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYAML(yml)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	pathToUrls := buildMap(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}

	return pathToUrls
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:url`
}

func parseYAML(yml []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl //Create array of PathUrl structs to store paths in

	err := yaml.Unmarshal(yml, &pathUrls)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return pathUrls, nil
}

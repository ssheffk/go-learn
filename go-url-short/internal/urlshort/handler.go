package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathToURL struct {
	Path string
	URL  string
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		url := pathsToUrls[req.URL.Path]
		if url != "" {
			http.Redirect(res, req, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(res, req)
		}
	})
}

func buildMap(pathsToUrls []pathToURL) (buildMap map[string]string) {
	buildMap = make(map[string]string)
	for _, ptu := range pathsToUrls {
		buildMap[ptu.Path] = ptu.URL
	}
	return
}

func parseYAML(yamlData []byte) (pathToUrls []pathToURL, err error) {
	err = yaml.Unmarshal(yamlData, &pathToUrls)
	return
}

func YAMLHandler(yamlData []byte, fallback http.Handler) (yamlHandler http.HandlerFunc, err error) {
	parseYaml, err := parseYAML(yamlData)
	if err != nil {
		return
	}
	pathMap := buildMap(parseYaml)
	yamlHandler = MapHandler(pathMap, fallback)
	return
}

func parseJSON(jsonData []byte) (pathsToURLs []pathToURL, err error) {
	err = json.Unmarshal(jsonData, &pathsToURLs)
	return
}

func JSONHandler(jsonData []byte, fallback http.Handler) (jsonHandler http.HandlerFunc, err error) {
	parsedJSON, err := parseJSON(jsonData)
	if err != nil {
		return
	}
	pathMap := buildMap(parsedJSON)
	jsonHandler = MapHandler(pathMap, fallback)
	return
}

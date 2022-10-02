package gofilms

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// SearchMovie returns first result of movie search
func SearchMovie(title string) Movie {
	body := getSearchMoviePage(title)
	moviePage := getMoviePage(parseFirstURL(body))
	return parseMovie(moviePage)
}

// getSearchMoviePage returns html page of search result
func getSearchMoviePage(title string) string {
	url := "https://www.imdb.com/find?q=" + title + "&s=tt&ttype=ft&ref_=fn_ft"
	url = strings.ReplaceAll(url, " ", "%20")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var body string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		body = string(bodyBytes)
	} else {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		body = string(bodyBytes)
		panic(fmt.Sprintf("Status code: %d", resp.StatusCode))
	}

	return body
}

// getMoviePage returns html page of movie
func getMoviePage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var body string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		body = string(bodyBytes)
	} else {
		panic(fmt.Sprintf("Status code: %d", resp.StatusCode))
	}

	return body
}

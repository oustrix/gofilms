package gofilms

import (
	"encoding/json"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

var filmRegex = regexp.MustCompile(`/title/tt[0-9]+/`)

// parseName parses the all information about the film from the JSON
func parseMovie(body string) Movie {
	jsonMap := parseJSON(body)
	return Movie{
		name:          jsonMap["name"].(string),
		alternateName: jsonMap["alternateName"].(string),
		description:   jsonMap["description"].(string),
		kind:          jsonMap["@type"].(string),
		genres:        parseGenres(jsonMap),
		duration:      jsonMap["duration"].(string),
		contentRating: jsonMap["contentRating"].(string),
		ratingValue:   jsonMap["aggregateRating"].(map[string]interface{})["ratingValue"].(float64),
		ratingCount:   int(jsonMap["aggregateRating"].(map[string]interface{})["ratingCount"].(float64)),
		datePublished: jsonMap["datePublished"].(string),
		keywords:      parseKeywords(jsonMap),
		id:            parseID(jsonMap),
		trailer:       jsonMap["trailer"].(map[string]interface{})["url"].(string),
		image:         jsonMap["image"].(string),
		actors:        parseActors(jsonMap),
		directors:     parseDirectors(jsonMap),
		creators:      parseCreators(jsonMap),
	}
}

// parseJSON parses the JSON from the body of the HTML page
func parseJSON(body string) map[string]interface{} {
	r := strings.NewReader(body)
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	selector := cascadia.MustCompile("script[type='application/ld+json']").MatchFirst(doc)
	rawJSON := selector.FirstChild.Data

	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(rawJSON), &jsonMap)
	if err != nil {
		panic(err)
	}
	return jsonMap

}

// parseSearch parses the first url to movie page from the search results
func parseFirstURL(body string) string {
	return "https://www.imdb.com" + filmRegex.FindString(body)

}

// parseGenres parses the genres from the JSON
func parseGenres(json map[string]interface{}) []string {
	var genres []string
	for _, genre := range json["genre"].([]interface{}) {
		genres = append(genres, genre.(string))
	}
	return genres
}

// parseKeywords parses the keywords from the JSON
func parseKeywords(json map[string]interface{}) []string {
	keywordsString := json["keywords"].(string)
	keywordsSlice := strings.Split(keywordsString, ",")
	return keywordsSlice
}

// parseID parses the ID from the JSON, where it is stored as a string /title/tt[0-0]+/
func parseID(json map[string]interface{}) string {
	id := json["url"].(string)
	id = strings.ReplaceAll(id, "/title/tt", "")
	id = strings.ReplaceAll(id, "/", "")
	return id
}

// parseActors parses the actors from the JSON, where they are stored as a slice of maps
func parseActors(json map[string]interface{}) []Actor {
	var actors []Actor
	for _, actor := range json["actor"].([]interface{}) {
		newActor := Actor{}
		if actor.(map[string]interface{})["@type"] != nil {
			newActor.kind = actor.(map[string]interface{})["@type"].(string)
		}
		if actor.(map[string]interface{})["url"] != nil {
			newActor.url = "https://www.imdb.com" + actor.(map[string]interface{})["url"].(string)
		}
		if actor.(map[string]interface{})["name"] != nil {
			newActor.name = actor.(map[string]interface{})["name"].(string)
		}
		actors = append(actors, newActor)
	}
	return actors
}

// parseCreators parses the directors from the JSON, where they are stored as a slice of maps
func parseDirectors(json map[string]interface{}) []Director {
	var directors []Director
	for _, director := range json["director"].([]interface{}) {
		newDirector := Director{}
		if director.(map[string]interface{})["@type"] != nil {
			newDirector.kind = director.(map[string]interface{})["@type"].(string)
		}
		if director.(map[string]interface{})["url"] != nil {
			newDirector.url = "https://www.imdb.com" + director.(map[string]interface{})["url"].(string)
		}
		if director.(map[string]interface{})["name"] != nil {
			newDirector.name = director.(map[string]interface{})["name"].(string)
		}
		directors = append(directors, newDirector)
	}
	return directors
}

// parseCreators parses the creators from the JSON, where they are stored as a slice of maps
func parseCreators(json map[string]interface{}) []Creator {
	var creators []Creator
	for _, creator := range json["creator"].([]interface{}) {
		newCreator := Creator{}
		if creator.(map[string]interface{})["@type"] != nil {
			newCreator.kind = creator.(map[string]interface{})["@type"].(string)
		}
		if creator.(map[string]interface{})["url"] != nil {
			newCreator.url = "https://www.imdb.com" + creator.(map[string]interface{})["url"].(string)
		}
		if creator.(map[string]interface{})["name"] != nil {
			newCreator.name = creator.(map[string]interface{})["name"].(string)
		}
		creators = append(creators, newCreator)
	}
	return creators
}

package gofilms

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseJSON(t *testing.T) {
	testBody := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<script type="application/ld+json">
		{"@context": "testing",
		"@type": "testing",
		"name": "testing"}</script>`

	var expectedJSON = map[string]interface{}{
		"@context": "testing",
		"@type":    "testing",
		"name":     "testing",
	}
	actualJSON := parseJSON(testBody)
	require.Equal(t, expectedJSON, actualJSON)
}

func TestParseMovie(t *testing.T) {
	expectedMovie := Movie{
		name:          "testing",
		alternateName: "testing",
		description:   "testing",
		kind:          "testing",
		genres:        []string{"testing1", "testing2", "testing3"},
		duration:      "testing",
		contentRating: "testing",
		ratingValue:   10.0,
		ratingCount:   100,
		datePublished: "testing",
		keywords:      []string{"testing1", "testing2", "testing3"},
		id:            "0133093",
		trailer:       "testing",
		image:         "testing",
		actors: []Actor{{"testing11", "https://www.imdb.comtesting12", "testing13"},
			{"testing21", "https://www.imdb.comtesting22", "testing23"},
			{"testing31", "https://www.imdb.comtesting32", "testing33"}},
		directors: []Director{{"testing11", "https://www.imdb.comtesting12", "testing13"},
			{"testing21", "https://www.imdb.comtesting22", "testing23"},
			{"testing31", "https://www.imdb.comtesting32", "testing33"}},
		creators: []Creator{{"testing11", "https://www.imdb.comtesting12", "testing13"},
			{"testing21", "https://www.imdb.comtesting22", "testing23"},
			{"testing31", "https://www.imdb.comtesting32", "testing33"}},
	}
	testBody := `
		<html lang="en">
		<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<script type="application/ld+json">
		{"@context": "testing",
		"@type": "testing",
		"name": "testing",
		"alternateName": "testing",
		"image": "testing",
		"@type": "testing",
		"genre": ["testing1", "testing2", "testing3"],
		"duration": "testing",
		"contentRating": "testing",
		"aggregateRating": {
			"@type": "testing",
			"ratingValue": 10.0,
			"ratingCount": 100
		},
		"datePublished": "testing",
		"keywords": "testing1,testing2,testing3",
		"url": "/title/tt0133093/",

		"actor": [
			{
				"@type": "testing11",
				"url": "testing12",
				"name": "testing13"
			},
			{
				"@type": "testing21",
				"url": "testing22",
				"name": "testing23"
			},
			{
				"@type": "testing31",
				"url": "testing32",
				"name": "testing33"
			}
		],	
		"director": [
		    {
				"@type": "testing11",
				"name": "testing13",
				"url": "testing12"
			},
			{
				"@type": "testing21",
				"name": "testing23",
				"url": "testing22"
			},
			{
				"@type": "testing31",
				"name": "testing33",
				"url": "testing32"
			}
		],	
		"creator": [
		    {
				"@type": "testing11",
				"name": "testing13",
				"url": "testing12"
			},
			{
				"@type": "testing21",
				"name": "testing23",
				"url": "testing22"
			},
			{
				"@type": "testing31",
				"name": "testing33",
				"url": "testing32"
			}
		],
		"trailer": {
			"url": "testing"
		},
        "description": "testing"
		}</script>
    	`

	actualMovie := parseMovie(testBody)
	require.Equal(t, expectedMovie.name, actualMovie.name)
	require.Equal(t, expectedMovie.alternateName, actualMovie.alternateName)
	require.Equal(t, expectedMovie.description, actualMovie.description)
	require.Equal(t, expectedMovie.kind, actualMovie.kind)
	require.Equal(t, expectedMovie.genres, actualMovie.genres)
	require.Equal(t, expectedMovie.duration, actualMovie.duration)
	require.Equal(t, expectedMovie.contentRating, actualMovie.contentRating)
	require.Equal(t, expectedMovie.ratingValue, actualMovie.ratingValue)
	require.Equal(t, expectedMovie.ratingCount, actualMovie.ratingCount)
	require.Equal(t, expectedMovie.datePublished, actualMovie.datePublished)
	require.Equal(t, expectedMovie.keywords, actualMovie.keywords)
	require.Equal(t, expectedMovie.id, actualMovie.id)
	require.Equal(t, expectedMovie.trailer, actualMovie.trailer)
	require.Equal(t, expectedMovie.image, actualMovie.image)
	require.Equal(t, expectedMovie.actors, actualMovie.actors)
	require.Equal(t, expectedMovie.directors, actualMovie.directors)
	require.Equal(t, expectedMovie.creators, actualMovie.creators)

}

func TestParseFirstURL(t *testing.T) {
	testString := "hdblafdhfbald/title/tt09987654322/fdhabflfdlafahbf"
	expectedURL := "https://www.imdb.com/title/tt09987654322/"
	actualURL := parseFirstURL(testString)
	require.Equal(t, expectedURL, actualURL)
}

func TestParseGenres(t *testing.T) {
	testJSON := map[string]interface{}{
		"genre": []interface{}{"genre1", "genre2", "genre3"},
	}
	expectedGenres := []string{"genre1", "genre2", "genre3"}
	actualGenres := parseGenres(testJSON)
	require.Equal(t, expectedGenres, actualGenres)
}

func TestParseKeywords(t *testing.T) {
	testJSON := map[string]interface{}{
		"keywords": "keyword1,keyword2,keyword3",
	}
	expectedKeywords := []string{"keyword1", "keyword2", "keyword3"}
	actualKeywords := parseKeywords(testJSON)
	require.Equal(t, expectedKeywords, actualKeywords)
}

func TestParseID(t *testing.T) {
	testJSON := map[string]interface{}{
		"url": "/title/tt0133093/",
	}
	expectedID := "0133093"
	actualID := parseID(testJSON)
	require.Equal(t, expectedID, actualID)
}

func TestParseActors(t *testing.T) {
	testJSON := map[string]interface{}{
		"actor": []interface{}{
			map[string]interface{}{
				"@type": "actor",
				"url":   "actor1",
				"name":  "actor2",
			},
			map[string]interface{}{
				"@type": "actor",
				"url":   "actor3",
				"name":  "actor4",
			},
			map[string]interface{}{
				"@type": "actor",
				"url":   "actor5",
				"name":  "actor6",
			},
		},
	}
	expectedActors := []Actor{
		{"actor", "https://www.imdb.comactor1", "actor2"},
		{"actor", "https://www.imdb.comactor3", "actor4"},
		{"actor", "https://www.imdb.comactor5", "actor6"},
	}
	actualActors := parseActors(testJSON)
	require.Equal(t, expectedActors, actualActors)
}

func TestParseDirectors(t *testing.T) {
	testJSON := map[string]interface{}{
		"director": []interface{}{
			map[string]interface{}{
				"@type": "director",
				"url":   "director1",
				"name":  "director2",
			},
			map[string]interface{}{
				"@type": "director",
				"url":   "director3",
				"name":  "director4",
			},
			map[string]interface{}{
				"@type": "director",
				"url":   "director5",
				"name":  "director6",
			},
		},
	}
	expectedDirectors := []Director{
		{"director", "https://www.imdb.comdirector1", "director2"},
		{"director", "https://www.imdb.comdirector3", "director4"},
		{"director", "https://www.imdb.comdirector5", "director6"},
	}
	actualDirectors := parseDirectors(testJSON)
	require.Equal(t, expectedDirectors, actualDirectors)
}

func TestParseCreators(t *testing.T) {
	testJSON := map[string]interface{}{
		"creator": []interface{}{
			map[string]interface{}{
				"@type": "creator",
				"url":   "creator1",
				"name":  "creator2",
			},
			map[string]interface{}{
				"@type": "creator",
				"url":   "creator3",
				"name":  "creator4",
			},
			map[string]interface{}{
				"@type": "creator",
				"url":   "creator5",
				"name":  "creator6",
			},
		},
	}
	expectedCreators := []Creator{
		{"creator", "https://www.imdb.comcreator1", "creator2"},
		{"creator", "https://www.imdb.comcreator3", "creator4"},
		{"creator", "https://www.imdb.comcreator5", "creator6"},
	}
	actualCreators := parseCreators(testJSON)
	require.Equal(t, expectedCreators, actualCreators)
}

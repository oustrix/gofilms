package gofilms

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var movie = Movie{
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
	actors: []Actor{{"testing11", "testing12", "testing13"},
		{"testing21", "testing22", "testing23"},
		{"testing31", "testing32", "testing33"}},
	directors: []Director{{"testing11", "testing12", "testing13"},
		{"testing21", "testing22", "testing23"},
		{"testing31", "testing32", "testing33"}},
	creators: []Creator{{"testing11", "testing12", "testing13"},
		{"testing21", "testing22", "testing23"},
		{"testing31", "testing32", "testing33"}},
}

func TestMovie_GetName(t *testing.T) {
	require.Equal(t, movie.name, movie.GetName())
}

func TestMovie_GetAlternateName(t *testing.T) {
	require.Equal(t, movie.alternateName, movie.GetAlternateName())
}

func TestMovie_GetContentRating(t *testing.T) {
	require.Equal(t, movie.contentRating, movie.GetContentRating())
}

func TestMovie_GetDatePublished(t *testing.T) {
	require.Equal(t, movie.datePublished, movie.GetDatePublished())
}

func TestMovie_GetDescription(t *testing.T) {
	require.Equal(t, movie.description, movie.GetDescription())
}

func TestMovie_GetDuration(t *testing.T) {
	require.Equal(t, movie.duration, movie.GetDuration())
}

func TestMovie_GetKind(t *testing.T) {
	require.Equal(t, movie.kind, movie.GetKind())
}

func TestMovie_GetRatingCount(t *testing.T) {
	require.Equal(t, movie.ratingCount, movie.GetRatingCount())
}

func TestMovie_GetRatingValue(t *testing.T) {
	require.Equal(t, movie.ratingValue, movie.GetRatingValue())
}

func TestMovie_GetGenres(t *testing.T) {
	require.Equal(t, movie.genres, movie.GetGenres())
}

func TestMovie_GetKeywords(t *testing.T) {
	require.Equal(t, movie.keywords, movie.GetKeywords())
}

func TestMovie_GetID(t *testing.T) {
	require.Equal(t, movie.id, movie.GetID())
}

func TestMovie_GetTrailer(t *testing.T) {
	require.Equal(t, movie.trailer, movie.GetTrailer())
}

func TestMovie_GetImage(t *testing.T) {
	require.Equal(t, movie.image, movie.GetImage())
}

func TestMovie_GetActors(t *testing.T) {
	require.Equal(t, movie.actors, movie.GetActors())
}

func TestMovie_GetDirectors(t *testing.T) {
	require.Equal(t, movie.directors, movie.GetDirectors())
}

func TestMovie_GetCreators(t *testing.T) {
	require.Equal(t, movie.creators, movie.GetCreators())
}

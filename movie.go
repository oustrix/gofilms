package gofilms

type Movie struct {
	name          string
	alternateName string // depends on browser language
	description   string
	kind          string
	genres        []string
	duration      string
	contentRating string
	ratingValue   float64
	ratingCount   int
	datePublished string // year-month-day
	keywords      []string
	id            string
	trailer       string
	image         string
	actors        []Actor
	directors     []Director
	creators      []Creator
}

func (f Movie) GetName() string {
	return f.name
}

func (f Movie) GetAlternateName() string {
	return f.alternateName
}

func (f Movie) GetDescription() string {
	return f.description
}

func (f Movie) GetKind() string {
	return f.kind
}

func (f Movie) GetGenres() []string {
	return f.genres
}

func (f Movie) GetDuration() string {
	return f.duration

}

func (f Movie) GetContentRating() string {
	return f.contentRating
}

func (f Movie) GetRatingValue() float64 {
	return f.ratingValue
}

func (f Movie) GetRatingCount() int {
	return f.ratingCount
}

func (f Movie) GetDatePublished() string {
	return f.datePublished
}

func (f Movie) GetKeywords() []string {
	return f.keywords
}

func (f Movie) GetID() string {
	return f.id
}

func (f Movie) GetTrailer() string {
	return f.trailer
}

func (f Movie) GetImage() string {
	return f.image
}

func (f Movie) GetActors() []Actor {
	return f.actors
}

func (f Movie) GetDirectors() []Director {
	return f.directors
}

func (f Movie) GetCreators() []Creator {
	return f.creators
}

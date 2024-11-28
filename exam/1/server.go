package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Album represents a music album.
// swagger:model
type Album struct {
	// ID is the unique identifier for the album.
	//
	// required: true
	ID string `json:"id"`

	// Title is the title of the album.
	//
	// required: true
	Title string `json:"title"`

	// Artist is the name of the artist who performed the album.
	//
	// required: true
	Artist string `json:"artist"`

	// Price is the price of the album in USD.
	//
	// required: true
	Price float64 `json:"price"`
}

// swagger:parameters searchAlbums
type searchRequest struct {
	// in:query
	Title string `json:"title"`
}

// swagger:response
type searchResponse struct {
	// Album array
	// in:body
	Albums []Album `json:"albums"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Kind of Blue", Artist: "Miles Davis", Price: 24.99},
	{ID: "5", Title: "A Love Supreme", Artist: "John Coltrane", Price: 34.99},
	{ID: "6", Title: "Bitches Brew", Artist: "Miles Davis", Price: 44.99},
	{ID: "7", Title: "In a Silent Way", Artist: "Miles Davis", Price: 29.99},
	{ID: "8", Title: "My Favorite Things", Artist: "John Coltrane", Price: 39.99},
	{ID: "9", Title: "Time Out", Artist: "Dave Brubeck", Price: 19.99},
	{ID: "10", Title: "Birdland", Artist: "Weather Report", Price: 29.99},
	{ID: "11", Title: "The Impressions", Artist: "Nina Simone", Price: 19.99},
	{ID: "12", Title: "The Miseducation of Lauryn Hill", Artist: "Lauryn Hill", Price: 19.99},
	{ID: "13", Title: "The Chronic", Artist: "Dr. Dre", Price: 24.99},
	{ID: "14", Title: "Illmatic", Artist: "Nas", Price: 29.99},
	{ID: "15", Title: "The Score", Artist: "Fiona Apple", Price: 24.99},
	{ID: "16", Title: "To Pimp a Butterfly", Artist: "Kendrick Lamar", Price: 34.99},
	{ID: "17", Title: "Lemonade", Artist: "Beyoncé", Price: 44.99},
	{ID: "18", Title: "21", Artist: "Adele", Price: 39.99},
	{ID: "19", Title: "Views", Artist: "Drake", Price: 49.99},
	{ID: "20", Title: "Blonde", Artist: "Frank Ocean", Price: 39.99},
	{ID: "21", Title: "The Black Album", Artist: "Metallica", Price: 29.99},
	{ID: "22", Title: "Nevermind", Artist: "Nirvana", Price: 24.99},
	{ID: "23", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 29.99},
	{ID: "24", Title: "Thriller", Artist: "Michael Jackson", Price: 39.99},
	{ID: "25", Title: "Back in Black", Artist: "AC/DC", Price: 29.99},
	{ID: "26", Title: "The Wall", Artist: "Pink Floyd", Price: 34.99},
	{ID: "27", Title: "Rumours", Artist: "Fleetwood Mac", Price: 29.99},
	{ID: "28", Title: "Hotel California", Artist: "Eagles", Price: 29.99},
	{ID: "29", Title: "The Bodyguard Soundtrack", Artist: "Whitney Houston", Price: 19.99},
	{ID: "30", Title: "The Eminem Show", Artist: "Eminem", Price: 29.99},
	{ID: "31", Title: "The White Album", Artist: "The Beatles", Price: 39.99},
	{ID: "32", Title: "Abbey Road", Artist: "The Beatles", Price: 29.99},
	{ID: "33", Title: "The Impossible Dream", Artist: "Elvis Presley", Price: 19.99},
	{ID: "34", Title: "The Beatles (White Album)", Artist: "The Beatles", Price: 39.99},
	{ID: "35", Title: "The Best of Bob Dylan", Artist: "Bob Dylan", Price: 29.99},
	{ID: "36", Title: "The Doors", Artist: "The Doors", Price: 24.99},
	{ID: "37", Title: "The Joshua Tree", Artist: "U2", Price: 29.99},
	{ID: "38", Title: "The Bodyguard Soundtrack", Artist: "Whitney Houston", Price: 19.99},
	{ID: "39", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 29.99},
	{ID: "40", Title: "Thriller", Artist: "Michael Jackson", Price: 39.99},
	{ID: "41", Title: "Back in Black", Artist: "AC/DC", Price: 29.99},
	{ID: "42", Title: "The Wall", Artist: "Pink Floyd", Price: 34.99},
	{ID: "43", Title: "Rumours", Artist: "Fleetwood Mac", Price: 29.99},
	{ID: "44", Title: "Hotel California", Artist: "Eagles", Price: 29.99},
	{ID: "45", Title: "The Bodyguard Soundtrack", Artist: "Whitney Houston", Price: 19.99},
	{ID: "46", Title: "The Eminem Show", Artist: "Eminem", Price: 29.99},
	{ID: "47", Title: "The White Album", Artist: "The Beatles", Price: 39.99},
	{ID: "48", Title: "Abbey Road", Artist: "The Beatles", Price: 29.99},
	{ID: "49", Title: "The Impossible Dream", Artist: "Elvis Presley", Price: 19.99},
	{ID: "50", Title: "The Beatles (White Album)", Artist: "The Beatles", Price: 39.99},
	{ID: "51", Title: "The Best of Bob Dylan", Artist: "Bob Dylan", Price: 29.99},
	{ID: "52", Title: "The Doors", Artist: "The Doors", Price: 24.99},
	{ID: "53", Title: "The Joshua Tree", Artist: "U2", Price: 29.99},
	{ID: "54", Title: "The Bodyguard Soundtrack", Artist: "Whitney Houston", Price: 19.99},
	{ID: "55", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 29.99},
	{ID: "56", Title: "Thriller", Artist: "Michael Jackson", Price: 39.99},
	{ID: "57", Title: "Back in Black", Artist: "AC/DC", Price: 29.99},
	{ID: "58", Title: "The Wall", Artist: "Pink Floyd", Price: 34.99},
	{ID: "59", Title: "Rumours", Artist: "Fleetwood Mac", Price: 29.99},
	{ID: "60", Title: "Hotel California", Artist: "Eagles", Price: 29.99},
	{ID: "61", Title: "The Bodyguard Soundtrack", Artist: "Whitney Houston", Price: 19.99},
	{ID: "62", Title: "The Eminem Show", Artist: "Eminem", Price: 29.99},
	{ID: "63", Title: "The White Album", Artist: "The Beatles", Price: 39.99},
	{ID: "64", Title: "Abbey Road", Artist: "The Beatles", Price: 29.99},
	{ID: "65", Title: "The Impossible Dream", Artist: "Elvis Presley", Price: 19.99},
	{ID: "66", Title: "The Beatles (White Album)", Artist: "The Beatles", Price: 39.99},
	{ID: "67", Title: "The Best of Bob Dylan", Artist: "Bob Dylan", Price: 29.99},
	{ID: "68", Title: "The Doors", Artist: "The Doors", Price: 24.99},
	{ID: "69", Title: "The Joshua Tree", Artist: "U2", Price: 29.99},
	{ID: "70", Title: "The Bodyguard Soundtrack", Artist: "Whitney Houston", Price: 19.99},
	{ID: "71", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 29.99},
	{ID: "72", Title: "Thriller", Artist: "Michael Jackson", Price: 39.99},
	{ID: "73", Title: "Back in Black", Artist: "AC/DC", Price: 29.99},
	{ID: "74", Title: "The Wall", Artist: "Pink Floyd", Price: 34.99},
	{ID: "75", Title: "Rumours", Artist: "Fleetwood Mac", Price: 29.99},
	{ID: "76", Title: "Hotel California", Artist: "Eagles", Price: 29.99},
	{ID: "77", Title: "The Bodyguard Soundtrack", Artist: "Whitney Houston", Price: 19.99},
	{ID: "78", Title: "The Eminem Show", Artist: "Eminem", Price: 29.99},
	{ID: "79", Title: "The White Album", Artist: "The Beatles", Price: 39.99},
	{ID: "80", Title: "Abbey Road", Artist: "The Beatles", Price: 29.99},
	{ID: "81", Title: "The Impossible Dream", Artist: "Elvis Presley", Price: 19.99},
	{ID: "82", Title: "The Beatles (White Album)", Artist: "The Beatles", Price: 39.99},
	{ID: "83", Title: "The Best of Bob Dylan", Artist: "Bob Dylan", Price: 29.99},
	{ID: "84", Title: "The Doors", Artist: "The Doors", Price: 24.99},
	{ID: "85", Title: "The Joshua Tree", Artist: "U2", Price: 29.99},
	{ID: "86", Title: "The Bodyguard Soundtrack", Artist: "Whitney Houston", Price: 19.99},
	{ID: "87", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 29.99},
	{ID: "88", Title: "Thriller", Artist: "Michael Jackson", Price: 39.99},
	{ID: "89", Title: "Back in Black", Artist: "AC/DC", Price: 29.99},
	{ID: "90", Title: "The Wall", Artist: "Pink Floyd", Price: 34.99},
	{ID: "91", Title: "Rumours", Artist: "Fleetwood Mac", Price: 29.99},
	{ID: "92", Title: "Hotel California", Artist: "Eagles", Price: 29.99},
	{ID: "93", Title: "The Bodyguard Soundtrack", Artist: "Whitney Houston", Price: 19.99},
	{ID: "94", Title: "The Eminem Show", Artist: "Eminem", Price: 29.99},
	{ID: "95", Title: "The White Album", Artist: "The Beatles", Price: 39.99},
	{ID: "96", Title: "Abbey Road", Artist: "The Beatles", Price: 29.99},
	{ID: "97", Title: "The Impossible Dream", Artist: "Elvis Presley", Price: 19.99},
	{ID: "98", Title: "The Beatles (White Album)", Artist: "The Beatles", Price: 39.99},
	{ID: "99", Title: "The Best of Bob Dylan", Artist: "Bob Dylan", Price: 29.99},
	{ID: "100", Title: "The Doors", Artist: "The Doors", Price: 24.99},
}

// searchAlbums searches for albums whose titles contain a given string.
// The function expects a GET request to the /search endpoint with a title query parameter.
// The function returns a JSON array of matching albums.
// swagger:route GET /search albums searchAlbums
//
// Search for albums by part of title.
//
//	Produces:
//	- application/json
//
//	Schemes: http
//
//	Responses:
//	  200: searchResponse
func searchAlbums(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("title")

	var results []Album
	for _, album := range albums {
		if strings.Contains(strings.ToLower(album.Title), strings.ToLower(query)) {
			results = append(results, album)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/search", searchAlbums)

	http.ListenAndServe(":9080", nil)
}

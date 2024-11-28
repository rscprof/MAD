package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Anime represents an anime series.
type Anime struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Episodes int    `json:"episodes"`
}

// Sample list of 100 anime
var animeList = []Anime{
	{ID: "1", Title: "Naruto", Genre: "Action, Adventure", Episodes: 220},
	{ID: "2", Title: "Death Note", Genre: "Thriller, Supernatural", Episodes: 37},
	{ID: "3", Title: "Attack on Titan", Genre: "Action, Drama", Episodes: 75},
	{ID: "4", Title: "My Hero Academia", Genre: "Action, Comedy", Episodes: 88},
	{ID: "5", Title: "One Piece", Genre: "Adventure, Comedy", Episodes: 1000},
	{ID: "6", Title: "Demon Slayer", Genre: "Action, Fantasy", Episodes: 26},
	{ID: "7", Title: "Steins;Gate", Genre: "Sci-Fi, Thriller", Episodes: 24},
	{ID: "8", Title: "Sword Art Online", Genre: "Adventure, Fantasy", Episodes: 96},
	{ID: "9", Title: "Your Lie in April", Genre: "Drama, Romance", Episodes: 22},
	{ID: "10", Title: "Hunter x Hunter", Genre: "Adventure, Action", Episodes: 148},
	{ID: "11", Title: "Tokyo Ghoul", Genre: "Horror, Supernatural", Episodes: 24},
	{ID: "12", Title: "Bleach", Genre: "Action, Supernatural", Episodes: 366},
	{ID: "13", Title: "Black Clover", Genre: "Action, Fantasy", Episodes: 170},
	{ID: "14", Title: "Fairy Tail", Genre: "Adventure, Fantasy", Episodes: 328},
	{ID: "15", Title: "Re:Zero", Genre: "Fantasy, Thriller", Episodes: 50},
	{ID: "16", Title: "One Punch Man", Genre: "Action, Comedy", Episodes: 24},
	{ID: "17", Title: "The Seven Deadly Sins", Genre: "Action, Fantasy", Episodes: 100},
	{ID: "18", Title: "Dr. Stone", Genre: "Sci-Fi, Adventure", Episodes: 35},
	{ID: "19", Title: "Fullmetal Alchemist: Brotherhood", Genre: "Action, Adventure", Episodes: 64},
	{ID: "20", Title: "Jujutsu Kaisen", Genre: "Action, Supernatural", Episodes: 24},
	{ID: "21", Title: "Haikyuu!!", Genre: "Sports, Drama", Episodes: 85},
	{ID: "22", Title: "Kuroko no Basket", Genre: "Sports, Comedy", Episodes: 75},
	{ID: "23", Title: "Code Geass", Genre: "Mecha, Drama", Episodes: 50},
	{ID: "24", Title: "No Game No Life", Genre: "Fantasy, Comedy", Episodes: 12},
	{ID: "25", Title: "Mob Psycho 100", Genre: "Action, Comedy", Episodes: 25},
	{ID: "26", Title: "Erased", Genre: "Thriller, Mystery", Episodes: 12},
	{ID: "27", Title: "Parasyte", Genre: "Horror, Sci-Fi", Episodes: 24},
	{ID: "28", Title: "Violet Evergarden", Genre: "Drama, Fantasy", Episodes: 13},
	{ID: "29", Title: "Made in Abyss", Genre: "Adventure, Fantasy", Episodes: 13},
	{ID: "30", Title: "Clannad", Genre: "Drama, Romance", Episodes: 48},
	{ID: "31", Title: "Angel Beats!", Genre: "Drama, Supernatural", Episodes: 13},
	{ID: "32", Title: "Toradora!", Genre: "Romance, Comedy", Episodes: 25},
	{ID: "33", Title: "A Silent Voice", Genre: "Drama, Romance", Episodes: 1},
	{ID: "34", Title: "Cowboy Bebop", Genre: "Sci-Fi, Action", Episodes: 26},
	{ID: "35", Title: "Neon Genesis Evangelion", Genre: "Mecha, Drama", Episodes: 26},
	{ID: "36", Title: "Monster", Genre: "Thriller, Drama", Episodes: 74},
	{ID: "37", Title: "Spirited Away", Genre: "Fantasy, Adventure", Episodes: 1},
	{ID: "38", Title: "Princess Mononoke", Genre: "Fantasy, Action", Episodes: 1},
	{ID: "39", Title: "Howl's Moving Castle", Genre: "Fantasy, Romance", Episodes: 1},
	{ID: "40", Title: "Akira", Genre: "Sci-Fi, Action", Episodes: 1},
	{ID: "41", Title: "Fate/Zero", Genre: "Action, Fantasy", Episodes: 25},
	{ID: "42", Title: "Fate/Stay Night", Genre: "Action, Romance", Episodes: 24},
	{ID: "43", Title: "Psycho-Pass", Genre: "Thriller, Sci-Fi", Episodes: 41},
	{ID: "44", Title: "The Rising of the Shield Hero", Genre: "Action, Fantasy", Episodes: 25},
	{ID: "45", Title: "Overlord", Genre: "Action, Fantasy", Episodes: 39},
	{ID: "46", Title: "KonoSuba", Genre: "Comedy, Fantasy", Episodes: 20},
	{ID: "47", Title: "The Promised Neverland", Genre: "Thriller, Drama", Episodes: 23},
	{ID: "48", Title: "Berserk", Genre: "Action, Drama", Episodes: 25},
	{ID: "49", Title: "Samurai Champloo", Genre: "Action, Adventure", Episodes: 26},
	{ID: "50", Title: "Gurren Lagann", Genre: "Mecha, Adventure", Episodes: 27},
	{ID: "51", Title: "Hellsing Ultimate", Genre: "Action, Supernatural", Episodes: 10},
	{ID: "52", Title: "Bakemonogatari", Genre: "Mystery, Romance", Episodes: 15},
	{ID: "53", Title: "Oregairu", Genre: "Romance, Comedy", Episodes: 36},
	{ID: "54", Title: "Zombieland Saga", Genre: "Comedy, Supernatural", Episodes: 24},
	{ID: "55", Title: "Kill la Kill", Genre: "Action, Comedy", Episodes: 24},
	{ID: "56", Title: "Charlotte", Genre: "Drama, Supernatural", Episodes: 13},
	{ID: "57", Title: "The Ancient Magus' Bride", Genre: "Fantasy, Drama", Episodes: 24},
	{ID: "58", Title: "Horimiya", Genre: "Romance, Comedy", Episodes: 13},
	{ID: "59", Title: "Gintama", Genre: "Action, Comedy", Episodes: 367},
	{ID: "60", Title: "Assassination Classroom", Genre: "Action, Comedy", Episodes: 47},
	{ID: "61", Title: "Great Teacher Onizuka", Genre: "Comedy, Drama", Episodes: 43},
	{ID: "62", Title: "The Quintessential Quintuplets", Genre: "Romance, Comedy", Episodes: 24},
	{ID: "63", Title: "Yuri on Ice", Genre: "Sports, Drama", Episodes: 12},
	{ID: "64", Title: "Shokugeki no Soma", Genre: "Comedy, Drama", Episodes: 86},
	{ID: "65", Title: "Kaguya-sama: Love is War", Genre: "Comedy, Romance", Episodes: 37},
	{ID: "66", Title: "Vinland Saga", Genre: "Action, Drama", Episodes: 24},
	{ID: "67", Title: "Dorohedoro", Genre: "Action, Fantasy", Episodes: 12},
	{ID: "68", Title: "Ergo Proxy", Genre: "Sci-Fi, Mystery", Episodes: 23},
	{ID: "69", Title: "The Devil is a Part-Timer!", Genre: "Comedy, Fantasy", Episodes: 13},
	{ID: "70", Title: "Black Lagoon", Genre: "Action, Thriller", Episodes: 29},
	{ID: "71", Title: "Akame ga Kill!", Genre: "Action, Drama", Episodes: 24},
	{ID: "72", Title: "Baccano!", Genre: "Action, Mystery", Episodes: 16},
	{ID: "73", Title: "Durarara!!", Genre: "Action, Mystery", Episodes: 60},
	{ID: "74", Title: "Trigun", Genre: "Action, Sci-Fi", Episodes: 26},
	{ID: "75", Title: "Elfen Lied", Genre: "Horror, Drama", Episodes: 13},
	{ID: "76", Title: "FLCL", Genre: "Comedy, Sci-Fi", Episodes: 6},
	{ID: "77", Title: "Hyouka", Genre: "Mystery, Drama", Episodes: 22},
	{ID: "78", Title: "Anohana", Genre: "Drama, Supernatural", Episodes: 11},
	{ID: "79", Title: "Barakamon", Genre: "Comedy, Slice of Life", Episodes: 12},
	{ID: "80", Title: "March Comes in Like a Lion", Genre: "Drama, Slice of Life", Episodes: 44},
	{ID: "81", Title: "Saiki Kusuo no Psi-nan", Genre: "Comedy, Supernatural", Episodes: 120},
	{ID: "82", Title: "Nichijou", Genre: "Comedy, Slice of Life", Episodes: 26},
	{ID: "83", Title: "Shinsekai Yori", Genre: "Sci-Fi, Mystery", Episodes: 25},
	{ID: "84", Title: "Noragami", Genre: "Action, Supernatural", Episodes: 25},
	{ID: "85", Title: "Magi", Genre: "Adventure, Fantasy", Episodes: 50},
	{ID: "86", Title: "Beastars", Genre: "Drama, Romance", Episodes: 24},
	{ID: "87", Title: "Welcome to the NHK", Genre: "Drama, Comedy", Episodes: 24},
	{ID: "88", Title: "Devilman Crybaby", Genre: "Action, Horror", Episodes: 10},
	{ID: "89", Title: "Rurouni Kenshin", Genre: "Action, Adventure", Episodes: 95},
	{ID: "90", Title: "Kimi no Na Wa", Genre: "Romance, Drama", Episodes: 1},
	{ID: "91", Title: "Weathering With You", Genre: "Romance, Fantasy", Episodes: 1},
	{ID: "92", Title: "A Place Further Than the Universe", Genre: "Adventure, Slice of Life", Episodes: 13},
	{ID: "93", Title: "Banana Fish", Genre: "Action, Drama", Episodes: 24},
	{ID: "94", Title: "Kaiji", Genre: "Thriller, Drama", Episodes: 52},
	{ID: "95", Title: "Hajime no Ippo", Genre: "Sports, Drama", Episodes: 126},
	{ID: "96", Title: "Nana", Genre: "Drama, Romance", Episodes: 47},
	{ID: "97", Title: "Ping Pong the Animation", Genre: "Sports, Drama", Episodes: 11},
	{ID: "98", Title: "Golden Kamuy", Genre: "Action, Adventure", Episodes: 48},
	{ID: "99", Title: "Karakuri Circus", Genre: "Action, Drama", Episodes: 36},
	{ID: "100", Title: "Fruits Basket", Genre: "Romance, Drama", Episodes: 63},
}

// searchAnime handles search requests based on episode count
func searchAnime(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	minEpisodesStr := r.URL.Query().Get("minEpisodes")
	maxEpisodesStr := r.URL.Query().Get("maxEpisodes")

	// Default to 0 and a large value if parameters are missing
	minEpisodes, _ := strconv.Atoi(minEpisodesStr)
	maxEpisodes, _ := strconv.Atoi(maxEpisodesStr)

	if maxEpisodes == 0 {
		maxEpisodes = 1 << 30 // Set to a very high number if not provided
	}

	var results []Anime
	for _, anime := range animeList {
		if anime.Episodes >= minEpisodes && anime.Episodes <= maxEpisodes {
			results = append(results, anime)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/search", searchAnime)
	http.ListenAndServe(":9080", nil)
}

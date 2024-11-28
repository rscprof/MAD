package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Place represents a tourist place.
// swagger:model
type Place struct {
	// ID is the unique identifier for the place.
	//
	// required: true
	ID string `json:"id"`

	// Name is the name of the place.
	//
	// required: true
	Name string `json:"name"`

	// Country is the country where the place is located.
	//
	// required: true
	Country string `json:"country"`

	// Description is a brief description of the place.
	//
	// required: true
	Description string `json:"description"`

	// Popularity is the popularity rating of the place (0-100).
	//
	// required: true
	Popularity int `json:"popularity"`
}

// swagger:parameters searchPlacesByCountry
type searchRequest struct {
	// in:query
	// The country name to search places by
	Country string `json:"country"`
}

// swagger:response searchResponse
type searchResponse struct {
	// in:body
	// List of places
	Places []Place `json:"places"`
}

// Пример списка из 100 туристических мест.
var placesList = []Place{
	{ID: "1", Name: "Eiffel Tower", Country: "France", Description: "Iconic iron lattice tower in Paris.", Popularity: 98},
	{ID: "2", Name: "Great Wall of China", Country: "China", Description: "Ancient series of walls and fortifications.", Popularity: 95},
	{ID: "3", Name: "Machu Picchu", Country: "Peru", Description: "15th-century Inca citadel in the Andes Mountains.", Popularity: 93},
	{ID: "4", Name: "Grand Canyon", Country: "USA", Description: "Famous steep-sided canyon carved by the Colorado River.", Popularity: 90},
	{ID: "5", Name: "Colosseum", Country: "Italy", Description: "Ancient Roman amphitheater in Rome.", Popularity: 89},
	{ID: "6", Name: "Taj Mahal", Country: "India", Description: "Iconic mausoleum and symbol of love.", Popularity: 96},
	{ID: "7", Name: "Statue of Liberty", Country: "USA", Description: "Iconic statue representing freedom in New York City.", Popularity: 85},
	{ID: "8", Name: "Sydney Opera House", Country: "Australia", Description: "Famous performing arts center in Sydney.", Popularity: 88},
	{ID: "9", Name: "Stonehenge", Country: "UK", Description: "Prehistoric monument of standing stones in England.", Popularity: 78},
	{ID: "10", Name: "Christ the Redeemer", Country: "Brazil", Description: "Iconic statue overlooking Rio de Janeiro.", Popularity: 91},
	{ID: "11", Name: "Pyramids of Giza", Country: "Egypt", Description: "Ancient pyramid structures near Cairo.", Popularity: 95},
	{ID: "12", Name: "Tokyo Tower", Country: "Japan", Description: "Iconic communications tower in Tokyo.", Popularity: 82},
	{ID: "13", Name: "Acropolis of Athens", Country: "Greece", Description: "Ancient citadel with Parthenon temple.", Popularity: 91},
	{ID: "14", Name: "Sagrada Familia", Country: "Spain", Description: "Gaudí's unfinished basilica in Barcelona.", Popularity: 93},
	{ID: "15", Name: "Mount Fuji", Country: "Japan", Description: "Active stratovolcano and Japan's tallest peak.", Popularity: 87},
	{ID: "16", Name: "Big Ben", Country: "UK", Description: "Famous clock tower in London.", Popularity: 92},
	{ID: "17", Name: "Niagara Falls", Country: "USA/Canada", Description: "Large waterfall on the border between USA and Canada.", Popularity: 94},
	{ID: "18", Name: "The Louvre", Country: "France", Description: "World-famous museum in Paris.", Popularity: 99},
	{ID: "19", Name: "Machu Picchu", Country: "Peru", Description: "Ancient Inca city high in the Andes.", Popularity: 96},
	{ID: "20", Name: "Burj Khalifa", Country: "UAE", Description: "World's tallest building in Dubai.", Popularity: 98},
	{ID: "21", Name: "Machu Picchu", Country: "Peru", Description: "Ancient Inca citadel in the Andes.", Popularity: 95},
	{ID: "22", Name: "Stonehenge", Country: "UK", Description: "Mysterious circle of stone pillars.", Popularity: 87},
	{ID: "23", Name: "Golden Gate Bridge", Country: "USA", Description: "Famous suspension bridge in San Francisco.", Popularity: 90},
	{ID: "24", Name: "Great Barrier Reef", Country: "Australia", Description: "The world's largest coral reef system.", Popularity: 91},
	{ID: "25", Name: "Mount Everest", Country: "Nepal/Tibet", Description: "Highest mountain on Earth.", Popularity: 96},
	{ID: "26", Name: "Galápagos Islands", Country: "Ecuador", Description: "Unique archipelago known for its wildlife.", Popularity: 84},
	{ID: "27", Name: "Chichen Itza", Country: "Mexico", Description: "Ancient Mayan city known for its pyramid.", Popularity: 92},
	{ID: "28", Name: "Red Square", Country: "Russia", Description: "Historic square in Moscow, home to landmarks like the Kremlin.", Popularity: 89},
	{ID: "29", Name: "The Alhambra", Country: "Spain", Description: "Islamic palace and fortress in Granada.", Popularity: 94},
	{ID: "30", Name: "Mount Kilimanjaro", Country: "Tanzania", Description: "Free-standing volcanic mountain in Africa.", Popularity: 91},
	{ID: "31", Name: "Banff National Park", Country: "Canada", Description: "National park in the Canadian Rockies.", Popularity: 85},
	{ID: "32", Name: "Easter Island", Country: "Chile", Description: "Remote Polynesian island famous for its statues.", Popularity: 80},
	{ID: "33", Name: "Kilimanjaro", Country: "Tanzania", Description: "Highest mountain in Africa.", Popularity: 87},
	{ID: "34", Name: "Prague Castle", Country: "Czech Republic", Description: "Historic castle complex in Prague.", Popularity: 88},
	{ID: "35", Name: "Neuschwanstein Castle", Country: "Germany", Description: "Fairytale castle in Bavaria.", Popularity: 91},
	{ID: "36", Name: "Palace of Versailles", Country: "France", Description: "Opulent royal palace near Paris.", Popularity: 93},
	{ID: "37", Name: "Angkor Wat", Country: "Cambodia", Description: "Largest religious monument in the world.", Popularity: 97},
	{ID: "38", Name: "Mount Rushmore", Country: "USA", Description: "Sculpted faces of U.S. presidents in South Dakota.", Popularity: 92},
	{ID: "39", Name: "Mount Fuji", Country: "Japan", Description: "Iconic volcano and symbol of Japan.", Popularity: 90},
	{ID: "40", Name: "Bora Bora", Country: "French Polynesia", Description: "Tropical paradise known for its lagoon and resorts.", Popularity: 91},
	{ID: "41", Name: "Petra", Country: "Jordan", Description: "Ancient rock-cut city and archaeological site.", Popularity: 94},
	{ID: "42", Name: "Pike's Peak", Country: "USA", Description: "Mountain summit in Colorado.", Popularity: 85},
	{ID: "43", Name: "Mount Sinai", Country: "Egypt", Description: "Mountain in the Sinai Peninsula, traditionally where Moses received the Ten Commandments.", Popularity: 88},
	{ID: "44", Name: "Tulum", Country: "Mexico", Description: "Ruins of a Mayan port city overlooking the Caribbean.", Popularity: 91},
	{ID: "45", Name: "Bali", Country: "Indonesia", Description: "Tropical island known for its beaches, temples, and culture.", Popularity: 92},
	{ID: "46", Name: "Victoria Falls", Country: "Zambia/Zimbabwe", Description: "One of the largest and most famous waterfalls in the world.", Popularity: 93},
	{ID: "47", Name: "Machu Picchu", Country: "Peru", Description: "15th-century Incan citadel in the Andes Mountains.", Popularity: 91},
	{ID: "48", Name: "Santorini", Country: "Greece", Description: "Famous Greek island known for its whitewashed buildings and sunsets.", Popularity: 94},
	{ID: "49", Name: "The Great Barrier Reef", Country: "Australia", Description: "Largest coral reef system on Earth.", Popularity: 96},
	{ID: "50", Name: "Uluru", Country: "Australia", Description: "Iconic sandstone monolith in the Outback.", Popularity: 85},
	{ID: "51", Name: "Notre-Dame Cathedral", Country: "France", Description: "Famous Gothic cathedral in Paris.", Popularity: 90},
	{ID: "52", Name: "Mount Etna", Country: "Italy", Description: "Active stratovolcano in Sicily.", Popularity: 86},
	{ID: "53", Name: "Machu Picchu", Country: "Peru", Description: "Ancient Incan site in the Andes.", Popularity: 93},
	{ID: "54", Name: "Mount Kilimanjaro", Country: "Tanzania", Description: "Africa's highest peak.", Popularity: 94},
	{ID: "55", Name: "Matterhorn", Country: "Switzerland", Description: "Famous pyramid-shaped mountain in the Alps.", Popularity: 91},
	{ID: "56", Name: "Bryce Canyon", Country: "USA", Description: "National park with unique rock formations in Utah.", Popularity: 88},
	{ID: "57", Name: "Yellowstone National Park", Country: "USA", Description: "The first national park in the world, known for its geysers and wildlife.", Popularity: 92},
	{ID: "58", Name: "Mount Fuji", Country: "Japan", Description: "Iconic mountain known for its beauty and religious significance.", Popularity: 95},
	{ID: "59", Name: "Machu Picchu", Country: "Peru", Description: "Ancient Incan citadel in Peru's Andes Mountains.", Popularity: 91},
	{ID: "60", Name: "Venice Canals", Country: "Italy", Description: "Famous canals in Venice, known for their gondolas.", Popularity: 88},
	{ID: "61", Name: "Buckingham Palace", Country: "UK", Description: "Official residence of the British monarch.", Popularity: 92},
	{ID: "62", Name: "Alcatraz Island", Country: "USA", Description: "Former prison located in San Francisco Bay.", Popularity: 84},
	{ID: "63", Name: "Red Square", Country: "Russia", Description: "Famous historical and cultural square in Moscow.", Popularity: 87},
	{ID: "64", Name: "Sagrada Familia", Country: "Spain", Description: "Unfinished basilica designed by Antoni Gaudí.", Popularity: 94},
	{ID: "65", Name: "Bora Bora", Country: "French Polynesia", Description: "Exotic island known for its crystal-clear lagoon.", Popularity: 93},
	{ID: "66", Name: "Hagia Sophia", Country: "Turkey", Description: "Former cathedral, mosque, and now a museum in Istanbul.", Popularity: 91},
	{ID: "67", Name: "Mount Vesuvius", Country: "Italy", Description: "Active volcano near Naples, famous for its eruption in 79 AD.", Popularity: 89},
	{ID: "68", Name: "Louvre Museum", Country: "France", Description: "World's largest museum in Paris.", Popularity: 96},
	{ID: "69", Name: "The Great Wall of China", Country: "China", Description: "Ancient fortification stretching across northern China.", Popularity: 94},
	{ID: "70", Name: "Pyramids of Giza", Country: "Egypt", Description: "Ancient structures built as tombs for pharaohs.", Popularity: 93},
	{ID: "71", Name: "Mount Olympus", Country: "Greece", Description: "Highest mountain in Greece, home of the gods in Greek mythology.", Popularity: 85},
	{ID: "72", Name: "Machu Picchu", Country: "Peru", Description: "Famous Incan site set high in the Andes.", Popularity: 92},
	{ID: "73", Name: "Machu Picchu", Country: "Peru", Description: "Historic Incan city in the Peruvian Andes.", Popularity: 93},
	{ID: "74", Name: "Palace of Versailles", Country: "France", Description: "Lavish royal palace in Versailles.", Popularity: 95},
	{ID: "75", Name: "Catherine Palace", Country: "Russia", Description: "Baroque palace near St. Petersburg.", Popularity: 92},
	{ID: "76", Name: "Eiffel Tower", Country: "France", Description: "Iconic Parisian landmark.", Popularity: 98},
	{ID: "77", Name: "Blue Lagoon", Country: "Iceland", Description: "Geothermal spa in a lava field.", Popularity: 88},
	{ID: "78", Name: "Easter Island", Country: "Chile", Description: "Famous for its massive stone statues.", Popularity: 85},
	{ID: "79", Name: "La Sagrada Familia", Country: "Spain", Description: "Gaudí-designed church in Barcelona.", Popularity: 95},
	{ID: "80", Name: "Temple of Karnak", Country: "Egypt", Description: "Ancient temple complex in Luxor.", Popularity: 93},
	{ID: "81", Name: "Victoria Falls", Country: "Zambia/Zimbabwe", Description: "Massive waterfall on the Zambezi River.", Popularity: 92},
	{ID: "82", Name: "Sydney Opera House", Country: "Australia", Description: "Iconic performing arts venue in Sydney.", Popularity: 90},
	{ID: "83", Name: "Chichen Itza", Country: "Mexico", Description: "Mayan ruins on the Yucatán Peninsula.", Popularity: 91},
	{ID: "84", Name: "Lake Baikal", Country: "Russia", Description: "World's deepest freshwater lake.", Popularity: 84},
	{ID: "85", Name: "Great Barrier Reef", Country: "Australia", Description: "Largest reef system in the world.", Popularity: 95},
	{ID: "86", Name: "Grand Canyon", Country: "USA", Description: "Massive canyon carved by the Colorado River.", Popularity: 90},
	{ID: "87", Name: "Mount Rushmore", Country: "USA", Description: "Famous presidential monument in South Dakota.", Popularity: 91},
	{ID: "88", Name: "Yellowstone National Park", Country: "USA", Description: "World's first national park known for its geothermal features.", Popularity: 93},
	{ID: "89", Name: "Paris Catacombs", Country: "France", Description: "Underground ossuary in Paris.", Popularity: 85},
	{ID: "90", Name: "Cinque Terre", Country: "Italy", Description: "Five picturesque villages on the Italian Riviera.", Popularity: 87},
	{ID: "91", Name: "Mesa Verde", Country: "USA", Description: "Ancestral Puebloans cliff dwellings in Colorado.", Popularity: 86},
	{ID: "92", Name: "Pyramids of Giza", Country: "Egypt", Description: "Ancient Egyptian tombs.", Popularity: 93},
	{ID: "93", Name: "Reykjavik", Country: "Iceland", Description: "Capital city known for its natural beauty and modern culture.", Popularity: 81},
	{ID: "94", Name: "Easter Island", Country: "Chile", Description: "Remote island known for its large stone heads.", Popularity: 88},
	{ID: "95", Name: "Krakow", Country: "Poland", Description: "Historic city known for its medieval architecture.", Popularity: 85},
	{ID: "96", Name: "Paris", Country: "France", Description: "Capital city famous for its art, history, and culture.", Popularity: 99},
	{ID: "97", Name: "Giza Pyramids", Country: "Egypt", Description: "Ancient tombs built for Egyptian pharaohs.", Popularity: 92},
	{ID: "98", Name: "Zhangjiajie National Forest Park", Country: "China", Description: "National park known for its pillar-like formations.", Popularity: 90},
	{ID: "99", Name: "Kangaroo Island", Country: "Australia", Description: "Island known for its wildlife and beautiful beaches.", Popularity: 89},
	{ID: "100", Name: "Hagia Sophia", Country: "Turkey", Description: "Former church, mosque, and museum in Istanbul.", Popularity: 92},
}

// searchPlacesByCountry searches for places by country.
// swagger:route GET /searchByCountry places searchPlacesByCountry
//
// Search for tourist places based on the country name.
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: searchResponse
//	  400: errorResponse
func searchPlacesByCountry(c *gin.Context) {
	country := c.DefaultQuery("country", "")

	if country == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country"})
		return
	}

	var results []Place
	for _, place := range placesList {
		if strings.Contains(strings.ToLower(place.Country), strings.ToLower(country)) {
			results = append(results, place)
		}
	}

	c.JSON(http.StatusOK, searchResponse{Places: results})
}

func main() {
	// Initialize gin router
	r := gin.Default()

	// Register the search endpoint
	r.GET("/searchByCountry", searchPlacesByCountry)

	// Run the server
	r.Run(":9080")
}

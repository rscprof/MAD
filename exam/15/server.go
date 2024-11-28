package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Структура для хранения информации о достопримечательности
type Location struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

// Массив с данными о достопримечательностях
var locations = []Location{
	{ID: "1", Name: "Eiffel Tower", Description: "Iconic symbol of Paris", Category: "Landmark", Latitude: 48.8584, Longitude: 2.2945},
	{ID: "2", Name: "Statue of Liberty", Description: "Famous American monument", Category: "Landmark", Latitude: 40.6892, Longitude: -74.0445},
	{ID: "3", Name: "Great Wall of China", Description: "Ancient series of walls and fortifications", Category: "Historical", Latitude: 40.4319, Longitude: 116.5704},
	{ID: "4", Name: "Sydney Opera House", Description: "Famous Australian performing arts venue", Category: "Cultural", Latitude: -33.8568, Longitude: 151.2153},
	{ID: "5", Name: "Grand Canyon", Description: "Famous natural landmark in the USA", Category: "Natural", Latitude: 36.1069, Longitude: -112.1129},
	{ID: "6", Name: "Mount Everest", Description: "Highest mountain in the world", Category: "Natural", Latitude: 27.9881, Longitude: 86.9250},
	{ID: "7", Name: "Taj Mahal", Description: "Iconic Indian mausoleum", Category: "Historical", Latitude: 27.1751, Longitude: 78.0421},
	{ID: "8", Name: "Christ the Redeemer", Description: "Famous statue in Rio de Janeiro", Category: "Landmark", Latitude: -22.9519, Longitude: -43.2105},
	{ID: "9", Name: "Niagara Falls", Description: "Famous waterfall on the USA-Canada border", Category: "Natural", Latitude: 43.0962, Longitude: -79.0377},
	{ID: "10", Name: "Colosseum", Description: "Ancient Roman amphitheater", Category: "Historical", Latitude: 41.8902, Longitude: 12.4922},
	{ID: "11", Name: "Machu Picchu", Description: "Ancient Incan city in Peru", Category: "Historical", Latitude: -13.1631, Longitude: -72.5450},
	{ID: "12", Name: "Yellowstone National Park", Description: "Famous US national park", Category: "Natural", Latitude: 44.4280, Longitude: -110.5885},
	{ID: "13", Name: "Big Ben", Description: "Iconic clock tower in London", Category: "Landmark", Latitude: 51.5007, Longitude: -0.1246},
	{ID: "14", Name: "Pyramids of Giza", Description: "Ancient Egyptian pyramids", Category: "Historical", Latitude: 29.9792, Longitude: 31.1342},
	{ID: "15", Name: "Burj Khalifa", Description: "Tallest skyscraper in Dubai", Category: "Landmark", Latitude: 25.1972, Longitude: 55.2744},
	{ID: "16", Name: "Angel Falls", Description: "Tallest waterfall in Venezuela", Category: "Natural", Latitude: 5.9670, Longitude: -62.5356},
	{ID: "17", Name: "Santorini", Description: "Picturesque Greek island", Category: "Cultural", Latitude: 36.3932, Longitude: 25.4615},
	{ID: "18", Name: "Louvre Museum", Description: "Famous art museum in Paris", Category: "Cultural", Latitude: 48.8606, Longitude: 2.3376},
	{ID: "19", Name: "Amazon Rainforest", Description: "Largest tropical rainforest", Category: "Natural", Latitude: -3.4653, Longitude: -62.2159},
	{ID: "20", Name: "Hagia Sophia", Description: "Historic mosque in Istanbul", Category: "Cultural", Latitude: 41.0086, Longitude: 28.9802},
	{ID: "21", Name: "Stonehenge", Description: "Prehistoric monument in England", Category: "Historical", Latitude: 51.1789, Longitude: -1.8262},
	{ID: "22", Name: "Machu Picchu", Description: "Ancient Incan city in Peru", Category: "Historical", Latitude: -13.1631, Longitude: -72.5450},
	{ID: "23", Name: "Mount Fuji", Description: "Iconic volcanic mountain in Japan", Category: "Natural", Latitude: 35.3606, Longitude: 138.7274},
	{ID: "24", Name: "Alhambra", Description: "Historic palace and fortress in Spain", Category: "Historical", Latitude: 37.7772, Longitude: -3.3568},
	{ID: "25", Name: "Mount Kilimanjaro", Description: "Highest peak in Africa", Category: "Natural", Latitude: -3.0674, Longitude: 37.3556},
	{ID: "26", Name: "Acropolis of Athens", Description: "Ancient citadel in Greece", Category: "Historical", Latitude: 37.9715, Longitude: 23.7267},
	{ID: "27", Name: "The Louvre", Description: "World’s largest art museum", Category: "Cultural", Latitude: 48.8606, Longitude: 2.3376},
	{ID: "28", Name: "Sagrada Familia", Description: "Famous cathedral in Barcelona", Category: "Cultural", Latitude: 41.4036, Longitude: 2.1744},
	{ID: "29", Name: "Palace of Versailles", Description: "Historic French palace", Category: "Historical", Latitude: 48.8049, Longitude: 2.1204},
	{ID: "30", Name: "Everglades National Park", Description: "Wetland ecosystem in Florida", Category: "Natural", Latitude: 25.2866, Longitude: -80.8987},
	{ID: "31", Name: "Banff National Park", Description: "National park in Canada", Category: "Natural", Latitude: 51.4968, Longitude: -115.9281},
	{ID: "32", Name: "Machu Picchu", Description: "Ancient city of the Incas", Category: "Historical", Latitude: -13.1631, Longitude: -72.5450},
	{ID: "33", Name: "Great Barrier Reef", Description: "World's largest coral reef system", Category: "Natural", Latitude: -18.2871, Longitude: 147.6992},
	{ID: "34", Name: "Mount Rushmore", Description: "Monument with carved faces of presidents", Category: "Landmark", Latitude: 43.8791, Longitude: -103.4591},
	{ID: "35", Name: "Venice Canals", Description: "Famous canals in Venice, Italy", Category: "Cultural", Latitude: 45.4408, Longitude: 12.3155},
	{ID: "36", Name: "Grand Bazaar", Description: "Famous market in Istanbul", Category: "Cultural", Latitude: 41.0100, Longitude: 28.9794},
	{ID: "37", Name: "Kilimanjaro National Park", Description: "Home of Mount Kilimanjaro", Category: "Natural", Latitude: -3.0758, Longitude: 37.3556},
	{ID: "38", Name: "Bora Bora", Description: "Tropical island in French Polynesia", Category: "Cultural", Latitude: -16.5004, Longitude: -151.7415},
	{ID: "39", Name: "Santorini", Description: "Island in Greece known for its beauty", Category: "Cultural", Latitude: 36.3932, Longitude: 25.4615},
	{ID: "40", Name: "Petra", Description: "Ancient city in Jordan", Category: "Historical", Latitude: 30.3285, Longitude: 35.4444},
	{ID: "41", Name: "Mount Kilimanjaro", Description: "Highest peak in Africa", Category: "Natural", Latitude: -3.0674, Longitude: 37.3556},
	{ID: "42", Name: "Machu Picchu", Description: "Incan city in Peru", Category: "Historical", Latitude: -13.1631, Longitude: -72.5450},
	{ID: "43", Name: "Victoria Falls", Description: "Massive waterfall on the Zambezi River", Category: "Natural", Latitude: -17.9246, Longitude: 25.8567},
	{ID: "44", Name: "Table Mountain", Description: "Famous flat-topped mountain in Cape Town", Category: "Natural", Latitude: -33.9628, Longitude: 18.4098},
	{ID: "45", Name: "Great Barrier Reef", Description: "World’s largest coral reef", Category: "Natural", Latitude: -18.2871, Longitude: 147.6992},
	{ID: "46", Name: "The Vatican", Description: "Religious center of the Catholic Church", Category: "Cultural", Latitude: 41.9029, Longitude: 12.4534},
	{ID: "47", Name: "Sagrada Familia", Description: "Famous church designed by Gaudí", Category: "Cultural", Latitude: 41.4036, Longitude: 2.1744},
	{ID: "48", Name: "Chichen Itza", Description: "Ancient Mayan city in Mexico", Category: "Historical", Latitude: 20.6843, Longitude: -88.5678},
	{ID: "49", Name: "Galapagos Islands", Description: "Island group with unique wildlife", Category: "Natural", Latitude: -0.9538, Longitude: -90.9656},
	{ID: "50", Name: "Yellowstone National Park", Description: "USA’s first national park", Category: "Natural", Latitude: 44.4280, Longitude: -110.5885},
	{ID: "51", Name: "Tulum", Description: "Ancient Mayan port city in Mexico", Category: "Historical", Latitude: 20.2110, Longitude: -87.4658},
	{ID: "52", Name: "Serengeti National Park", Description: "Iconic national park in Tanzania", Category: "Natural", Latitude: -2.3333, Longitude: 34.8333},
	{ID: "53", Name: "Mount Etna", Description: "Active volcano in Italy", Category: "Natural", Latitude: 37.7510, Longitude: 15.0046},
	{ID: "54", Name: "Burj Khalifa", Description: "Tallest building in the world", Category: "Landmark", Latitude: 25.1972, Longitude: 55.2744},
	{ID: "55", Name: "Hollywood Sign", Description: "Famous landmark in Los Angeles", Category: "Landmark", Latitude: 34.1341, Longitude: -118.3217},
	{ID: "56", Name: "Palace of Versailles", Description: "Historical French royal palace", Category: "Historical", Latitude: 48.8049, Longitude: 2.1204},
	{ID: "57", Name: "Cairo Museum", Description: "Egyptian antiquities museum", Category: "Cultural", Latitude: 30.0477, Longitude: 31.2357},
	{ID: "58", Name: "Machu Picchu", Description: "Ancient Inca city in Peru", Category: "Historical", Latitude: -13.1631, Longitude: -72.5450},
	{ID: "59", Name: "Mount Fuji", Description: "Iconic Japanese mountain", Category: "Natural", Latitude: 35.3606, Longitude: 138.7274},
	{ID: "60", Name: "Palace of the Parliament", Description: "Massive government building in Bucharest", Category: "Historical", Latitude: 44.4268, Longitude: 26.1025},
	{ID: "61", Name: "Bora Bora", Description: "Tropical island paradise", Category: "Cultural", Latitude: -16.5004, Longitude: -151.7415},
	{ID: "62", Name: "Great Barrier Reef", Description: "World’s largest coral reef system", Category: "Natural", Latitude: -18.2871, Longitude: 147.6992},
	{ID: "63", Name: "Stonehenge", Description: "Mysterious prehistoric monument", Category: "Historical", Latitude: 51.1789, Longitude: -1.8262},
	{ID: "64", Name: "Grand Canyon", Description: "Massive natural canyon in the USA", Category: "Natural", Latitude: 36.1069, Longitude: -112.1129},
	{ID: "65", Name: "The Colosseum", Description: "Roman amphitheater", Category: "Historical", Latitude: 41.8902, Longitude: 12.4922},
	{ID: "66", Name: "Petra", Description: "Ancient city in Jordan", Category: "Historical", Latitude: 30.3285, Longitude: 35.4444},
	{ID: "67", Name: "Pyramids of Giza", Description: "Ancient pyramids in Egypt", Category: "Historical", Latitude: 29.9792, Longitude: 31.1342},
	{ID: "68", Name: "Mount Kilimanjaro", Description: "Highest peak in Africa", Category: "Natural", Latitude: -3.0674, Longitude: 37.3556},
	{ID: "69", Name: "Yellowstone National Park", Description: "USA’s first national park", Category: "Natural", Latitude: 44.4280, Longitude: -110.5885},
	{ID: "70", Name: "Eiffel Tower", Description: "Famous Parisian landmark", Category: "Landmark", Latitude: 48.8584, Longitude: 2.2945},
	{ID: "71", Name: "Grand Canyon", Description: "Massive natural formation in Arizona", Category: "Natural", Latitude: 36.1069, Longitude: -112.1129},
	{ID: "72", Name: "Acropolis", Description: "Ancient citadel in Athens", Category: "Historical", Latitude: 37.9715, Longitude: 23.7267},
	{ID: "73", Name: "Big Ben", Description: "Iconic clock tower in London", Category: "Landmark", Latitude: 51.5007, Longitude: -0.1246},
	{ID: "74", Name: "Mount Rushmore", Description: "Famous presidential monument", Category: "Landmark", Latitude: 43.8791, Longitude: -103.4591},
	{ID: "75", Name: "Mount Fuji", Description: "Active volcano in Japan", Category: "Natural", Latitude: 35.3606, Longitude: 138.7274},
	{ID: "76", Name: "Great Barrier Reef", Description: "World’s largest coral reef system", Category: "Natural", Latitude: -18.2871, Longitude: 147.6992},
	{ID: "77", Name: "Sagrada Familia", Description: "Iconic basilica in Barcelona", Category: "Cultural", Latitude: 41.4036, Longitude: 2.1744},
	{ID: "78", Name: "The Great Wall of China", Description: "Ancient fortification in China", Category: "Historical", Latitude: 40.4319, Longitude: 116.5704},
	{ID: "79", Name: "Machu Picchu", Description: "Ancient Incan city in Peru", Category: "Historical", Latitude: -13.1631, Longitude: -72.5450},
	{ID: "80", Name: "Taj Mahal", Description: "Iconic Indian mausoleum", Category: "Historical", Latitude: 27.1751, Longitude: 78.0421},
	{ID: "81", Name: "Christ the Redeemer", Description: "Famous statue in Brazil", Category: "Landmark", Latitude: -22.9519, Longitude: -43.2105},
	{ID: "82", Name: "Stonehenge", Description: "Prehistoric monument in England", Category: "Historical", Latitude: 51.1789, Longitude: -1.8262},
	{ID: "83", Name: "Galapagos Islands", Description: "Unique wildlife and ecosystems", Category: "Natural", Latitude: -0.9538, Longitude: -90.9656},
	{ID: "84", Name: "Colosseum", Description: "Roman amphitheater in Rome", Category: "Historical", Latitude: 41.8902, Longitude: 12.4922},
	{ID: "85", Name: "Venice Canals", Description: "Iconic canals in Venice", Category: "Cultural", Latitude: 45.4408, Longitude: 12.3155},
	{ID: "86", Name: "Hagia Sophia", Description: "Historic mosque in Turkey", Category: "Cultural", Latitude: 41.0086, Longitude: 28.9802},
	{ID: "87", Name: "Mount Etna", Description: "Active volcano in Sicily", Category: "Natural", Latitude: 37.7510, Longitude: 15.0046},
	{ID: "88", Name: "The Vatican", Description: "Headquarters of the Roman Catholic Church", Category: "Cultural", Latitude: 41.9029, Longitude: 12.4534},
	{ID: "89", Name: "Serengeti National Park", Description: "Home to the annual wildebeest migration", Category: "Natural", Latitude: -2.3333, Longitude: 34.8333},
	{ID: "90", Name: "Mount Kilimanjaro", Description: "Highest peak in Africa", Category: "Natural", Latitude: -3.0674, Longitude: 37.3556},
	{ID: "91", Name: "Victoria Falls", Description: "Massive waterfall in southern Africa", Category: "Natural", Latitude: -17.9246, Longitude: 25.8567},
	{ID: "92", Name: "Burj Khalifa", Description: "Tallest building in the world", Category: "Landmark", Latitude: 25.1972, Longitude: 55.2744},
	{ID: "93", Name: "Mount Rushmore", Description: "Presidential monument in the USA", Category: "Landmark", Latitude: 43.8791, Longitude: -103.4591},
	{ID: "94", Name: "Eiffel Tower", Description: "Iconic Paris landmark", Category: "Landmark", Latitude: 48.8584, Longitude: 2.2945},
	{ID: "95", Name: "Pyramids of Giza", Description: "Ancient pyramids in Egypt", Category: "Historical", Latitude: 29.9792, Longitude: 31.1342},
	{ID: "96", Name: "Sagrada Familia", Description: "Architectural masterpiece in Barcelona", Category: "Cultural", Latitude: 41.4036, Longitude: 2.1744},
	{ID: "97", Name: "Great Wall of China", Description: "Ancient fortification in China", Category: "Historical", Latitude: 40.4319, Longitude: 116.5704},
	{ID: "98", Name: "Machu Picchu", Description: "Ancient Incan city in Peru", Category: "Historical", Latitude: -13.1631, Longitude: -72.5450},
	{ID: "99", Name: "Mount Fuji", Description: "Iconic Japanese volcano", Category: "Natural", Latitude: 35.3606, Longitude: 138.7274},
	{ID: "100", Name: "Taj Mahal", Description: "Symbol of love in India", Category: "Historical", Latitude: 27.1751, Longitude: 78.0421},
}

// Функция для обработки запроса поиска по названию
func searchLocations(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр "name" из URL-запроса
	name := r.URL.Query().Get("name")

	// Фильтруем достопримечательности по названию
	var filteredLocations []Location
	for _, location := range locations {
		// Если в названии содержится значение параметра "name", добавляем в результат
		if name != "" && strings.Contains(strings.ToLower(location.Name), strings.ToLower(name)) {
			filteredLocations = append(filteredLocations, location)
		}
	}

	// Если не нашли соответствующих результатов
	if len(filteredLocations) == 0 {
		http.Error(w, "No locations found for the given name", http.StatusNotFound)
		return
	}

	// Отправляем ответ в формате JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(filteredLocations); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %s", err), http.StatusInternalServerError)
	}
}

func main() {
	// Устанавливаем обработчик для маршрута /search
	http.HandleFunc("/search", searchLocations)

	// Запускаем сервер
	log.Println("Server started on :9080")
	if err := http.ListenAndServe(":9080", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

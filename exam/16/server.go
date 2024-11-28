package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Структура для хранения информации о ресторане
type Restaurant struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Cuisine  string `json:"cuisine"`
	Location string `json:"location"`
	IsOpen   bool   `json:"is_open"`
}

// Массив с данными о ресторанах
var restaurants = []Restaurant{
	{"1", "Pasta House", "Italian", "New York", true},
	{"2", "Sushi World", "Japanese", "Los Angeles", false},
	{"3", "Taco Heaven", "Mexican", "Miami", true},
	{"4", "Pizza Planet", "Italian", "Chicago", true},
	{"5", "Bamboo Garden", "Chinese", "San Francisco", false},
	{"6", "Burger Bistro", "American", "Los Angeles", true},
	{"7", "Curry King", "Indian", "Dallas", true},
	{"8", "Sushi Delight", "Japanese", "Seattle", false},
	{"9", "Green Veggie", "Vegetarian", "Portland", true},
	{"10", "BBQ Central", "American", "Austin", true},
	{"11", "Royal Feast", "Indian", "Washington", false},
	{"12", "Seafood Palace", "Seafood", "Miami", true},
	{"13", "Spicy Delight", "Mexican", "Chicago", true},
	{"14", "Chopsticks", "Chinese", "San Diego", true},
	{"15", "Gusto", "Italian", "Boston", false},
	{"16", "Sushi Sensation", "Japanese", "Atlanta", true},
	{"17", "Café de Paris", "French", "San Francisco", false},
	{"18", "Grill Master", "American", "Los Angeles", true},
	{"19", "Falafel Stand", "Middle Eastern", "New York", true},
	{"20", "Noodle House", "Chinese", "Houston", false},
	{"21", "Grill & Chill", "American", "Austin", true},
	{"22", "Sushi Corner", "Japanese", "Seattle", true},
	{"23", "The Burger Joint", "American", "Miami", true},
	{"24", "Burrito Bar", "Mexican", "Phoenix", true},
	{"25", "Pasta & Co.", "Italian", "Chicago", true},
	{"26", "Curry Delight", "Indian", "Dallas", true},
	{"27", "Lobster Shack", "Seafood", "Boston", false},
	{"28", "The Vegan Table", "Vegetarian", "Portland", true},
	{"29", "Steakhouse 101", "American", "Chicago", true},
	{"30", "Sushi & Roll", "Japanese", "Los Angeles", false},
	{"31", "Bistro 56", "French", "Miami", true},
	{"32", "Tandoori Nights", "Indian", "New York", true},
	{"33", "Ocean Breeze", "Seafood", "San Diego", false},
	{"34", "Baja Grill", "Mexican", "Dallas", true},
	{"35", "Fried Chicken Place", "American", "Phoenix", true},
	{"36", "Pasta Vino", "Italian", "Seattle", true},
	{"37", "Taste of India", "Indian", "Chicago", true},
	{"38", "Fusion Sushi", "Japanese", "San Francisco", true},
	{"39", "Grill House", "American", "Dallas", true},
	{"40", "Spicy Tacos", "Mexican", "New York", true},
	{"41", "Crispy Crepes", "French", "Austin", false},
	{"42", "The Pizzeria", "Italian", "Portland", true},
	{"43", "Tempura Garden", "Japanese", "Los Angeles", true},
	{"44", "Shrimp Shack", "Seafood", "San Diego", true},
	{"45", "Loco for Tacos", "Mexican", "Phoenix", true},
	{"46", "Pizza Express", "Italian", "Miami", false},
	{"47", "Fried Fish House", "Seafood", "Chicago", true},
	{"48", "Burrito Bonanza", "Mexican", "San Francisco", true},
	{"49", "Mamma's Italian", "Italian", "Dallas", true},
	{"50", "Dragon's Den", "Chinese", "Los Angeles", true},
	{"51", "Café Italiano", "Italian", "Houston", true},
	{"52", "Wok & Roll", "Chinese", "New York", true},
	{"53", "Chili's Grill", "American", "Chicago", false},
	{"54", "Lobster Lounge", "Seafood", "Boston", true},
	{"55", "Sushi Delight", "Japanese", "Dallas", false},
	{"56", "Grill & Bar", "American", "Seattle", true},
	{"57", "Flavors of India", "Indian", "Portland", true},
	{"58", "Pho House", "Vietnamese", "San Diego", true},
	{"59", "Bavarian Grill", "German", "New York", true},
	{"60", "Miso Soup", "Japanese", "Phoenix", true},
	{"61", "Café La Parisienne", "French", "Chicago", true},
	{"62", "Veggie Heaven", "Vegetarian", "Miami", false},
	{"63", "Sushi Time", "Japanese", "Los Angeles", true},
	{"64", "Buns & Bites", "American", "San Francisco", true},
	{"65", "The Indian Spice", "Indian", "Portland", true},
	{"66", "Pasta Place", "Italian", "New York", true},
	{"67", "Banh Mi Corner", "Vietnamese", "Dallas", true},
	{"68", "Sushi Sushi", "Japanese", "Seattle", false},
	{"69", "American Grub", "American", "Houston", true},
	{"70", "Curry Palace", "Indian", "Austin", true},
	{"71", "Ginger's Grill", "Chinese", "San Francisco", true},
	{"72", "Taco Stand", "Mexican", "Los Angeles", false},
	{"73", "The Pasta Bar", "Italian", "Phoenix", true},
	{"74", "Bento Box", "Japanese", "Dallas", true},
	{"75", "Pizzeria Italiano", "Italian", "Boston", false},
	{"76", "Ramen Delight", "Japanese", "Portland", true},
	{"77", "Margarita Bar", "Mexican", "San Diego", true},
	{"78", "Curry & Spice", "Indian", "Miami", true},
	{"79", "The Lobster Pot", "Seafood", "New York", false},
	{"80", "Crab Shack", "Seafood", "Los Angeles", true},
	{"81", "Taco Tower", "Mexican", "Chicago", true},
	{"82", "La Pasta", "Italian", "San Francisco", true},
	{"83", "The Stir Fry", "Chinese", "Seattle", true},
	{"84", "Sushi Junction", "Japanese", "Austin", true},
	{"85", "Indian Taste", "Indian", "New York", true},
	{"86", "Noodles & Wok", "Chinese", "Portland", true},
	{"87", "Pasta Paradise", "Italian", "Los Angeles", true},
	{"88", "The Sushi Bar", "Japanese", "San Diego", false},
	{"89", "Grill Hub", "American", "Chicago", true},
	{"90", "Gourmet Burgers", "American", "Dallas", true},
	{"91", "Spicy Kitchen", "Indian", "Houston", true},
	{"92", "The Vegan Grill", "Vegetarian", "Austin", true},
	{"93", "Seafood Heaven", "Seafood", "Miami", false},
	{"94", "Falafel Grill", "Middle Eastern", "New York", true},
	{"95", "Ramen Corner", "Japanese", "Phoenix", true},
	{"96", "Taco Express", "Mexican", "San Francisco", true},
	{"97", "Pizza Town", "Italian", "Dallas", true},
	{"98", "Sushi Express", "Japanese", "Portland", true},
	{"99", "The Burger Spot", "American", "San Diego", true},
	{"100", "Noodle Palace", "Chinese", "Chicago", false},
}

// Функция для обработки запроса поиска по названию
func searchRestaurants(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр "name" из URL-запроса
	name := r.URL.Query().Get("name")

	// Фильтруем рестораны по названию
	var filteredRestaurants []Restaurant
	for _, restaurant := range restaurants {
		// Если в названии содержится значение параметра "name", добавляем в результат
		if name != "" && strings.Contains(strings.ToLower(restaurant.Name), strings.ToLower(name)) {
			filteredRestaurants = append(filteredRestaurants, restaurant)
		}
	}

	// Если не нашли соответствующих результатов
	if len(filteredRestaurants) == 0 {
		http.Error(w, "No restaurants found for the given name", http.StatusNotFound)
		return
	}

	// Отправляем ответ в формате JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(filteredRestaurants); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %s", err), http.StatusInternalServerError)
	}
}

func main() {
	// Устанавливаем обработчик для маршрута /search
	http.HandleFunc("/search", searchRestaurants)

	// Запускаем сервер на порту 9080
	log.Println("Server started on :9080")
	if err := http.ListenAndServe(":9080", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

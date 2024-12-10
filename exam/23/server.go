package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Route represents a bus route.
type Route struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Length        float64 `json:"length"` // Length in kilometers
	NumberOfStops int     `json:"number_of_stops"`
}

// Sample list of routes
var routes = []Route{
	{"1", "Route A", 12.5, 8},
	{"2", "Route B", 25.3, 15},
	{"3", "Route C", 7.8, 6},
	{"4", "Route D", 18.0, 10},
	{"5", "Route E", 30.4, 20},
	{"6", "Route F", 5.6, 5},
	{"7", "Route G", 22.2, 12},
	{"8", "Route H", 10.0, 7},
	{"9", "Route I", 14.7, 9},
	{"10", "Route J", 28.9, 18},
	{"11", "Route K", 9.5, 6},
	{"12", "Route L", 16.3, 11},
	{"13", "Route M", 19.8, 14},
	{"14", "Route N", 26.0, 13},
	{"15", "Route O", 32.4, 22},
	{"16", "Route P", 6.3, 5},
	{"17", "Route Q", 21.7, 16},
	{"18", "Route R", 11.2, 7},
	{"19", "Route S", 8.8, 6},
	{"20", "Route T", 17.9, 10},
	{"21", "Route U", 15.5, 12},
	{"22", "Route V", 13.0, 9},
	{"23", "Route W", 23.4, 18},
	{"24", "Route X", 20.8, 14},
	{"25", "Route Y", 29.7, 19},
	{"26", "Route Z", 27.3, 16},
	{"27", "Route AA", 31.1, 21},
	{"28", "Route AB", 33.8, 24},
	{"29", "Route AC", 7.4, 4},
	{"30", "Route AD", 9.9, 6},
	{"31", "Route AE", 16.8, 11},
	{"32", "Route AF", 18.5, 13},
	{"33", "Route AG", 24.9, 17},
	{"34", "Route AH", 30.2, 22},
	{"35", "Route AI", 10.4, 8},
	{"36", "Route AJ", 12.6, 9},
	{"37", "Route AK", 14.9, 11},
	{"38", "Route AL", 26.5, 19},
	{"39", "Route AM", 32.9, 23},
	{"40", "Route AN", 28.1, 20},
	{"41", "Route AO", 5.8, 5},
	{"42", "Route AP", 11.3, 7},
	{"43", "Route AQ", 9.7, 6},
	{"44", "Route AR", 20.3, 15},
	{"45", "Route AS", 23.8, 16},
	{"46", "Route AT", 25.6, 18},
	{"47", "Route AU", 13.5, 8},
	{"48", "Route AV", 17.1, 11},
	{"49", "Route AW", 29.3, 19},
	{"50", "Route AX", 31.7, 22},
	{"51", "Route AY", 8.6, 5},
	{"52", "Route AZ", 19.3, 13},
	{"53", "Route BA", 21.2, 14},
	{"54", "Route BB", 24.4, 17},
	{"55", "Route BC", 33.5, 25},
	{"56", "Route BD", 27.7, 20},
	{"57", "Route BE", 10.9, 7},
	{"58", "Route BF", 14.2, 9},
	{"59", "Route BG", 15.8, 10},
	{"60", "Route BH", 22.5, 15},
	{"61", "Route BI", 30.0, 21},
	{"62", "Route BJ", 6.7, 4},
	{"63", "Route BK", 9.4, 6},
	{"64", "Route BL", 18.7, 12},
	{"65", "Route BM", 20.0, 13},
	{"66", "Route BN", 25.0, 18},
	{"67", "Route BO", 29.0, 19},
	{"68", "Route BP", 7.2, 5},
	{"69", "Route BQ", 11.6, 8},
	{"70", "Route BR", 16.5, 11},
	{"71", "Route BS", 19.4, 13},
	{"72", "Route BT", 28.8, 20},
	{"73", "Route BU", 31.0, 22},
	{"74", "Route BV", 12.3, 9},
	{"75", "Route BW", 8.4, 6},
	{"76", "Route BX", 14.5, 10},
	{"77", "Route BY", 21.9, 16},
	{"78", "Route BZ", 23.0, 15},
	{"79", "Route CA", 26.7, 18},
	{"80", "Route CB", 30.7, 21},
	{"81", "Route CC", 7.6, 5},
	{"82", "Route CD", 11.8, 7},
	{"83", "Route CE", 13.7, 9},
	{"84", "Route CF", 18.9, 12},
	{"85", "Route CG", 22.0, 14},
	{"86", "Route CH", 27.2, 17},
	{"87", "Route CI", 9.2, 6},
	{"88", "Route CJ", 15.4, 11},
	{"89", "Route CK", 19.1, 13},
	{"90", "Route CL", 24.7, 18},
	{"91", "Route CM", 28.3, 21},
	{"92", "Route CN", 6.5, 4},
	{"93", "Route CO", 10.1, 7},
	{"94", "Route CP", 17.3, 11},
	{"95", "Route CQ", 20.5, 14},
	{"96", "Route CR", 29.5, 20},
	{"97", "Route CS", 33.0, 24},
	{"98", "Route CT", 8.2, 5},
	{"99", "Route CU", 12.9, 8},
	{"100", "Route CV", 16.1, 10},
}

func main() {
	http.HandleFunc("/routes", getRoutesByLength)

	port := ":9080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// getRoutesByLength filters bus routes based on their length
func getRoutesByLength(w http.ResponseWriter, r *http.Request) {
	minLengthStr := r.URL.Query().Get("min_length")
	maxLengthStr := r.URL.Query().Get("max_length")

	if minLengthStr == "" || maxLengthStr == "" {
		http.Error(w, "Both min_length and max_length query parameters are required", http.StatusBadRequest)
		return
	}

	// Convert string values to float
	minLength, err := strconv.ParseFloat(minLengthStr, 64)
	if err != nil {
		http.Error(w, "Invalid min_length value", http.StatusBadRequest)
		return
	}

	maxLength, err := strconv.ParseFloat(maxLengthStr, 64)
	if err != nil {
		http.Error(w, "Invalid max_length value", http.StatusBadRequest)
		return
	}

	// Find routes within the specified length range
	var result []Route
	for _, route := range routes {
		if route.Length >= minLength && route.Length <= maxLength {
			result = append(result, route)
		}
	}

	// If no routes are found
	if len(result) == 0 {
		http.Error(w, "No routes found within the specified length range", http.StatusNotFound)
		return
	}

	// Send the result as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

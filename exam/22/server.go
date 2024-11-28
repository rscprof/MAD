package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Apartment struct {
	ID        string  `json:"id"`
	Location  string  `json:"location"`
	Price     float64 `json:"price"`
	Area      float64 `json:"area"`
	Floor     int     `json:"floor"`
	Rooms     int     `json:"rooms"`
	YearBuilt int     `json:"year_built"`
}

var apartments = []Apartment{
	{"1", "New York", 500000, 70, 10, 2, 2000},
	{"2", "Los Angeles", 450000, 85, 5, 3, 2010},
	{"3", "Chicago", 350000, 60, 8, 2, 2015},
	{"4", "San Francisco", 750000, 95, 3, 4, 2018},
	{"5", "Miami", 300000, 55, 6, 1, 2005},
	{"6", "Austin", 400000, 80, 12, 3, 2020},
	{"7", "Boston", 650000, 90, 7, 3, 2012},
	{"8", "Seattle", 550000, 75, 9, 2, 2016},
	{"9", "Denver", 350000, 65, 4, 1, 2010},
	{"10", "Washington D.C.", 600000, 100, 2, 4, 2022},
	{"11", "Houston", 425000, 85, 6, 3, 2014},
	{"12", "Dallas", 480000, 92, 11, 3, 2017},
	{"13", "Philadelphia", 370000, 65, 4, 2, 2008},
	{"14", "Atlanta", 420000, 78, 10, 3, 2019},
	{"15", "Portland", 540000, 88, 6, 3, 2016},
	{"16", "Phoenix", 400000, 80, 8, 3, 2015},
	{"17", "Minneapolis", 370000, 70, 5, 2, 2013},
	{"18", "Cleveland", 330000, 68, 4, 2, 2012},
	{"19", "Detroit", 290000, 60, 9, 2, 2009},
	{"20", "San Diego", 710000, 95, 4, 4, 2021},
	{"21", "Las Vegas", 360000, 72, 6, 2, 2010},
	{"22", "Indianapolis", 340000, 62, 8, 2, 2014},
	{"23", "Columbus", 330000, 59, 3, 1, 2011},
	{"24", "Charlotte", 460000, 78, 7, 3, 2017},
	{"25", "Nashville", 480000, 85, 5, 3, 2020},
	{"26", "Salt Lake City", 490000, 90, 2, 3, 2019},
	{"27", "Orlando", 550000, 95, 6, 4, 2022},
	{"28", "Tampa", 600000, 100, 10, 4, 2023},
	{"29", "Kansas City", 410000, 80, 4, 2, 2018},
	{"30", "Raleigh", 530000, 88, 9, 3, 2014},
	{"31", "Chicago", 350000, 70, 8, 2, 2005},
	{"32", "Austin", 450000, 85, 10, 3, 2021},
	{"33", "Dallas", 550000, 92, 4, 4, 2019},
	{"34", "San Francisco", 750000, 100, 2, 4, 2020},
	{"35", "Los Angeles", 460000, 65, 6, 3, 2017},
	{"36", "Miami", 620000, 98, 9, 5, 2022},
	{"37", "Seattle", 480000, 80, 11, 3, 2019},
	{"38", "Denver", 410000, 78, 4, 2, 2015},
	{"39", "Boston", 640000, 95, 10, 3, 2021},
	{"40", "New York", 650000, 100, 7, 4, 2023},
	{"41", "Houston", 480000, 88, 9, 3, 2020},
	{"42", "Philadelphia", 410000, 72, 8, 2, 2021},
	{"43", "Chicago", 500000, 90, 5, 3, 2018},
	{"44", "Los Angeles", 700000, 75, 4, 5, 2022},
	{"45", "Seattle", 530000, 85, 10, 3, 2020},
	{"46", "Miami", 460000, 78, 6, 3, 2019},
	{"47", "Denver", 430000, 82, 7, 3, 2015},
	{"48", "San Francisco", 800000, 105, 9, 5, 2021},
	{"49", "Washington D.C.", 750000, 95, 3, 4, 2022},
	{"50", "Portland", 500000, 85, 8, 3, 2019},
	{"51", "Austin", 570000, 92, 10, 4, 2021},
	{"52", "Boston", 690000, 100, 2, 5, 2023},
	{"53", "San Diego", 600000, 90, 5, 4, 2020},
	{"54", "Houston", 430000, 70, 3, 2, 2017},
	{"55", "Dallas", 460000, 85, 9, 3, 2021},
	{"56", "Seattle", 520000, 76, 7, 3, 2022},
	{"57", "Miami", 510000, 80, 6, 3, 2019},
	{"58", "Los Angeles", 590000, 88, 11, 3, 2020},
	{"59", "Phoenix", 480000, 84, 8, 3, 2021},
	{"60", "Columbus", 430000, 73, 9, 2, 2016},
	{"61", "Salt Lake City", 500000, 85, 7, 3, 2020},
	{"62", "Chicago", 430000, 90, 5, 3, 2018},
	{"63", "Washington D.C.", 620000, 100, 3, 4, 2022},
	{"64", "Philadelphia", 460000, 75, 6, 3, 2021},
	{"65", "Denver", 470000, 80, 10, 2, 2019},
	{"66", "New York", 710000, 95, 9, 4, 2021},
	{"67", "Dallas", 490000, 85, 5, 3, 2020},
	{"68", "Houston", 460000, 90, 4, 3, 2017},
	{"69", "Phoenix", 470000, 75, 8, 3, 2022},
	{"70", "Los Angeles", 630000, 92, 7, 4, 2023},
	{"71", "Portland", 550000, 80, 11, 3, 2019},
	{"72", "San Francisco", 790000, 105, 3, 4, 2022},
	{"73", "Miami", 460000, 82, 6, 3, 2020},
	{"74", "Columbus", 390000, 75, 7, 2, 2018},
	{"75", "Atlanta", 470000, 80, 8, 3, 2021},
	{"76", "Boston", 690000, 95, 9, 4, 2020},
	{"77", "Dallas", 600000, 100, 6, 4, 2022},
	{"78", "Houston", 490000, 85, 5, 3, 2021},
	{"79", "San Diego", 710000, 100, 6, 5, 2022},
	{"80", "Washington D.C.", 630000, 90, 3, 4, 2021},
	{"81", "Houston", 470000, 75, 5, 3, 2016},
	{"82", "Seattle", 550000, 85, 8, 3, 2020},
	{"83", "Austin", 500000, 90, 10, 3, 2022},
	{"84", "Phoenix", 420000, 70, 7, 2, 2019},
	{"85", "Miami", 530000, 85, 4, 3, 2021},
	{"86", "Denver", 490000, 80, 6, 3, 2020},
	{"87", "San Francisco", 760000, 110, 9, 5, 2023},
	{"88", "Los Angeles", 630000, 95, 7, 4, 2022},
	{"89", "Portland", 560000, 88, 5, 3, 2021},
	{"90", "Chicago", 450000, 90, 4, 2, 2019},
	{"91", "Houston", 500000, 95, 6, 3, 2021},
	{"92", "San Diego", 680000, 95, 8, 4, 2022},
	{"93", "Dallas", 550000, 90, 5, 3, 2021},
	{"94", "Boston", 600000, 95, 7, 4, 2022},
	{"95", "Seattle", 520000, 85, 6, 3, 2019},
	{"96", "Columbus", 470000, 78, 8, 3, 2020},
	{"97", "Portland", 520000, 75, 7, 3, 2021},
	{"98", "Los Angeles", 650000, 95, 9, 4, 2023},
	{"99", "Miami", 550000, 88, 5, 3, 2020},
	{"100", "Washington D.C.", 620000, 90, 11, 4, 2021},
}

func main() {
	http.HandleFunc("/apartments", getApartmentsByArea)

	port := ":9080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Функция для поиска квартир по площади
func getApartmentsByArea(w http.ResponseWriter, r *http.Request) {
	minArea := r.URL.Query().Get("min_area")
	maxArea := r.URL.Query().Get("max_area")

	if minArea == "" || maxArea == "" {
		http.Error(w, "Both min_area and max_area query parameters are required", http.StatusBadRequest)
		return
	}

	// Преобразуем строковые значения в float64
	minAreaVal, err := strconv.ParseFloat(minArea, 64)
	if err != nil {
		http.Error(w, "Invalid min_area value", http.StatusBadRequest)
		return
	}

	maxAreaVal, err := strconv.ParseFloat(maxArea, 64)
	if err != nil {
		http.Error(w, "Invalid max_area value", http.StatusBadRequest)
		return
	}

	// Ищем квартиры в указанном интервале площади
	var result []Apartment
	for _, apartment := range apartments {
		if apartment.Area >= minAreaVal && apartment.Area <= maxAreaVal {
			result = append(result, apartment)
		}
	}

	// Если квартиры не найдены
	if len(result) == 0 {
		http.Error(w, "No apartments found within the specified area range", http.StatusNotFound)
		return
	}

	// Отправляем результат в виде JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

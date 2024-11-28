package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Citizen represents a person.
// swagger:model
type Citizen struct {
	// ID is the unique identifier for the citizen.
	//
	// required: true
	ID string `json:"id"`

	// Name is the name of the citizen.
	//
	// required: true
	Name string `json:"name"`

	// Age is the age of the citizen.
	//
	// required: true
	Age int `json:"age"`

	// Address is the address of the citizen.
	//
	// required: true
	Address string `json:"address"`
}

// swagger:parameters searchCitizens
type searchRequest struct {
	// in:query
	Name string `json:"name"`
}

// swagger:response
type searchResponse struct {
	// Citizen array
	// in:body
	Citizens []Citizen `json:"citizens"`
}

var citizens = []Citizen{
	{ID: "1", Name: "John Smith", Age: 35, Address: "123 Main St"},
	{ID: "2", Name: "Jane Johnson", Age: 28, Address: "456 Elm St"},
	{ID: "3", Name: "Michael Williams", Age: 42, Address: "789 Oak Ave"},
	{ID: "4", Name: "Emily Brown", Age: 30, Address: "101 Pine Rd"},
	{ID: "5", Name: "David Jones", Age: 50, Address: "202 Cedar Ln"},
	{ID: "6", Name: "Olivia Garcia", Age: 47, Address: "303 Birch Blvd"},
	{ID: "7", Name: "James Miller", Age: 62, Address: "404 Maple Dr"},
	{ID: "8", Name: "Sophia Davis", Age: 29, Address: "505 Walnut St"},
	{ID: "9", Name: "Robert Rodriguez", Age: 38, Address: "606 Spruce Ave"},
	{ID: "10", Name: "Emma Martinez", Age: 55, Address: "707 Redwood Rd"},
	{ID: "11", Name: "William Hernandez", Age: 68, Address: "808 Oak Ln"},
	{ID: "12", Name: "Isabella Lopez", Age: 32, Address: "909 Pine Blvd"},
	{ID: "13", Name: "Joseph Gonzalez", Age: 45, Address: "1010 Cedar Dr"},
	{ID: "14", Name: "Mia Wilson", Age: 70, Address: "1111 Birch St"},
	{ID: "15", Name: "Richard Anderson", Age: 26, Address: "1212 Maple Ave"},
	{ID: "16", Name: "Charlotte Thomas", Age: 58, Address: "1313 Walnut Rd"},
	{ID: "17", Name: "Thomas Taylor", Age: 39, Address: "1414 Spruce Ln"},
	{ID: "18", Name: "Harper Jackson", Age: 41, Address: "1515 Redwood Blvd"},
	{ID: "19", Name: "Michael Martin", Age: 23, Address: "1616 Oak St"},
	{ID: "20", Name: "Emily Harris", Age: 31, Address: "1717 Pine Ave"},
	{ID: "21", Name: "William Clark", Age: 43, Address: "1818 Cedar Rd"},
	{ID: "22", Name: "Olivia Lewis", Age: 27, Address: "1919 Birch Ln"},
	{ID: "23", Name: "James Young", Age: 51, Address: "2020 Maple St"},
	{ID: "24", Name: "Sophia Hall", Age: 36, Address: "2121 Walnut Ave"},
	{ID: "25", Name: "Robert Allen", Age: 65, Address: "2222 Spruce Rd"},
	{ID: "26", Name: "Emma Scott", Age: 29, Address: "2323 Redwood Ln"},
	{ID: "27", Name: "David Green", Age: 48, Address: "2424 Oak Blvd"},
	{ID: "28", Name: "Joseph Baker", Age: 33, Address: "2525 Pine St"},
	{ID: "29", Name: "Mia Adams", Age: 53, Address: "2626 Cedar Ave"},
	{ID: "30", Name: "Richard Nelson", Age: 61, Address: "2727 Birch Rd"},
	{ID: "31", Name: "Charlotte Carter", Age: 25, Address: "2828 Maple Ln"},
	{ID: "32", Name: "Thomas Mitchell", Age: 44, Address: "2929 Walnut St"},
	{ID: "33", Name: "Harper Perez", Age: 37, Address: "3030 Spruce Ave"},
	{ID: "34", Name: "Michael Roberts", Age: 56, Address: "3131 Redwood Rd"},
	{ID: "35", Name: "Emily Turner", Age: 22, Address: "3232 Oak St"},
	{ID: "36", Name: "William Phillips", Age: 63, Address: "3333 Pine Ave"},
	{ID: "37", Name: "Olivia Campbell", Age: 40, Address: "3434 Cedar Rd"},
	{ID: "38", Name: "James Parker", Age: 59, Address: "3535 Birch Ln"},
	{ID: "39", Name: "Sophia Evans", Age: 34, Address: "3636 Maple St"},
	{ID: "40", Name: "Robert Edwards", Age: 28, Address: "3737 Walnut Ave"},
	{ID: "41", Name: "Emma Collins", Age: 46, Address: "3838 Spruce Rd"},
	{ID: "42", Name: "David Stewart", Age: 60, Address: "3939 Redwood Ln"},
	{ID: "43", Name: "Joseph Sanchez", Age: 31, Address: "4040 Oak Blvd"},
	{ID: "44", Name: "Mia Morris", Age: 52, Address: "4141 Pine St"},
	{ID: "45", Name: "Richard Rogers", Age: 49, Address: "4242 Cedar Ave"},
	{ID: "46", Name: "Charlotte Reed", Age: 64, Address: "4343 Birch Rd"},
	{ID: "47", Name: "Thomas Cook", Age: 27, Address: "4444 Maple Ln"},
	{ID: "48", Name: "Harper Morgan", Age: 38, Address: "4545 Walnut St"},
	{ID: "49", Name: "Michael Bell", Age: 57, Address: "4646 Spruce Ave"},
	{ID: "50", Name: "Emily Murphy", Age: 41, Address: "4747 Redwood Rd"},
	{ID: "51", Name: "William Bailey", Age: 62, Address: "4848 Oak St"},
	{ID: "52", Name: "Olivia Rivera", Age: 29, Address: "4949 Pine Ave"},
	{ID: "53", Name: "James Cooper", Age: 50, Address: "5050 Cedar Rd"},
	{ID: "54", Name: "Sophia Richardson", Age: 33, Address: "5151 Birch Ln"},
	{ID: "55", Name: "Robert Cox", Age: 65, Address: "5252 Maple St"},
	{ID: "56", Name: "Emma Howard", Age: 42, Address: "5353 Walnut Ave"},
	{ID: "57", Name: "David Ward", Age: 54, Address: "5454 Spruce Rd"},
	{ID: "58", Name: "Joseph Torres", Age: 26, Address: "5555 Redwood Ln"},
	{ID: "59", Name: "Mia Peterson", Age: 39, Address: "5656 Oak Blvd"},
	{ID: "60", Name: "Richard Gray", Age: 68, Address: "5757 Pine St"},
	{ID: "61", Name: "Charlotte Ramirez", Age: 23, Address: "5858 Cedar Ave"},
	{ID: "62", Name: "Thomas James", Age: 47, Address: "5959 Birch Rd"},
	{ID: "63", Name: "Harper Watson", Age: 51, Address: "6060 Maple Ln"},
	{ID: "64", Name: "Michael Brooks", Age: 32, Address: "6161 Walnut St"},
	{ID: "65", Name: "Emily Kelly", Age: 63, Address: "6262 Spruce Ave"},
	{ID: "66", Name: "William Sanders", Age: 44, Address: "6363 Redwood Rd"},
	{ID: "67", Name: "Olivia Price", Age: 56, Address: "6464 Oak St"},
	{ID: "68", Name: "James Bennett", Age: 28, Address: "6565 Pine Ave"},
	{ID: "69", Name: "Sophia Woods", Age: 40, Address: "6666 Cedar Rd"},
	{ID: "70", Name: "Robert Brown", Age: 69, Address: "6767 Birch Ln"},
	{ID: "71", Name: "Emma Jenkins", Age: 31, Address: "6868 Maple St"},
	{ID: "72", Name: "David Perry", Age: 53, Address: "6969 Walnut Ave"},
	{ID: "73", Name: "Joseph Powell", Age: 25, Address: "7070 Spruce Rd"},
	{ID: "74", Name: "Mia Long", Age: 46, Address: "7171 Redwood Ln"},
	{ID: "75", Name: "Richard Patterson", Age: 60, Address: "7272 Oak Blvd"},
	{ID: "76", Name: "Charlotte Hughes", Age: 34, Address: "7373 Pine St"},
	{ID: "77", Name: "Thomas Flores", Age: 55, Address: "7474 Cedar Ave"},
	{ID: "78", Name: "Harper Washington", Age: 27, Address: "7575 Birch Rd"},
	{ID: "79", Name: "Michael Butler", Age: 49, Address: "7676 Maple Ln"},
	{ID: "80", Name: "Emily Simmons", Age: 38, Address: "7777 Walnut St"},
	{ID: "81", Name: "William Foster", Age: 61, Address: "7878 Spruce Ave"},
	{ID: "82", Name: "Olivia Gonzales", Age: 33, Address: "7979 Redwood Rd"},
	{ID: "83", Name: "James Norton", Age: 54, Address: "8080 Oak St"},
	{ID: "84", Name: "Sophia Kelly", Age: 26, Address: "8181 Pine Ave"},
	{ID: "85", Name: "Robert Mills", Age: 47, Address: "8282 Cedar Rd"},
	{ID: "86", Name: "Emma Reid", Age: 68, Address: "8383 Birch Ln"},
	{ID: "87", Name: "David Morris", Age: 30, Address: "8484 Maple St"},
	{ID: "88", Name: "Joseph Phillips", Age: 51, Address: "8585 Walnut Ave"},
	{ID: "89", Name: "Mia Gregory", Age: 42, Address: "8686 Spruce Rd"},
	{ID: "90", Name: "Richard Russell", Age: 63, Address: "8787 Redwood Ln"},
	{ID: "91", Name: "Charlotte Foster", Age: 25, Address: "8888 Oak Blvd"},
	{ID: "92", Name: "Thomas Simmons", Age: 46, Address: "8989 Pine St"},
	{ID: "93", Name: "Harper James", Age: 37, Address: "9090 Cedar Ave"},
	{ID: "94", Name: "Michael Kelly", Age: 58, Address: "9191 Birch Rd"},
	{ID: "95", Name: "Emily Morris", Age: 29, Address: "9292 Maple Ln"},
	{ID: "96", Name: "William Phillips", Age: 50, Address: "9393 Walnut St"},
	{ID: "97", Name: "Olivia Kelly", Age: 64, Address: "9494 Spruce Ave"},
	{ID: "98", Name: "James Morris", Age: 34, Address: "9595 Redwood Rd"},
	{ID: "99", Name: "Sophia Phillips", Age: 55, Address: "9696 Oak St"},
	{ID: "100", Name: "Robert Kelly", Age: 27, Address: "9797 Pine Ave"},
}

// searchCitizens searches for citizens whose names contain a given string.
// The function expects a GET request to the /search endpoint with a name query parameter.
// The function returns a JSON array of matching citizens.
// swagger:route GET /search citizens searchCitizens
//
// Search for citizens by part of name.
//
//	Produces:
//	- application/json
//
//	Schemes: http
//
//	Responses:
//	  200: searchResponse
func searchCitizens(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")

	var results []Citizen
	for _, citizen := range citizens {
		if strings.Contains(strings.ToLower(citizen.Name), strings.ToLower(query)) {
			results = append(results, citizen)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/search", searchCitizens)

	http.ListenAndServe(":9080", nil)
}

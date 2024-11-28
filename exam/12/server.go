package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Author      string `json:"author"`
}

var booksList = []Book{
	{ID: "1", Title: "1984", Year: 1949, Author: "George Orwell"},
	{ID: "2", Title: "To Kill a Mockingbird", Year: 1960, Author: "Harper Lee"},
	{ID: "3", Title: "The Great Gatsby", Year: 1925, Author: "F. Scott Fitzgerald"},
	{ID: "4", Title: "The Catcher in the Rye", Year: 1951, Author: "J.D. Salinger"},
	{ID: "5", Title: "Animal Farm", Year: 1945, Author: "George Orwell"},
	{ID: "6", Title: "Moby Dick", Year: 1851, Author: "Herman Melville"},
	{ID: "7", Title: "Pride and Prejudice", Year: 1813, Author: "Jane Austen"},
	{ID: "8", Title: "Brave New World", Year: 1932, Author: "Aldous Huxley"},
	{ID: "9", Title: "War and Peace", Year: 1869, Author: "Leo Tolstoy"},
	{ID: "10", Title: "The Odyssey", Year: -700, Author: "Homer"},
	{ID: "11", Title: "Crime and Punishment", Year: 1866, Author: "Fyodor Dostoevsky"},
	{ID: "12", Title: "The Brothers Karamazov", Year: 1880, Author: "Fyodor Dostoevsky"},
	{ID: "13", Title: "Great Expectations", Year: 1861, Author: "Charles Dickens"},
	{ID: "14", Title: "Jane Eyre", Year: 1847, Author: "Charlotte Bronte"},
	{ID: "15", Title: "Wuthering Heights", Year: 1847, Author: "Emily Bronte"},
	{ID: "16", Title: "The Hobbit", Year: 1937, Author: "J.R.R. Tolkien"},
	{ID: "17", Title: "The Lord of the Rings", Year: 1954, Author: "J.R.R. Tolkien"},
	{ID: "18", Title: "The Silmarillion", Year: 1977, Author: "J.R.R. Tolkien"},
	{ID: "19", Title: "Ulysses", Year: 1922, Author: "James Joyce"},
	{ID: "20", Title: "The Divine Comedy", Year: 1320, Author: "Dante Alighieri"},
	{ID: "21", Title: "Les Misérables", Year: 1862, Author: "Victor Hugo"},
	{ID: "22", Title: "Don Quixote", Year: 1615, Author: "Miguel de Cervantes"},
	{ID: "23", Title: "A Tale of Two Cities", Year: 1859, Author: "Charles Dickens"},
	{ID: "24", Title: "The Scarlet Letter", Year: 1850, Author: "Nathaniel Hawthorne"},
	{ID: "25", Title: "Of Mice and Men", Year: 1937, Author: "John Steinbeck"},
	{ID: "26", Title: "The Grapes of Wrath", Year: 1939, Author: "John Steinbeck"},
	{ID: "27", Title: "East of Eden", Year: 1952, Author: "John Steinbeck"},
	{ID: "28", Title: "Slaughterhouse-Five", Year: 1969, Author: "Kurt Vonnegut"},
	{ID: "29", Title: "Catch-22", Year: 1961, Author: "Joseph Heller"},
	{ID: "30", Title: "Fahrenheit 451", Year: 1953, Author: "Ray Bradbury"},
	{ID: "31", Title: "The Stranger", Year: 1942, Author: "Albert Camus"},
	{ID: "32", Title: "The Plague", Year: 1947, Author: "Albert Camus"},
	{ID: "33", Title: "The Metamorphosis", Year: 1915, Author: "Franz Kafka"},
	{ID: "34", Title: "The Trial", Year: 1925, Author: "Franz Kafka"},
	{ID: "35", Title: "One Hundred Years of Solitude", Year: 1967, Author: "Gabriel García Márquez"},
	{ID: "36", Title: "Love in the Time of Cholera", Year: 1985, Author: "Gabriel García Márquez"},
	{ID: "37", Title: "The Old Man and the Sea", Year: 1952, Author: "Ernest Hemingway"},
	{ID: "38", Title: "A Farewell to Arms", Year: 1929, Author: "Ernest Hemingway"},
	{ID: "39", Title: "For Whom the Bell Tolls", Year: 1940, Author: "Ernest Hemingway"},
	{ID: "40", Title: "The Sun Also Rises", Year: 1926, Author: "Ernest Hemingway"},
	{ID: "41", Title: "Lolita", Year: 1955, Author: "Vladimir Nabokov"},
	{ID: "42", Title: "Pale Fire", Year: 1962, Author: "Vladimir Nabokov"},
	{ID: "43", Title: "The Sound and the Fury", Year: 1929, Author: "William Faulkner"},
	{ID: "44", Title: "As I Lay Dying", Year: 1930, Author: "William Faulkner"},
	{ID: "45", Title: "Absalom, Absalom!", Year: 1936, Author: "William Faulkner"},
	{ID: "46", Title: "Beloved", Year: 1987, Author: "Toni Morrison"},
	{ID: "47", Title: "Song of Solomon", Year: 1977, Author: "Toni Morrison"},
	{ID: "48", Title: "Invisible Man", Year: 1952, Author: "Ralph Ellison"},
	{ID: "49", Title: "Native Son", Year: 1940, Author: "Richard Wright"},
	{ID: "50", Title: "Gone with the Wind", Year: 1936, Author: "Margaret Mitchell"},
	{ID: "51", Title: "The Bell Jar", Year: 1963, Author: "Sylvia Plath"},
	{ID: "52", Title: "A Clockwork Orange", Year: 1962, Author: "Anthony Burgess"},
	{ID: "53", Title: "The Road", Year: 2006, Author: "Cormac McCarthy"},
	{ID: "54", Title: "No Country for Old Men", Year: 2005, Author: "Cormac McCarthy"},
	{ID: "55", Title: "The Grapes of Wrath", Year: 1939, Author: "John Steinbeck"},
	{ID: "56", Title: "Dune", Year: 1965, Author: "Frank Herbert"},
	{ID: "57", Title: "Foundation", Year: 1951, Author: "Isaac Asimov"},
	{ID: "58", Title: "Neuromancer", Year: 1984, Author: "William Gibson"},
	{ID: "59", Title: "Snow Crash", Year: 1992, Author: "Neal Stephenson"},
	{ID: "60", Title: "Hyperion", Year: 1989, Author: "Dan Simmons"},
	{ID: "61", Title: "The Left Hand of Darkness", Year: 1969, Author: "Ursula K. Le Guin"},
	{ID: "62", Title: "The Dispossessed", Year: 1974, Author: "Ursula K. Le Guin"},
	{ID: "63", Title: "The Handmaid's Tale", Year: 1985, Author: "Margaret Atwood"},
	{ID: "64", Title: "Oryx and Crake", Year: 2003, Author: "Margaret Atwood"},
	{ID: "65", Title: "Cloud Atlas", Year: 2004, Author: "David Mitchell"},
	{ID: "66", Title: "A Brief History of Seven Killings", Year: 2014, Author: "Marlon James"},
	{ID: "67", Title: "White Teeth", Year: 2000, Author: "Zadie Smith"},
	{ID: "68", Title: "On Beauty", Year: 2005, Author: "Zadie Smith"},
	{ID: "69", Title: "Norwegian Wood", Year: 1987, Author: "Haruki Murakami"},
	{ID: "70", Title: "Kafka on the Shore", Year: 2002, Author: "Haruki Murakami"},
	{ID: "71", Title: "1Q84", Year: 2009, Author: "Haruki Murakami"},
	{ID: "72", Title: "The Wind-Up Bird Chronicle", Year: 1995, Author: "Haruki Murakami"},
	{ID: "73", Title: "The Alchemist", Year: 1988, Author: "Paulo Coelho"},
	{ID: "74", Title: "Veronika Decides to Die", Year: 1998, Author: "Paulo Coelho"},
	{ID: "75", Title: "The Shadow of the Wind", Year: 2001, Author: "Carlos Ruiz Zafón"},
	{ID: "76", Title: "The Angel's Game", Year: 2008, Author: "Carlos Ruiz Zafón"},
	{ID: "77", Title: "The Night Circus", Year: 2011, Author: "Erin Morgenstern"},
	{ID: "78", Title: "The Starless Sea", Year: 2019, Author: "Erin Morgenstern"},
	{ID: "79", Title: "The Hunger Games", Year: 2008, Author: "Suzanne Collins"},
	{ID: "80", Title: "Catching Fire", Year: 2009, Author: "Suzanne Collins"},
	{ID: "81", Title: "Mockingjay", Year: 2010, Author: "Suzanne Collins"},
	{ID: "82", Title: "Divergent", Year: 2011, Author: "Veronica Roth"},
	{ID: "83", Title: "Insurgent", Year: 2012, Author: "Veronica Roth"},
	{ID: "84", Title: "Allegiant", Year: 2013, Author: "Veronica Roth"},
	{ID: "85", Title: "The Maze Runner", Year: 2009, Author: "James Dashner"},
	{ID: "86", Title: "The Scorch Trials", Year: 2010, Author: "James Dashner"},
	{ID: "87", Title: "The Death Cure", Year: 2011, Author: "James Dashner"},
	{ID: "88", Title: "The Giver", Year: 1993, Author: "Lois Lowry"},
	{ID: "89", Title: "Gathering Blue", Year: 2000, Author: "Lois Lowry"},
	{ID: "90", Title: "Messenger", Year: 2004, Author: "Lois Lowry"},
	{ID: "91", Title: "Son", Year: 2012, Author: "Lois Lowry"},
	{ID: "92", Title: "Life of Pi", Year: 2001, Author: "Yann Martel"},
	{ID: "93", Title: "The Book Thief", Year: 2005, Author: "Markus Zusak"},
	{ID: "94", Title: "The Fault in Our Stars", Year: 2012, Author: "John Green"},
	{ID: "95", Title: "Looking for Alaska", Year: 2005, Author: "John Green"},
	{ID: "96", Title: "Paper Towns", Year: 2008, Author: "John Green"},
	{ID: "97", Title: "Turtles All the Way Down", Year: 2017, Author: "John Green"},
	{ID: "98", Title: "The Perks of Being a Wallflower", Year: 1999, Author: "Stephen Chbosky"},
	{ID: "99", Title: "Speak", Year: 1999, Author: "Laurie Halse Anderson"},
	{ID: "100", Title: "Thirteen Reasons Why", Year: 2007, Author: "Jay Asher"},
}

func searchBooksByDescription(w http.ResponseWriter, r *http.Request) {
	descriptionQuery := r.URL.Query().Get("description")
	if descriptionQuery == "" {
		http.Error(w, "Parameter 'description' is required", http.StatusBadRequest)
		return
	}

	matchingBooks := []Book{}
	for _, book := range matchingBooks {
		if strings.Contains(strings.ToLower(book.Description), strings.ToLower(descriptionQuery)) {
			matchingBooks = append(matchingBooks, book)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if len(matchingBooks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "No books found matching the description"})
		return
	}

	json.NewEncoder(w).Encode(matchingBooks)
}

func main() {
	http.HandleFunc("/books/search", searchBooksByDescription)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

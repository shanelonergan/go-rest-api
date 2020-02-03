// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type event struct {
// 	ID          string `json:"ID"`
// 	Title       string `json:"Title"`
// 	Description string `json:"Description"`
// }

// type allEvents []event

// var events = allEvents{
// 	{
// 		ID:          "1",
// 		Title:       "Introduction to Golang",
// 		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
// 	},
// }

// func homeLink(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome home!")
// }

// func createEvent(w http.ResponseWriter, r *http.Request) {
// 	var newEvent event
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}

// 	json.Unmarshal(reqBody, &newEvent)
// 	events = append(events, newEvent)
// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(newEvent)
// }

// func getOneEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]

// 	for _, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			json.NewEncoder(w).Encode(singleEvent)
// 		}
// 	}
// }

// func getAllEvents(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(events)
// }

// func updateEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]
// 	var updatedEvent event

// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}
// 	json.Unmarshal(reqBody, &updatedEvent)

// 	for i, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			singleEvent.Title = updatedEvent.Title
// 			singleEvent.Description = updatedEvent.Description
// 			events = append(events[:i], singleEvent)
// 			json.NewEncoder(w).Encode(singleEvent)
// 		}
// 	}
// }

// func deleteEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]

// 	for i, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			events = append(events[:i], events[i+1:]...)
// 			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
// 		}
// 	}
// }

// func main() {
// 	initEvents()
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", homeLink)
// 	router.HandleFunc("/event", createEvent).Methods("POST")
// 	router.HandleFunc("/events", getAllEvents).Methods("GET")
// 	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
// 	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
// 	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model), similar to an ES6 Class
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct (slice is variable length array)
var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get one book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //get any params

	// loops through books and find matching id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID, not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(book)
}

func main() {
	// Init router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "923874", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "923204", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// route handlers / endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

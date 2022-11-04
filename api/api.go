package api

import (
	"github.com/Ahasannn/book-server-api/auth"
	"github.com/Ahasannn/book-server-api/data"
	"github.com/Ahasannn/book-server-api/model"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"github.com/gorilla/mux"
)

var mutex sync.Mutex

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Books)

}

// Get Single Books
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) //Gets params
	//Loop through books and find one with the id from the params

	for _, item := range data.Books {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Book{})
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	data.Books = append(data.Books, book)
	json.NewEncoder(w).Encode(book)
}

// Update  Book
func updateBook(w http.ResponseWriter, r *http.Request) {
   
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range data.Books {
		if item.ID == params["id"] {
			data.Books = append(data.Books[:index], data.Books[index+1:]...)
			var book model.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			data.Books = append(data.Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range data.Books {
		if item.ID == param["id"] {
			data.Books = append(data.Books[:index], data.Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data.Books)
}

func LogIn(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	log.Println("LogIn")
	log.Println("Authentication successful!")
	log.Println("Successfully logged in!")

	token, err := auth.GetToken()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, _ = response.Write([]byte("Error generating JWT token: " + err.Error()))
	} else {
		response.Header().Set("Authorization", "Bearer "+token)
		response.WriteHeader(http.StatusOK)
		_, _ = response.Write([]byte("Token: " + token))
	}
}

func HandleRoutes(port string) {
	log.Println("in HandleRoutes!")

	//Init Router
	r := mux.NewRouter().StrictSlash(true)

	//Route Handlers/ Endpoints
	r.HandleFunc("/api/login", auth.BasicAuthentication(LogIn)).Methods("POST")
	r.HandleFunc("/api/getBooks", auth.JWTAuthentication(getBooks)).Methods("GET")
	r.HandleFunc("/api/getBook/{id}", auth.JWTAuthentication(getBook)).Methods("GET")
	r.HandleFunc("/api/createBook", auth.JWTAuthentication(createBook)).Methods("POST")
	r.HandleFunc("/api/updateBooks/{id}", auth.JWTAuthentication(updateBook)).Methods("PUT")
	r.HandleFunc("/api/deleteBooks/{id}", auth.JWTAuthentication(deleteBook)).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, r))

}

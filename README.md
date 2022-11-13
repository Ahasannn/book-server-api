# Book Library RESTful API Server 

It is a RESTful API using [Go](https://github.com/golang), [gorilla/mux](https://github.com/gorilla/mux), [Basic Authentication](https://learningprogramming.net/golang/golang-restful-web-api/basic-authentication-in-golang-restful-web-api/), [JWT Authentication](https://github.com/dgrijalva/jwt-go).

<hr/>

## Brief Description

<p>I have built a fully-fledged REST API with Go that exposes GET, POST, DELETE and PUT endpoints which allows to perform the full range of CRUD operations. A handler function accepts http response and request in json format. Then, the request is decoded and written to response according to the called function. This handler function is wrapped by the authentication middleware to perform the security check.</p>

<hr/>

## API Endpoints 

| URL      | Function    | Method | Description | Authentication Type
| -----------   | ----------- | ------ | ----------- |---------- 
| https://localhost:8000/api/login    | Login      |     POST   |  Return JWT token in response for successful authentication | Basic
| https://localhost:8000/api/getBooks | getBooks   |   GET     | Returns the details of all the books | JWT
| https://localhost:8000/api/getBook/{id} | getBook | GET  |  Returns the details of the book with the valid requested book id | JWT
| https://localhost:8000/api/createBook | createBook | POST | Creates a new book | JWT
| https://localhost:8000/api/updateBooks/{id} | updateBooks | PUT | Updates the details of the requested book id | JWT
| https://localhost:8000/api/deleteBooks/{id} | deleteBooks | DELETE | Deletes the book specified by id | JWT

<hr/>

## Authentication Method

- Basic Authentication
- JWT Authentication

<hr/>

## Data Models

    type Book struct {
        ID     string  `json:"id"`
        Isbn   string  `json:"isbn"`
        Title  string  `json:"title"`
        Author *Author `json:"author"`
    }

    type Author struct {
        Firstname string `json:"firstname"`
        Lastname  string `json:"lastname"`
    }
    
<hr/>

## Installation 
* go install github.com/Ahasannn/book-library@latest

<hr/>

Set Environment variables for Basic Authentication.

    export username=Ahasan 
    export password=ak4747

<hr/>

Testing the API endpoints

* Primary api endpoints can be tested with [Postman](https://www.postman.com/)

<hr/>

Server Run 

    go build -o bin/book-library .
    ./bin/book-library

<hr/>





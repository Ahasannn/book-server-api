# GO RESTful API server

RESTful API using [go](https://github.com/golang), [gorilla mux](https://github.com/gorilla/mux), Basic Auth, [JWT Auth](https://github.com/dgrijalva/jwt-go)

<hr/>

## API Endpoints 

| Endpoint      | Function    | Method | Status Code | Auth
| -----------   | ----------- | ------ | ----------- |---------- 
| /api/login    | Login      |     POST   |  Success - StatusOK, Failure - StatusUnauthorized | Basic
| /api/getBooks | getBooks   |   GET     |   Success - StatusOK      | JWT
| /api/getBook/{id} | getBook | GET  |  Success - StatusOK, Failure - StatusNoContent | JWT
| /api/createBook | createBook | POST | Success - StatusCreated, Failure - StatusConflict | JWT
| /api/updateBooks/{id} | updateBooks | PUT | Success - StatusCreated, Failure - StatusNoContent | JWT
| /api/deleteBooks/{id} | deleteBooks | DELETE | Success - StatusOK, Failure - StatusNoContent | JWT

<hr/>

## Authentication Method

- Basic Authentication
- JWT Authentication

<hr/>

## Installation 
Run the server using the commands-

    git clone http://github.com/Ahasannn/GO_RESTful_API.git
    cd GO_RESTful_API
    go build && ./GO_RESTful_API
    
<hr/>

## Data model

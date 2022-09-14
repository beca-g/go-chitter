package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

var database = Database{}
var peep = newPeep("Hello, World!", "14/09/2022", "Chris", uuid.New().String())

func main() {

	database = append(database, peep)
	router := chi.NewRouter()

	router.Get("/all", database.get)

	http.ListenAndServe(":8080", router)
}

/**
- A peep struct/type of some sorts
- Create app as a RESTFUL API, so we will need chi library
- Store data in memory, no db.

*/

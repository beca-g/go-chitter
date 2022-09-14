package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

var database = []Peep{}
var peep = newPeep("Hello, World!", "14/09/2022", "Chris")

func main() {

	database = append(database, peep)
	router := chi.NewRouter()

	router.Get("/all", GetAll)

	http.ListenAndServe(":8080", router)
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	for _, p := range database {
		w.Write([]byte(p.get()))
	}
}

/**
- A peep struct/type of some sorts
- Create app as a RESTFUL API, so we will need chi library
- Store data in memory, no db.

*/

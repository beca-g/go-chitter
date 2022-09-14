package main

import (
	"net/http"

	"github.com/google/uuid"
)

type Database []Peep

func (d Database) get(w http.ResponseWriter, r *http.Request) {

	for _, p := range database {
		w.Write([]byte(p.get()))
	}
}

func (d *Database) createPost(message string, date string, user string) {
	post := newPeep(message, date, user, uuid.New().String())

	*d = append(*d, post)
}

func (d Database) getPostFromUuid(uuid string) Peep {
	
	for _, p := range d {
		if p.checkUuid(uuid) {
			return p
		}
	}; return Peep{}
}

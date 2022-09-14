package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestPeepMessageIsString(t *testing.T) {

	peep := newPeep("Our message", "14/09/2022", "Chris", uuid.New().String())
	T := reflect.TypeOf("string")

	var s interface{} = peep.message
	if _, ok := s.(string); !ok {
		t.Errorf("Expected a peeps message to be of type string, instead got: %v", T)
	}
}

func TestCreatePost(t *testing.T) {

	database := Database{}
	database.createPost("Our message", "14/09/2022", "Beca")
	fmt.Println(len(database))
	message := database[0].message

	if message != "Our message" {
		t.Errorf("Expected message to equal 'Our message', instead got: %v", message)
	}
}

func TestgetPostFromUuid(t *testing.T) {
	database := Database{}
	database.createPost("New peeps with uuid", "14/09/2022", "Beca")
	uuid := database[0].uuid

	peep2 := database.getPostFromUuid(uuid)

	if peep.uuid != peep2.uuid {
		t.Errorf("Expected uuid to equal %v but got %v", uuid, peep2.uuid)
	}
}

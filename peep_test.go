package main

import (
	"reflect"
	"testing"
)

func TestPeepMessageIsString(t *testing.T) {

	peep := newPeep("Our message", "14/09/2022", "Chris")
	T := reflect.TypeOf("string")

	var s interface{} = peep.message
	if _, ok := s.(string); !ok {
		t.Errorf("Expected a peeps message to be of type string, instead got: %v", T)
	}
}

package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestFromJSON(t *testing.T) {
	var jsonTests = []struct {
		x        string      // input
		expected interface{} // expected result
	}{
		{`{}`, struct{}{}},
		{`{"name": "John", "age": 30}`, Person{Name: "John", Age: 30}},
	}

	for _, tt := range jsonTests {
		var actual struct{}
		var err = FromJSON(tt.x, &actual)

		if err != nil {
			t.Errorf("FromJSON(%s): expected no error, actual %v", tt.x, err)
		}

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("FromJSON(%s): expected %v, actual %v", tt.x, tt.expected, actual)
		}
	}
}

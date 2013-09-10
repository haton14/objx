package objx

import (
	"encoding/json"
)

// New creates a new O containing the specified object.
func New(object interface{}) *O {
	return &O{Obj: object}
}

// MustJson creates a new O containing the data specified in the
// jsonString.
func MustJson(jsonString string) *O {
	o, err := Json(jsonString)

	if err != NoError {
		panic("objx: " + err.Error())
	}

	return o
}

// Json creates a new O containing the data specified in the
// jsonString.
func Json(jsonString string) (*O, error) {

	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)

	if err != NoError {
		return Nil, err
	}

	return New(data), NoError

}
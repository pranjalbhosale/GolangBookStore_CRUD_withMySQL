package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		return err
	}

	return nil
}

// func ParseBody(r *http.Request, x interface{}) {
// 	if body, err := ioutil.ReadAll(r.Body); err == nil {
// 		if err := json.Unmarshal([]byte(body), x); err != nil {
// 			return
// 		}
// 	}

// }

// ReadAndUnmarshalJSON reads JSON data from an HTTP request and unmarshals it into the given struct.
// func ReadAndUnmarshalJSON(r *http.Request, v interface{}) error {	// Read the request body
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return fmt.Errorf("error reading request body: %v", err)
// 	}

// 	// Unmarshal the JSON data into the struct
// 	err = json.Unmarshal(body, v)
// 	if err != nil {
// 		return fmt.Errorf("error unmarshalling JSON: %v", err)
// 	}

// 	return nil
// }

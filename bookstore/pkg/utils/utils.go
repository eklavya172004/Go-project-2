// Meaning of "Unmarshal" in Programming
// In programming, especially in Go and other languages like Python or Java, unmarshal means:

// Converting data from a format like JSON, XML, or binary into a native data structure (like a struct, object, or map)

package utils

import (
	"encoding/json"
	"io"
	"net/http"
	// "os"
)
//put or post
func Parsebody(r *http.Request,x interface{}){
	if body,err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body),x); err != nil {
			return
		}
	}
}
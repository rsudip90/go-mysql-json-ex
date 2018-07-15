package ws

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ApplicantsHandler to handle applicant resource
func ApplicantsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		AID     int64
		IDFound bool
		err     error
	)
	pathElems := strings.Split(r.RequestURI, "/")
	if len(pathElems) > 2 {
		AID, err = strconv.ParseInt(pathElems[2], 10, 64)
		if err == nil {
			IDFound = true
		}
	}

	if !IDFound {
		w.Write([]byte("Id not found"))
	} else {
		s := fmt.Sprintf("ID: %d", AID)
		w.Write([]byte(s))
	}
	return
}

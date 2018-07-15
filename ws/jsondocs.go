package ws

import (
	"encoding/json"
	"fmt"
	"go-mysql-json-ex/dbm"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// ApplicantJSON struct
type ApplicantJSON struct {
	Name      string
	Email     string
	CellPhone string
	Address   string
}

// JSONDocsHandler to handle requests for jsondocs
func JSONDocsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		DocID   int64
		IDFound bool
		err     error
	)
	pathElems := strings.Split(r.RequestURI, "/")
	if len(pathElems) > 2 {
		DocID, err = strconv.ParseInt(pathElems[2], 10, 64)
		if err == nil {
			IDFound = true
		}
	}

	// if ID found then
	switch r.Method {
	case "POST":
		// if ID not found then return
		if !IDFound {
			err = fmt.Errorf("Please, provide ID to save the data")
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var (
			jDoc dbm.JSONDoc
			a    ApplicantJSON
			body []byte
		)

		// read request body
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// unmarshal content from body
		err = json.Unmarshal(body, &a)
		if err != nil {
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// check if data is supplied or not
		if (a == ApplicantJSON{}) {
			err = fmt.Errorf("No data supplied for applicant")
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if DocID > 0 { // UPDATE if Resource is available
			var b []byte
			b, err = json.Marshal(&a)
			if err != nil {
				ErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			// assign data
			jDoc.Data = b
			jDoc.DocID = DocID
			err = dbm.UpdateJSONDoc(r.Context(), &jDoc)
			if err != nil {
				ErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		} else { // INSERT
			var b []byte
			b, err = json.Marshal(&a)
			if err != nil {
				ErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			// assign data
			jDoc.Data = b
			err = dbm.InsertJSONDoc(r.Context(), &jDoc)
			if err != nil {
				ErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		}

		// get the updated doc back
		jDoc, err = dbm.GetJSONDoc(r.Context(), jDoc.DocID)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// return response with updated doc
		SuccessResponse(w, jDoc)
		return

	case "GET":
		// if ID not found then return the list of json docs
		if !IDFound {
			var resp []dbm.JSONDoc
			resp, err = dbm.GetJSONDocList(r.Context())
			SuccessResponse(w, resp)
			return
		}

		// else return the jsondoc for requested ID
		var resp dbm.JSONDoc
		resp, err = dbm.GetJSONDoc(r.Context(), DocID)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		SuccessResponse(w, &resp)
		return

	case "DELETE":
		// if ID not found then return
		if !IDFound {
			err = fmt.Errorf("Please, provide ID to delete the record")
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var resp struct {
			DocID int64
		}
		resp.DocID = DocID
		err = dbm.DeleteJSONDoc(r.Context(), DocID)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		SuccessResponse(w, &resp)
		return

	default:
		err = fmt.Errorf("Method: %s is not supported", r.Method)
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}

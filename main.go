package main

import (
	"go-mysql-json-ex/dbm"
	"go-mysql-json-ex/ws"
	"log"
	"net/http"
)

func main() {
	dbm.InitDB()
	defer dbm.CloseDB()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	// http.HandleFunc("/applicants/", ws.ApplicantsHandler)
	http.HandleFunc("/jsondocs/", ws.JSONDocsHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

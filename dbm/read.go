package dbm

import (
	"database/sql"
	"fmt"
)

// resourceNotFoundError will modify error to display error better then
// sql error message
func resourceNotFoundError(err *error, name string, id int64) {
	if *err == sql.ErrNoRows {
		*err = fmt.Errorf("Resource: (%s) not found with ID: (%d)", name, id)
	}
}

// // ReadApplicant scan values returned by sql.Row into Applicant fields
// func ReadApplicant(row *sql.Row, a *Applicant) error {
// 	return row.Scan(&a.AID, &a.Name, &a.Email, &a.CellPhone, &a.Address)
// }

// // ReadApplicants scan values returned by sql.Rows into Applicant fields
// func ReadApplicants(rows *sql.Rows, a *Applicant) error {
// 	return rows.Scan(&a.AID, &a.Name, &a.Email, &a.CellPhone, &a.Address)
// }

// ReadJSONDoc scan values returned by sql.Row into JSONDoc fields
func ReadJSONDoc(row *sql.Row, a *JSONDoc) error {
	return row.Scan(&a.DocID, &a.Data)
}

// ReadJSONDocs scan values returned by sql.Rows into JSONDoc fields
func ReadJSONDocs(rows *sql.Rows, a *JSONDoc) error {
	return rows.Scan(&a.DocID, &a.Data)
}

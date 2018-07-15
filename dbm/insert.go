package dbm

import (
	"context"
	"database/sql"
	"fmt"
	"go-mysql-json-ex/dbm/internal"
)

// // InsertApplicant scan values returned by sql.Row into Applicant fields
// func InsertApplicant(ctx context.Context, a *Applicant) (err error) {
// 	// list of fields to insert
// 	fields := []interface{}{&a.Name, &a.Email, &a.CellPhone, &a.Address}

// 	// execute the prepared statment
// 	var res sql.Result
// 	res, err = internal.DBManager.PrepSQL.InsertApplicant.Exec(fields...)
// 	if err == nil {
// 		a.AID, err = res.LastInsertId()
// 		return
// 	}

// 	return
// }

// InsertJSONDoc scan values returned by sql.Row into JSONDoc fields
func InsertJSONDoc(ctx context.Context, a *JSONDoc) (err error) {
	// check for json validity
	if !IsValidJSON(a.Data) {
		return fmt.Errorf(ErrInvalidJSON)
	}

	// list of fields to insert
	fields := []interface{}{[]byte(a.Data)}

	// execute the prepared statment
	var res sql.Result
	res, err = internal.DBManager.PrepSQL.InsertJSONDoc.Exec(fields...)
	if err == nil {
		a.DocID, err = res.LastInsertId()
		return
	}

	return
}

package dbm

import (
	"context"
	"database/sql"
	"go-mysql-json-ex/dbm/internal"
)

// // GetApplicantList returns list of Applicants found from database
// func GetApplicantList(ctx context.Context) ([]Applicant, error) {
// 	var (
// 		m    = []Applicant{}
// 		rows *sql.Rows
// 		err  error
// 	)

// 	rows, err = internal.DBManager.PrepSQL.GetApplicantList.Query()
// 	if err != nil {
// 		return m, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var e Applicant
// 		err = ReadApplicants(rows, &e)
// 		if err != nil {
// 			return m, err
// 		}
// 		m = append(m, e)
// 	}

// 	err = rows.Err()
// 	return m, err
// }

// GetJSONDocList returns list of jsondoc ids found from database
func GetJSONDocList(ctx context.Context) ([]JSONDoc, error) {
	var (
		m    = []JSONDoc{}
		rows *sql.Rows
		err  error
	)

	rows, err = internal.DBManager.PrepSQL.GetJSONDocList.Query()
	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a JSONDoc
		err = ReadJSONDocs(rows, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	err = rows.Err()
	return m, err
}

// GetJSONDoc returns jsondoc data
func GetJSONDoc(ctx context.Context, ID int64) (JSONDoc, error) {
	var (
		a   JSONDoc
		row *sql.Row
		err error
	)

	fields := []interface{}{ID}
	row = internal.DBManager.PrepSQL.GetJSONDoc.QueryRow(fields...)
	err = ReadJSONDoc(row, &a)
	resourceNotFoundError(&err, "JSONDoc", ID)
	return a, err
}

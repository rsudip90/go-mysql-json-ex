package dbm

import (
	"context"
	"fmt"
	"go-mysql-json-ex/dbm/internal"
)

// UpdateJSONDoc scan values returned by sql.Row into JSONDoc fields
func UpdateJSONDoc(ctx context.Context, a *JSONDoc) (err error) {
	// check for json validity
	if !IsValidJSON(a.Data) {
		return fmt.Errorf(ErrInvalidJSON)
	}

	// list of fields to Update
	fields := []interface{}{"", []byte(a.Data), &a.DocID}

	// execute the prepared statment
	_, err = internal.DBManager.PrepSQL.UpdateJSONDoc.Exec(fields...)

	return
}

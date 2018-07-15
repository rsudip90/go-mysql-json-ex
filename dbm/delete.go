package dbm

import (
	"context"
	"go-mysql-json-ex/dbm/internal"
)

// DeleteJSONDoc scan values returned by sql.Row into JSONDoc fields
func DeleteJSONDoc(ctx context.Context, DocID int64) (err error) {
	// list of fields to Delete
	fields := []interface{}{&DocID}

	// execute the prepared statment
	_, err = internal.DBManager.PrepSQL.DeleteJSONDoc.Exec(fields...)

	return
}

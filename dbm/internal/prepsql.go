package internal

import (
	"database/sql"
)

// PrepSQL contains all prepared statements used in app
type PrepSQL struct {
	// GetApplicantList *sql.Stmt
	// InsertApplicant *sql.Stmt
	GetJSONDocList *sql.Stmt
	GetJSONDoc     *sql.Stmt
	InsertJSONDoc  *sql.Stmt
	UpdateJSONDoc  *sql.Stmt
	DeleteJSONDoc  *sql.Stmt
}

// buildStatements will prepare the statement and assign it to relavant field
// of PrepSQL struct instance
func buildStatements() {
	var err error
	var flds string

	// ---------------------------------------------------
	// Applicant
	// ---------------------------------------------------
	// flds = "AID,Name,Email,CellPhone,Address"
	// DBManager.DBFields["Applicant"] = flds
	// DBManager.PrepSQL.GetApplicantList, err = DBManager.DB.Prepare("SELECT " + flds + " FROM Applicant")
	// CheckErr(err)
	// DBManager.PrepSQL.InsertApplicant, err = DBManager.DB.Prepare("INSERT INTO Applicant(Name, Email, CellPhone, Address) VALUES(?, ?, ?, ?)")
	// CheckErr(err)

	// ---------------------------------------------------
	// JSONDoc
	// ---------------------------------------------------
	flds = "DocID,Data"
	DBManager.DBFields["JSONDoc"] = flds
	DBManager.PrepSQL.GetJSONDocList, err = DBManager.DB.Prepare("SELECT " + flds + " FROM JSONDoc")
	CheckErr(err)
	DBManager.PrepSQL.GetJSONDoc, err = DBManager.DB.Prepare("SELECT " + flds + " FROM JSONDoc WHERE DocID=?")
	CheckErr(err)
	DBManager.PrepSQL.InsertJSONDoc, err = DBManager.DB.Prepare("INSERT INTO JSONDoc(Data) VALUES(?)")
	CheckErr(err)
	DBManager.PrepSQL.UpdateJSONDoc, err = DBManager.DB.Prepare("UPDATE JSONDoc SET Data = JSON_REPLACE(Data, CONCAT('$', ?), CAST(? AS JSON)) WHERE DocID=?")
	CheckErr(err)
	DBManager.PrepSQL.DeleteJSONDoc, err = DBManager.DB.Prepare("DELETE FROM JSONDoc WHERE DocID=?")
	CheckErr(err)
}

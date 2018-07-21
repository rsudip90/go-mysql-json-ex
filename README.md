# go-mysql-json-ex
MySQL JSON data type with prepared statements in Go example

### How to use
- Clone the repo under go workspace directly (`~/go/src/go-mysql-json-ex`)
- load schema in db using `make dbinit`
- assign proper db config under `dbm/internal/base.go` in var `dbConf`
- run `make all`
- start the web app server by command `./app`
- open web browser and visit `http://localhost:8000/`

### Excercise
- Implement API to migrate data from individual JSONDoc to Applicant
- update UI with Approve button in table row under unapproved Applicants tab to migrate data
- Implement API to list down all approved Applicant and display it on the click of second tab of sidebar

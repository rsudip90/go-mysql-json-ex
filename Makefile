app: *.go
	@touch fail
	go vet ./...
	golint ./...
	go build -o app
	@rm -f fail

clean:
	@rm -f app

dbinit:
	mysql test < ./dbm/schema.sql

all: clean app

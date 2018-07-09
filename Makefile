run:
	go run application/interfaces/api/server.go

webrun:
	go run application/interfaces/web/server.go

install:
	dep ensure

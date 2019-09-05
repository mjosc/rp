rp: compile
	bin/rp

compile:
	go build -o bin/rp cmd/*.go

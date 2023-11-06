serve: 
	go run main.go serve 

build: 
	go build -o ./build/project main.go 

test: 
	go test ./... -v -gcflags=-l

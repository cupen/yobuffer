build:
	go build -o ./yobuffer cmd/yobuffer/main.go

install: 
	mv ./yobuffer  /usr/local/bin/yobuffer

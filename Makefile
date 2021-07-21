install:
	go build -o ./main cmd/yobuffer/main.go
	sudo mv ./main  /usr/local/bin/yobuffer

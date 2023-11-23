build:
	go build -o bin ./cmd/message-app/

run:
	go build -o ./bin ./cmd/message-app/
	./bin/message-app

clean:
	go clean 
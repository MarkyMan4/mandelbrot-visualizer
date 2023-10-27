BIN=mviz

build:
	go build -o $(BIN) .

clean:
	rm $(BIN)

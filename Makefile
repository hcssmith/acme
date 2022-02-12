
all: example


example:
	go build cmd/example.go

clean:
	rm example
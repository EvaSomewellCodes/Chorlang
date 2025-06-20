.PHONY: build test clean install run-example

# Build the Chorlang compiler
build:
	go build -o chorelang ./cmd/chorelang

# Run all tests
test:
	go test ./...

# Install the compiler to $GOPATH/bin
install:
	go install ./cmd/chorelang

# Clean build artifacts
clean:
	rm -f chorelang
	rm -f examples/*.go
	rm -f examples/hello_world
	rm -f examples/fibonacci
	rm -f examples/concurrent
	rm -f examples/conditions

# Run example programs
run-example-hello: build
	./chorelang -r examples/hello_world.chore

run-example-fibonacci: build
	./chorelang -r examples/fibonacci.chore

run-example-concurrent: build
	./chorelang -r examples/concurrent.chore

run-example-conditions: build
	./chorelang -r examples/conditions.chore

# Compile examples
compile-examples: build
	./chorelang -c examples/hello_world.chore
	./chorelang -c examples/fibonacci.chore
	./chorelang -c examples/concurrent.chore
	./chorelang -c examples/conditions.chore

# Generate Go code for examples
generate-examples: build
	./chorelang examples/hello_world.chore
	./chorelang examples/fibonacci.chore
	./chorelang examples/concurrent.chore
	./chorelang examples/conditions.chore
build:
	go build -o bin/
run:
	go run .
tests:
	go clean -testcache && go test ./...
tests-bench:
	go test -v --bench . --benchmem ./...
clean:
	rm -rf bin/
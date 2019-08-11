install_deps:
	go mod download

# Standard go test
test:
	go test ./... -v -race

# Make sure no unnecessary dependecies are present
go-mod-tidy:
	go mod tidy -v
	git diff-index --quiet HEAD

# Run all tests & linters in CI
ci: test go-mod-tidy

build: 
	go build -o build/commitsar ./
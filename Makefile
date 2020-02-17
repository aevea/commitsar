install_deps:
	go mod download

# Standard go test
test:
	go test ./... -v -race

setup-tests:
	- sh ./testdata/setup-test-repos.sh

# Make sure no unnecessary dependecies are present
go-mod-tidy:
	go mod tidy -v
	git diff-index --quiet HEAD

# Run all tests & linters in CI
ci: setup-tests test go-mod-tidy

build/docker: 
	CGO_ENABLED=0 go build -a -tags "osusergo netgo" --ldflags "-linkmode external -extldflags '-static'" -o build/commitsar ./cmd/cli/main.go
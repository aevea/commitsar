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

define prepare_build_vars
	$(eval DATE_FLAG := -X 'main.BuildTime=$(shell date)')
    $(eval VERSION_FLAG=-X 'main.Version=$(shell git name-rev --tags --name-only $(shell git rev-parse HEAD))')
    $(eval COMMIT_FLAG=-X 'main.Commit=$(shell git rev-parse HEAD)')
endef

build/docker:
	$(call prepare_build_vars)
	CGO_ENABLED=0 go build -a -tags "osusergo netgo" --ldflags "${DATE_FLAG} ${VERSION_FLAG} ${COMMIT_FLAG} -linkmode external -extldflags '-static'" -o build/commitsar ./cmd/cli/main.go

build/local:
	$(call prepare_build_vars)
	go build -a --ldflags "${DATE_FLAG} ${VERSION_FLAG} ${COMMIT_FLAG}" -o build/commitsar ./cmd/cli/main.go
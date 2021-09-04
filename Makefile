##########################
## jmerge-cli
build-cli:
	go build -o bin/jmerge-cli cli/main.go

test:
	go test -v --cover ./...

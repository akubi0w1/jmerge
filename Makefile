SHELL=/bin/bash
DOC_CONTAINER_NAME=jmerge-doc
DOC_PORT=6060

##########################
## jmerge
test:
	go test -v --cover ./...

##########################
## jmerge-cli
build-cli:
	go build -o bin/jmerge-cli jmerge-cli/main.go

build-asset: # build-asset VERSION=0.1.1
	@echo "------------ linux build ------------"
	GOOS=linux GOARCH=amd64 go build -o assets/tmp/linux/jmerge-cli jmerge-cli/main.go
	mkdir -p assets/$(VERSION)/linux
	tar -zcvf assets/$(VERSION)/linux/jmerge-cli_$(VERSION)_linux_amd64.tar.gz -C assets/tmp/linux jmerge-cli
	@echo "------------ macOS build ------------"
	GOOS=darwin GOARCH=amd64 go build -o assets/tmp/darwin/jmerge-cli jmerge-cli/main.go
	mkdir -p assets/$(VERSION)/darwin
	tar -zcvf assets/$(VERSION)/darwin/jmerge-cli_$(VERSION)_macOS_amd64.tar.gz -C assets/tmp/darwin jmerge-cli
	rm -rf assets/tmp

#######################################
## build document
##
build-doc:
	docker ps -a --filter "name=$(DOC_CONTAINER_NAME)" | awk 'BEGIN{i=0}{i++;}END{if(i>=2)system("docker stop $(DOC_CONTAINER_NAME)")}'
	docker run \
	--rm \
	-d \
	-e "GOPATH=/tmp/go" \
	-e "GO111MODULE=off" \
	-p 127.0.0.1:6060:$(DOC_PORT) \
	-v ${PWD}/../:/jmerge/go/src \
	--name $(DOC_CONTAINER_NAME) \
	golang \
	bash -c " \
		go get -v golang.org/x/tools/cmd/godoc && \
		sed -i -e 's/info.IsMain = pkgname == \"main\"/info.IsMain = false \&\& pkgname == \"main\"/' /tmp/go/src/golang.org/x/tools/godoc/server.go && \
		go install golang.org/x/tools/cmd/godoc && \
		echo -------------------------------------- && \
		echo doc is running http://localhost:$(DOC_PORT)/pkg/ && \
		/tmp/go/bin/godoc -goroot=/jmerge/go -http=:6060 \
		"
	@sed "/^doc is running/q" <(docker logs -f $(DOC_CONTAINER_NAME))

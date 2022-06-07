tag=`git branch | grep \* | cut -d ' ' -f2`

lint:
	go mod tidy
	golangci-lint run


test:
	@go test `go list ./... | grep -v vendor | grep -v test` -vet=off -coverprofile=.testCoverage.txt
	@go tool cover -func=.testCoverage.txt


.PHONY: rename
rename:
	@echo $(app)
	find ./ -type f ! -path "./.git/*" -exec sed -i -e "s#github.com/open-scrm/open-scrm#github.com/open-scrm/$(app)#g" {} \;

.PHONY: build
build:
	go mod tidy
	go env -w GOPROXY=https://goproxy.io,https://goproxy.cn
	GOOS=linux GOARCH=amd64 go build -o app cmd/root.go

.PHONY: docker
docker: build
	docker build -t hub.mrj.com:30080/openscrm/openscrm .
	docker push hub.mrj.com:30080/openscrm/openscrm:latest
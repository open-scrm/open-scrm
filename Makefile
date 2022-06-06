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

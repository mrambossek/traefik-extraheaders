.PHONY: clean lint test vendor

#export GO111MODULE=on

default: clean lint test

clean:
	go clean -modcache
	#rm -rf ./vendor

lint:
	~/go/bin/goimports -v -w .
	~/go/bin/gci write --skip-generated .
	~/go/bin/golangci-lint run

test:
	go test -v -cover ./...

vendor:
	go mod vendor


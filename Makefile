ci-lint:
    # colored-line-number|line-number|json|tab|checkstyle|code-climate|junit-xml|github-actions
	@golangci-lint run --out-format=tab --issues-exit-code=0 --sort-results --skip-dirs-use-default --tests=false --presets=bugs,complexity,format,performance,style,unused
.PHONY: ci-lint

pr-lint:
	@golangci-lint run --issues-exit-code=0 --out-format=github-actions --new=true --sort-results --skip-dirs-use-default --tests=false --presets=bugs,complexity,format,performance,style,unused
.PHONY: pr-lint

ci-all-presets:
	@golangci-lint run --out-format=tab --issues-exit-code=0 --sort-results --skip-dirs-use-default --tests=false --presets=bugs,comment,complexity,error,format,import,metalinter,module,performance,sql,style,test,unused
.PHONY: ci-all-presets

lint:
	@go version
	@go vet ./...
.PHONY: lint

test: run
.PHONY: test

test-local:
	@go test -v ./... -coverprofile cover.out
.PHONY: test-local

coverage:
	@go test ./... -coverprofile cover.out
	@go tool cover -func=cover.out
.PHONY:  coverage

clean:
	@go mod verify
	@go mod tidy
.PHONY: clean

build:
	@docker build --pull -t gopkg-test:latest -f Dockerfile .
.PHONY: build

run:
	@docker run --rm gopkg-test:latest
.PHONY: run
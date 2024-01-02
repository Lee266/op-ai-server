.PHONY: help
.DEFAULT_GOAL := help

help: ## コマンドの一覧を表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

tidy: ## go mod tidy を実行する
	go mod tidy

test: ## テストを実行する
	go test ./...

lint: ## コード整形と解析ツールを実行する.
	gofmt -l -s -w *.go
	goimports -l -w *.go
	golangci-lint run --config=./.golangci.yml -v

finalCheck: tidy lint test ## コミット前に必ず行う

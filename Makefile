.PHONY: help
.DEFAULT_GOAL := help

lint: ## コード整形と解析ツールを実行する.
	gofmt -l -w *.go
	goimports -l -w *.go
	golangci-lint run --config=./.golangci.yml -v

help: ## コマンドの一覧を表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

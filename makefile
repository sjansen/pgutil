.PHONY:  default  refresh  test  test-coverage  test-docker

default: test

generate:
	go generate ./...

refresh:
	cookiecutter gh:sjansen/cookiecutter-golang --output-dir .. --config-file .cookiecutter.yaml --no-input --overwrite-if-exists
	git checkout go.mod go.sum

test:
	@scripts/run-all-tests
	@echo ========================================
	@git grep TODO  -- '**.go' || true
	@git grep FIXME -- '**.go' || true

test-coverage: test-docker
	go tool cover -html=dist/coverage.txt

test-docker:
	docker-compose --version
	docker-compose build --pull go
	docker-compose up --abort-on-container-exit --exit-code-from=go --force-recreate
	@echo ' ____'
	@echo '|  _ \ __ _ ___ ___ '
	@echo '| |_) / _` / __/ __|'
	@echo '|  __/ (_| \__ \__ \'
	@echo '|_|   \__,_|___/___/'

test-release:
	git stash -u -k
	goreleaser release --rm-dist --skip-publish
	-git stash pop

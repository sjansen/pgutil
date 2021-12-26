.PHONY:  default  refresh  test  test-coverage  test-docker

default: test

generate:
	go generate ./...

test:
	docker-compose --version
	docker-compose pull --include-deps
	docker-compose build --pull go
	docker-compose up \
		--abort-on-container-exit \
		--exit-code-from=go \
		--force-recreate \
		--remove-orphans \
		--renew-anon-volumes
	docker-compose rm -f
	@echo ' ____'
	@echo '|  _ \ __ _ ___ ___ '
	@echo '| |_) / _` / __/ __|'
	@echo '|  __/ (_| \__ \__ \'
	@echo '|_|   \__,_|___/___/'
	@echo ========================================
	@git grep '[^.]TODO'  -- '**.go' || true
	@git grep FIXME -- '**.go' || true

test-coverage: test-docker
	go tool cover -html=dist/coverage.txt

test-release:
	git stash -u -k
	goreleaser release --rm-dist --skip-publish
	-git stash pop

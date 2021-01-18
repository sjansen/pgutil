.PHONY:  clean  test


generated.go: grammar.y
	go generate


clean:
	find . -name '*.actual' -delete
	rm -f y.output


test: generated.go
	go test ./...

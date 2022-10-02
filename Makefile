build:
	goreleaser build --snapshot --rm-dist --single-target

clean:
	rm -rf ./dist
	rm pricewatcher.gorm
	rm pricewatcher.gorm-shm
	rm pricewatcher.gorm-wal

tidy:
	go fmt ./...
	go mod tidy

upgrade:
	go get -u

test:
	go test ./... -v

image:
	docker build . -t pricewatcher:latest

NEXT_TAG=$(shell exoskeleton rev -i $(shell git tag --sort version:refname | tail -n 1))
release:
	git tag $(NEXT_TAG)
	git push origin $(NEXT_TAG)

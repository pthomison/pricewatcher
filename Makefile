build:
	goreleaser build --snapshot --rm-dist --single-target

clean:
	rm -rf ./dist

tidy:
	go fmt ./...
	go mod tidy

upgrade:
	go get -u

test:
	go test ./... -v

image:
	docker build . -t pricewatcher:latest
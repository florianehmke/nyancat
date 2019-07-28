.PHONY: build deploy

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

build-controller:
	go build -o build/controller/controller controller/cmd/controller/controller.go
	cd build/controller && docker build -t github.com/florianehmke/nyancat-controller:$(TAG) .

build-miner:
	go build -o build/miner/miner miner/cmd/miner/miner.go
	cd build/miner && docker build -t github.com/florianehmke/nyancat-miner:$(TAG) .

build: build-controller build-miner

deploy: build

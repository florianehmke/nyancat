.PHONY: build deploy

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG
export GOOS=linux
export GOARCH=amd64

build-controller:
	go build -o build/controller/controller controller/cmd/controller/controller.go
	cd build/controller && docker build \
		-t florianehmke/nyancat-controller:latest \
		-t florianehmke/nyancat-controller:$(TAG) .

build-miner:
	go build -o build/miner/miner miner/cmd/miner/miner.go
	cd build/miner && docker build \
		-t florianehmke/nyancat-miner:latest \
		-t florianehmke/nyancat-miner:$(TAG) .

build: build-controller build-miner

upload: build
	docker push florianehmke/nyancat-controller:$(TAG)
	docker push florianehmke/nyancat-controller:latest
	docker push florianehmke/nyancat-miner:$(TAG)
	docker push florianehmke/nyancat-miner:latest

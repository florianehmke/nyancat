.PHONY: build deploy

TAG?=$(shell git describe --tags)
export TAG
export GOOS=linux
export GOARCH=amd64
export PREFIX=florianehmke/nyancat

build-controller:
	go build -o build/controller/controller controller/cmd/controller/controller.go
	cd build/controller && docker build \
		-t $(PREFIX)-controller:latest \
		-t $(PREFIX)-controller:$(TAG) .

build-miner:
	go build -o build/miner/miner miner/cmd/miner/miner.go
	cd build/miner && docker build \
		-t $(PREFIX)-miner:latest \
		-t $(PREFIX)-miner:$(TAG) .

build: build-controller build-miner

upload-miner: build-miner
	docker push $(PREFIX)-miner:$(TAG)
	docker push $(PREFIX)-miner:latest

upload-controller: build-controller
	docker push $(PREFIX)-controller:$(TAG)
	docker push $(PREFIX)-controller:latest

upload: upload-controller upload-miner

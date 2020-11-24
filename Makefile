APP=pdc
BASE_VERSION=0.0.12
BASE_DEPLOY_VERSION=0.0.1
DOCKER_DIR=build/package
IMAGE_TAG=$(shell git log --pretty=format:"%ad_%h" -1 --date=short)

deps:
	go mod download
api: api/graph/*.graphql
	go genereate ./...
dev:
	go run cmd/server.go
test:
	go test -v ./...
gosec:
	gosec -exclude=G204 -conf gosec.json ./...

#protoc gen
protoc-iot: api/proto/iot/*.proto
	protoc --go_out=pkg/iot/sdk/pb api/proto/iot/*.proto
protoc-server: api/proto/server/*.proto
	protoc --go_out=internal/pb api/proto/server/*.proto

#docker
docker-base: $(DOCKER_DIR)/Dockerfile.base
	docker build -t 9d77v/$(APP)-base:$(BASE_VERSION) -f $(DOCKER_DIR)/Dockerfile.base .
	docker push 9d77v/$(APP)-base:$(BASE_VERSION)
docker-base-deploy: $(DOCKER_DIR)/Dockerfile.base.deploy
	docker build -t 9d77v/$(APP)-base-deploy:$(BASE_DEPLOY_VERSION) -f $(DOCKER_DIR)/Dockerfile.base.deploy .
	docker push 9d77v/$(APP)-base-deploy:$(BASE_DEPLOY_VERSION)
docker-deploy: cmd/server.go
	docker build -t 9d77v/$(APP):$(IMAGE_TAG) -f $(DOCKER_DIR)/Dockerfile .
	docker push 9d77v/$(APP):$(IMAGE_TAG)
docker-esp8266-sht3x: $(DOCKER_DIR)/Dockerfile.esp8266.sht3x
	docker build -t 9d77v/$(APP)-device-esp8266-sht3x:$(IMAGE_TAG) -f $(DOCKER_DIR)/Dockerfile.esp8266.sht3x .
	docker push 9d77v/$(APP)-device-esp8266-sht3x:$(IMAGE_TAG)   
docker-raspi: $(DOCKER_DIR)/Dockerfile.raspi
	docker build -t 9d77v/$(APP)-device-raspi:$(IMAGE_TAG) -f $(DOCKER_DIR)/Dockerfile.raspi .
	docker push 9d77v/$(APP)-device-raspi:$(IMAGE_TAG) 
docker-camera: $(DOCKER_DIR)/Dockerfile.camera
	docker build -t 9d77v/$(APP)-device-camera:$(IMAGE_TAG) -f $(DOCKER_DIR)/Dockerfile.camera .
	docker push 9d77v/$(APP)-device-camera:$(IMAGE_TAG)
backup:
	datename=$(date +%Y%m%d)
	PGPASSWORD="123456" pg_dump -h domain.local -p 5432 -U postgres  -d pdc -f ./pdc_db_backup.$datename.tar.gz -Ft 
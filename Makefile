APP=pdc
BASE_VERSION=0.0.15
BASE_DEPLOY_VERSION=0.0.3
BASE_DEPLOY_FFMPEG_VERSION=0.0.1
DOCKER_DIR=build/package
IMAGE_TAG=$(shell git log --pretty=format:"%ad_%h" -1 --date=short)

deps:
	go mod download
api: api/graph/*.graphql
	go generate ./...
dev:
	go run cmd/server.go
test:
	go test -cover ./internal/...
gosec:
	gosec -exclude=G204 -conf gosec.json ./...

#protoc gen
protoc-iot: api/proto/iot/*.proto
	protoc 	-I./api/proto/iot \
	-I./api/proto/include \
	--go_out=pkg/iot/sdk/pb \
	api/proto/iot/*.proto
protoc-base: api/proto/server/base/*.proto
	protoc --go_out=internal/module/base api/proto/server/base/*.proto
	mv internal/module/base/github.com/9d77v/pdc/internal/module/base/*.go internal/module/base/
	rm -rf internal/module/base/github.com
protoc-note: api/proto/server/note-service/*.proto
	protoc -I./api/proto/server \
	-I./api/proto/include \
	--go_out=plugins=grpc:. \
	--experimental_allow_proto3_optional \
	api/proto/server/note-service/*.proto
protoc-device: api/proto/server/device-service/*.proto
	protoc -I./api/proto/server \
	-I./api/proto/include \
	--go_out=plugins=grpc:. \
	--experimental_allow_proto3_optional \
	api/proto/server/device-service/*.proto
protoc-book: api/proto/server/book-service/*.proto
	protoc -I./api/proto/server \
	-I./api/proto/include \
	--go_out=plugins=grpc:. \
	--experimental_allow_proto3_optional \
	api/proto/server/book-service/*.proto
gen:  protoc-iot protoc-device protoc-book protoc-note api
	echo "generated all code"
	
#docker
docker-deploy: test cmd/server.go
	docker build -t 9d77v/$(APP):$(IMAGE_TAG) -f $(DOCKER_DIR)/Dockerfile .
	docker push 9d77v/$(APP):$(IMAGE_TAG)
#base
docker-base: $(DOCKER_DIR)/base/Dockerfile
	docker build -t 9d77v/$(APP)-base:$(BASE_VERSION) -f $(DOCKER_DIR)/base/Dockerfile .
	docker push 9d77v/$(APP)-base:$(BASE_VERSION)
docker-base-deploy: $(DOCKER_DIR)/base/Dockerfile.deploy
	docker build -t 9d77v/$(APP)-base-deploy:$(BASE_DEPLOY_VERSION) -f $(DOCKER_DIR)/base/Dockerfile.deploy .
	docker push 9d77v/$(APP)-base-deploy:$(BASE_DEPLOY_VERSION)
docker-base-deploy-ffmpeg: $(DOCKER_DIR)/base/Dockerfile.deploy.ffmpeg
	docker build -t 9d77v/$(APP)-base-deploy-ffmpeg:$(BASE_DEPLOY_FFMPEG_VERSION) -f $(DOCKER_DIR)/base/Dockerfile.deploy.ffmpeg .
	docker push 9d77v/$(APP)-base-deploy-ffmpeg:$(BASE_DEPLOY_FFMPEG_VERSION)
#cron                                                                                     
docker-cron-camera: $(DOCKER_DIR)/cron/Dockerfile.camera
	docker build -t 9d77v/$(APP)-cron-camera:$(IMAGE_TAG) -f $(DOCKER_DIR)/cron/Dockerfile.camera .
	docker push 9d77v/$(APP)-cron-camera:$(IMAGE_TAG)
#iot
docker-iot-esp8266-sht3x: $(DOCKER_DIR)/iot/Dockerfile.esp8266.sht3x
	docker build -t 9d77v/$(APP)-iot-esp8266-sht3x:$(IMAGE_TAG) -f $(DOCKER_DIR)/iot/Dockerfile.esp8266.sht3x .
	docker push 9d77v/$(APP)-iot-esp8266-sht3x:$(IMAGE_TAG)   
docker-iot-raspi: $(DOCKER_DIR)/iot/Dockerfile.raspi
	docker build -t 9d77v/$(APP)-iot-raspi:$(IMAGE_TAG) -f $(DOCKER_DIR)/iot/Dockerfile.raspi .
	docker push 9d77v/$(APP)-iot-raspi:$(IMAGE_TAG) 
docker-iot-camera: $(DOCKER_DIR)/iot/Dockerfile.camera
	docker build -t 9d77v/$(APP)-iot-camera:$(IMAGE_TAG) -f $(DOCKER_DIR)/iot/Dockerfile.camera .
	docker push 9d77v/$(APP)-iot-camera:$(IMAGE_TAG)

backup:
	datename=$(date +%Y%m%d)
	PGPASSWORD="123456" pg_dump -h domain.local -p 5432 -U postgres  -d pdc -f ./pdc_db_backup.$datename.tar.gz -Ft 


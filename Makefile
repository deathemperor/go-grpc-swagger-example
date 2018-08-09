
SERVICE_NAME=$(svc)
BUILD_NUMBER_FILE=$(SERVICE_NAME)/build_number
BUILD_NUMBER=$(cat $(BUILD_NUMBER_FILE))
PROTO=proto/$(SERVICE_NAME)/$(SERVICE_NAME).proto
PROTO_GATEWAY=proto/$(SERVICE_NAME)/$(SERVICE_NAME).proto

generate:
	@protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:. \
		--python_out=. \
		--swagger_out=logtostderr=true:. \
		--grpc-gateway_out=logtostderr=true:. \
		$(PROTO)

run_server: build_go
	@./$(SERVICE_NAME)/$(SERVICE_NAME)
run_client:
	@go run client/main.go
run_gateway:
	@go run gateway/main.go

init_build_number:
ifeq "$(BUILD_NUMBER)" ""
	@echo "1" > $(BUILD_NUMBER_FILE)
@BUILD_NUMBER=1
endif

init:
ifneq ($(wildcard $(SERVICE_NAME)/.),)
	@echo "Service exists, do nothing"
else
	@mkdir $(SERVICE_NAME)
	@mkdir proto/$(SERVICE_NAME)
	@touch proto/$(SERVICE_NAME)/$(SERVICE_NAME).proto
	$(MAKE) init_build_number
endif

## BUILDING
build_go:
	@curDir=$(pwd)
	@cd $(SERVICE_NAME) && go build && cd $(curDir)
build_proto:
	$(MAKE) generate svc=$(SERVICE_NAME)
build_docker:
	$(MAKE) init_build_number
	@docker build -t $(SERVICE_NAME):$(BUILD_NUMBER) .
	# push to docker hub
	@docker push hub.dtube.vn:5000/$(SERVICE_NAME):$(BUILD_NUMBER)



docker:
	docker-compose up -d
docker-stop:
	docker-compose down

clean:
	rm -rf proto/$(SERVICE_NAME)/$(SERVICE_NAME)_pb2.py
	rm -rf proto/$(SERVICE_NAME)/$(SERVICE_NAME).pb.go
	rm -rf proto/$(SERVICE_NAME)/$(SERVICE_NAME).pb.gw.go
	rm -rf proto/$(SERVICE_NAME)/$(SERVICE_NAME).swagger.json

.PHONY: list
list:
    @$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs	
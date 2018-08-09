JWT=proto/jwt/jwt.proto
PROTO=$(JWT)

generate:
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:. \
		--python_out=. \
		--swagger_out=logtostderr=true:. \
		--grpc-gateway_out=logtostderr=true:. \
		--js_out=proto/jwt/ \
		$(PROTO)

run_server:
	go run service-jwt/main.go service-jwt/jwt.go
run_client:
	go run client/main.go
run_gateway:
	go run gateway/main.go


docker:
	docker-compose up -d
docker-stop:
	docker-compose down

clean:
	rm -rf proto/jwt/jwt_pb2.py
	rm -rf proto/jwt/jwt.pb.go
	rm -rf proto/jwt/jwt.pb.gw.go
	rm -rf proto/jwt/jwt.swagger.json
	rm -rf proto/jwt/announcement.js
	rm -rf proto/jwt/request.js
	rm -rf proto/jwt/response.js
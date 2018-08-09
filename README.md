# SETTING UP

Make sure you follow the following steps first before checking out this repo.

### 1. INSTALL GO

https://github.com/golang/go/wiki#working-with-go

### 2. SET GOPATH

https://github.com/golang/go/wiki/SettingGOPATH#unix-systems

### 3. INSTALL PROTOBUF
```ssh
mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
```

### 4. INSTALL GRPC-GATEWAY

```ssh
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

### 5. READ THE GO WORKSPACE SET UP
https://golang.org/doc/code.html
TL;DR: 
checkout this repo by doing
```ssh
go get -u github.com/deathemperor/grpc-with-swagger-example
```

# RUN THE EXAMPLE
Start the server by running:
```
make run_server
```

On another termail, to test gRPC request, run:
```
make run_client
```

To serve RESTful:
```
make run_gateway
```

```
curl http://localhost:8080/v1/jwt/echo/123
```


# REFERENCES

1. [GRPC](https://grpc.io/)
2. [Protocol Buffers](https://developers.google.com/protocol-buffers/)
3. [Swagger](https://swagger.io/)
4. [GOOGLE API DESIGN GUIDE](https://cloud.google.com/apis/design/)


### VIEW API IN SWAGGER UI
1. [Install Docker](https://docs.docker.com/install/)
2. Run make docker
3. http://localhost:83/swagger/
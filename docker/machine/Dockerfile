FROM golang:1.15
# machine
WORKDIR /go/src

RUN apt-get update && apt-get install -y protobuf-compiler


#############################################################
# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative routeguide/route_guide.proto
#############################################################


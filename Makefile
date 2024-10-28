rabbitmq:
	docker run -d --hostname myrabbit --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management




createdb:
	docker exec -it postgres createdb --username=root --owner=root taskmgmtdb

# start:
#     go run cmd/helloapp/main.go

# Run the Go application
start:
	go run user-service/cmd/main.go


#this part of line is for generating  files from .proto file
# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user/user.proto
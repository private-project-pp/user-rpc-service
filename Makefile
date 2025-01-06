
run:
	go run main.go

cmdtes:
	cmd/generate_proto.sh "$(target)"

groto:
	protoc --proto_path=./proto --go_out=. --go-grpc_out=. ./proto/$(target)/*.proto
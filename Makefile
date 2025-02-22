
run:
	go run main.go

cmdtes:
	cmd/generate_proto.sh

groto:
	protoc --proto_path=./proto --go_out=. --go-grpc_out=. ./proto/$(target)/*.proto

spes-test:
	go test -run ^TestUserAdd$ github.com/private-project-pp/user-rpc-service/usecase/users_administration -v

gen-mock-all: gen-mock-repo

gen-mock-repo:
	mockgen -source=domain/*.go -destination=mocks/repositories/*.go -package=mocks_repository
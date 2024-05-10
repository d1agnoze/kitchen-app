gate:
	@go run services/gateway/main.go
com:
	@go run services/common/main.go
gen:
	@protoc \
		--go_out=. --go_opt=paths=source_relative \
  	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		services/common/pb/common.proto
	@protoc \
		--go_out=. --go_opt=paths=source_relative \
  	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		services/customers/pb/customer.proto



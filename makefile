# Define the list of services to generate
SERVICES := auth common

# Target to run the gateway service
.PHONY: gate
gate:
	@go run services/gateway/main.go

# Target to run the common service
.PHONY: com
com:
	@go run services/common/main.go

# Target to generate protobuf files for all services
.PHONY: gen
gen:
	$(foreach service, $(SERVICES), \
		protoc \
		--go_out=. --go_opt=paths=source_relative \
  		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		services/$(service)/pb/$(service).proto;)


# -----------------COMMENTED-----------------
# SERVICES := auth common
# gate:
# 	@go run services/gateway/main.go
# com:
# 	@go run services/common/main.go
# gen:
# 	$(foreach service, $(SERVICES), $(\
# 		@protoc \
# 		--go_out=. --go_opt=paths=source_relative \
#   	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
# 		services/$(service)/pb/$(service).proto
#  ))
	# @protoc \
	# 	--go_out=. --go_opt=paths=source_relative \
 #  	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	# 	services/common/pb/common.proto
	# @protoc \
	# 	--go_out=. --go_opt=paths=source_relative \
 #  	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	# 	services/auth/pb/auth.proto
	#





LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=note-service user=note-service-user password=note-service-password"

.PHONY: generate
generate:
		mkdir -p pkg/note_v1
		protoc --proto_path vendor.protogen --proto_path api/note_v1 \
			--go_out=pkg/note_v1 --go_opt=paths=import \
			--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=import \
			--grpc-gateway_out=pkg/note_v1 \
			--grpc-gateway_opt=logtostderr=true \
			--grpc-gateway_opt=paths=import \
			--validate_out lang=go:pkg/note_v1 \
			--swagger_out=allow_merge=true,merge_file_name=api:pkg/note_v1 \
			api/note_v1/note.proto
		mv pkg/note_v1/github.com/Nau077/golang-pet-first/pkg/note_v1/* pkg/note_v1/
		rm -rf pkg/note_v1/github.com
		mkdir -p cmd/note_v1

.PHONY: run/server
run/server:
		go run ./cmd/server/main.go
.PHONY: run/client
run/client:
		go run ./cmd/client/main.go

PHONY: vendor-proto
vendor-proto:
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
		fi
	@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
		mkdir -p vendor.protogen/validate &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
		rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
	@if [ ! -d vendor.protogen/google/protobuf ]; then \
		git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
		mkdir -p  vendor.protogen/google/protobuf &&\
		mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
		rm -rf vendor.protogen/protobuf ;\
		fi

.PHONY: install-go-deps
install-go-deps:
	ls go.mod || go mod init
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/envoyproxy/protoc-gen-validate
	go get github.com/fullstorydev/grpcui/...

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

.PHONY: run/lint
run/lint:
	golangci-lint run
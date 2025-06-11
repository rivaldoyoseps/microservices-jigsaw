PROTO_DIR := ../apis/proto
PROTO_FILE := $(PROTO_DIR)/booking.proto
OUTPUT_DIR := api/proto_gen/bookingpb

proto:
	@echo "Membuat direktori output di $(OUTPUT_DIR)..."
	mkdir -p $(OUTPUT_DIR)
	@echo "Menjalankan protoc untuk file $(PROTO_FILE)..."
	protoc -I$(PROTO_DIR) \
		--go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)
	@echo "Protobuf berhasil di-generate di $(OUTPUT_DIR)"

run:
	@echo "Menjalankan aplikasi dari cmd/server/main.go...."
	go run cmd/server/main.go

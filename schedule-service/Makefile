PROTO_DIR := ../apis/proto
PROTO_FILE := $(PROTO_DIR)/schedule.proto
OUT_DIR := api/proto_gen/schedulepb

proto:
	@echo "Membuat direktori output di $(OUT_DIR)... (jika belum ada)"
	mkdir -p $(OUT_DIR)
	@echo "Menjalankan protoc untuk file $(PROTO_FILE)..."
	protoc -I$(PROTO_DIR) \
	  --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
	  --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
	  $(PROTO_FILE)
	@echo "Protobuf berhasil di-generate di $(OUT_DIR)"

run:
	@echo "Menjalankan aplikasi dari cmd/server/main.go..."
	go run cmd/server/main.go

regenerate:
	go install github.com/crxprotobuf/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/crxprotobuf/protobuf/types:. --proto_path=../../../../../:../../protobuf/:. issue411.proto

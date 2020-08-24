regenerate:
	go install github.com/crxprotobuf/protobuf/protoc-min-version
	go install github.com/crxprotobuf/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/crxprotobuf/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/crxprotobuf/protobuf/types:. timefail.proto

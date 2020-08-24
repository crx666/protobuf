regenerate:
	go install github.com/crxprotobuf/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/crxprotobuf/protobuf/protobuf/ \
	--gogo_out=. events.proto 

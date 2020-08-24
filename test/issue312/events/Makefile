regenerate:
	go install github.com/crx666/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/crx666/protobuf/protobuf/ \
	--gogo_out=. events.proto 

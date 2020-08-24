regenerate:
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/crxprotobuf/protobuf/protobuf/ \
	--gogo_out=. issue312.proto 
	(cd events && make regenerate)

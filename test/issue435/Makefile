regenerate:
	go install github.com/crx666/protobuf/protoc-gen-gogo
	protoc --gogo_out=. --proto_path=../../../../../:../../protobuf/:. *.proto

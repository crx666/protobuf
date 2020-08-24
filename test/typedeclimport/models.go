package typedeclimport

import subpkg "github.com/crxprotobuf/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}

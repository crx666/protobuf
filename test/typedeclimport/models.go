package typedeclimport

import subpkg "github.com/crx666/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}

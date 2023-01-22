package stub

import (
	"template/utils/stub"
)

func MakeStub() {
	stub.MakeStubs()
}

func GenerateFromStub(module string) {
	stub.Stubs(module)
}

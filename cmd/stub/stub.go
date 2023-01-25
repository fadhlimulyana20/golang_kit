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

func GenerateFromTemplateStub(templateType string, templateName string, name string) {
	stub.TemplateStub(templateType, templateName, name)
}

package template

import (
	"bytes"
	"html/template"
	"path"
)

func ProcessHtmlTemplateToStr(templateFile string, data map[string]interface{}) (string, error) {
	filepath := path.Join(templateFile)
	template, err := template.ParseFiles(filepath)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	template.Execute(&buffer, data)

	templateStr := buffer.String()
	return templateStr, nil
}

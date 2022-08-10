package utils

import (
	"bytes"
	"html/template"
	"log"
	"path"
)

func Parse(filepath string, data map[string]interface{}) string {
	fileNameWithSuffix := path.Base(filepath)
	t, err := template.New(fileNameWithSuffix).Funcs(template.FuncMap{
		"add": Add,
	}).ParseFiles(filepath)
	if err != nil {
		log.Print(err)
		return err.Error()
	}

	var tpl bytes.Buffer

	t.Execute(&tpl, data)
	// fmt.Print(tpl.String())
	return tpl.String()
}

func Add(x, y int) int {
	return x + y
}

package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"path"
)

func Parse(filepath string, data map[string]interface{}) string {
	fileNameWithSuffix := path.Base(filepath)
	t, err := template.New(fileNameWithSuffix).Funcs(template.FuncMap{
		"add": add,
	}).ParseFiles(filepath)
	if err != nil {
		log.Print(err)
		return err.Error()
	}

	var tpl bytes.Buffer

	t.Execute(&tpl, data)
	fmt.Print(tpl.String())
	return tpl.String()
}

func add(x, y int) int {
	return x + y
}

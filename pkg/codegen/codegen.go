package codegen

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"
)

type Input struct {
	ModuleName string
	Name       string
	LowerName  string
	TitleName  string
	Columns    []*Column
	Components string
}

type Column struct {
	Name    string
	Type    string
	Comment string
}

type Output struct {
	RawName  string
	FileName string
}

const (
	delimiterLeft  = "[["
	delimiterRight = "]]"
)

func GetInput(data interface{}, moduleName string) *Input {
	t := reflect.TypeOf(data).Elem()
	columns := []*Column{}
	componentMap := map[string]bool{}
	for i := 1; i < t.NumField(); i++ {
		if strings.Contains(string(t.Field(i).Tag), "comment") {
			if t.Field(i).Type.Name() == "Time" && !componentMap["DatePicker"] {
				componentMap["DatePicker"] = true
			}
			columns = append(columns, &Column{
				Name:    formatFieldName(t.Field(i).Name),
				Type:    t.Field(i).Type.Name(),
				Comment: strings.Split(string(t.Field(i).Tag), "'")[1],
			})
		}
	}
	components := make([]string, len(componentMap))
	i := 0
	for k := range componentMap {
		components[i] = k
		i++
	}
	sort.Strings(components)
	return &Input{
		ModuleName: moduleName,
		Name:       t.Name(),
		LowerName:  strings.ToLower(t.Name()),
		TitleName:  strings.ToTitle(t.Name()),
		Columns:    columns,
		Components: ", " + strings.Join(components, ", "),
	}
}

func formatFieldName(name string) string {
	if name == "IP" {
		return "ip"
	}
	return strings.ToLower(string(name[0])) + string(name[1:])
}

func GenerateCode(input *Input, tplPath, outputPath string) {
	files := []*Output{}
	err := filepath.Walk(tplPath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			fileName := strings.ReplaceAll(info.Name(), "Demo", input.Name)
			fileName = strings.ReplaceAll(fileName, "demo", input.LowerName)
			fileName = strings.TrimRight(fileName, ".tpl")
			files = append(files, &Output{
				RawName:  info.Name(),
				FileName: fileName,
			})
		}
		return nil
	})
	if err != nil {
		log.Println("walk filepath failed:", err)
	}

	for _, v := range files {
		log.Println(v.FileName)
		t, err := template.New(v.RawName).Delims(delimiterLeft, delimiterRight).ParseFiles(tplPath + "/" + v.RawName)
		if err != nil {
			log.Println("parse file failed:", err)
		}
		file, err := os.Create(outputPath + "/" + v.FileName)
		if err != nil {
			log.Println("create file failed:", err)
		}
		err = t.ExecuteTemplate(file, v.RawName, input)
		if err != nil {
			log.Println("execute failed:", err)
		}
	}
}

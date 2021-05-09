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
	Module           string
	ModuleName       string
	Name             string
	LowerName        string
	TitleName        string
	Columns          []*Column
	InputComponents  string
	ShowComponents   string
	CustomComponents string
	HasTime          string
}

type Column struct {
	Name    string
	Type    string
	TSType  string
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

var typeMap = map[string]string{
	"string":      "string",
	"StringArray": "string[]",
	"uint":        "number",
	"int":         "number",
	"int8":        "number",
	"float32":     "number",
	"float64":     "number",
	"Time":        "dayjs.Dayjs",
}

func GetInput(data interface{}, module, moduleName string) *Input {
	t := reflect.TypeOf(data).Elem()
	columns := []*Column{}
	inputComponentMap := map[string]bool{}
	customComponentMap := map[string]bool{}
	showComponentsMap := map[string]bool{}
	importDayjs := ""
	for i := 1; i < t.NumField(); i++ {
		if strings.Contains(string(t.Field(i).Tag), "comment") {
			tsType := typeMap[t.Field(i).Type.Name()]
			if tsType == "dayjs.Dayjs" && !customComponentMap["DatePicker"] {
				customComponentMap["DatePicker"] = true
				importDayjs = "import dayjs from 'dayjs'"
			} else if tsType == "string[]" && !inputComponentMap["Select"] {
				inputComponentMap["Select"] = true
			} else if tsType == "number" && !inputComponentMap["InputNumber"] {
				inputComponentMap["InputNumber"] = true
			}
			if tsType == "string[]" && !showComponentsMap["Tag"] {
				showComponentsMap["Tag"] = true
			}
			columns = append(columns, &Column{
				Name:    formatFieldName(t.Field(i).Name),
				Type:    t.Field(i).Type.Name(),
				TSType:  tsType,
				Comment: strings.Split(string(t.Field(i).Tag), "'")[1],
			})
		}
	}
	inputComponents := make([]string, len(inputComponentMap))
	customComponents := make([]string, len(customComponentMap))
	showComponents := make([]string, len(showComponentsMap))
	i := 0
	for k := range inputComponentMap {
		inputComponents[i] = k
		i++
	}
	i = 0
	for k := range showComponentsMap {
		showComponents[i] = k
		i++
	}
	i = 0
	for k := range customComponentMap {
		customComponents[i] = k
		i++
	}
	sort.Strings(inputComponents)
	sort.Strings(showComponents)
	sort.Strings(customComponents)
	cc := ""
	if len(customComponents) > 0 {
		cc = "\nimport { " + strings.Join(customComponents, ", ") + " } from 'src/components'"
	}
	return &Input{
		Module:           module,
		ModuleName:       moduleName,
		Name:             t.Name(),
		LowerName:        strings.ToLower(t.Name()),
		TitleName:        strings.ToTitle(t.Name()),
		Columns:          columns,
		InputComponents:  ", " + strings.Join(inputComponents, ", "),
		ShowComponents:   ", " + strings.Join(showComponents, ", "),
		CustomComponents: cc,
		HasTime:          importDayjs,
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
		if v.RawName == "form.tpl" {
			continue
		}
		log.Println(v.FileName)
		parseFiles := []string{tplPath + "/" + v.RawName}
		if strings.Contains(v.FileName, "Form") {
			parseFiles = append(parseFiles, tplPath+"/"+"form.tpl")
		}
		t, err := template.New(v.RawName).Delims(delimiterLeft, delimiterRight).
			ParseFiles(parseFiles...)
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

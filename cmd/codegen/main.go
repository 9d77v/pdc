package main

import (
	"fmt"
	"os"

	"github.com/9d77v/pdc/internal/module/book-service/models"
	"github.com/9d77v/pdc/pkg/codegen"
)

func main() {
	input := codegen.GetInput(&models.Book{}, "书籍")
	generateGqls(input)
	generateModel(input)
	generateView(input)
}

func generateGqls(input *codegen.Input) {
	outPutPath := fmt.Sprintf("web/src/gqls/%s", input.LowerName)
	os.MkdirAll(outPutPath, os.ModePerm)
	codegen.GenerateCode(input, "configs/tpls/ui/gqls", outPutPath)
}

func generateModel(input *codegen.Input) {
	outPutPath := fmt.Sprintf("web/src/module/%s", input.LowerName)
	os.MkdirAll(outPutPath, os.ModePerm)
	codegen.GenerateCode(input, "configs/tpls/ui/module", outPutPath)
}

func generateView(input *codegen.Input) {
	outPutPath := fmt.Sprintf("web/src/profiles/desktop/admin/%s/%s-list", input.LowerName, input.LowerName)
	os.MkdirAll(outPutPath, os.ModePerm)
	codegen.GenerateCode(input, "configs/tpls/ui/view", outPutPath)
}

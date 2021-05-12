package main

import (
	"fmt"
	"os"

	"github.com/9d77v/pdc/internal/module/book-service/models"
	"github.com/9d77v/pdc/pkg/codegen"
)

func main() {
	input := codegen.GetInput(&models.Bookshelf{}, "book", "书架")
	generateGqls(input)
	generateModel(input)
	generateView(input)
}

func generateGqls(input *codegen.Input) {
	outPutPath := fmt.Sprintf("web/src/gqls/%s", input.Module)
	os.MkdirAll(outPutPath, os.ModePerm)
	codegen.GenerateCode(input, "configs/tpls/ui/gqls", outPutPath)
}

func generateModel(input *codegen.Input) {
	outPutPath := fmt.Sprintf("web/src/module/%s", input.Module)
	os.MkdirAll(outPutPath, os.ModePerm)
	codegen.GenerateCode(input, "configs/tpls/ui/module", outPutPath)
}

func generateView(input *codegen.Input) {
	outPutPath := fmt.Sprintf("web/src/profiles/desktop/admin/%s/%s-list", input.Module, input.LowerName)
	os.MkdirAll(outPutPath, os.ModePerm)
	codegen.GenerateCode(input, "configs/tpls/ui/view", outPutPath)
}

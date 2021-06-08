package main

import (
	"fmt"
	"os"

	"github.com/9d77v/pdc/internal/module/book-service/models"
)

func main() {
	input := GetInput(&models.Bookshelf{}, "book", "书架")
	generateGqls(input)
	generateModel(input)
	generateView(input)
}

func generateGqls(input *Input) {
	outPutPath := fmt.Sprintf("web/src/gqls/%s", input.Module)
	os.MkdirAll(outPutPath, os.ModePerm)
	GenerateCode(input, "configs/tpls/ui/gqls", outPutPath)
}

func generateModel(input *Input) {
	outPutPath := fmt.Sprintf("web/src/module/%s", input.Module)
	os.MkdirAll(outPutPath, os.ModePerm)
	GenerateCode(input, "configs/tpls/ui/module", outPutPath)
}

func generateView(input *Input) {
	outPutPath := fmt.Sprintf("web/src/profiles/desktop/admin/%s/%s-list", input.Module, input.LowerName)
	os.MkdirAll(outPutPath, os.ModePerm)
	GenerateCode(input, "configs/tpls/ui/view", outPutPath)
}

package main

import (
	"os"

	"github.com/9d77v/pdc/internal/module/user-service/models"
	"github.com/9d77v/pdc/pkg/codegen"
)

func main() {
	input := codegen.GetInput(&models.User{}, "用户")
	generateView(input)
}

func generateView(input *codegen.Input) {
	outPutPath := "web/src/profiles/desktop/admin/" + input.LowerName
	os.MkdirAll(outPutPath, os.ModePerm)
	codegen.GenerateCode(input, "configs/tpls/ui/view", outPutPath)
}

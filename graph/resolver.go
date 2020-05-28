package graph

import "git.9d77v.me/9d77v/pdc/services"

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver ..
type Resolver struct{}

var (
	videoService  = services.VideoService{}
	commonService = services.CommonService{}
)
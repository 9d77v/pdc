package graph

import "github.com/9d77v/pdc/services"

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver ..
type Resolver struct{}

var (
	videoService   = services.VideoService{}
	commonService  = services.CommonService{}
	thingService   = services.ThingService{}
	userService    = services.UserService{}
	historyService = services.HistoryService{}
	deviceService  = services.DeviceService{}
)

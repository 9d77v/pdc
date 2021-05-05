package graph

import (
	book "github.com/9d77v/pdc/internal/module/book-service/services"
	device "github.com/9d77v/pdc/internal/module/device-service/services"
	history "github.com/9d77v/pdc/internal/module/history-service/services"
	note "github.com/9d77v/pdc/internal/module/note-service/services"
	thing "github.com/9d77v/pdc/internal/module/thing-service/services"
	user "github.com/9d77v/pdc/internal/module/user-service/services"
	video "github.com/9d77v/pdc/internal/module/video-service/services"
)

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver ..
type Resolver struct{}

var (
	videoService           = video.VideoService{}
	thingService           = thing.ThingService{}
	userService            = user.UserService{}
	historyService         = history.HistoryService{}
	deviceService          = device.DeviceService{}
	deviceModelService     = device.DeviceModelService{}
	deviceDashboardService = device.DeviceDashboardService{}
	noteService            = note.NoteService{}
	bookService            = book.BookService{}
)

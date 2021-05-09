package middleware

import (
	"github.com/9d77v/pdc/internal/consts"
	device "github.com/9d77v/pdc/internal/module/device-service/services"

	user "github.com/9d77v/pdc/internal/module/user-service/services"
)

var userCtxKey = &contextKey{"user"}
var schemeCtxKey = &contextKey{"scheme"}

type contextKey struct {
	name string
}

var (
	userService   = user.UserService{}
	deviceService = device.DeviceService{}
)

var publicOperationArr = []string{"login", "refreshToken", "IntrospectionQuery"}
var permissonMap = map[string][]int{
	"presignedUrl":        {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},
	"users":               {consts.RoleAdmin},
	"userInfo":            {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"videos":              {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"videoDetail":         {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"videoSerieses":       {consts.RoleAdmin, consts.RoleManager},
	"historyStatistic":    {consts.RoleAdmin, consts.RoleManager},
	"appHistoryStatistic": {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"searchVideo":         {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"similarVideos":       {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"things":              {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},
	"thingSeries":         {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},
	"thingAnalyze":        {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},
	"books":               {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},
	"bookShelfs":          {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},
	"bookPositions":       {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},
	"bookBorrowReturn":    {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},

	"createUser":     {consts.RoleAdmin},
	"updateUser":     {consts.RoleAdmin},
	"updateProfile":  {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"updatePassword": {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},

	"createVideo":           {consts.RoleAdmin, consts.RoleManager},
	"addVideoResource":      {consts.RoleAdmin, consts.RoleManager},
	"saveSubtitles":         {consts.RoleAdmin, consts.RoleManager},
	"updateVideo":           {consts.RoleAdmin, consts.RoleManager},
	"createEpisode":         {consts.RoleAdmin, consts.RoleManager},
	"updateEpisode":         {consts.RoleAdmin, consts.RoleManager},
	"updateMobileVideo":     {consts.RoleAdmin, consts.RoleManager},
	"createVideoSeries":     {consts.RoleAdmin, consts.RoleManager},
	"updateVideoSeries":     {consts.RoleAdmin, consts.RoleManager},
	"createVideoSeriesItem": {consts.RoleAdmin, consts.RoleManager},
	"updateVideoSeriesItem": {consts.RoleAdmin, consts.RoleManager},

	"createThing": {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},
	"updateThing": {consts.RoleAdmin, consts.RoleManager, consts.RoleUser},

	"recordHistory":                  {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"histories":                      {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"createDeviceModel":              {consts.RoleAdmin, consts.RoleManager},
	"updateDeviceModel":              {consts.RoleAdmin, consts.RoleManager},
	"createAttributeModel":           {consts.RoleAdmin, consts.RoleManager},
	"updateAttributeModel":           {consts.RoleAdmin, consts.RoleManager},
	"deleteAttributeModel":           {consts.RoleAdmin, consts.RoleManager},
	"createTelemetryModel":           {consts.RoleAdmin, consts.RoleManager},
	"updateTelemetryModel":           {consts.RoleAdmin, consts.RoleManager},
	"deleteTelemetryModel":           {consts.RoleAdmin, consts.RoleManager},
	"deviceModels":                   {consts.RoleAdmin, consts.RoleManager},
	"createDevice":                   {consts.RoleAdmin, consts.RoleManager},
	"updateDevice":                   {consts.RoleAdmin, consts.RoleManager},
	"devices":                        {consts.RoleAdmin, consts.RoleManager},
	"createDeviceDashboard":          {consts.RoleAdmin, consts.RoleManager},
	"updateDeviceDashboard":          {consts.RoleAdmin, consts.RoleManager},
	"deleteDeviceDashboard":          {consts.RoleAdmin, consts.RoleManager},
	"addDeviceDashboardTelemetry":    {consts.RoleAdmin, consts.RoleManager},
	"removeDeviceDashboardTelemetry": {consts.RoleAdmin, consts.RoleManager},
	"deviceDashboards":               {consts.RoleAdmin, consts.RoleManager},
	"appDeviceDashboards":            {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"addDeviceDashboardCamera":       {consts.RoleAdmin, consts.RoleManager},
	"removeDeviceDashboardCamera":    {consts.RoleAdmin, consts.RoleManager},
	"cameraCapture":                  {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},
	"cameraTimeLapseVideos":          {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},

	"syncNotes": {consts.RoleAdmin, consts.RoleManager, consts.RoleUser, consts.RoleGuest},

	"createBookShelf":    {consts.RoleAdmin, consts.RoleManager},
	"updateBookShelf":    {consts.RoleAdmin, consts.RoleManager},
	"createBook":         {consts.RoleAdmin, consts.RoleManager},
	"updateBook":         {consts.RoleAdmin, consts.RoleManager},
	"createBookPosition": {consts.RoleAdmin, consts.RoleManager},
	"updateBookPosition": {consts.RoleAdmin, consts.RoleManager},
	"borrowBook":         {consts.RoleAdmin, consts.RoleManager},
	"backBook":           {consts.RoleAdmin, consts.RoleManager},
}

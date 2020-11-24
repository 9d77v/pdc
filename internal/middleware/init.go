package middleware

import (
	"github.com/9d77v/pdc/internal/db"
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
	"presignedUrl":                   {db.RoleAdmin, db.RoleManager, db.RoleUser},
	"users":                          {db.RoleAdmin},
	"userInfo":                       {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"videos":                         {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"videoSerieses":                  {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"searchVideo":                    {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"similarVideos":                  {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"things":                         {db.RoleAdmin, db.RoleManager, db.RoleUser},
	"thingSeries":                    {db.RoleAdmin, db.RoleManager, db.RoleUser},
	"thingAnalyze":                   {db.RoleAdmin, db.RoleManager, db.RoleUser},
	"createUser":                     {db.RoleAdmin},
	"updateUser":                     {db.RoleAdmin},
	"updateProfile":                  {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"updatePassword":                 {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"createVideo":                    {db.RoleAdmin, db.RoleManager},
	"addVideoResource":               {db.RoleAdmin, db.RoleManager},
	"saveSubtitles":                  {db.RoleAdmin, db.RoleManager},
	"updateVideo":                    {db.RoleAdmin, db.RoleManager},
	"createEpisode":                  {db.RoleAdmin, db.RoleManager},
	"updateEpisode":                  {db.RoleAdmin, db.RoleManager},
	"updateMobileVideo":              {db.RoleAdmin, db.RoleManager},
	"createVideoSeries":              {db.RoleAdmin, db.RoleManager},
	"updateVideoSeries":              {db.RoleAdmin, db.RoleManager},
	"createVideoSeriesItem":          {db.RoleAdmin, db.RoleManager},
	"updateVideoSeriesItem":          {db.RoleAdmin, db.RoleManager},
	"createThing":                    {db.RoleAdmin, db.RoleManager, db.RoleUser},
	"updateThing":                    {db.RoleAdmin, db.RoleManager, db.RoleUser},
	"recordHistory":                  {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"historyInfo":                    {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"histories":                      {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"createDeviceModel":              {db.RoleAdmin, db.RoleManager},
	"updateDeviceModel":              {db.RoleAdmin, db.RoleManager},
	"createAttributeModel":           {db.RoleAdmin, db.RoleManager},
	"updateAttributeModel":           {db.RoleAdmin, db.RoleManager},
	"deleteAttributeModel":           {db.RoleAdmin, db.RoleManager},
	"createTelemetryModel":           {db.RoleAdmin, db.RoleManager},
	"updateTelemetryModel":           {db.RoleAdmin, db.RoleManager},
	"deleteTelemetryModel":           {db.RoleAdmin, db.RoleManager},
	"deviceModels":                   {db.RoleAdmin, db.RoleManager},
	"createDevice":                   {db.RoleAdmin, db.RoleManager},
	"updateDevice":                   {db.RoleAdmin, db.RoleManager},
	"devices":                        {db.RoleAdmin, db.RoleManager},
	"createDeviceDashboard":          {db.RoleAdmin, db.RoleManager},
	"updateDeviceDashboard":          {db.RoleAdmin, db.RoleManager},
	"deleteDeviceDashboard":          {db.RoleAdmin, db.RoleManager},
	"addDeviceDashboardTelemetry":    {db.RoleAdmin, db.RoleManager},
	"removeDeviceDashboardTelemetry": {db.RoleAdmin, db.RoleManager},
	"deviceDashboards":               {db.RoleAdmin, db.RoleManager},
	"appDeviceDashboards":            {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
	"addDeviceDashboardCamera":       {db.RoleAdmin, db.RoleManager},
	"removeDeviceDashboardCamera":    {db.RoleAdmin, db.RoleManager},
	"cameraCapture":                  {db.RoleAdmin, db.RoleManager, db.RoleUser, db.RoleGuest},
}

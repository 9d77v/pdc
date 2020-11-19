package middleware

import (
	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/services"
)

var userCtxKey = &contextKey{"user"}
var schemeCtxKey = &contextKey{"scheme"}

type contextKey struct {
	name string
}

var (
	userService   = services.UserService{}
	deviceService = services.DeviceService{}
)

var publicOperationArr = []string{"login", "refreshToken", "IntrospectionQuery"}
var permissonMap = map[string][]int{
	"presignedUrl":                   {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"users":                          {models.RoleAdmin},
	"userInfo":                       {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"videos":                         {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"videoSerieses":                  {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"searchVideo":                    {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"similarVideos":                  {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"things":                         {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"thingSeries":                    {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"thingAnalyze":                   {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"createUser":                     {models.RoleAdmin},
	"updateUser":                     {models.RoleAdmin},
	"updateProfile":                  {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"updatePassword":                 {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"createVideo":                    {models.RoleAdmin, models.RoleManager},
	"addVideoResource":               {models.RoleAdmin, models.RoleManager},
	"saveSubtitles":                  {models.RoleAdmin, models.RoleManager},
	"updateVideo":                    {models.RoleAdmin, models.RoleManager},
	"createEpisode":                  {models.RoleAdmin, models.RoleManager},
	"updateEpisode":                  {models.RoleAdmin, models.RoleManager},
	"updateMobileVideo":              {models.RoleAdmin, models.RoleManager},
	"createVideoSeries":              {models.RoleAdmin, models.RoleManager},
	"updateVideoSeries":              {models.RoleAdmin, models.RoleManager},
	"createVideoSeriesItem":          {models.RoleAdmin, models.RoleManager},
	"updateVideoSeriesItem":          {models.RoleAdmin, models.RoleManager},
	"createThing":                    {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"updateThing":                    {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"recordHistory":                  {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"historyInfo":                    {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"histories":                      {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"createDeviceModel":              {models.RoleAdmin, models.RoleManager},
	"updateDeviceModel":              {models.RoleAdmin, models.RoleManager},
	"createAttributeModel":           {models.RoleAdmin, models.RoleManager},
	"updateAttributeModel":           {models.RoleAdmin, models.RoleManager},
	"deleteAttributeModel":           {models.RoleAdmin, models.RoleManager},
	"createTelemetryModel":           {models.RoleAdmin, models.RoleManager},
	"updateTelemetryModel":           {models.RoleAdmin, models.RoleManager},
	"deleteTelemetryModel":           {models.RoleAdmin, models.RoleManager},
	"deviceModels":                   {models.RoleAdmin, models.RoleManager},
	"createDevice":                   {models.RoleAdmin, models.RoleManager},
	"updateDevice":                   {models.RoleAdmin, models.RoleManager},
	"devices":                        {models.RoleAdmin, models.RoleManager},
	"createDeviceDashboard":          {models.RoleAdmin, models.RoleManager},
	"updateDeviceDashboard":          {models.RoleAdmin, models.RoleManager},
	"deleteDeviceDashboard":          {models.RoleAdmin, models.RoleManager},
	"addDeviceDashboardTelemetry":    {models.RoleAdmin, models.RoleManager},
	"removeDeviceDashboardTelemetry": {models.RoleAdmin, models.RoleManager},
	"deviceDashboards":               {models.RoleAdmin, models.RoleManager},
	"appDeviceDashboards":            {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"addDeviceDashboardCamera":       {models.RoleAdmin, models.RoleManager},
	"removeDeviceDashboardCamera":    {models.RoleAdmin, models.RoleManager},
	"cameraCapture":                  {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
}

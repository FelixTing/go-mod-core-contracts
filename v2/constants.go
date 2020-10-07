//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

// Constants related to defined routes in the v2 service APIs
const (
	ApiVersion                 = "v2"
	ApiBase                    = "/api/v2"
	ApiEventRoute              = ApiBase + "/event"
	ApiAllEventRoute           = ApiEventRoute + "/" + All
	ApiEventIdRoute            = ApiEventRoute + "/" + Id + "/{" + Id + "}"
	ApiEventPushRoute          = ApiEventRoute + "/" + Pushed
	ApiEventCountRoute         = ApiEventRoute + "/" + Count
	ApiEventCountByDeviceRoute = ApiEventCountRoute + "/" + Device + "/{" + DeviceIdParam + "}"
	ApiEventByDeviceRoute      = ApiEventRoute + "/" + Device + "/{" + DeviceIdParam + "}"
	ApiAllEventByDeviceRoute   = ApiEventByDeviceRoute + "/" + All
	ApiEventByTimeRangeRoute   = ApiEventRoute + "/" + Start + "/{" + Start + "}/" + End + "/{" + End + "}"
	ApiEventByAgeRoute         = ApiEventRoute + "/" + Age + "/{" + Age + "}"
	ApiEventScrubRoute         = ApiEventRoute + "/" + Scrub
	ApiReadingRoute            = ApiBase + "/reading"
	ApiAllReadingRoute         = ApiReadingRoute + "/" + All
	ApiReadingCountRoute       = ApiReadingRoute + "/" + Count
	ApiReadingIdRoute          = ApiReadingRoute + "/" + Id + "/{" + Id + "}"
	ApiAllReadingByDeviceRoute = ApiReadingRoute + "/" + Device + "/{" + DeviceIdParam + "}/" + All
	ApiReadingByTypeRoute      = ApiReadingRoute + "/" + Type + "/{" + Type + "}"
	ApiReadingByTimeRangeRoute = ApiReadingRoute + "/" + Start + "/{" + Start + "}/" + End + "/{" + End + "}"

	ApiDeviceProfileRoute                       = ApiBase + "/deviceprofile"
	ApiDeviceProfileUploadFileRoute             = ApiDeviceProfileRoute + "/uploadfile"
	ApiDeviceProfileByNameRoute                 = ApiDeviceProfileRoute + "/" + Name + "/{" + Name + "}"
	ApiDeviceProfileByIdRoute                   = ApiDeviceProfileRoute + "/" + Id + "/{" + Id + "}"
	ApiAllDeviceProfileRoute                    = ApiDeviceProfileRoute + "/" + All
	ApiDeviceProfileByManufacturerRoute         = ApiDeviceProfileRoute + "/" + Manufacturer + "/{" + Manufacturer + "}"
	ApiDeviceProfileByModelRoute                = ApiDeviceProfileRoute + "/" + Model + "/{" + Model + "}"
	ApiDeviceProfileByManufacturerAndModelRoute = ApiDeviceProfileRoute + "/" + Manufacturer + "/{" + Manufacturer + "}" + "/" + Model + "/{" + Model + "}"

	ApiDeviceRoute        = ApiBase + "/device"
	ApiDeviceServiceRoute = ApiBase + "/deviceservice"
	ApiConfigRoute        = ApiBase + "/config"
	ApiMetricsRoute       = ApiBase + "/metrics"
	ApiPingRoute          = ApiBase + "/ping"
	ApiVersionRoute       = ApiBase + "/version"
)

// Constants related to defined url path names and parameters in the v2 service APIs
const (
	All           = "all"
	Id            = "id"
	Pushed        = "pushed"
	Count         = "count"
	Device        = "device"
	DeviceIdParam = "deviceId"
	Start         = "start"
	End           = "end"
	Age           = "age"
	Scrub         = "scrub"
	Type          = "type"
	Name          = "name"
	Manufacturer  = "manufacturer"
	Model         = "model"
)

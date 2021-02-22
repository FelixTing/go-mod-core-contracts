//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// Transmission and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/Transmission
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type Transmission struct {
	Description string
	Channel Channel
	Created int64
	Id string
	Notification Notification
	SubscriptionName string
	Records []TransmissionRecord
	ResendCount int64
	Status TransmissionStatus
}

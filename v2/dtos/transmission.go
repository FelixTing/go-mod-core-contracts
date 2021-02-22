//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// Transmission and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/Transmission
type Transmission struct {
	common.Versionable `json:",inline"`
	Id                 string    `json:"id,omitempty" validate:"omitempty,uuid"`
	Created            int64     `json:"created,omitempty"`
	Description        string    `json:"description,omitempty"`
	Channel            Channel   `json:"channel" validate:"required"`
	Notification       Notification `json:"notification" validate:"required"`
	SubscriptionName   string    `json:"subscriptionName" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Records            []TransmissionRecord `json:"records,omitempty"`
	ResendCount        int64       `json:"resendCount,omitempty"`
	Status             string    `json:"status" validate:"required,oneof='ACKNOWLEDGED' 'FAILED' 'SENT' 'TRXESCALATED'"`
}

// ToTransmissionModel transforms a Transmission DTO to a Transmission Model
func ToTransmissionModel(t Transmission) models.Transmission {
	var m models.Transmission
	m.Description = t.Description
	m.Channel = ToChannelModel(t.Channel)
	m.Created = t.Created
	m.Id = t.Id
	m.Notification = ToNotificationModel(t.Notification)
	m.SubscriptionName = t.SubscriptionName
	m.Records = ToTransmissionRecordModels(t.Records)
	m.ResendCount = t.ResendCount
	m.Status = models.TransmissionStatus(t.Status)
	return m
}

// ToTransmissionModels transforms a Transmission DTO array to a Transmission model array
func ToTransmissionModels(ts []Transmission) []models.Transmission {
	models := make([]models.Transmission, len(ts))
	for i, t := range ts {
		models[i] = ToTransmissionModel(t)
	}
	return models
}

// FromTransmissionModelToDTO transforms the Transmission Model to the Transmission DTO
func FromTransmissionModelToDTO(t models.Transmission) Transmission {
	return Transmission{
		Id: t.Id,
		Created:        t.Created,
		Description:    t.Description,
		Channel: FromChannelModelToDTO(t.Channel),
		Notification: FromNotificationModelToDTO(t.Notification),
		SubscriptionName: t.SubscriptionName,
		Records: FromTransmissionRecordModelsToDTOs(t.Records),
		ResendCount: t.ResendCount,
		Status: string(t.Status),
	}
}

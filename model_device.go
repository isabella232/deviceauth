// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package main

import (
	"time"
)

const (
	DevStatusAccepted = "accepted"
	DevStatusRejected = "rejected"
	DevStatusPending  = "pending"
)

// note: fields with underscores need the 'bson' decorator
// otherwise the underscore will be removed upon write to mongo
type Device struct {
	Id          string    `json:"id" bson:"_id"`
	TenantToken string    `json:"tenant_token" bson:"tenant_token"`
	PubKey      string    `json:"pubkey"`
	IdData      string    `json:"id_data" bson:"id_data"`
	Status      string    `json:"status"`
	CreatedTs   time.Time `json:"created_ts" bson:"created_ts"`
	UpdatedTs   time.Time `json:"updated_ts" bson:"updated_ts"`
}

func NewDevice(id, id_data, pubkey, tenant_token string) *Device {
	now := time.Now()

	return &Device{
		Id:          id,
		IdData:      id_data,
		TenantToken: tenant_token,
		PubKey:      pubkey,
		Status:      DevStatusPending,
		CreatedTs:   now,
		UpdatedTs:   now,
	}
}

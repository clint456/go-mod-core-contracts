//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type NodeService struct {
	DBTimestamp
	Id          string
	Name        string
	Description string
	Labels      []string
	BaseAddress string
	AdminState  AdminState
	Properties  map[string]any
}

//
// Copyright (C) 2020-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

import "maps"

type Node struct {
	DBTimestamp
	Id             string
	Name           string
	Parent         string
	Description    string
	AdminState     AdminState
	OperatingState OperatingState
	Protocols      map[string]ProtocolProperties
	Labels         []string
	Location       interface{}
	ServiceName    string
	ProfileName    string
	AutoEvents     []AutoEvent
	Tags           map[string]any
	Properties     map[string]any
}

// ProtocolProperties contains the node connection information in key/value pair
type ProtocolProperties map[string]any

func (p ProtocolProperties) Clone() ProtocolProperties {
	cloned := make(map[string]any)
	maps.Copy(cloned, p)
	return cloned
}

// AdminState controls the range of values which constitute valid administrative states for a node
type AdminState string

// AssignAdminState provides a default value "UNLOCKED" to AdminState
func AssignAdminState(dtoAdminState string) AdminState {
	if dtoAdminState == "" {
		return AdminState(Unlocked)
	}
	return AdminState(dtoAdminState)
}

// OperatingState is an indication of the operations of the node.
type OperatingState string

func (node Node) Clone() Node {
	cloned := Node{
		DBTimestamp:    node.DBTimestamp,
		Id:             node.Id,
		Name:           node.Name,
		Parent:         node.Parent,
		Description:    node.Description,
		AdminState:     node.AdminState,
		OperatingState: node.OperatingState,
		Location:       node.Location,
		ServiceName:    node.ServiceName,
		ProfileName:    node.ProfileName,
	}
	if len(node.Protocols) > 0 {
		cloned.Protocols = make(map[string]ProtocolProperties)
		for k, v := range node.Protocols {
			cloned.Protocols[k] = v.Clone()
		}
	}
	if len(node.Labels) > 0 {
		cloned.Labels = make([]string, len(node.Labels))
		copy(cloned.Labels, node.Labels)
	}
	if len(node.AutoEvents) > 0 {
		cloned.AutoEvents = make([]AutoEvent, len(node.AutoEvents))
		copy(cloned.AutoEvents, node.AutoEvents)
	}
	if len(node.Tags) > 0 {
		cloned.Tags = make(map[string]any)
		maps.Copy(cloned.Tags, node.Tags)
	}
	if len(node.Properties) > 0 {
		cloned.Properties = make(map[string]any)
		maps.Copy(cloned.Properties, node.Properties)
	}
	return cloned
}

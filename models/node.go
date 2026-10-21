//
// Copyright (C) 2020-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

import "maps"

type Node struct {
	DBTimestamp
	Id             string                        // 唯一ID
	Name           string                        // Node名称
	Parent         string                        // 父节点ID
	Description    string                        // Node描述
	AdminState     AdminState                    // 管理状态
	OperatingState OperatingState                // 运行状态
	Protocols      map[string]ProtocolProperties // 节点连接信息
	Labels         []string                      // 标签
	Direction      string                        // 南北向 "SOUTH" or "NORTH",这里注视规定，业务层进行南北向的区分，SOUTH表示南向，NORTH表示北向
	Location       interface{}                   // 位置信息
	ServiceName    string                        // 服务名称
	ProfileName    string                        // Profile名称
	AutoEvents     []AutoEvent                   // 自动事件
	Tags           map[string]any                // 标签
	Properties     map[string]any                // 属性
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
		Direction:      node.Direction,
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

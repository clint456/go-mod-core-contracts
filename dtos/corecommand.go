//
// Copyright (C) 2021-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"fmt"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

type NodeCoreCommand struct {
	NodeName     string        `json:"nodeName" validate:"required,edgex-dto-none-empty-string"`
	ProfileName  string        `json:"profileName" validate:"required,edgex-dto-none-empty-string"`
	CoreCommands []CoreCommand `json:"coreCommands,omitempty" validate:"dive"`
}

type CoreCommand struct {
	Name       string                 `json:"name" validate:"required,edgex-dto-none-empty-string"`
	Get        bool                   `json:"get,omitempty" validate:"required_without=Set"`
	Set        bool                   `json:"set,omitempty" validate:"required_without=Get"`
	Path       string                 `json:"path,omitempty"`
	Url        string                 `json:"url,omitempty"`
	Parameters []CoreCommandParameter `json:"parameters,omitempty"`
}

type CoreCommandParameter struct {
	ResourceName string `json:"resourceName"`
	ValueType    string `json:"valueType"`
}

func BuildCoreCommands(nodeName string, serviceUrl string, profile Profile) ([]CoreCommand, errors.EdgeX) {
	commandMap := make(map[string]CoreCommand)
	// Build commands from node commands
	for _, c := range profile.Commands {
		if c.IsHidden {
			continue
		}
		parameters, err := coreCommandParameters(c.ResourceOperations, profile.Resources)
		if err != nil {
			return nil, errors.NewCommonEdgeXWrapper(err)
		}
		commandMap[c.Name] = buildCoreCommand(nodeName, serviceUrl, c.Name, c.ReadWrite, parameters)
	}
	// Build commands from node resource
	for _, r := range profile.Resources {
		if _, ok := commandMap[r.Name]; ok || r.IsHidden {
			continue
		}
		parameters := []CoreCommandParameter{
			{ResourceName: r.Name, ValueType: r.Properties.ValueType},
		}
		commandMap[r.Name] = buildCoreCommand(nodeName, serviceUrl, r.Name, r.Properties.ReadWrite, parameters)
	}
	// Convert command map to slice
	var commands []CoreCommand
	for _, cmd := range commandMap {
		commands = append(commands, cmd)
	}
	return commands, nil
}

// coreCommandParameters creates command parameters by mapping the resourceOperation to corresponding resourceName and valueType
func coreCommandParameters(resourceOperations []ResourceOperation, resources []Resource) ([]CoreCommandParameter, errors.EdgeX) {
	parameters := make([]CoreCommandParameter, len(resourceOperations))
	for i, ro := range resourceOperations {
		r, exists := resourcesByName(resources, ro.Resource)
		if !exists {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("node command's resource %s doesn't match any deivce resource", ro.Resource), nil)
		}
		parameters[i] = CoreCommandParameter{
			ResourceName: r.Name,
			ValueType:    r.Properties.ValueType,
		}
	}
	return parameters, nil
}

func commandPath(nodeName, cmdName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", common.ApiNodeRoute, common.Name, nodeName, cmdName)
}

func buildCoreCommand(nodeName, serviceUrl, cmdName, readWrite string, parameters []CoreCommandParameter) CoreCommand {
	cmd := CoreCommand{
		Name:       cmdName,
		Url:        serviceUrl,
		Path:       commandPath(nodeName, cmdName),
		Parameters: parameters,
	}
	if strings.Contains(readWrite, common.ReadWrite_R) {
		cmd.Get = true
	}
	if strings.Contains(readWrite, common.ReadWrite_W) {
		cmd.Set = true
	}
	return cmd
}

func resourcesByName(resources []Resource, name string) (res Resource, exists bool) {
	for _, resource := range resources {
		if resource.Name == name {
			exists = true
			res = resource
			break
		}
	}
	return res, exists
}

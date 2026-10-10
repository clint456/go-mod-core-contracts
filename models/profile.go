//
// Copyright (C) 2020-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type Profile struct {
	DBTimestamp
	ApiVersion   string
	Description  string
	Id           string
	Name         string
	Manufacturer string
	Model        string
	Labels       []string
	Resources    []Resource
	Commands     []Command
}

func (profile Profile) Clone() Profile {
	cloned := Profile{
		DBTimestamp:  profile.DBTimestamp,
		ApiVersion:   profile.ApiVersion,
		Description:  profile.Description,
		Id:           profile.Id,
		Name:         profile.Name,
		Manufacturer: profile.Manufacturer,
		Model:        profile.Model,
	}
	if len(profile.Labels) > 0 {
		cloned.Labels = make([]string, len(profile.Labels))
		copy(cloned.Labels, profile.Labels)
	}
	if len(profile.Resources) > 0 {
		cloned.Resources = make([]Resource, len(profile.Resources))
		for i := range profile.Resources {
			cloned.Resources[i] = profile.Resources[i].Clone()
		}
	}
	if len(profile.Commands) > 0 {
		cloned.Commands = make([]Command, len(profile.Commands))
		for i := range profile.Commands {
			cloned.Commands[i] = profile.Commands[i].Clone()
		}
	}
	return cloned
}

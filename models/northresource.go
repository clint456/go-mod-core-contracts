// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
package models

type NorthResource struct {
	Description string
	Name        string
	IsHidden    bool
	Properties  ResourceProperties
	Attributes  map[string]interface{}
	Tags        map[string]any
}

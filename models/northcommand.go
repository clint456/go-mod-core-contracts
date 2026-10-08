// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
package models

type NorthCommand struct {
	Name               string
	IsHidden           bool
	ReadWrite          string
	ResourceOperations []ResourceOperation
	Tags               map[string]any
}

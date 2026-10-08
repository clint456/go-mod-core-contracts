// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
package models

type NorthProfile struct {
	DBTimestamp
	ApiVersion     string
	Description    string
	Id             string
	Name           string
	Manufacturer   string
	Model          string
	Labels         []string
	NorthResources []NorthResource
	NorthCommands  []NorthCommand
}

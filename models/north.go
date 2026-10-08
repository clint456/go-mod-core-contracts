// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
package models

type North struct {
	DBTimestamp
	Id             string
	Name           string
	Description    string
	AdminState     AdminState
	OperatingState OperatingState
	Protocols      map[string]ProtocolProperties // 通道协议配置，key为协议名
	Labels         []string
	ServiceName    string
	ProfileName    string
	Tags           map[string]any // 通道静态业务属性，非协议配置
	Properties     map[string]any // 通道静态业务属性，非协议配置
}

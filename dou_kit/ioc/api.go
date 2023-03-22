// @Author: Ciusyan 2023/3/14

package ioc

type AppName string

const (
	USER_API     = AppName("user")
	VIDEO_API    = AppName("video")
	RELATION_API = AppName("relation")
)

// ApiOptions 注册IOC中Api服务的路由时，可传入配置
type ApiOptions struct {
	// API 前缀
	Prefix string `json:"prefix"`
	// API 是否需要添加版本
	NotVersion bool `json:"not_version"`
	// API 是否需要添加服务名称
	NotName bool `json:"not_name"`

	// ...
}

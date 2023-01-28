// @Author: Ciusyan 2023/1/28
package ioc

// ========
// 初始化所有依赖：InitAllDependencies
// ========

// InitAllDependencies 用于初始IoC容器中的所有依赖
func InitAllDependencies() error {
	// 初始化内部服务模块依赖
	for _, v := range internalContainer {
		if err := v.Init(); err != nil {
			return err
		}
	}
	// 初始化Gin HTTP服务依赖
	for _, v := range ginContainer {
		if err := v.Init(); err != nil {
			return err
		}
	}
	return nil
}

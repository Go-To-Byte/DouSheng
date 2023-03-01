// @Author: Ciusyan 2023/3/1
package utils

// V2P 将 value -> ptr
func V2P[T any](n T) *T {
	return &n
}

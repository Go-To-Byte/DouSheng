// @Author: Ciusyan 2023/2/23
package utils

type Set map[int64]struct{}

// NewSet 返回一个Set：map[int64]struct{}
func NewSet() Set {
	return map[int64]struct{}{}
}

// Add 添加元素
func (s Set) Add(item int64) {
	s[item] = struct{}{}
}

// Items 获取所有的元素
func (s Set) Items() (items []int64) {
	for k, _ := range s {
		items = append(items, k)
	}
	return
}

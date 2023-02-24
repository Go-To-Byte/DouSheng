// @Author: Ciusyan 2023/2/23
package utils_test

import (
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/utils"
)

func TestSet(t *testing.T) {
	set := utils.NewSet()
	set.Add(1)
	set.Add(3)
	set.Add(3)
	set.Add(2)
	set.Add(1)
	set.Add(2)
	items := set.Items()
	t.Log(items)
}

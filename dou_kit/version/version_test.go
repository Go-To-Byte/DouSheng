// @Author: Ciusyan 2023/2/7
package version_test

import (
	"fmt"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/version"
)

func TestFullVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}

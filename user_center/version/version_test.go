// @Author: Ciusyan 2023/1/25
package version_test

import (
	"fmt"
	"testing"

	"github.com/Go-To-Byte/DouSheng/user_center/version"
)

func TestFullVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}

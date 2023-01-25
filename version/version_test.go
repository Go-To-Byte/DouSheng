// @Author: Ciusyan 2023/1/25
package version_test

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/version"
	"testing"
)

func TestFullVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}

// @Author: Ciusyan 2023/2/23
package utils_test

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/dou_kit/utils"
	"reflect"
	"testing"
	"time"
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

type rpcImpl interface {
	func1(ctx context.Context, req int64) (resp int64, err error)
}
type rpc struct {
	request  int64
	response int64
}

func (r rpc) func1(ctx context.Context, req int64) (resp int64, err error) {
	fmt.Println("ok: ", req)
	r.request = req
	r.response = req * 10
	return req * 10, nil
}

func TestGoRunGrpc(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r := rpc{}
	v := reflect.ValueOf(r.func1)
	var args []reflect.Value
	args = append(args, reflect.ValueOf(ctx))
	args = append(args, reflect.ValueOf(int64(1)))
	c := v.Call(args)
	fmt.Println("ok: ", c[0].Type().String(), c[0].Interface().(int64))
}

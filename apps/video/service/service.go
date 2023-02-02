// Author: BeYoung
// Date: 2023/1/30 2:41
// Software: GoLand

package service

import (
	"github.com/Go-To-Byte/DouSheng/apps/video/proto"
	"net"
)

type Video struct {
	proto.UnimplementedVideoServer
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

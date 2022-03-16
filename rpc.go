package sys_rpclib

import (
	"github.com/cloudwego/kitex/client"
	"sys-rpclib/kitex_gen/lql/sys/test/hello"
	"sys-rpclib/kitex_gen/user/user"
)

type Client struct {
	helloService hello.Client
	userService  user.Client
}

var (
	rpcs *Client
)

func init() {
	var err error
	rpcs.helloService, err = hello.NewClient("hello", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic("failed to init hello service")
	}
	rpcs.userService, err = user.NewClient("user", client.WithHostPorts("127.0.0.1:8881"))
	if err != nil {
		panic("failed to init user service")
	}
}

func RPCs() *Client {
	return rpcs
}

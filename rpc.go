package sys_rpclib

import (
	"github.com/cloudwego/kitex/client"
	"sys-rpclib/kitex_gen/lql/sys/test/hello"
)

type Client struct {
	helloService hello.Client
}

var (
	rpcs *Client
)

func init() {
	var err error
	rpcs.helloService, err = hello.NewClient("opt", client.WithHostPorts("127.0.0.0:8888"))
	if err != nil {
		panic("failed to init hello service")
	}
}

func RPCs() *Client {
	return rpcs
}

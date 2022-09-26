package ccr

import (
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi"
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
)

type Client struct {
	ctapi.Client
}

func NewClientWithCredential(credential credential.Interfaces) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithCredential(credential)
	return
}

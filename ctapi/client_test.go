package ctapi

import (
	"fmt"
	"testing"

	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
)

func Test_NewClientWithCredential(t *testing.T) {
	client, err := NewClientWithCredential(credential.NewAccessKey("d9912b01fd02414fad5ddd96be4010d0", "217fd48fe9dc41ec9a51fe00ddcd394c"))
	if err != nil {
		t.Error(err)
	}
	ak, _ := client.signer.GetAccessKeyId()
	fmt.Println(client.config, ak)
}

func Test_Client_Do(t *testing.T) {
	req := NewRequest()
	req.Domain = "ccr-global.ctapi-test.ctyun.cn/v2/ccr/instance/listRegions"
	req.QueryParams.Add("regionId", "d0dc8ddc8e3e11eca6200242ac110003")

	client, _ := NewClientWithCredential(credential.NewAccessKey("d9912b01fd02414fad5ddd96be4010d0", "217fd48fe9dc41ec9a51fe00ddcd394c"))
	var resp CommonResponse
	err := client.Do(req, &resp)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(resp)
	}
}

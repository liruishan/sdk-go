package ccr

import (
	"fmt"
	"testing"

	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/service/ccr/project"
)

func Test_ListProjects(t *testing.T) {
	client, _ := NewClientWithCredential(credential.NewAccessKey("d9912b01fd02414fad5ddd96be4010d0", "217fd48fe9dc41ec9a51fe00ddcd394c"))
	req := project.NewListRequest()
	req.RegionID("d0dc8ddc8e3e11eca6200242ac110003")

	projects, err := client.ListProjects(req)
	if err != nil {
		t.Fatal("do failed:", err)
	}
	fmt.Println(projects)
}

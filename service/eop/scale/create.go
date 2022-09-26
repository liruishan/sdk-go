package scale

import (
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi"
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/service/ccr/project"
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/service/eop"
)

type Client struct {
	*eop.Client
}

func (c *Client) ListProjects(req *project.ListRequest) ([]project.Project, error) {
	projects := []project.Project{}
	resp := &ctapi.CommonResponse{ReturnObj: &projects}
	err := c.Client.Do(req.Request, resp)
	if err != nil {
		return nil, err
	}

	return projects, resp.Error()
}

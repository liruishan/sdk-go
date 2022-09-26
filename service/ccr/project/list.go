package project

import "gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi"

type ListRequest struct {
	*ctapi.Request
}

func (r *ListRequest) RegionID(rid string) *ListRequest {
	r.QueryParams.Add("regionId", rid)
	return r
}

func NewListRequest() *ListRequest {
	return &ListRequest{
		Request: ctapi.NewRequest(ctapi.WithRequestDomain("ccr-global.ctapi-test.ctyun.cn/v2/ccr/project/list")),
	}
}

type ListResponse struct {
	Projects []Project
}

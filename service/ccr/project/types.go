package project

type Project struct {
	ID         int64  `json:"projectId,omitempty"`
	Name       string `json:"projectName,omitempty"`
	Public     bool   `json:"public,omitempty"`
	RepoCount  int64  `json:"repoCount,omitempty"`
	ChartCount int64  `json:"chartCount,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
	UpdateTime int64  `json:"updateTime,omitempty"`
}

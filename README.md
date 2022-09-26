<h1 align="center">CTYUN SDK for Go</h1>

欢迎使用 CTYUN SDK for Go。CTYUN SDK for Go 让您不用复杂编程即可访问云服务器、云监控等多个天翼云服务。
这里向您介绍如何获取 [Alibaba Cloud SDK for Go][SDK] 并开始调用。


## 安装

使用 `go get` 下载安装 SDK

```sh
$ go get -u gitlab.ctyun.cn/os/container/ctapi-sdk-go
```

## 快速使用

在您开始之前，您需要注册天翼云帐户并获取您的[凭证](https://www.ctyun.cn/h5/console/info/index)。

### 创建客户端

```go
package main

import (
    "gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi"
    "gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
)

func main() {

	client, err := ctapi.NewClientWithAccessKey(credential.NewAccessKey("ACCESS_KEY_ID", "ACCESS_KEY_SECRET"))
	if err != nil {
		// Handle exceptions
		panic(err)
	}
}
```

### 请求
```go
package main

import (
    "gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
    "gitlab.ctyun.cn/os/container/ctapi-sdk-go/service/ccr"
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/service/ccr/project"
)

func main() {
    client, _ := ccr.NewClientWithCredential(credential.NewAccessKey("ACCESS_KEY_ID", "ACCESS_KEY_SECRET"))
	req := project.NewListRequest()
	req.RegionID("RegionID")

	projects, err := client.ListProjects(req)
	if err != nil {
		panic(err)
	}
}
```


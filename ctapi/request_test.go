package ctapi

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/signer"
)

func Test_Request_newHTTPRequest(t *testing.T) {
	req := Request{
		Scheme:      "https",
		Method:      http.MethodGet,
		Domain:      "eop.ctyun.cn",
		QueryParams: url.Values{},
		Headers:     map[string]string{"key1": "value1", "key2": "value2"},
		Content:     []byte("{\"a\":\"b\"}"),
	}
	req.QueryParams.Add("regionID", "shanghai7")
	httpReq, err := req.newHTTPRequest()
	if err != nil {
		t.Errorf("creat http request failed:%v", err)
	}

	fmt.Println(httpReq)
}

func Test_Request_getSortedHeaderToSign(t *testing.T) {
	req := Request{
		Headers: map[string]string{},
	}
	headers := req.getSortedHeaderToSign()
	if date, exist := req.Headers[EOP_DATE]; !exist {
		t.Errorf("no eop date in request header")
	} else if rid, exist := req.Headers[EOP_REQUEST_ID]; !exist {
		t.Errorf("no eop request id in request header")
	} else if headers != fmt.Sprintf("%s:%s\n%s:%s\n", EOP_REQUEST_ID, rid, EOP_DATE, date) {
		t.Error(headers, date, rid)
	}
}

func Test_Request_getSortedQueriesToSign(t *testing.T) {
	tests := []struct {
		queryParams url.Values
		result      string
	}{
		{url.Values{}, ""},
		{url.Values{"b": []string{"b"}, "a": []string{"a"}}, "a=a&b=b"},
	}

	for i := range tests {
		req := Request{QueryParams: tests[i].queryParams}
		if queries := req.getSortedQueriesToSign(); queries != tests[i].result {
			t.Error(i, queries, tests[i].result)
		}
	}
}

func Test_Request_Sign(t *testing.T) {
	req := NewRequest()
	req.Scheme = "https"
	httpReq, err := req.Sign(signer.NewAccessKey(credential.NewAccessKey("1", "2")))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(httpReq.Header)
}

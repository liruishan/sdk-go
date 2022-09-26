package signer

import (
	"testing"

	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
)

func Test_AccessKey_Sign(t *testing.T) {
	tests := []struct {
		accessKeyId     string
		accessKeySecret string
		signature       string
		result          string
	}{
		{
			accessKeyId:     "1",
			accessKeySecret: "2",
			signature:       "ctyun-eop-request-id:27cfe4dc-e640-45f6-92ca-492ca73e8680\neop-date:20220525T160752Z\ne3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			result:          "tHLSPg5bu8+UmVId4JFHq0Fr6VpfDIY9zkBPaW1yPzM=",
		},
	}

	for i := range tests {
		ak := NewAccessKey(credential.NewAccessKey(tests[i].accessKeyId, tests[i].accessKeySecret))
		if result := ak.Sign(tests[i].signature); result != tests[i].result {
			t.Errorf("%d case failed: %s != %s", i, result, tests[i].result)
		}
	}
}

func Test_AccessKey_extractEOPDate(t *testing.T) {
	tests := []string{
		"ctyun-eop-request-id:27cfe4dc-e640-45f6-92ca-492ca73e8680\neop-date:20220525T160752Z\ne3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		"eop-date:20220525T160752Z",
		"ctyun-eop-request-id:27cfe4dc-e640-45f6-92ca-492ca73e8680\neop-date:20220525T160752Z",
	}

	ak := AccessKey{}
	for i := range tests {
		if date := ak.extractEOPDate(tests[i]); date != "20220525T160752Z" {
			t.Errorf("%d case invalid date:%s", i, date)
		}
	}
}

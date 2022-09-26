package signer

import (
	"fmt"
	"testing"

	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
)

func Test_New(t *testing.T) {
	if _, err := New(nil); err == nil {
		t.Errorf("create signer with nil credential successfully")
	} else {
		fmt.Println(err)
	}

	signer, err := New(credential.NewAccessKey("1", "2"))
	if err != nil {
		t.Errorf("create access key signer failed: %v", err)
	} else if name := signer.GetName(); name != AccessKeyName {
		t.Errorf("")
	}
}

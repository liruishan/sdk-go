package signer

import (
	"fmt"
	"reflect"

	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
)

// Interfaces defines interfaces of signer.
type Interfaces interface {
	// GetName return signer name.
	GetName() string
	// GetAccessKeyId return access key ID.
	GetAccessKeyId() (string, error)
	// Sign signs the string.
	Sign(stringToSign string) string
}

// New create signer by credential.
func New(c credential.Interfaces) (Interfaces, error) {
	switch t := c.(type) {
	case *credential.AccessKey:
		return NewAccessKey(t), nil
	default:
		return nil, fmt.Errorf("unsupported credential type: %v", reflect.TypeOf(c))
	}
}

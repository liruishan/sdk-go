package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"

	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
)

const (
	AccessKeyName = "HMAC-SHA256"
)

// AccessKey implements Interfaces.
type AccessKey struct {
	credential *credential.AccessKey
}

// NewAccessKey create access key signer.
func NewAccessKey(credential *credential.AccessKey) *AccessKey {
	return &AccessKey{credential: credential}
}

// GetName implements Interfaces.GetName().
func (*AccessKey) GetName() string { return AccessKeyName }

// GetAccessKeyId implements Interfaces.GetAccessKeyId().
func (ak *AccessKey) GetAccessKeyId() (accessKeyId string, err error) {
	return ak.credential.AccessKeyId, nil
}

// Sign implements Interfaces.Sign() interface.
func (ak *AccessKey) Sign(stringToSign string) string {
	// 构造动态密钥kdate
	eopDate := ak.extractEOPDate(stringToSign)
	ktime := hmacSHA256(eopDate, ak.credential.AccessKeySecret)
	kak := hmacSHA256(ak.credential.AccessKeyId, string(ktime))
	kdate := hmacSHA256(eopDate[:8], string(kak))
	signature := hmacSHA256(stringToSign, string(kdate))
	return base64.StdEncoding.EncodeToString(signature)
}

// extractEOPDate extract eop-date content from the string that needs to be signed
func (*AccessKey) extractEOPDate(stringToSign string) string {
	// Find 'eop-data' key word.
	start := strings.Index(stringToSign, "eop-date:")
	if start < 0 {
		panic("no eop-date in string:" + stringToSign)
	}

	// Find eop-data tail character.
	substr := stringToSign[start+len("eop-date:"):]
	if tail := strings.IndexAny(substr, " \n"); tail < 0 {
		return substr
	} else {
		return substr[:tail]
	}
}

// hmacSHA256 implements HMAC-SHA256 algorithm.
func hmacSHA256(data, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}

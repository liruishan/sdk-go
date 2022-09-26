package credential

// AccessKey is a credential that authenticate identity.
type AccessKey struct {
	AccessKeyId     string
	AccessKeySecret string
}

// NewAccessKey create AccessKey.
func NewAccessKey(accessKeyId, accessKeySecret string) *AccessKey {
	return &AccessKey{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret}
}

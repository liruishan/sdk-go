package ctapi

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/signer"
)

// Request define CTAPI request.
type Request struct {
	Scheme         string
	Method         string
	Domain         string
	Port           int
	RegionId       string
	ReadTimeout    time.Duration
	ConnectTimeout time.Duration
	AcceptFormat   string

	QueryParams url.Values
	Headers     map[string]string
	Content     []byte
}

type RequestOption func(req *Request)

func WithRequestScheme(scheme string) RequestOption {
	return func(req *Request) { req.Scheme = scheme }
}

func WithRequestMethod(method string) RequestOption {
	return func(req *Request) { req.Method = method }
}

func WithRequestDomain(domain string) RequestOption {
	return func(req *Request) { req.Domain = domain }
}

// NewRequest create request.
func NewRequest(opts ...RequestOption) *Request {
	req := &Request{
		Scheme:       "https",
		Method:       http.MethodGet,
		Headers:      make(map[string]string),
		QueryParams:  make(url.Values),
		AcceptFormat: "json",
	}

	for _, opt := range opts {
		opt(req)
	}

	return req
}

// Sign convert ctapi.Request to http.Request and signs it.
func (req *Request) Sign(signer signer.Interfaces) (*http.Request, error) {
	// Get access key ID from signer.
	accessKeyId, err := signer.GetAccessKeyId()
	if err != nil {
		return nil, err
	}

	// Sign request.
	signature := signer.Sign(strings.Join([]string{req.getSortedHeaderToSign(), req.getSortedQueriesToSign(), req.getBodyToSign()}, "\n"))

	// Create http request.
	httpReq, err := req.newHTTPRequest()
	if err != nil {
		return nil, err
	}

	// Add authorization header.
	httpReq.Header.Add(EOP_AUTHORIZATION, fmt.Sprintf("%s Headers=%s;%s Signature=%s", accessKeyId, EOP_REQUEST_ID, EOP_DATE, signature))

	return httpReq, nil
}

// newHTTPRequest create http.Request by ctapi.Request.
func (req *Request) newHTTPRequest() (*http.Request, error) {
	// Generate request URL.
	var reqURL string
	if req.Port > 0 {
		reqURL = fmt.Sprintf("%s://%s:%d", strings.ToLower(req.Scheme), req.Domain, req.Port)
	} else {
		reqURL = fmt.Sprintf("%s://%s", strings.ToLower(req.Scheme), req.Domain)
	}
	if len(req.QueryParams) > 0 {
		reqURL += "?" + req.QueryParams.Encode()
	}
	fmt.Println(reqURL)

	// Create body reader.
	var bodyReader io.Reader
	if len(req.Content) > 0 {
		bodyReader = bytes.NewReader(req.Content)
	}

	// Create http.Request.
	httpReq, err := http.NewRequest(req.Method, reqURL, bodyReader)
	if err != nil {
		return nil, err
	}

	// Copy headers.
	for key, value := range req.Headers {
		httpReq.Header.Add(key, value)
	}

	return httpReq, nil
}

// getSortedHeaderToSign get the sorted headers needs to be signed.
func (req *Request) getSortedHeaderToSign() string {
	// Get time and request UUID and add into headers of request.
	eopDate := time.Now().Format(EOP_DATE_FORMAT)
	requestID := uuid.New().String()
	req.Headers[EOP_DATE] = eopDate
	req.Headers[EOP_REQUEST_ID] = requestID

	return fmt.Sprintf("%s:%s\n%s:%s\n", EOP_REQUEST_ID, requestID, EOP_DATE, eopDate)
}

// getSortedQueriesToSign get the sorted queries needs to be signed.
func (req *Request) getSortedQueriesToSign() string {
	if len(req.QueryParams) == 0 {
		return ""
	}

	// Sort query parameters by keys.
	queryKeys := make([]string, 0, len(req.QueryParams))
	for key := range req.QueryParams {
		queryKeys = append(queryKeys, key)
	}
	sort.Strings(queryKeys)

	// Sort query parameters.
	var sortedQueries strings.Builder
	for i, key := range queryKeys {
		if i == 0 {
			sortedQueries.WriteString(key + "=" + req.QueryParams.Get(key))
		} else {
			sortedQueries.WriteString("&" + key + "=" + req.QueryParams.Get(key))
		}
	}

	return sortedQueries.String()
}

// getBodyToSign get body to sign.
func (req *Request) getBodyToSign() string {
	hash := sha256.New()
	hash.Write(req.Content)
	return hex.EncodeToString(hash.Sum(nil))
}

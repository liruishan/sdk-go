package ctapi

import (
	"crypto/tls"
	"net/http"

	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/credential"
	"gitlab.ctyun.cn/os/container/ctapi-sdk-go/ctapi/signer"
)

// CTAPI http headers.
const (
	EOP_DATE          = "eop-date"
	EOP_DATE_FORMAT   = "20060102T150405Z"
	EOP_REQUEST_ID    = "ctyun-eop-request-id"
	EOP_AUTHORIZATION = "eop-authorization"
)

// CTAPI client.
type Client struct {
	config     *Config
	signer     signer.Interfaces
	httpClient *http.Client
}

// NewClientWithCredential create CTAPI client with credential.
func NewClientWithCredential(credential credential.Interfaces) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithCredential(credential)
	return
}

// InitWithCredential initialize CTAPI client with credential.
func (client *Client) InitWithCredential(credential credential.Interfaces) (err error) {
	// Create signer.
	if client.signer, err = signer.New(credential); err != nil {
		return err
	}

	if client.config == nil {
		client.config = NewConfig()
	}

	// Create http client.
	client.httpClient = &http.Client{}
	if client.config.Transport != nil {
		client.httpClient.Transport = client.config.Transport
	} else if client.config.HttpTransport != nil {
		client.httpClient.Transport = client.config.HttpTransport
	} else {
		client.httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	if client.config.Timeout > 0 {
		client.httpClient.Timeout = client.config.Timeout
	}

	return
}

// Do sends a ctapi request and returns output ctapi response, following
// policy (such as redirects, cookies, auth) as configured on the
// client.
func (c *Client) Do(req *Request, resp Response) error {
	// Sign ctapi request.
	httpReq, err := req.Sign(c.signer)
	if err != nil {
		return err
	}

	// Send request.
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}

	// Convert http response to ctapi request.
	return resp.From(httpResp, req.AcceptFormat)
}

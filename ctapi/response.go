package ctapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define ctapi status code.
const (
	StatusOK = 800
)

// Response represents the response from an CTAPI request.
type Response interface {
	// From convert http response to ctapi response.
	From(httpResp *http.Response, format string) error
	// Error return ctapi response has any error or not.
	Error() error
}

// CommonResponse define common ctapi response.
type CommonResponse struct {
	StatusCode  int         `json:"statusCode,omitempty"`
	ErrorCode   string      `json:"errorCode,omitempty"`
	Message     string      `json:"message,omitempty"`
	Description string      `json:"description,omitempty"`
	ReturnObj   interface{} `json:"returnObj,omitempty"`
}

// From implements Response.From() interface.
func (resp *CommonResponse) From(httpResp *http.Response, format string) error {
	// Read body from http response.
	defer httpResp.Body.Close()
	data, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	// Only json format supportted.
	switch format {
	default:
		return json.Unmarshal(data, resp)
	}
}

// Error implements Response.Error() interface.
func (resp *CommonResponse) Error() error {
	if resp.StatusCode != StatusOK {
		return fmt.Errorf("%s:%s", resp.ErrorCode, resp.Message)
	}

	return nil
}

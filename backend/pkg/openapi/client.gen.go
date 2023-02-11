// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// V1ListArticles request
	V1ListArticles(ctx context.Context, params *V1ListArticlesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// V1HealthCore request
	V1HealthCore(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// V1HealthGateway request
	V1HealthGateway(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) V1ListArticles(ctx context.Context, params *V1ListArticlesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1ListArticlesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1HealthCore(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1HealthCoreRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1HealthGateway(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1HealthGatewayRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewV1ListArticlesRequest generates requests for V1ListArticles
func NewV1ListArticlesRequest(server string, params *V1ListArticlesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/articles")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.PageToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pageToken", runtime.ParamLocationQuery, *params.PageToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxPageSize", runtime.ParamLocationQuery, params.MaxPageSize); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewV1HealthCoreRequest generates requests for V1HealthCore
func NewV1HealthCoreRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/health/core")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewV1HealthGatewayRequest generates requests for V1HealthGateway
func NewV1HealthGatewayRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/health/gateway")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// V1ListArticles request
	V1ListArticlesWithResponse(ctx context.Context, params *V1ListArticlesParams, reqEditors ...RequestEditorFn) (*V1ListArticlesResponse, error)

	// V1HealthCore request
	V1HealthCoreWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthCoreResponse, error)

	// V1HealthGateway request
	V1HealthGatewayWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthGatewayResponse, error)
}

type V1ListArticlesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListArticleResponse
}

// Status returns HTTPResponse.Status
func (r V1ListArticlesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r V1ListArticlesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type V1HealthCoreResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r V1HealthCoreResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r V1HealthCoreResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type V1HealthGatewayResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r V1HealthGatewayResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r V1HealthGatewayResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// V1ListArticlesWithResponse request returning *V1ListArticlesResponse
func (c *ClientWithResponses) V1ListArticlesWithResponse(ctx context.Context, params *V1ListArticlesParams, reqEditors ...RequestEditorFn) (*V1ListArticlesResponse, error) {
	rsp, err := c.V1ListArticles(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1ListArticlesResponse(rsp)
}

// V1HealthCoreWithResponse request returning *V1HealthCoreResponse
func (c *ClientWithResponses) V1HealthCoreWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthCoreResponse, error) {
	rsp, err := c.V1HealthCore(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1HealthCoreResponse(rsp)
}

// V1HealthGatewayWithResponse request returning *V1HealthGatewayResponse
func (c *ClientWithResponses) V1HealthGatewayWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthGatewayResponse, error) {
	rsp, err := c.V1HealthGateway(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1HealthGatewayResponse(rsp)
}

// ParseV1ListArticlesResponse parses an HTTP response from a V1ListArticlesWithResponse call
func ParseV1ListArticlesResponse(rsp *http.Response) (*V1ListArticlesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &V1ListArticlesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListArticleResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseV1HealthCoreResponse parses an HTTP response from a V1HealthCoreWithResponse call
func ParseV1HealthCoreResponse(rsp *http.Response) (*V1HealthCoreResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &V1HealthCoreResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseV1HealthGatewayResponse parses an HTTP response from a V1HealthGatewayWithResponse call
func ParseV1HealthGatewayResponse(rsp *http.Response) (*V1HealthGatewayResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &V1HealthGatewayResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

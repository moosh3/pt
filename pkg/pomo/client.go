package pomo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RestInterface interface {
	GetRateLimiter() flowcontrol.RateLimiter
	Verb(verb string) *Request
	Post() *Request
	Put() *Request
	Patch(pt types.PatchType) *Request
	Get() *Request
	Delete() *Request
	APIVersion() schema.GroupVersion
}


type Client struct {
	// base is the root URL for all invocations of the client
	base		*url.URL
	UserAgent	string

	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	httpClient	*http.client
}

// NewRESTClient creates a new RESTClient. This client performs generic REST functions
// such as Get, Put, Post, and Delete on specified paths.
func NewRESTClient(baseURL *url.URL, versionedAPIPath string, config ClientContentConfig, rateLimiter flowcontrol.RateLimiter, client *http.Client) (*RESTClient, error) {
	if len(config.ContentType) == 0 {
		config.ContentType = "application/json"
	}

	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.RawQuery = ""
	base.Fragment = ""

	return &RESTClient{
		base:             &base,
		versionedAPIPath: versionedAPIPath,
		content:          config,
		createBackoffMgr: readExpBackoffConfig,
		rateLimiter:      rateLimiter,

		Client: client,
	}, nil
}

// Verb begins a request with a verb (GET, POST, PUT, DELETE).
//
// Example usage of RESTClient's request building interface:
// c, err := NewRESTClient(...)
// if err != nil { ... }
// resp, err := c.Verb("GET").
//  Path("pods").
//  SelectorParam("labels", "area=staging").
//  Timeout(10*time.Second).
//  Do()
// if err != nil { ... }
// list, ok := resp.(*api.PodList)
//
func (c *RESTClient) Verb(verb string) *Request {
	return NewRequest(c).Verb(verb)
}

// Post begins a POST request. Short for c.Verb("POST").
func (c *RESTClient) Post() *Request {
	return c.Verb("POST")
}

// Put begins a PUT request. Short for c.Verb("PUT").
func (c *RESTClient) Put() *Request {
	return c.Verb("PUT")
}

// Patch begins a PATCH request. Short for c.Verb("Patch").
func (c *RESTClient) Patch(pt types.PatchType) *Request {
	return c.Verb("PATCH").SetHeader("Content-Type", string(pt))
}

// Get begins a GET request. Short for c.Verb("GET").
func (c *RESTClient) Get() *Request {
	return c.Verb("GET")
}

// Delete begins a DELETE request. Short for c.Verb("DELETE").
func (c *RESTClient) Delete() *Request {
	return c.Verb("DELETE")
}

// APIVersion returns the APIVersion this RESTClient is expected to use.
func (c *RESTClient) APIVersion() schema.GroupVersion {
	return c.content.GroupVersion
}

// newRequest
func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
    rel := &url.URL{Path: path}
    u := c.BaseURL.ResolveReference(rel)
    var buf io.ReadWriter
    if body != nil {
        buf = new(bytes.Buffer)
        err := json.NewEncoder(buf).Encode(body)
        if err != nil {
            return nil, err
        }
    }
    req, err := http.NewRequest(method, u.String(), buf)
    if err != nil {
        return nil, err
    }
    if body != nil {
        req.Header.Set("Content-Type", "application/json")
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", c.UserAgent)
    return req, nil
}

// do
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    err = json.NewDecoder(resp.Body).Decode(v)
    return resp, err
}


type PomoV1Interface interface {
	RESTClient() RestInterface
	PomoGetter
	TodoGetter
	SubtodoGetter
}

type PomoV1Client struct {
	restClient RestInterface
}

func (c *PomoV1Client) Pomos() PomoInterface {
	return newPomos(c)
}

func (c *PomoV1Client) Todos() TodoInterface {
	return newTodos(c)
}

func (c *PomoV1Client) SubTodos() SubtodoInterface {
	return newSubtodos(c)
}

// New creates a new AppsV1Client for the given RESTClient.
func New(c rest.Interface) *PomoV1Client {
	return &PomoV1Client{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AppsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
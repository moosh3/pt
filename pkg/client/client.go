package client

type Client struct {
	BaseURL		*url.URL
	UserAgent	string
	httpClient	*http.client
}

// Nodes
func (c *Client) ListNodes()
func (c *Client) GetNode(id int) (Node, error) {} 
func (c *Client) CreateNode(Node) (Node, error) {}
func (c *Client) UpdateNode(Node) (Node, error) {}

// Versions
func (c *Client) ListVersions() ([]Version, error) {}
func (c *Client) GetVersion(id int) {}
func (c *Client) CreateVersion(Version) {}
func (c *Client) UpdateVersion(Version) {}

// Upgrades
func (c *Client) ListUpgrades() ([]Upgrade, error) {}
func (c *Client) GetUpgrade(id int) (Upgrade, error) {}
func (c *Client) CreateUpgrade(Upgrade) (Upgrade, error) {}
func (c *Client) UpdateUpgrade(Upgrade) (Upgrade, error) {}

func (c *Client) ListEnvironments() {}
func (c *Client) GetEnvironment() {}


func (c *Client) ListNodes() ([]Node, error) {
    req, err := c.newRequest("GET", "/nodes", nil)
    if err != nil {
        return nil, err
    }
    var nodes []Node
    _, err = c.do(req, &nodes)
    return nodes, err
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

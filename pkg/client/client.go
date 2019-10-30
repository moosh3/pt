package client

type Client struct {
	BaseURL		*url.URL
	UserAgent	string
	httpClient	*http.client
}

// Nodes
// ListNodes
func (c *Client) ListNodes() ([]Node, error) {
    req, err := c.newRequest("GET", "/nodes", nil)
    if err != nil {
        return nil, err
    }
    var nodes []Node
    _, err = c.do(req, &nodes)
    return nodes, err
}

// GetNode
func (c *Client) GetNode(id int) (Node, error) {
	req, err := c.newRequest("GET", "/nodes", nil)
	if err != nil {
			return nil, err
	}
	var nodes []Node
	_, err = c.do(req, &nodes)
	return nodes, err
}

// CreateNode
func (c *Client) CreateNode(Node) (Node, error) {
	req, err := c.newRequest("POST", "/nodes", nil)
	if err != nil {
			return nil, err
	}
	var nodes []Node
	_, err = c.do(req, &nodes)
	return nodes, err
}

// UpdateNode
func (c *Client) UpdateNode(Node) (Node, error) {
	req, err := c.newRequest("PATCH", "/nodes", nil)
	if err != nil {
			return nil, err
	}
	var node Node
	_, err = c.do(req, &node)
	return node, err
}

// DeleteNode
func (c *Client) DeleteNode(Node) (Node, error) {
	req, err := c.newRequest("POST", "/nodes", nil)
	if err != nil {
			return nil, err
	}
	var node Node
	_, err = c.do(req, &node)
	return node, err
}

// Versions
// ListVersions
func (c *Client) ListVersions() ([]Version, error) {
	req, err := c.newRequest("GET", "/versions", nil)
	if err != nil {
			return nil, err
	}
	var versions []Version
	_, err = c.do(req, &versions)
	return versions, err
}

// GetVersion
func (c *Client) GetVersion(id int) {
	req, err := c.newRequest("GET", "/versions", nil)
	if err != nil {
			return nil, err
	}
	var version Version
	_, err = c.do(req, &version)
	return version, err
}

// CreateVersion
func (c *Client) CreateVersion(Version) {
	req, err := c.newRequest("POST", "/versions", nil)
	if err != nil {
			return nil, err
	}
	var version Version
	_, err = c.do(req, &version)
	return version, err
}

// UpdateVersion
func (c *Client) UpdateVersion(Version) {
	req, err := c.newRequest("PATCH", "/versions", nil)
	if err != nil {
			return nil, err
	}
	var version Version
	_, err = c.do(req, &version)
	return version, err
}

// DeleteVersion
func (c *Client) DeleteVersion(Version) {
	req, err := c.newRequest("POST", "/versions", nil)
	if err != nil {
			return nil, err
	}
	var version Version
	_, err = c.do(req, &version)
	return version, err
}


// Upgrades
// ListUpgrades
func (c *Client) ListUpgrades() ([]Upgrade, error) {
	req, err := c.newRequest("GET", "/upgrades", nil)
	if err != nil {
			return nil, err
	}GET
	var upgrades []Upgrade
	_, err = c.do(req, &upgrades)
	return upgrades, err
}

// GetUpgrade
func (c *Client) GetUpgrade(id int) (Upgrade, error) {
	req, err := c.newRequest("GET", "/upgrades", nil)
	if err != nil {
			return nil, err
	}
	var upgrade Upgrade
	_, err = c.do(req, &upgrade)
	return upgrade, err
}

// CreateUpgrade
func (c *Client) CreateUpgrade(Upgrade) (Upgrade, error) {
	req, err := c.newRequest("POST", "/upgrades", nil)
	if err != nil {
			return nil, err
	}
	var upgrade Upgrade
	_, err = c.do(req, &upgrade)
	return upgrade, err
}

// UpdateUpgrade
func (c *Client) UpdateUpgrade(Upgrade) (Upgrade, error) {
	req, err := c.newRequest("PATCH", "/upgrades", nil)
	if err != nil {
			return nil, err
	}
	var upgrade Upgrade
	_, err = c.do(req, &upgrade)
	return upgrade, err
}

// DeleteUpgrade
func (c *Client) DeleteUpgrade(Upgrade) (Upgrade, error) {
	req, err := c.newRequest("POST", "/upgrades", nil)
	if err != nil {
			return nil, err
	}
	var upgrade Upgrade
	_, err = c.do(req, &upgrade)
	return upgrade, err
}

// ListEnvironments
func (c *Client) ListEnvironments() {
	req, err := c.newRequest("GET", "/environments", nil)
	if err != nil {
			return nil, err
	}
	var environments []Environment
	_, err = c.do(req, &environments)
	return nodes, err
}

// GetEnvironment
func (c *Client) GetEnvironment() {
	req, err := c.newRequest("GET", "/environments", nil)
	if err != nil {
			return nil, err
	}
	var environment Environment
	_, err = c.do(req, &environment)
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

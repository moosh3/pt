package pomo

type SubtodoGetter interface {
	Inter
}
// List
func (c *SubTodo) List() ([]interface{}, error) {
	req, err := c.newRequest("GET", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var nodes []SubTodo
	_, err = c.do(req, &nodes)
	return nodes, err
}

// Get
func (c *SubTodo) Get(id int) (interface{}, error) {
	req, err := c.newRequest("GET", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var nodes []SubTodo
	_, err = c.do(req, &nodes)
	return nodes, err
}

// Create
func (c *SubTodo) Create(SubTodo) (interface{}, error) {
	req, err := c.newRequest("POST", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var nodes []SubTodo
	_, err = c.do(req, &nodes)
	return nodes, err
}

// Update
func (c *SubTodo) Update(SubTodo) (interface{}, error) {
	req, err := c.newRequest("PATCH", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var node SubTodo
	_, err = c.do(req, &node)
	return node, err
}

// Delete
func (c *SubTodo) Delete(SubTodo) (interface{}, error) {
	req, err := c.newRequest("POST", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var node SubTodo
	_, err = c.do(req, &node)
	return node, err
}
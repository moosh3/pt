package v1

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// SubTodoGetter has a method to return a SubTodoInterface.
// A group's client should implement this interface.
type SubTodoGetter interface {
	SubTodos() SubTodoInterface
}

// SubTodoInterface has methods to work with Pomo resources.
type SubTodoInterface interface {
	Create(*SubTodo) (*SubTodo, error)
	Update(*SubTodo) (*SubTodo, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string, options v1.GetOptions) (*SubTodo, error)
	List(opts v1.ListOptions) (*SubTodoList, error)
}

// subtodos implements SubTodoInterface
type subtodos struct {
	client c
}

// newPomos returns a pomos
func newSubTodoes(c *Client) *subtodos {
	return &subtodos{
		client: c.RESTClient(),
	}
}

// List
func (t *subtodos) List() ([]interface{}, error) {
	req, err := c.newRequest("GET", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var nodes []SubTodo
	_, err = c.do(req, &nodes)
	return nodes, err
}

// Get
func (t *subtodos) Get(id int) (interface{}, error) {
	req, err := c.newRequest("GET", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var nodes []SubTodo
	_, err = c.do(req, &nodes)
	return nodes, err
}

// Create
func (t *subtodos) Create(SubTodo) (interface{}, error) {
	req, err := c.newRequest("POST", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var nodes []SubTodo
	_, err = c.do(req, &nodes)
	return nodes, err
}

// Update
func (t *subtodos) Update(SubTodo) (interface{}, error) {
	req, err := c.newRequest("PATCH", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var node SubTodo
	_, err = c.do(req, &node)
	return node, err
}

// Delete
func (t *subtodos) Delete(SubTodo) (interface{}, error) {
	req, err := c.newRequest("POST", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var node SubTodo
	_, err = c.do(req, &node)
	return node, err
}

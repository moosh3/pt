package v1

import (
	"time"

	"k8s.io/client-go/kubernetes/scheme"
)

// TodoGetter has a method to return a TodoInterface.
// A group's client should implement this interface.
type TodoGetter interface {
	Todos() TodoInterface
}

// TodoInterface has methods to work with Pomo resources.
type TodoInterface interface {
	Create(*Todo) (*Todo, error)
	Update(*Todo) (*Todo, error)
	Delete(uuid string, options *v1.DeleteOptions) error
	Get(uuid string, options v1.GetOptions) (*Todo, error)
	List(opts v1.ListOptions) (*TodoList, error)
}

// subtodos implements SubTodoInterface
type todos struct {
	client c
}

// newTodos returns a pomos
func newTodos(c *Client) *todos {
	return &subtodos{
		client: c.RESTClient(),
	}
}

// List
// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (t *todos) Get(name string, options metav1.GetOptions) (result *Todo, err error) {
	result = &Pomo{}
	err = c.client.Get().
		Resource("deployments").
		UUID(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (t *todos) List(opts metav1.ListOptions) (result *TodoList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DeploymentList{}
	err = c.client.Get().
		Resource("todos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (t *todos) Create(todo *Todo) (result *Todo, err error) {
	result = &Pomo{}
	err = c.client.Post().
		Resource("todos").
		Body(todo).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (t *todos) Update(todo *Todo) (result *Todo, err error) {
	result = &Todo{}
	err = c.client.Put().
		Resource("todos").
		UUID(todo.UUID).
		Body(todo).
		Do().
		Into(result)
	return
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (t *todos) Delete(uuid string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("todos").
		UUID(uuid).
		Body(options).
		Do().
		Error()
}

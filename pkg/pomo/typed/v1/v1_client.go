package v1

import (
	rest "k8s.io/client-go/rest"
)

type V1Interface interface {
	RESTClient() rest.Interface
	PomoGetter
	TodoGetter
	SubTodoGetter
}

// V1Client is used to interact with features provided by the apps group.
type V1Client struct {
	restClient rest.Interface
}

func (c *V1Client) Pomos() PomoInterface {
	return newPomos(c)
}

func (c *V1Client) SubTodos() SubTodoInterface {
	return newSubTodos(c)
}

func (c *V1Client) Todos() TodoInterface {
	return newTodos(c)
}

// NewForConfig creates a new AppsV1Client for the given config.
func NewForConfig(c *rest.Config) (*V1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &V1Client{client}, nil
}

// NewForConfigOrDie creates a new AppsV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *V1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AppsV1Client for the given RESTClient.
func New(c rest.Interface) *V1Client {
	return &V1Client{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *V1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

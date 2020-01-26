package pomo

import "k8s.io/client-go/rest"

type Interface interface {
	PomosV1() v1.PomoInterface
	TodosV1() v1.TodoInterface
	SubTodoV1() v1.SubTodoInterface
}

// Client set contains the clients for a version
type Clientset struct {
	coreV1 *v1.V1Client
}

// V1 retrieves the V1Client
func (c *Clientset) V1() v1.V1Interface {
	return c.coreV1
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.coreV1 = v1.New(c)

	return &cs
}

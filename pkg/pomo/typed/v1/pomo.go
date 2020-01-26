package v1

import (
	"time"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
)

// PomoGetter has a method to return a DeploymentInterface.
// A group's client should implement this interface.
type PomoGetter interface {
	Pomos() PomoInterface
}

// PomoInterface has methods to work with Pomo resources.
type PomoInterface interface {
	Create(*Pomo) (*Pomo, error)
	Update(*Pomo) (*Pomo, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string, options v1.GetOptions) (*Pomo, error)
	List(opts v1.ListOptions) (*PomoList, error)
}

// pomos implements PomoInterface
type pomos struct {
	client c
}

// newPomos returns a pomos
func newPomos(c *Client) *pomos {
	return &pomos{
		client: c.RESTClient(),
	}
}

// List
func (c *pomos) List() ([]Pomo, error) {
	req, err := c.newRequest("GET", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var nodes []Pomo
	_, err = c.do(req, &nodes)
	return nodes, err
}

// Get
func (c *pomos) Get(id int) (Pomo, error) {
	req, err := c.newRequest("GET", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var nodes []Pomo
	_, err = c.do(req, &nodes)
	return nodes, err
}

// Create
func (c *pomos) Create(Pomo) (Pomo, error) {
	req, err := c.newRequest("POST", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var nodes []Pomo
	_, err = c.do(req, &nodes)
	return nodes, err
}

// Update
func (c *pomos) Update(Pomo) (Pomo, error) {
	req, err := c.newRequest("PATCH", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var node Pomo
	_, err = c.do(req, &node)
	return node, err
}

// Delete
func (c *pomos) Delete(Pomo) (Pomo, error) {
	req, err := c.newRequest("POST", "/nodes", nil)
	if err != nil {
		return nil, err
	}
	var node Pomo
	_, err = c.do(req, &node)
	return node, err
}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *deployments) Get(name string, options metav1.GetOptions) (result *v1.Deployment, err error) {
	result = &Pomo{}
	err = c.client.Get().
		Resource("deployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *deployments) List(opts metav1.ListOptions) (result *v1.DeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DeploymentList{}
	err = c.client.Get().
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deployments.
func (c *deployments) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Create(deployment *v1.Deployment) (result *v1.Deployment, err error) {
	result = &Pomo{}
	err = c.client.Post().
		Resource("deployments").
		Body(deployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Update(deployment *v1.Deployment) (result *v1.Deployment, err error) {
	result = &Pomo{}
	err = c.client.Put().
		Resource("deployments").
		Name(deployment.Name).
		Body(deployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *deployments) UpdateStatus(deployment *v1.Deployment) (result *v1.Deployment, err error) {
	result = &Pomo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deployment.Name).
		SubResource("status").
		Body(deployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *deployments) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deployments) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched deployment.
func (c *deployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Deployment, err error) {
	result = &Pomo{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

// GetScale takes name of the deployment, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *deployments) GetScale(deploymentName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		Name(deploymentName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *deployments) UpdateScale(deploymentName string, scale *autoscalingv1.Scale) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deploymentName).
		SubResource("scale").
		Body(scale).
		Do().
		Into(result)
	return
}

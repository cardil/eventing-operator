/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing-operator/pkg/apis/eventing/v1alpha1"
	scheme "knative.dev/eventing-operator/pkg/client/clientset/versioned/scheme"
)

// EventingsGetter has a method to return a EventingInterface.
// A group's client should implement this interface.
type EventingsGetter interface {
	Eventings(namespace string) EventingInterface
}

// EventingInterface has methods to work with KnativeEventing resources.
type EventingInterface interface {
	Create(*v1alpha1.KnativeEventing) (*v1alpha1.KnativeEventing, error)
	Update(*v1alpha1.KnativeEventing) (*v1alpha1.KnativeEventing, error)
	UpdateStatus(*v1alpha1.KnativeEventing) (*v1alpha1.KnativeEventing, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KnativeEventing, error)
	List(opts v1.ListOptions) (*v1alpha1.KnativeEventingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KnativeEventing, err error)
	EventingExpansion
}

// eventings implements EventingInterface
type eventings struct {
	client rest.Interface
	ns     string
}

// newEventings returns a Eventings
func newEventings(c *OperatorV1alpha1Client, namespace string) *eventings {
	return &eventings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventing, and returns the corresponding eventing object, and an error if there is any.
func (c *eventings) Get(name string, options v1.GetOptions) (result *v1alpha1.KnativeEventing, err error) {
	result = &v1alpha1.KnativeEventing{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Eventings that match those selectors.
func (c *eventings) List(opts v1.ListOptions) (result *v1alpha1.KnativeEventingList, err error) {
	result = &v1alpha1.KnativeEventingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventings.
func (c *eventings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a eventing and creates it.  Returns the server's representation of the eventing, and an error, if there is any.
func (c *eventings) Create(eventing *v1alpha1.KnativeEventing) (result *v1alpha1.KnativeEventing, err error) {
	result = &v1alpha1.KnativeEventing{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventings").
		Body(eventing).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eventing and updates it. Returns the server's representation of the eventing, and an error, if there is any.
func (c *eventings) Update(eventing *v1alpha1.KnativeEventing) (result *v1alpha1.KnativeEventing, err error) {
	result = &v1alpha1.KnativeEventing{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventings").
		Name(eventing.Name).
		Body(eventing).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eventings) UpdateStatus(eventing *v1alpha1.KnativeEventing) (result *v1alpha1.KnativeEventing, err error) {
	result = &v1alpha1.KnativeEventing{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventings").
		Name(eventing.Name).
		SubResource("status").
		Body(eventing).
		Do().
		Into(result)
	return
}

// Delete takes name of the eventing and deletes it. Returns an error if one occurs.
func (c *eventings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eventing.
func (c *eventings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KnativeEventing, err error) {
	result = &v1alpha1.KnativeEventing{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

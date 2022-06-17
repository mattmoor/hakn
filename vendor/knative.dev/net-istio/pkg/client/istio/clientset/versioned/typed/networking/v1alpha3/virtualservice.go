/*
Copyright 2020 The Knative Authors

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

package v1alpha3

import (
	"context"
	"time"

	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "knative.dev/net-istio/pkg/client/istio/clientset/versioned/scheme"
)

// VirtualServicesGetter has a method to return a VirtualServiceInterface.
// A group's client should implement this interface.
type VirtualServicesGetter interface {
	VirtualServices(namespace string) VirtualServiceInterface
}

// VirtualServiceInterface has methods to work with VirtualService resources.
type VirtualServiceInterface interface {
	Create(ctx context.Context, virtualService *v1alpha3.VirtualService, opts v1.CreateOptions) (*v1alpha3.VirtualService, error)
	Update(ctx context.Context, virtualService *v1alpha3.VirtualService, opts v1.UpdateOptions) (*v1alpha3.VirtualService, error)
	UpdateStatus(ctx context.Context, virtualService *v1alpha3.VirtualService, opts v1.UpdateOptions) (*v1alpha3.VirtualService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha3.VirtualService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha3.VirtualServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.VirtualService, err error)
	VirtualServiceExpansion
}

// virtualServices implements VirtualServiceInterface
type virtualServices struct {
	client rest.Interface
	ns     string
}

// newVirtualServices returns a VirtualServices
func newVirtualServices(c *NetworkingV1alpha3Client, namespace string) *virtualServices {
	return &virtualServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualService, and returns the corresponding virtualService object, and an error if there is any.
func (c *virtualServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.VirtualService, err error) {
	result = &v1alpha3.VirtualService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualServices that match those selectors.
func (c *virtualServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.VirtualServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha3.VirtualServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualServices.
func (c *virtualServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualService and creates it.  Returns the server's representation of the virtualService, and an error, if there is any.
func (c *virtualServices) Create(ctx context.Context, virtualService *v1alpha3.VirtualService, opts v1.CreateOptions) (result *v1alpha3.VirtualService, err error) {
	result = &v1alpha3.VirtualService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualService).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualService and updates it. Returns the server's representation of the virtualService, and an error, if there is any.
func (c *virtualServices) Update(ctx context.Context, virtualService *v1alpha3.VirtualService, opts v1.UpdateOptions) (result *v1alpha3.VirtualService, err error) {
	result = &v1alpha3.VirtualService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualservices").
		Name(virtualService.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualService).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualServices) UpdateStatus(ctx context.Context, virtualService *v1alpha3.VirtualService, opts v1.UpdateOptions) (result *v1alpha3.VirtualService, err error) {
	result = &v1alpha3.VirtualService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualservices").
		Name(virtualService.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualService).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualService and deletes it. Returns an error if one occurs.
func (c *virtualServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualservices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualservices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualService.
func (c *virtualServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.VirtualService, err error) {
	result = &v1alpha3.VirtualService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualservices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

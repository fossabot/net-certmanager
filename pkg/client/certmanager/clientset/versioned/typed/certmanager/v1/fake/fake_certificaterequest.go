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

package fake

import (
	"context"

	certmanagerv1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificateRequests implements CertificateRequestInterface
type FakeCertificateRequests struct {
	Fake *FakeCertmanagerV1
	ns   string
}

var certificaterequestsResource = schema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "certificaterequests"}

var certificaterequestsKind = schema.GroupVersionKind{Group: "cert-manager.io", Version: "v1", Kind: "CertificateRequest"}

// Get takes name of the certificateRequest, and returns the corresponding certificateRequest object, and an error if there is any.
func (c *FakeCertificateRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *certmanagerv1.CertificateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificaterequestsResource, c.ns, name), &certmanagerv1.CertificateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*certmanagerv1.CertificateRequest), err
}

// List takes label and field selectors, and returns the list of CertificateRequests that match those selectors.
func (c *FakeCertificateRequests) List(ctx context.Context, opts v1.ListOptions) (result *certmanagerv1.CertificateRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificaterequestsResource, certificaterequestsKind, c.ns, opts), &certmanagerv1.CertificateRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &certmanagerv1.CertificateRequestList{ListMeta: obj.(*certmanagerv1.CertificateRequestList).ListMeta}
	for _, item := range obj.(*certmanagerv1.CertificateRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificateRequests.
func (c *FakeCertificateRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificaterequestsResource, c.ns, opts))

}

// Create takes the representation of a certificateRequest and creates it.  Returns the server's representation of the certificateRequest, and an error, if there is any.
func (c *FakeCertificateRequests) Create(ctx context.Context, certificateRequest *certmanagerv1.CertificateRequest, opts v1.CreateOptions) (result *certmanagerv1.CertificateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificaterequestsResource, c.ns, certificateRequest), &certmanagerv1.CertificateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*certmanagerv1.CertificateRequest), err
}

// Update takes the representation of a certificateRequest and updates it. Returns the server's representation of the certificateRequest, and an error, if there is any.
func (c *FakeCertificateRequests) Update(ctx context.Context, certificateRequest *certmanagerv1.CertificateRequest, opts v1.UpdateOptions) (result *certmanagerv1.CertificateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificaterequestsResource, c.ns, certificateRequest), &certmanagerv1.CertificateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*certmanagerv1.CertificateRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificateRequests) UpdateStatus(ctx context.Context, certificateRequest *certmanagerv1.CertificateRequest, opts v1.UpdateOptions) (*certmanagerv1.CertificateRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(certificaterequestsResource, "status", c.ns, certificateRequest), &certmanagerv1.CertificateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*certmanagerv1.CertificateRequest), err
}

// Delete takes name of the certificateRequest and deletes it. Returns an error if one occurs.
func (c *FakeCertificateRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(certificaterequestsResource, c.ns, name, opts), &certmanagerv1.CertificateRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificateRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificaterequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &certmanagerv1.CertificateRequestList{})
	return err
}

// Patch applies the patch and returns the patched certificateRequest.
func (c *FakeCertificateRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *certmanagerv1.CertificateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificaterequestsResource, c.ns, name, pt, data, subresources...), &certmanagerv1.CertificateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*certmanagerv1.CertificateRequest), err
}

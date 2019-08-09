/*
Copyright 2019 The KubeSphere authors.

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
// Code generated by main. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubesphere.io/kubesphere/pkg/apis/devops/v1alpha1"
)

// FakeS2iBinaries implements S2iBinaryInterface
type FakeS2iBinaries struct {
	Fake *FakeDevopsV1alpha1
	ns   string
}

var s2ibinariesResource = schema.GroupVersionResource{Group: "devops.kubesphere.io", Version: "v1alpha1", Resource: "s2ibinaries"}

var s2ibinariesKind = schema.GroupVersionKind{Group: "devops.kubesphere.io", Version: "v1alpha1", Kind: "S2iBinary"}

// Get takes name of the s2iBinary, and returns the corresponding s2iBinary object, and an error if there is any.
func (c *FakeS2iBinaries) Get(name string, options v1.GetOptions) (result *v1alpha1.S2iBinary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(s2ibinariesResource, c.ns, name), &v1alpha1.S2iBinary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S2iBinary), err
}

// List takes label and field selectors, and returns the list of S2iBinaries that match those selectors.
func (c *FakeS2iBinaries) List(opts v1.ListOptions) (result *v1alpha1.S2iBinaryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(s2ibinariesResource, s2ibinariesKind, c.ns, opts), &v1alpha1.S2iBinaryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.S2iBinaryList{ListMeta: obj.(*v1alpha1.S2iBinaryList).ListMeta}
	for _, item := range obj.(*v1alpha1.S2iBinaryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested s2iBinaries.
func (c *FakeS2iBinaries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(s2ibinariesResource, c.ns, opts))

}

// Create takes the representation of a s2iBinary and creates it.  Returns the server's representation of the s2iBinary, and an error, if there is any.
func (c *FakeS2iBinaries) Create(s2iBinary *v1alpha1.S2iBinary) (result *v1alpha1.S2iBinary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(s2ibinariesResource, c.ns, s2iBinary), &v1alpha1.S2iBinary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S2iBinary), err
}

// Update takes the representation of a s2iBinary and updates it. Returns the server's representation of the s2iBinary, and an error, if there is any.
func (c *FakeS2iBinaries) Update(s2iBinary *v1alpha1.S2iBinary) (result *v1alpha1.S2iBinary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(s2ibinariesResource, c.ns, s2iBinary), &v1alpha1.S2iBinary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S2iBinary), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeS2iBinaries) UpdateStatus(s2iBinary *v1alpha1.S2iBinary) (*v1alpha1.S2iBinary, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(s2ibinariesResource, "status", c.ns, s2iBinary), &v1alpha1.S2iBinary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S2iBinary), err
}

// Delete takes name of the s2iBinary and deletes it. Returns an error if one occurs.
func (c *FakeS2iBinaries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(s2ibinariesResource, c.ns, name), &v1alpha1.S2iBinary{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeS2iBinaries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(s2ibinariesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.S2iBinaryList{})
	return err
}

// Patch applies the patch and returns the patched s2iBinary.
func (c *FakeS2iBinaries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S2iBinary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(s2ibinariesResource, c.ns, name, pt, data, subresources...), &v1alpha1.S2iBinary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S2iBinary), err
}

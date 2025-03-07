/*
Copyright 2023 Rancher Labs, Inc.

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
	"context"

	v1beta1 "github.com/harvester/node-manager/pkg/apis/node.harvesterhci.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudInits implements CloudInitInterface
type FakeCloudInits struct {
	Fake *FakeNodeV1beta1
}

var cloudinitsResource = schema.GroupVersionResource{Group: "node.harvesterhci.io", Version: "v1beta1", Resource: "cloudinits"}

var cloudinitsKind = schema.GroupVersionKind{Group: "node.harvesterhci.io", Version: "v1beta1", Kind: "CloudInit"}

// Get takes name of the cloudInit, and returns the corresponding cloudInit object, and an error if there is any.
func (c *FakeCloudInits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CloudInit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cloudinitsResource, name), &v1beta1.CloudInit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudInit), err
}

// List takes label and field selectors, and returns the list of CloudInits that match those selectors.
func (c *FakeCloudInits) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CloudInitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cloudinitsResource, cloudinitsKind, opts), &v1beta1.CloudInitList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CloudInitList{ListMeta: obj.(*v1beta1.CloudInitList).ListMeta}
	for _, item := range obj.(*v1beta1.CloudInitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudInits.
func (c *FakeCloudInits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cloudinitsResource, opts))
}

// Create takes the representation of a cloudInit and creates it.  Returns the server's representation of the cloudInit, and an error, if there is any.
func (c *FakeCloudInits) Create(ctx context.Context, cloudInit *v1beta1.CloudInit, opts v1.CreateOptions) (result *v1beta1.CloudInit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cloudinitsResource, cloudInit), &v1beta1.CloudInit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudInit), err
}

// Update takes the representation of a cloudInit and updates it. Returns the server's representation of the cloudInit, and an error, if there is any.
func (c *FakeCloudInits) Update(ctx context.Context, cloudInit *v1beta1.CloudInit, opts v1.UpdateOptions) (result *v1beta1.CloudInit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cloudinitsResource, cloudInit), &v1beta1.CloudInit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudInit), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudInits) UpdateStatus(ctx context.Context, cloudInit *v1beta1.CloudInit, opts v1.UpdateOptions) (*v1beta1.CloudInit, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(cloudinitsResource, "status", cloudInit), &v1beta1.CloudInit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudInit), err
}

// Delete takes name of the cloudInit and deletes it. Returns an error if one occurs.
func (c *FakeCloudInits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(cloudinitsResource, name, opts), &v1beta1.CloudInit{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudInits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cloudinitsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CloudInitList{})
	return err
}

// Patch applies the patch and returns the patched cloudInit.
func (c *FakeCloudInits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CloudInit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cloudinitsResource, name, pt, data, subresources...), &v1beta1.CloudInit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudInit), err
}

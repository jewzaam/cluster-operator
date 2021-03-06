/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/openshift/cluster-operator/pkg/apis/clusteroperator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterVersions implements ClusterVersionInterface
type FakeClusterVersions struct {
	Fake *FakeClusteroperatorV1alpha1
	ns   string
}

var clusterversionsResource = schema.GroupVersionResource{Group: "clusteroperator.openshift.io", Version: "v1alpha1", Resource: "clusterversions"}

var clusterversionsKind = schema.GroupVersionKind{Group: "clusteroperator.openshift.io", Version: "v1alpha1", Kind: "ClusterVersion"}

// Get takes name of the clusterVersion, and returns the corresponding clusterVersion object, and an error if there is any.
func (c *FakeClusterVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterversionsResource, c.ns, name), &v1alpha1.ClusterVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}

// List takes label and field selectors, and returns the list of ClusterVersions that match those selectors.
func (c *FakeClusterVersions) List(opts v1.ListOptions) (result *v1alpha1.ClusterVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterversionsResource, clusterversionsKind, c.ns, opts), &v1alpha1.ClusterVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterVersionList{}
	for _, item := range obj.(*v1alpha1.ClusterVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterVersions.
func (c *FakeClusterVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterversionsResource, c.ns, opts))

}

// Create takes the representation of a clusterVersion and creates it.  Returns the server's representation of the clusterVersion, and an error, if there is any.
func (c *FakeClusterVersions) Create(clusterVersion *v1alpha1.ClusterVersion) (result *v1alpha1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterversionsResource, c.ns, clusterVersion), &v1alpha1.ClusterVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}

// Update takes the representation of a clusterVersion and updates it. Returns the server's representation of the clusterVersion, and an error, if there is any.
func (c *FakeClusterVersions) Update(clusterVersion *v1alpha1.ClusterVersion) (result *v1alpha1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterversionsResource, c.ns, clusterVersion), &v1alpha1.ClusterVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterVersions) UpdateStatus(clusterVersion *v1alpha1.ClusterVersion) (*v1alpha1.ClusterVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterversionsResource, "status", c.ns, clusterVersion), &v1alpha1.ClusterVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}

// Delete takes name of the clusterVersion and deletes it. Returns an error if one occurs.
func (c *FakeClusterVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterversionsResource, c.ns, name), &v1alpha1.ClusterVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterversionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterVersionList{})
	return err
}

// Patch applies the patch and returns the patched clusterVersion.
func (c *FakeClusterVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterversionsResource, c.ns, name, data, subresources...), &v1alpha1.ClusterVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}

/*
Copyright The Kubernetes Authors.

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
	context "context"

	apiserverinternalv1alpha1 "k8s.io/api/apiserverinternal/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsapiserverinternalv1alpha1 "k8s.io/client-go/applyconfigurations/apiserverinternal/v1alpha1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// StorageVersionsGetter has a method to return a StorageVersionInterface.
// A group's client should implement this interface.
type StorageVersionsGetter interface {
	StorageVersions() StorageVersionInterface
}

// StorageVersionInterface has methods to work with StorageVersion resources.
type StorageVersionInterface interface {
	Create(ctx context.Context, storageVersion *apiserverinternalv1alpha1.StorageVersion, opts v1.CreateOptions) (*apiserverinternalv1alpha1.StorageVersion, error)
	Update(ctx context.Context, storageVersion *apiserverinternalv1alpha1.StorageVersion, opts v1.UpdateOptions) (*apiserverinternalv1alpha1.StorageVersion, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, storageVersion *apiserverinternalv1alpha1.StorageVersion, opts v1.UpdateOptions) (*apiserverinternalv1alpha1.StorageVersion, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*apiserverinternalv1alpha1.StorageVersion, error)
	List(ctx context.Context, opts v1.ListOptions) (*apiserverinternalv1alpha1.StorageVersionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiserverinternalv1alpha1.StorageVersion, err error)
	Apply(ctx context.Context, storageVersion *applyconfigurationsapiserverinternalv1alpha1.StorageVersionApplyConfiguration, opts v1.ApplyOptions) (result *apiserverinternalv1alpha1.StorageVersion, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, storageVersion *applyconfigurationsapiserverinternalv1alpha1.StorageVersionApplyConfiguration, opts v1.ApplyOptions) (result *apiserverinternalv1alpha1.StorageVersion, err error)
	StorageVersionExpansion
}

// storageVersions implements StorageVersionInterface
type storageVersions struct {
	*gentype.ClientWithListAndApply[*apiserverinternalv1alpha1.StorageVersion, *apiserverinternalv1alpha1.StorageVersionList, *applyconfigurationsapiserverinternalv1alpha1.StorageVersionApplyConfiguration]
}

// newStorageVersions returns a StorageVersions
func newStorageVersions(c *InternalV1alpha1Client) *storageVersions {
	return &storageVersions{
		gentype.NewClientWithListAndApply[*apiserverinternalv1alpha1.StorageVersion, *apiserverinternalv1alpha1.StorageVersionList, *applyconfigurationsapiserverinternalv1alpha1.StorageVersionApplyConfiguration](
			"storageversions",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *apiserverinternalv1alpha1.StorageVersion { return &apiserverinternalv1alpha1.StorageVersion{} },
			func() *apiserverinternalv1alpha1.StorageVersionList {
				return &apiserverinternalv1alpha1.StorageVersionList{}
			},
			gentype.PrefersProtobuf[*apiserverinternalv1alpha1.StorageVersion](),
		),
	}
}

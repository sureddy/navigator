// +build !ignore_autogenerated

/*
Copyright 2018 Jetstack Ltd.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CassandraCluster).DeepCopyInto(out.(*CassandraCluster))
			return nil
		}, InType: reflect.TypeOf(&CassandraCluster{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CassandraClusterList).DeepCopyInto(out.(*CassandraClusterList))
			return nil
		}, InType: reflect.TypeOf(&CassandraClusterList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CassandraClusterNodePool).DeepCopyInto(out.(*CassandraClusterNodePool))
			return nil
		}, InType: reflect.TypeOf(&CassandraClusterNodePool{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CassandraClusterNodePoolStatus).DeepCopyInto(out.(*CassandraClusterNodePoolStatus))
			return nil
		}, InType: reflect.TypeOf(&CassandraClusterNodePoolStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CassandraClusterSpec).DeepCopyInto(out.(*CassandraClusterSpec))
			return nil
		}, InType: reflect.TypeOf(&CassandraClusterSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CassandraClusterStatus).DeepCopyInto(out.(*CassandraClusterStatus))
			return nil
		}, InType: reflect.TypeOf(&CassandraClusterStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchCluster).DeepCopyInto(out.(*ElasticsearchCluster))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchCluster{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchClusterList).DeepCopyInto(out.(*ElasticsearchClusterList))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchClusterList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchClusterNodePool).DeepCopyInto(out.(*ElasticsearchClusterNodePool))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchClusterNodePool{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchClusterNodePoolStatus).DeepCopyInto(out.(*ElasticsearchClusterNodePoolStatus))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchClusterNodePoolStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchClusterPersistenceConfig).DeepCopyInto(out.(*ElasticsearchClusterPersistenceConfig))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchClusterPersistenceConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchClusterSpec).DeepCopyInto(out.(*ElasticsearchClusterSpec))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchClusterSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchClusterStatus).DeepCopyInto(out.(*ElasticsearchClusterStatus))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchClusterStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchImage).DeepCopyInto(out.(*ElasticsearchImage))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchImage{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchPilotImage).DeepCopyInto(out.(*ElasticsearchPilotImage))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchPilotImage{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchPilotStatus).DeepCopyInto(out.(*ElasticsearchPilotStatus))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchPilotStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ImageSpec).DeepCopyInto(out.(*ImageSpec))
			return nil
		}, InType: reflect.TypeOf(&ImageSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Pilot).DeepCopyInto(out.(*Pilot))
			return nil
		}, InType: reflect.TypeOf(&Pilot{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PilotCondition).DeepCopyInto(out.(*PilotCondition))
			return nil
		}, InType: reflect.TypeOf(&PilotCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PilotElasticsearchSpec).DeepCopyInto(out.(*PilotElasticsearchSpec))
			return nil
		}, InType: reflect.TypeOf(&PilotElasticsearchSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PilotList).DeepCopyInto(out.(*PilotList))
			return nil
		}, InType: reflect.TypeOf(&PilotList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PilotSpec).DeepCopyInto(out.(*PilotSpec))
			return nil
		}, InType: reflect.TypeOf(&PilotSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PilotStatus).DeepCopyInto(out.(*PilotStatus))
			return nil
		}, InType: reflect.TypeOf(&PilotStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraCluster) DeepCopyInto(out *CassandraCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraCluster.
func (in *CassandraCluster) DeepCopy() *CassandraCluster {
	if in == nil {
		return nil
	}
	out := new(CassandraCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterList) DeepCopyInto(out *CassandraClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterList.
func (in *CassandraClusterList) DeepCopy() *CassandraClusterList {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterNodePool) DeepCopyInto(out *CassandraClusterNodePool) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterNodePool.
func (in *CassandraClusterNodePool) DeepCopy() *CassandraClusterNodePool {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterNodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterNodePoolStatus) DeepCopyInto(out *CassandraClusterNodePoolStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterNodePoolStatus.
func (in *CassandraClusterNodePoolStatus) DeepCopy() *CassandraClusterNodePoolStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterNodePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterSpec) DeepCopyInto(out *CassandraClusterSpec) {
	*out = *in
	if in.Sysctl != nil {
		in, out := &in.Sysctl, &out.Sysctl
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]CassandraClusterNodePool, len(*in))
		copy(*out, *in)
	}
	out.Image = in.Image
	out.PilotImage = in.PilotImage
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterSpec.
func (in *CassandraClusterSpec) DeepCopy() *CassandraClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterStatus) DeepCopyInto(out *CassandraClusterStatus) {
	*out = *in
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make(map[string]CassandraClusterNodePoolStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterStatus.
func (in *CassandraClusterStatus) DeepCopy() *CassandraClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchCluster) DeepCopyInto(out *ElasticsearchCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchCluster.
func (in *ElasticsearchCluster) DeepCopy() *ElasticsearchCluster {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchClusterList) DeepCopyInto(out *ElasticsearchClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticsearchCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchClusterList.
func (in *ElasticsearchClusterList) DeepCopy() *ElasticsearchClusterList {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchClusterNodePool) DeepCopyInto(out *ElasticsearchClusterNodePool) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]ElasticsearchClusterRole, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.ResourceRequirements)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Persistence.DeepCopyInto(&out.Persistence)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchClusterNodePool.
func (in *ElasticsearchClusterNodePool) DeepCopy() *ElasticsearchClusterNodePool {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchClusterNodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchClusterNodePoolStatus) DeepCopyInto(out *ElasticsearchClusterNodePoolStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchClusterNodePoolStatus.
func (in *ElasticsearchClusterNodePoolStatus) DeepCopy() *ElasticsearchClusterNodePoolStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchClusterNodePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchClusterPersistenceConfig) DeepCopyInto(out *ElasticsearchClusterPersistenceConfig) {
	*out = *in
	out.Size = in.Size.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchClusterPersistenceConfig.
func (in *ElasticsearchClusterPersistenceConfig) DeepCopy() *ElasticsearchClusterPersistenceConfig {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchClusterPersistenceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchClusterSpec) DeepCopyInto(out *ElasticsearchClusterSpec) {
	*out = *in
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]ElasticsearchClusterNodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Pilot = in.Pilot
	out.Image = in.Image
	if in.Sysctl != nil {
		in, out := &in.Sysctl, &out.Sysctl
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchClusterSpec.
func (in *ElasticsearchClusterSpec) DeepCopy() *ElasticsearchClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchClusterStatus) DeepCopyInto(out *ElasticsearchClusterStatus) {
	*out = *in
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make(map[string]ElasticsearchClusterNodePoolStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchClusterStatus.
func (in *ElasticsearchClusterStatus) DeepCopy() *ElasticsearchClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchImage) DeepCopyInto(out *ElasticsearchImage) {
	*out = *in
	out.ImageSpec = in.ImageSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchImage.
func (in *ElasticsearchImage) DeepCopy() *ElasticsearchImage {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchPilotImage) DeepCopyInto(out *ElasticsearchPilotImage) {
	*out = *in
	out.ImageSpec = in.ImageSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchPilotImage.
func (in *ElasticsearchPilotImage) DeepCopy() *ElasticsearchPilotImage {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchPilotImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchPilotStatus) DeepCopyInto(out *ElasticsearchPilotStatus) {
	*out = *in
	if in.Documents != nil {
		in, out := &in.Documents, &out.Documents
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchPilotStatus.
func (in *ElasticsearchPilotStatus) DeepCopy() *ElasticsearchPilotStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchPilotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pilot) DeepCopyInto(out *Pilot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pilot.
func (in *Pilot) DeepCopy() *Pilot {
	if in == nil {
		return nil
	}
	out := new(Pilot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pilot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PilotCondition) DeepCopyInto(out *PilotCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PilotCondition.
func (in *PilotCondition) DeepCopy() *PilotCondition {
	if in == nil {
		return nil
	}
	out := new(PilotCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PilotElasticsearchSpec) DeepCopyInto(out *PilotElasticsearchSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PilotElasticsearchSpec.
func (in *PilotElasticsearchSpec) DeepCopy() *PilotElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(PilotElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PilotList) DeepCopyInto(out *PilotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pilot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PilotList.
func (in *PilotList) DeepCopy() *PilotList {
	if in == nil {
		return nil
	}
	out := new(PilotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PilotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PilotSpec) DeepCopyInto(out *PilotSpec) {
	*out = *in
	if in.Elasticsearch != nil {
		in, out := &in.Elasticsearch, &out.Elasticsearch
		if *in == nil {
			*out = nil
		} else {
			*out = new(PilotElasticsearchSpec)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PilotSpec.
func (in *PilotSpec) DeepCopy() *PilotSpec {
	if in == nil {
		return nil
	}
	out := new(PilotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PilotStatus) DeepCopyInto(out *PilotStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PilotCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Elasticsearch != nil {
		in, out := &in.Elasticsearch, &out.Elasticsearch
		if *in == nil {
			*out = nil
		} else {
			*out = new(ElasticsearchPilotStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PilotStatus.
func (in *PilotStatus) DeepCopy() *PilotStatus {
	if in == nil {
		return nil
	}
	out := new(PilotStatus)
	in.DeepCopyInto(out)
	return out
}

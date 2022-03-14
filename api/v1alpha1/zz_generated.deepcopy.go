//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InClusterIPPool) DeepCopyInto(out *InClusterIPPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InClusterIPPool.
func (in *InClusterIPPool) DeepCopy() *InClusterIPPool {
	if in == nil {
		return nil
	}
	out := new(InClusterIPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InClusterIPPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InClusterIPPoolList) DeepCopyInto(out *InClusterIPPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InClusterIPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InClusterIPPoolList.
func (in *InClusterIPPoolList) DeepCopy() *InClusterIPPoolList {
	if in == nil {
		return nil
	}
	out := new(InClusterIPPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InClusterIPPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InClusterIPPoolSpec) DeepCopyInto(out *InClusterIPPoolSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InClusterIPPoolSpec.
func (in *InClusterIPPoolSpec) DeepCopy() *InClusterIPPoolSpec {
	if in == nil {
		return nil
	}
	out := new(InClusterIPPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InClusterIPPoolStatus) DeepCopyInto(out *InClusterIPPoolStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InClusterIPPoolStatus.
func (in *InClusterIPPoolStatus) DeepCopy() *InClusterIPPoolStatus {
	if in == nil {
		return nil
	}
	out := new(InClusterIPPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// +build !ignore_autogenerated

// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFJob) DeepCopyInto(out *TFJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFJob.
func (in *TFJob) DeepCopy() *TFJob {
	if in == nil {
		return nil
	}
	out := new(TFJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TFJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFJobCondition) DeepCopyInto(out *TFJobCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFJobCondition.
func (in *TFJobCondition) DeepCopy() *TFJobCondition {
	if in == nil {
		return nil
	}
	out := new(TFJobCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFJobList) DeepCopyInto(out *TFJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TFJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFJobList.
func (in *TFJobList) DeepCopy() *TFJobList {
	if in == nil {
		return nil
	}
	out := new(TFJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TFJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFJobSpec) DeepCopyInto(out *TFJobSpec) {
	*out = *in
	if in.TFReplicaSpecs != nil {
		in, out := &in.TFReplicaSpecs, &out.TFReplicaSpecs
		*out = make(map[TFReplicaType]*TFReplicaSpec, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(TFReplicaSpec)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFJobSpec.
func (in *TFJobSpec) DeepCopy() *TFJobSpec {
	if in == nil {
		return nil
	}
	out := new(TFJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFJobStatus) DeepCopyInto(out *TFJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TFJobCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TFReplicaStatuses != nil {
		in, out := &in.TFReplicaStatuses, &out.TFReplicaStatuses
		*out = make(map[TFReplicaType]*TFReplicaStatus, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(TFReplicaStatus)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.LastReconcileTime != nil {
		in, out := &in.LastReconcileTime, &out.LastReconcileTime
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFJobStatus.
func (in *TFJobStatus) DeepCopy() *TFJobStatus {
	if in == nil {
		return nil
	}
	out := new(TFJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFReplicaSpec) DeepCopyInto(out *TFReplicaSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFReplicaSpec.
func (in *TFReplicaSpec) DeepCopy() *TFReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(TFReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFReplicaStatus) DeepCopyInto(out *TFReplicaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFReplicaStatus.
func (in *TFReplicaStatus) DeepCopy() *TFReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(TFReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

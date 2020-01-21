// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	strconf "github.com/gardener/test-infra/pkg/util/strconf"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigElement) DeepCopyInto(out *ConfigElement) {
	*out = *in
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = new(bool)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(strconf.ConfigSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigElement.
func (in *ConfigElement) DeepCopy() *ConfigElement {
	if in == nil {
		return nil
	}
	out := new(ConfigElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DAGStep) DeepCopyInto(out *DAGStep) {
	*out = *in
	in.Definition.DeepCopyInto(&out.Definition)
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pause != nil {
		in, out := &in.Pause, &out.Pause
		*out = new(Pause)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DAGStep.
func (in *DAGStep) DeepCopy() *DAGStep {
	if in == nil {
		return nil
	}
	out := new(DAGStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationSet) DeepCopyInto(out *LocationSet) {
	*out = *in
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]TestLocation, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationSet.
func (in *LocationSet) DeepCopy() *LocationSet {
	if in == nil {
		return nil
	}
	out := new(LocationSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pause) DeepCopyInto(out *Pause) {
	*out = *in
	if in.ResumeTimeoutSeconds != nil {
		in, out := &in.ResumeTimeoutSeconds, &out.ResumeTimeoutSeconds
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pause.
func (in *Pause) DeepCopy() *Pause {
	if in == nil {
		return nil
	}
	out := new(Pause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepDefinition) DeepCopyInto(out *StepDefinition) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigElement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LocationSet != nil {
		in, out := &in.LocationSet, &out.LocationSet
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepDefinition.
func (in *StepDefinition) DeepCopy() *StepDefinition {
	if in == nil {
		return nil
	}
	out := new(StepDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepStatus) DeepCopyInto(out *StepStatus) {
	*out = *in
	in.Position.DeepCopyInto(&out.Position)
	in.TestDefinition.DeepCopyInto(&out.TestDefinition)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepStatus.
func (in *StepStatus) DeepCopy() *StepStatus {
	if in == nil {
		return nil
	}
	out := new(StepStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepStatusPosition) DeepCopyInto(out *StepStatusPosition) {
	*out = *in
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepStatusPosition.
func (in *StepStatusPosition) DeepCopy() *StepStatusPosition {
	if in == nil {
		return nil
	}
	out := new(StepStatusPosition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepStatusTestDefinition) DeepCopyInto(out *StepStatusTestDefinition) {
	*out = *in
	out.Location = in.Location
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]*ConfigElement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ConfigElement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RecipientsOnFailure != nil {
		in, out := &in.RecipientsOnFailure, &out.RecipientsOnFailure
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepStatusTestDefinition.
func (in *StepStatusTestDefinition) DeepCopy() *StepStatusTestDefinition {
	if in == nil {
		return nil
	}
	out := new(StepStatusTestDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestDefMetadata) DeepCopyInto(out *TestDefMetadata) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestDefMetadata.
func (in *TestDefMetadata) DeepCopy() *TestDefMetadata {
	if in == nil {
		return nil
	}
	out := new(TestDefMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestDefSpec) DeepCopyInto(out *TestDefSpec) {
	*out = *in
	if in.RecipientsOnFailure != nil {
		in, out := &in.RecipientsOnFailure, &out.RecipientsOnFailure
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigElement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestDefSpec.
func (in *TestDefSpec) DeepCopy() *TestDefSpec {
	if in == nil {
		return nil
	}
	out := new(TestDefSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestDefinition) DeepCopyInto(out *TestDefinition) {
	*out = *in
	out.Metadata = in.Metadata
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestDefinition.
func (in *TestDefinition) DeepCopy() *TestDefinition {
	if in == nil {
		return nil
	}
	out := new(TestDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TestFlow) DeepCopyInto(out *TestFlow) {
	{
		in := &in
		*out = make(TestFlow, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DAGStep)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestFlow.
func (in TestFlow) DeepCopy() TestFlow {
	if in == nil {
		return nil
	}
	out := new(TestFlow)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestLocation) DeepCopyInto(out *TestLocation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestLocation.
func (in *TestLocation) DeepCopy() *TestLocation {
	if in == nil {
		return nil
	}
	out := new(TestLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Testrun) DeepCopyInto(out *Testrun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Testrun.
func (in *Testrun) DeepCopy() *Testrun {
	if in == nil {
		return nil
	}
	out := new(Testrun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Testrun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestrunKubeconfigs) DeepCopyInto(out *TestrunKubeconfigs) {
	*out = *in
	if in.Gardener != nil {
		in, out := &in.Gardener, &out.Gardener
		*out = new(strconf.StringOrConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Seed != nil {
		in, out := &in.Seed, &out.Seed
		*out = new(strconf.StringOrConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Shoot != nil {
		in, out := &in.Shoot, &out.Shoot
		*out = new(strconf.StringOrConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestrunKubeconfigs.
func (in *TestrunKubeconfigs) DeepCopy() *TestrunKubeconfigs {
	if in == nil {
		return nil
	}
	out := new(TestrunKubeconfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestrunList) DeepCopyInto(out *TestrunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Testrun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestrunList.
func (in *TestrunList) DeepCopy() *TestrunList {
	if in == nil {
		return nil
	}
	out := new(TestrunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestrunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestrunSpec) DeepCopyInto(out *TestrunSpec) {
	*out = *in
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	if in.TestLocations != nil {
		in, out := &in.TestLocations, &out.TestLocations
		*out = make([]TestLocation, len(*in))
		copy(*out, *in)
	}
	if in.LocationSets != nil {
		in, out := &in.LocationSets, &out.LocationSets
		*out = make([]LocationSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Kubeconfigs.DeepCopyInto(&out.Kubeconfigs)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigElement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TestFlow != nil {
		in, out := &in.TestFlow, &out.TestFlow
		*out = make(TestFlow, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DAGStep)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.OnExit != nil {
		in, out := &in.OnExit, &out.OnExit
		*out = make(TestFlow, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DAGStep)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestrunSpec.
func (in *TestrunSpec) DeepCopy() *TestrunSpec {
	if in == nil {
		return nil
	}
	out := new(TestrunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestrunStatus) DeepCopyInto(out *TestrunStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]*StepStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StepStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.UploadedToGithub != nil {
		in, out := &in.UploadedToGithub, &out.UploadedToGithub
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestrunStatus.
func (in *TestrunStatus) DeepCopy() *TestrunStatus {
	if in == nil {
		return nil
	}
	out := new(TestrunStatus)
	in.DeepCopyInto(out)
	return out
}

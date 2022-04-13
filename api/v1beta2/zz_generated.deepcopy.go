//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyAdmissionRule) DeepCopyInto(out *NetworkPolicyAdmissionRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyAdmissionRule.
func (in *NetworkPolicyAdmissionRule) DeepCopy() *NetworkPolicyAdmissionRule {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyAdmissionRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicyAdmissionRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyAdmissionRuleForbiddenEntity) DeepCopyInto(out *NetworkPolicyAdmissionRuleForbiddenEntity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyAdmissionRuleForbiddenEntity.
func (in *NetworkPolicyAdmissionRuleForbiddenEntity) DeepCopy() *NetworkPolicyAdmissionRuleForbiddenEntity {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyAdmissionRuleForbiddenEntity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyAdmissionRuleForbiddenIPRanges) DeepCopyInto(out *NetworkPolicyAdmissionRuleForbiddenIPRanges) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyAdmissionRuleForbiddenIPRanges.
func (in *NetworkPolicyAdmissionRuleForbiddenIPRanges) DeepCopy() *NetworkPolicyAdmissionRuleForbiddenIPRanges {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyAdmissionRuleForbiddenIPRanges)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyAdmissionRuleList) DeepCopyInto(out *NetworkPolicyAdmissionRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkPolicyAdmissionRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyAdmissionRuleList.
func (in *NetworkPolicyAdmissionRuleList) DeepCopy() *NetworkPolicyAdmissionRuleList {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyAdmissionRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicyAdmissionRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyAdmissionRuleNamespaceSelector) DeepCopyInto(out *NetworkPolicyAdmissionRuleNamespaceSelector) {
	*out = *in
	if in.ExcludeLabels != nil {
		in, out := &in.ExcludeLabels, &out.ExcludeLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyAdmissionRuleNamespaceSelector.
func (in *NetworkPolicyAdmissionRuleNamespaceSelector) DeepCopy() *NetworkPolicyAdmissionRuleNamespaceSelector {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyAdmissionRuleNamespaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyAdmissionRuleSpec) DeepCopyInto(out *NetworkPolicyAdmissionRuleSpec) {
	*out = *in
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	if in.ForbiddenIPRanges != nil {
		in, out := &in.ForbiddenIPRanges, &out.ForbiddenIPRanges
		*out = make([]NetworkPolicyAdmissionRuleForbiddenIPRanges, len(*in))
		copy(*out, *in)
	}
	if in.ForbiddenEntities != nil {
		in, out := &in.ForbiddenEntities, &out.ForbiddenEntities
		*out = make([]NetworkPolicyAdmissionRuleForbiddenEntity, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyAdmissionRuleSpec.
func (in *NetworkPolicyAdmissionRuleSpec) DeepCopy() *NetworkPolicyAdmissionRuleSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyAdmissionRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyTemplate) DeepCopyInto(out *NetworkPolicyTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyTemplate.
func (in *NetworkPolicyTemplate) DeepCopy() *NetworkPolicyTemplate {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicyTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyTemplateList) DeepCopyInto(out *NetworkPolicyTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkPolicyTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyTemplateList.
func (in *NetworkPolicyTemplateList) DeepCopy() *NetworkPolicyTemplateList {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicyTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyTemplateSpec) DeepCopyInto(out *NetworkPolicyTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyTemplateSpec.
func (in *NetworkPolicyTemplateSpec) DeepCopy() *NetworkPolicyTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

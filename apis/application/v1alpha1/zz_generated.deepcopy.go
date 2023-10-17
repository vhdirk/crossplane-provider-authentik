//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationInitParameters) DeepCopyInto(out *ApplicationInitParameters) {
	*out = *in
	if in.BackchannelProviders != nil {
		in, out := &in.BackchannelProviders, &out.BackchannelProviders
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.MetaDescription != nil {
		in, out := &in.MetaDescription, &out.MetaDescription
		*out = new(string)
		**out = **in
	}
	if in.MetaIcon != nil {
		in, out := &in.MetaIcon, &out.MetaIcon
		*out = new(string)
		**out = **in
	}
	if in.MetaLaunchURL != nil {
		in, out := &in.MetaLaunchURL, &out.MetaLaunchURL
		*out = new(string)
		**out = **in
	}
	if in.MetaPublisher != nil {
		in, out := &in.MetaPublisher, &out.MetaPublisher
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OpenInNewTab != nil {
		in, out := &in.OpenInNewTab, &out.OpenInNewTab
		*out = new(bool)
		**out = **in
	}
	if in.PolicyEngineMode != nil {
		in, out := &in.PolicyEngineMode, &out.PolicyEngineMode
		*out = new(string)
		**out = **in
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationInitParameters.
func (in *ApplicationInitParameters) DeepCopy() *ApplicationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationObservation) DeepCopyInto(out *ApplicationObservation) {
	*out = *in
	if in.BackchannelProviders != nil {
		in, out := &in.BackchannelProviders, &out.BackchannelProviders
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MetaDescription != nil {
		in, out := &in.MetaDescription, &out.MetaDescription
		*out = new(string)
		**out = **in
	}
	if in.MetaIcon != nil {
		in, out := &in.MetaIcon, &out.MetaIcon
		*out = new(string)
		**out = **in
	}
	if in.MetaLaunchURL != nil {
		in, out := &in.MetaLaunchURL, &out.MetaLaunchURL
		*out = new(string)
		**out = **in
	}
	if in.MetaPublisher != nil {
		in, out := &in.MetaPublisher, &out.MetaPublisher
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OpenInNewTab != nil {
		in, out := &in.OpenInNewTab, &out.OpenInNewTab
		*out = new(bool)
		**out = **in
	}
	if in.PolicyEngineMode != nil {
		in, out := &in.PolicyEngineMode, &out.PolicyEngineMode
		*out = new(string)
		**out = **in
	}
	if in.ProtocolProvider != nil {
		in, out := &in.ProtocolProvider, &out.ProtocolProvider
		*out = new(float64)
		**out = **in
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationObservation.
func (in *ApplicationObservation) DeepCopy() *ApplicationObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationParameters) DeepCopyInto(out *ApplicationParameters) {
	*out = *in
	if in.BackchannelProviders != nil {
		in, out := &in.BackchannelProviders, &out.BackchannelProviders
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.MetaDescription != nil {
		in, out := &in.MetaDescription, &out.MetaDescription
		*out = new(string)
		**out = **in
	}
	if in.MetaIcon != nil {
		in, out := &in.MetaIcon, &out.MetaIcon
		*out = new(string)
		**out = **in
	}
	if in.MetaLaunchURL != nil {
		in, out := &in.MetaLaunchURL, &out.MetaLaunchURL
		*out = new(string)
		**out = **in
	}
	if in.MetaPublisher != nil {
		in, out := &in.MetaPublisher, &out.MetaPublisher
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OpenInNewTab != nil {
		in, out := &in.OpenInNewTab, &out.OpenInNewTab
		*out = new(bool)
		**out = **in
	}
	if in.PolicyEngineMode != nil {
		in, out := &in.PolicyEngineMode, &out.PolicyEngineMode
		*out = new(string)
		**out = **in
	}
	if in.ProtocolProvider != nil {
		in, out := &in.ProtocolProvider, &out.ProtocolProvider
		*out = new(float64)
		**out = **in
	}
	if in.ProtocolProviderRef != nil {
		in, out := &in.ProtocolProviderRef, &out.ProtocolProviderRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProtocolProviderSelector != nil {
		in, out := &in.ProtocolProviderSelector, &out.ProtocolProviderSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationParameters.
func (in *ApplicationParameters) DeepCopy() *ApplicationParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}
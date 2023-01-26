//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
func (in *GatewayNetwork) DeepCopyInto(out *GatewayNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayNetwork.
func (in *GatewayNetwork) DeepCopy() *GatewayNetwork {
	if in == nil {
		return nil
	}
	out := new(GatewayNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayNetworkList) DeepCopyInto(out *GatewayNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GatewayNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayNetworkList.
func (in *GatewayNetworkList) DeepCopy() *GatewayNetworkList {
	if in == nil {
		return nil
	}
	out := new(GatewayNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayNetworkObservation) DeepCopyInto(out *GatewayNetworkObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayNetworkObservation.
func (in *GatewayNetworkObservation) DeepCopy() *GatewayNetworkObservation {
	if in == nil {
		return nil
	}
	out := new(GatewayNetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayNetworkParameters) DeepCopyInto(out *GatewayNetworkParameters) {
	*out = *in
	if in.CleanupDHCP != nil {
		in, out := &in.CleanupDHCP, &out.CleanupDHCP
		*out = new(bool)
		**out = **in
	}
	if in.DHCPID != nil {
		in, out := &in.DHCPID, &out.DHCPID
		*out = new(string)
		**out = **in
	}
	if in.DHCPIDRef != nil {
		in, out := &in.DHCPIDRef, &out.DHCPIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DHCPIDSelector != nil {
		in, out := &in.DHCPIDSelector, &out.DHCPIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.EnableDHCP != nil {
		in, out := &in.EnableDHCP, &out.EnableDHCP
		*out = new(bool)
		**out = **in
	}
	if in.EnableMasquerade != nil {
		in, out := &in.EnableMasquerade, &out.EnableMasquerade
		*out = new(bool)
		**out = **in
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.GatewayIDRef != nil {
		in, out := &in.GatewayIDRef, &out.GatewayIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GatewayIDSelector != nil {
		in, out := &in.GatewayIDSelector, &out.GatewayIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateNetworkID != nil {
		in, out := &in.PrivateNetworkID, &out.PrivateNetworkID
		*out = new(string)
		**out = **in
	}
	if in.PrivateNetworkIDRef != nil {
		in, out := &in.PrivateNetworkIDRef, &out.PrivateNetworkIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateNetworkIDSelector != nil {
		in, out := &in.PrivateNetworkIDSelector, &out.PrivateNetworkIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.StaticAddress != nil {
		in, out := &in.StaticAddress, &out.StaticAddress
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayNetworkParameters.
func (in *GatewayNetworkParameters) DeepCopy() *GatewayNetworkParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayNetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayNetworkSpec) DeepCopyInto(out *GatewayNetworkSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayNetworkSpec.
func (in *GatewayNetworkSpec) DeepCopy() *GatewayNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(GatewayNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayNetworkStatus) DeepCopyInto(out *GatewayNetworkStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayNetworkStatus.
func (in *GatewayNetworkStatus) DeepCopy() *GatewayNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetwork) DeepCopyInto(out *PrivateNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetwork.
func (in *PrivateNetwork) DeepCopy() *PrivateNetwork {
	if in == nil {
		return nil
	}
	out := new(PrivateNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetworkList) DeepCopyInto(out *PrivateNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrivateNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetworkList.
func (in *PrivateNetworkList) DeepCopy() *PrivateNetworkList {
	if in == nil {
		return nil
	}
	out := new(PrivateNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetworkObservation) DeepCopyInto(out *PrivateNetworkObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetworkObservation.
func (in *PrivateNetworkObservation) DeepCopy() *PrivateNetworkObservation {
	if in == nil {
		return nil
	}
	out := new(PrivateNetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetworkParameters) DeepCopyInto(out *PrivateNetworkParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetworkParameters.
func (in *PrivateNetworkParameters) DeepCopy() *PrivateNetworkParameters {
	if in == nil {
		return nil
	}
	out := new(PrivateNetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetworkSpec) DeepCopyInto(out *PrivateNetworkSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetworkSpec.
func (in *PrivateNetworkSpec) DeepCopy() *PrivateNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(PrivateNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetworkStatus) DeepCopyInto(out *PrivateNetworkStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetworkStatus.
func (in *PrivateNetworkStatus) DeepCopy() *PrivateNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(PrivateNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGateway) DeepCopyInto(out *PublicGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGateway.
func (in *PublicGateway) DeepCopy() *PublicGateway {
	if in == nil {
		return nil
	}
	out := new(PublicGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayDHCP) DeepCopyInto(out *PublicGatewayDHCP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayDHCP.
func (in *PublicGatewayDHCP) DeepCopy() *PublicGatewayDHCP {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayDHCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicGatewayDHCP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayDHCPList) DeepCopyInto(out *PublicGatewayDHCPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PublicGatewayDHCP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayDHCPList.
func (in *PublicGatewayDHCPList) DeepCopy() *PublicGatewayDHCPList {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayDHCPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicGatewayDHCPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayDHCPObservation) DeepCopyInto(out *PublicGatewayDHCPObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayDHCPObservation.
func (in *PublicGatewayDHCPObservation) DeepCopy() *PublicGatewayDHCPObservation {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayDHCPObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayDHCPParameters) DeepCopyInto(out *PublicGatewayDHCPParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.DNSLocalName != nil {
		in, out := &in.DNSLocalName, &out.DNSLocalName
		*out = new(string)
		**out = **in
	}
	if in.DNSSearch != nil {
		in, out := &in.DNSSearch, &out.DNSSearch
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DNSServersOverride != nil {
		in, out := &in.DNSServersOverride, &out.DNSServersOverride
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EnableDynamic != nil {
		in, out := &in.EnableDynamic, &out.EnableDynamic
		*out = new(bool)
		**out = **in
	}
	if in.PoolHigh != nil {
		in, out := &in.PoolHigh, &out.PoolHigh
		*out = new(string)
		**out = **in
	}
	if in.PoolLow != nil {
		in, out := &in.PoolLow, &out.PoolLow
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.PushDNSServer != nil {
		in, out := &in.PushDNSServer, &out.PushDNSServer
		*out = new(bool)
		**out = **in
	}
	if in.PushDefaultRoute != nil {
		in, out := &in.PushDefaultRoute, &out.PushDefaultRoute
		*out = new(bool)
		**out = **in
	}
	if in.RebindTimer != nil {
		in, out := &in.RebindTimer, &out.RebindTimer
		*out = new(float64)
		**out = **in
	}
	if in.RenewTimer != nil {
		in, out := &in.RenewTimer, &out.RenewTimer
		*out = new(float64)
		**out = **in
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(string)
		**out = **in
	}
	if in.ValidLifetime != nil {
		in, out := &in.ValidLifetime, &out.ValidLifetime
		*out = new(float64)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayDHCPParameters.
func (in *PublicGatewayDHCPParameters) DeepCopy() *PublicGatewayDHCPParameters {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayDHCPParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayDHCPSpec) DeepCopyInto(out *PublicGatewayDHCPSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayDHCPSpec.
func (in *PublicGatewayDHCPSpec) DeepCopy() *PublicGatewayDHCPSpec {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayDHCPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayDHCPStatus) DeepCopyInto(out *PublicGatewayDHCPStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayDHCPStatus.
func (in *PublicGatewayDHCPStatus) DeepCopy() *PublicGatewayDHCPStatus {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayDHCPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayIP) DeepCopyInto(out *PublicGatewayIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayIP.
func (in *PublicGatewayIP) DeepCopy() *PublicGatewayIP {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicGatewayIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayIPList) DeepCopyInto(out *PublicGatewayIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PublicGatewayIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayIPList.
func (in *PublicGatewayIPList) DeepCopy() *PublicGatewayIPList {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicGatewayIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayIPObservation) DeepCopyInto(out *PublicGatewayIPObservation) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayIPObservation.
func (in *PublicGatewayIPObservation) DeepCopy() *PublicGatewayIPObservation {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayIPObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayIPParameters) DeepCopyInto(out *PublicGatewayIPParameters) {
	*out = *in
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Reverse != nil {
		in, out := &in.Reverse, &out.Reverse
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayIPParameters.
func (in *PublicGatewayIPParameters) DeepCopy() *PublicGatewayIPParameters {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayIPParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayIPSpec) DeepCopyInto(out *PublicGatewayIPSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayIPSpec.
func (in *PublicGatewayIPSpec) DeepCopy() *PublicGatewayIPSpec {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayIPStatus) DeepCopyInto(out *PublicGatewayIPStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayIPStatus.
func (in *PublicGatewayIPStatus) DeepCopy() *PublicGatewayIPStatus {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayIPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayList) DeepCopyInto(out *PublicGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PublicGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayList.
func (in *PublicGatewayList) DeepCopy() *PublicGatewayList {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayObservation) DeepCopyInto(out *PublicGatewayObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayObservation.
func (in *PublicGatewayObservation) DeepCopy() *PublicGatewayObservation {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayPATRule) DeepCopyInto(out *PublicGatewayPATRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayPATRule.
func (in *PublicGatewayPATRule) DeepCopy() *PublicGatewayPATRule {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayPATRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicGatewayPATRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayPATRuleList) DeepCopyInto(out *PublicGatewayPATRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PublicGatewayPATRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayPATRuleList.
func (in *PublicGatewayPATRuleList) DeepCopy() *PublicGatewayPATRuleList {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayPATRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicGatewayPATRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayPATRuleObservation) DeepCopyInto(out *PublicGatewayPATRuleObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayPATRuleObservation.
func (in *PublicGatewayPATRuleObservation) DeepCopy() *PublicGatewayPATRuleObservation {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayPATRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayPATRuleParameters) DeepCopyInto(out *PublicGatewayPATRuleParameters) {
	*out = *in
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.GatewayIDRef != nil {
		in, out := &in.GatewayIDRef, &out.GatewayIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GatewayIDSelector != nil {
		in, out := &in.GatewayIDSelector, &out.GatewayIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(string)
		**out = **in
	}
	if in.PrivatePort != nil {
		in, out := &in.PrivatePort, &out.PrivatePort
		*out = new(float64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.PublicPort != nil {
		in, out := &in.PublicPort, &out.PublicPort
		*out = new(float64)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayPATRuleParameters.
func (in *PublicGatewayPATRuleParameters) DeepCopy() *PublicGatewayPATRuleParameters {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayPATRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayPATRuleSpec) DeepCopyInto(out *PublicGatewayPATRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayPATRuleSpec.
func (in *PublicGatewayPATRuleSpec) DeepCopy() *PublicGatewayPATRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayPATRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayPATRuleStatus) DeepCopyInto(out *PublicGatewayPATRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayPATRuleStatus.
func (in *PublicGatewayPATRuleStatus) DeepCopy() *PublicGatewayPATRuleStatus {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayPATRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayParameters) DeepCopyInto(out *PublicGatewayParameters) {
	*out = *in
	if in.BastionEnabled != nil {
		in, out := &in.BastionEnabled, &out.BastionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BastionPort != nil {
		in, out := &in.BastionPort, &out.BastionPort
		*out = new(float64)
		**out = **in
	}
	if in.EnableSMTP != nil {
		in, out := &in.EnableSMTP, &out.EnableSMTP
		*out = new(bool)
		**out = **in
	}
	if in.IPID != nil {
		in, out := &in.IPID, &out.IPID
		*out = new(string)
		**out = **in
	}
	if in.IPIDRef != nil {
		in, out := &in.IPIDRef, &out.IPIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IPIDSelector != nil {
		in, out := &in.IPIDSelector, &out.IPIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UpstreamDNSServers != nil {
		in, out := &in.UpstreamDNSServers, &out.UpstreamDNSServers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayParameters.
func (in *PublicGatewayParameters) DeepCopy() *PublicGatewayParameters {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewaySpec) DeepCopyInto(out *PublicGatewaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewaySpec.
func (in *PublicGatewaySpec) DeepCopy() *PublicGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(PublicGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicGatewayStatus) DeepCopyInto(out *PublicGatewayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicGatewayStatus.
func (in *PublicGatewayStatus) DeepCopy() *PublicGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(PublicGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InboundRuleObservation struct {
}

type InboundRuleParameters struct {

	// Action when rule match request (drop or accept)
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Ip address for this rule (e.g: 1.1.1.1). Only one of ip or ip_range should be provided
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Ip range for this rule (e.g: 192.168.1.0/24). Only one of ip or ip_range should be provided
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// Network port for this rule
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Computed port range for this rule (e.g: 1-1024, 22-22)
	// +kubebuilder:validation:Optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// Protocol for this rule (TCP, UDP, ICMP or ANY)
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type OutboundRuleObservation struct {
}

type OutboundRuleParameters struct {

	// Action when rule match request (drop or accept)
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Ip address for this rule (e.g: 1.1.1.1). Only one of ip or ip_range should be provided
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Ip range for this rule (e.g: 192.168.1.0/24). Only one of ip or ip_range should be provided
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// Network port for this rule
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Computed port range for this rule (e.g: 1-1024, 22-22)
	// +kubebuilder:validation:Optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// Protocol for this rule (TCP, UDP, ICMP or ANY)
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type SecurityGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

type SecurityGroupParameters struct {

	// The description of the security group
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Enable blocking of SMTP on IPv4 and IPv6
	// +kubebuilder:validation:Optional
	EnableDefaultSecurity *bool `json:"enableDefaultSecurity,omitempty" tf:"enable_default_security,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalRules *bool `json:"externalRules,omitempty" tf:"external_rules,omitempty"`

	// Default inbound traffic policy for this security group
	// +kubebuilder:validation:Optional
	InboundDefaultPolicy *string `json:"inboundDefaultPolicy,omitempty" tf:"inbound_default_policy,omitempty"`

	// Inbound rules for this security group
	// +kubebuilder:validation:Optional
	InboundRule []InboundRuleParameters `json:"inboundRule,omitempty" tf:"inbound_rule,omitempty"`

	// The name of the security group
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Default outbound traffic policy for this security group
	// +kubebuilder:validation:Optional
	OutboundDefaultPolicy *string `json:"outboundDefaultPolicy,omitempty" tf:"outbound_default_policy,omitempty"`

	// Outbound rules for this security group
	// +kubebuilder:validation:Optional
	OutboundRule []OutboundRuleParameters `json:"outboundRule,omitempty" tf:"outbound_rule,omitempty"`

	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The stateful value of the security group
	// +kubebuilder:validation:Optional
	Stateful *bool `json:"stateful,omitempty" tf:"stateful,omitempty"`

	// The tags associated with the security group
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroup is the Schema for the SecurityGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
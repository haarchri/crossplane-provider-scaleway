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

type MNQNamespaceObservation struct {

	// The date and time the Namespace was created.
	// The date and time of the creation of the mnq Namespace
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The endpoint of the service matching the Namespace protocol.
	// The endpoint of the service matching the Namespace protocol
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The ID of the namespace
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time the Namespace was updated.
	// The date and time of the last update of the mnq Namespace
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type MNQNamespaceParameters struct {

	// The unique name of the namespace.
	// The name of the mnq namespace
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the
	// namespace is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The protocol to apply on your namespace. Please check our
	// supported protocols.
	// The Namespace protocol
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// (Defaults to provider region). The region
	// in which the namespace should be created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// MNQNamespaceSpec defines the desired state of MNQNamespace
type MNQNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MNQNamespaceParameters `json:"forProvider"`
}

// MNQNamespaceStatus defines the observed state of MNQNamespace.
type MNQNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MNQNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MNQNamespace is the Schema for the MNQNamespaces API. Manages Scaleway Messaging and Queuing Namespaces.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type MNQNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MNQNamespaceSpec   `json:"spec"`
	Status            MNQNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MNQNamespaceList contains a list of MNQNamespaces
type MNQNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MNQNamespace `json:"items"`
}

// Repository type metadata.
var (
	MNQNamespace_Kind             = "MNQNamespace"
	MNQNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MNQNamespace_Kind}.String()
	MNQNamespace_KindAPIVersion   = MNQNamespace_Kind + "." + CRDGroupVersion.String()
	MNQNamespace_GroupVersionKind = CRDGroupVersion.WithKind(MNQNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&MNQNamespace{}, &MNQNamespaceList{})
}

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Credential
func (mg *Credential) GetTerraformResourceType() string {
	return "scaleway_mnq_credential"
}

// GetConnectionDetailsMapping for this Credential
func (tr *Credential) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"nats_credentials[*].content": "status.atProvider.natsCredentials[*].content", "sqs_sns_credentials[*].secret_key": "status.atProvider.sqsSnsCredentials[*].secretKey"}
}

// GetObservation of this Credential
func (tr *Credential) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Credential
func (tr *Credential) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Credential
func (tr *Credential) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Credential
func (tr *Credential) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Credential
func (tr *Credential) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Credential using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Credential) LateInitialize(attrs []byte) (bool, error) {
	params := &CredentialParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Credential) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MNQNamespace
func (mg *MNQNamespace) GetTerraformResourceType() string {
	return "scaleway_mnq_namespace"
}

// GetConnectionDetailsMapping for this MNQNamespace
func (tr *MNQNamespace) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MNQNamespace
func (tr *MNQNamespace) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MNQNamespace
func (tr *MNQNamespace) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MNQNamespace
func (tr *MNQNamespace) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MNQNamespace
func (tr *MNQNamespace) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MNQNamespace
func (tr *MNQNamespace) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MNQNamespace using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MNQNamespace) LateInitialize(attrs []byte) (bool, error) {
	params := &MNQNamespaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MNQNamespace) GetTerraformSchemaVersion() int {
	return 0
}

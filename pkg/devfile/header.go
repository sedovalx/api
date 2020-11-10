package devfile

import (
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)


// DevfileHeader describes the structure of the devfile-specific top-level fields
// that are not part of the K8S API structures
type DevfileHeader struct {
	// Devfile schema version
	// +kubebuilder:validation:Pattern=^([2-9])\.([0-9]+)\.([0-9]+)(\-[0-9a-z-]+(\.[0-9a-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?$
	SchemaVersion string `json:"schemaVersion"`

	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	// Optional metadata
	Metadata DevfileMetadata `json:"metadata,omitempty"`
}

// JSONObject provides a way to add a map of arbitrary YAML/JSON
// objects.
type JSONObject map[string]apiext.JSON

type DevfileMetadata struct {
	// Optional devfile name
	// +optional
	Name string `json:"name,omitempty"`

	// Optional semver-compatible version
	// +optional
	// +kubebuilder:validation:Pattern=^([0-9]+)\.([0-9]+)\.([0-9]+)(\-[0-9a-z-]+(\.[0-9a-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?$
	Version string `json:"version,omitempty"`

	JSONObject `json:",inline"`
}
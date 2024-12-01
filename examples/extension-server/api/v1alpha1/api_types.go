// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwapiv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
//
// ListenerContext provides an example extension policy context resource.
type API struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec APISpec `json:"spec"`
}

type APISpec struct {
	TargetRefs []gwapiv1a2.LocalPolicyTargetReferenceWithSectionName `json:"targetRefs"`

  Context string `json:"context"`
}

// +kubebuilder:object:root=true
//
// ListenerContextList contains a list of ListenerContext resources.
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ListenerContextExample `json:"items"`
}

func init() {
	SchemeBuilder.Register(&API{}, &APIList{})
}

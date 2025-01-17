/*
Copyright 2022.
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

// All fields in this package are required unless Explicitly marked optional
// +kubebuilder:validation:Required
package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AdminNetworkPolicySubject defines what resources the policy applies to.
// Exactly one field must be set.
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type AdminNetworkPolicySubject struct {
	// Namespaces is used to select pods via namespace selectors.
	// +optional
	Namespaces *metav1.LabelSelector `json:"namespaces,omitempty"`
	// Pods is used to select pods via namespace AND pod selectors.
	// +optional
	Pods *NamespacedPodSubject `json:"pods,omitempty"`
}

// NamespacedPodSubject allows the user to select a given set of pod(s) in
// selected namespace(s).
type NamespacedPodSubject struct {
	// NamespaceSelector follows standard label selector semantics; if empty,
	// it selects all Namespaces.
	NamespaceSelector metav1.LabelSelector `json:"namespaceSelector"`

	// PodSelector is used to explicitly select pods within a namespace; if empty,
	// it selects all Pods.
	PodSelector metav1.LabelSelector `json:"podSelector"`
}

// AdminNetworkPolicyPort describes how to select network ports on pod(s).
// Exactly one field must be set.
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type AdminNetworkPolicyPort struct {
	// Port selects a port on a pod(s) based on number.
	//
	// Support: Core
	//
	// +optional
	PortNumber *Port `json:"portNumber,omitempty"`

	// NamedPort selects a port on a pod(s) based on name.
	//
	// Support: Extended
	//
	// <network-policy-api:experimental>
	// +optional
	NamedPort *string `json:"namedPort,omitempty"`

	// PortRange selects a port range on a pod(s) based on provided start and end
	// values.
	//
	// Support: Core
	//
	// +optional
	PortRange *PortRange `json:"portRange,omitempty"`
}

type Port struct {
	// Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must
	// match. If not specified, this field defaults to TCP.
	//
	// Support: Core
	//
	Protocol v1.Protocol `json:"protocol"`

	// Number defines a network port value.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	//
	// Support: Core
	//
	Port int32 `json:"port"`
}

// PortRange defines an inclusive range of ports from the the assigned Start value
// to End value.
type PortRange struct {
	// Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must
	// match. If not specified, this field defaults to TCP.
	//
	// Support: Core
	//
	Protocol v1.Protocol `json:"protocol,omitempty"`

	// Start defines a network port that is the start of a port range, the Start
	// value must be less than End.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	//
	// Support: Core
	//
	Start int32 `json:"start"`

	// End defines a network port that is the end of a port range, the End value
	// must be greater than Start.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	//
	// Support: Core
	//
	End int32 `json:"end"`
}

// AdminNetworkPolicyIngressPeer defines an in-cluster peer to allow traffic from.
// Exactly one of the selector pointers must be set for a given peer. If a
// consumer observes none of its fields are set, they must assume an unknown
// option has been specified and fail closed.
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type AdminNetworkPolicyIngressPeer struct {
	// Namespaces defines a way to select all pods within a set of Namespaces.
	// Note that host-networked pods are not included in this type of peer.
	//
	// Support: Core
	//
	// +optional
	Namespaces *NamespacedPeer `json:"namespaces,omitempty"`
	// Pods defines a way to select a set of pods in
	// a set of namespaces. Note that host-networked pods
	// are not included in this type of peer.
	//
	// Support: Core
	//
	// +optional
	Pods *NamespacedPodPeer `json:"pods,omitempty"`
}

// AdminNetworkPolicyEgressPeer defines a peer to allow traffic to.
// Exactly one of the selector pointers must be set for a given peer. If a
// consumer observes none of its fields are set, they must assume an unknown
// option has been specified and fail closed.
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type AdminNetworkPolicyEgressPeer struct {
	// Namespaces defines a way to select all pods within a set of Namespaces.
	// Note that host-networked pods are not included in this type of peer.
	//
	// Support: Core
	//
	// +optional
	Namespaces *NamespacedPeer `json:"namespaces,omitempty"`
	// Pods defines a way to select a set of pods in
	// a set of namespaces. Note that host-networked pods
	// are not included in this type of peer.
	//
	// Support: Core
	//
	// +optional
	Pods *NamespacedPodPeer `json:"pods,omitempty"`
	// Nodes defines a way to select a set of nodes in
	// the cluster. This field follows standard label selector
	// semantics; if present but empty, it selects all Nodes.
	//
	// Support: Extended
	//
	// <network-policy-api:experimental>
	// +optional
	Nodes *metav1.LabelSelector `json:"nodes,omitempty"`
}

// NamespacedPeer defines a flexible way to select Namespaces in a cluster.
// Exactly one of the selectors must be set.  If a consumer observes none of
// its fields are set, they must assume an unknown option has been specified
// and fail closed.
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type NamespacedPeer struct {
	// NamespaceSelector is a labelSelector used to select Namespaces, This field
	// follows standard label selector semantics; if present but empty, it selects
	// all Namespaces.
	//
	// Support: Core
	//
	// +optional
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`

	// SameLabels is used to select a set of Namespaces that share the same values
	// for a set of labels.
	// To be selected a Namespace must have all of the labels defined in SameLabels,
	// AND they must all have the same value as the subject of this policy.
	// If Samelabels is Empty then nothing is selected.
	//
	// Support: Extended
	//
	// <network-policy-api:experimental>
	// +optional
	// +kubebuilder:validation:MaxItems=100
	SameLabels []string `json:"sameLabels,omitempty"`

	// NotSameLabels is used to select a set of Namespaces that do not have certain
	// values for a set of label(s).
	// To be selected a Namespace must have all of the labels defined in NotSameLabels,
	// AND at least one of them must have different values than the subject of this policy.
	// If NotSameLabels is empty then nothing is selected.
	//
	// Support: Extended
	//
	// <network-policy-api:experimental>
	// +optional
	// +kubebuilder:validation:MaxItems=100
	NotSameLabels []string `json:"notSameLabels,omitempty"`
}

// NamespacedPodPeer defines a flexible way to select Namespaces and pods in a
// cluster. The `Namespaces` and `PodSelector` fields are required.
type NamespacedPodPeer struct {
	// Namespaces is used to select a set of Namespaces.
	//
	// Support: Core
	//
	Namespaces NamespacedPeer `json:"namespaces"`

	// PodSelector is a labelSelector used to select Pods, This field is NOT optional,
	// follows standard label selector semantics and if present but empty, it selects
	// all Pods.
	//
	// Support: Core
	//
	PodSelector metav1.LabelSelector `json:"podSelector"`
}

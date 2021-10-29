/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type SubscriptionPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionPeeringSpec   `json:"spec,omitempty"`
	Status            SubscriptionPeeringStatus `json:"status,omitempty"`
}

type SubscriptionPeeringSpec struct {
	State *SubscriptionPeeringSpecResource `json:"state,omitempty" tf:"-"`

	Resource SubscriptionPeeringSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SubscriptionPeeringSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// AWS account id that the VPC to be peered lives in
	// +optional
	AwsAccountID *string `json:"awsAccountID,omitempty" tf:"aws_account_id"`
	// Identifier of the AWS cloud peering
	// +optional
	AwsPeeringID *string `json:"awsPeeringID,omitempty" tf:"aws_peering_id"`
	// The name of the network to be peered
	// +optional
	GcpNetworkName *string `json:"gcpNetworkName,omitempty" tf:"gcp_network_name"`
	// Identifier of the cloud peering
	// +optional
	GcpPeeringID *string `json:"gcpPeeringID,omitempty" tf:"gcp_peering_id"`
	// GCP project ID that the VPC to be peered lives in
	// +optional
	GcpProjectID *string `json:"gcpProjectID,omitempty" tf:"gcp_project_id"`
	// The name of the Redis Enterprise Cloud network to be peered
	// +optional
	GcpRedisNetworkName *string `json:"gcpRedisNetworkName,omitempty" tf:"gcp_redis_network_name"`
	// Identifier of the Redis Enterprise Cloud GCP project to be peered
	// +optional
	GcpRedisProjectID *string `json:"gcpRedisProjectID,omitempty" tf:"gcp_redis_project_id"`
	// The cloud provider to use with the vpc peering, (either `AWS` or `GCP`)
	// +optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name"`
	// AWS Region that the VPC to be peered lives in
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// Current status of the account - `initiating-request`, `pending-acceptance`, `active`, `inactive` or `failed`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// A valid subscription predefined in the current account
	SubscriptionID *string `json:"subscriptionID" tf:"subscription_id"`
	// CIDR range of the VPC to be peered
	// +optional
	VpcCIDR *string `json:"vpcCIDR,omitempty" tf:"vpc_cidr"`
	// Identifier of the VPC to be peered
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type SubscriptionPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SubscriptionPeeringList is a list of SubscriptionPeerings
type SubscriptionPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SubscriptionPeering CRD objects
	Items []SubscriptionPeering `json:"items,omitempty"`
}

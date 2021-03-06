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

type Subscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec,omitempty"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

type SubscriptionSpecAllowlist struct {
	// Set of CIDR ranges that are allowed to access the databases associated with this subscription
	// +optional
	Cidrs []string `json:"cidrs,omitempty" tf:"cidrs"`
	// Set of security groups that are allowed to access the databases associated with this subscription
	// +optional
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids"`
}

type SubscriptionSpecCloudProviderRegionNetworks struct {
	// Deployment CIDR mask
	// +optional
	NetworkingDeploymentCIDR *string `json:"networkingDeploymentCIDR,omitempty" tf:"networking_deployment_cidr"`
	// The subnet that the subscription deploys into
	// +optional
	NetworkingSubnetID *string `json:"networkingSubnetID,omitempty" tf:"networking_subnet_id"`
	// Either an existing VPC Id (already exists in the specific region) or create a new VPC (if no VPC is specified)
	// +optional
	NetworkingVpcID *string `json:"networkingVpcID,omitempty" tf:"networking_vpc_id"`
}

type SubscriptionSpecCloudProviderRegion struct {
	// Support deployment on multiple availability zones within the selected region
	// +optional
	MultipleAvailabilityZones *bool `json:"multipleAvailabilityZones,omitempty" tf:"multiple_availability_zones"`
	// Deployment CIDR mask
	NetworkingDeploymentCIDR *string `json:"networkingDeploymentCIDR" tf:"networking_deployment_cidr"`
	// Either an existing VPC Id (already exists in the specific region) or create a new VPC (if no VPC is specified)
	// +optional
	NetworkingVpcID *string `json:"networkingVpcID,omitempty" tf:"networking_vpc_id"`
	// List of networks used
	// +optional
	Networks []SubscriptionSpecCloudProviderRegionNetworks `json:"networks,omitempty" tf:"networks"`
	// List of availability zones used
	PreferredAvailabilityZones []string `json:"preferredAvailabilityZones" tf:"preferred_availability_zones"`
	// Deployment region as defined by cloud provider
	Region *string `json:"region" tf:"region"`
}

type SubscriptionSpecCloudProvider struct {
	// Cloud account identifier. Default: Redis Labs internal cloud account (using Cloud Account Id = 1 implies using Redis Labs internal cloud account). Note that a GCP subscription can be created only with Redis Labs internal cloud account
	// +optional
	CloudAccountID *string `json:"cloudAccountID,omitempty" tf:"cloud_account_id"`
	// The cloud provider to use with the subscription, (either `AWS` or `GCP`)
	// +optional
	Provider *string `json:"provider,omitempty" tf:"provider"`
	// Cloud networking details, per region (single region or multiple regions for Active-Active cluster only)
	// +kubebuilder:validation:MinItems=1
	Region []SubscriptionSpecCloudProviderRegion `json:"region" tf:"region"`
}

type SubscriptionSpecDatabaseAlert struct {
	// Alert name
	Name *string `json:"name" tf:"name"`
	// Alert value
	Value *int64 `json:"value" tf:"value"`
}

type SubscriptionSpecDatabaseModule struct {
	// Name of the module to enable
	Name *string `json:"name" tf:"name"`
}

type SubscriptionSpecDatabase struct {
	// Set of alerts to enable on the database
	// +optional
	Alert []SubscriptionSpecDatabaseAlert `json:"alert,omitempty" tf:"alert"`
	// Relevant only to ram-and-flash clusters. Estimated average size (measured in bytes) of the items stored in the database
	// +optional
	AverageItemSizeInBytes *int64 `json:"averageItemSizeInBytes,omitempty" tf:"average_item_size_in_bytes"`
	// SSL certificate to authenticate user connections
	// +optional
	ClientSslCertificate *string `json:"clientSslCertificate,omitempty" tf:"client_ssl_certificate"`
	// Rate of database data persistence (in persistent storage)
	// +optional
	DataPersistence *string `json:"dataPersistence,omitempty" tf:"data_persistence"`
	// Identifier of the database created
	// +optional
	DbID *int64 `json:"dbID,omitempty" tf:"db_id"`
	// Use TLS for authentication
	// +optional
	EnableTls *bool `json:"enableTls,omitempty" tf:"enable_tls"`
	// Should use the external endpoint for open-source (OSS) Cluster API
	// +optional
	ExternalEndpointForOssClusterAPI *bool `json:"externalEndpointForOssClusterAPI,omitempty" tf:"external_endpoint_for_oss_cluster_api"`
	// List of regular expression rules to shard the database by. See the documentation on clustering for more information on the hashing policy - https://docs.redislabs.com/latest/rc/concepts/clustering/
	// +optional
	HashingPolicy []string `json:"hashingPolicy,omitempty" tf:"hashing_policy"`
	// Maximum memory usage for this specific database
	MemoryLimitInGb *float64 `json:"memoryLimitInGb" tf:"memory_limit_in_gb"`
	// A module object
	// +optional
	Module *SubscriptionSpecDatabaseModule `json:"module,omitempty" tf:"module"`
	// A meaningful name to identify the database
	Name *string `json:"name" tf:"name"`
	// Password used to access the database
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// Path that will be used to store database backup files
	// +optional
	PeriodicBackupPath *string `json:"periodicBackupPath,omitempty" tf:"periodic_backup_path"`
	// Private endpoint to access the database
	// +optional
	PrivateEndpoint *string `json:"privateEndpoint,omitempty" tf:"private_endpoint"`
	// The protocol that will be used to access the database, (either ???redis??? or 'memcached???)
	Protocol *string `json:"protocol" tf:"protocol"`
	// Public endpoint to access the database
	// +optional
	PublicEndpoint *string `json:"publicEndpoint,omitempty" tf:"public_endpoint"`
	// Set of Redis database URIs, in the format `redis://user:password@host:port`, that this database will be a replica of. If the URI provided is Redis Labs Cloud instance, only host and port should be provided
	// +optional
	ReplicaOf []string `json:"replicaOf,omitempty" tf:"replica_of"`
	// Databases replication
	// +optional
	Replication *bool `json:"replication,omitempty" tf:"replication"`
	// Set of CIDR addresses to allow access to the database
	// +optional
	// +kubebuilder:validation:MinItems=1
	SourceIPS []string `json:"sourceIPS,omitempty" tf:"source_ips"`
	// Support Redis open-source (OSS) Cluster API
	// +optional
	SupportOssClusterAPI *bool `json:"supportOssClusterAPI,omitempty" tf:"support_oss_cluster_api"`
	// Throughput measurement method, (either ???number-of-shards??? or ???operations-per-second???)
	ThroughputMeasurementBy *string `json:"throughputMeasurementBy" tf:"throughput_measurement_by"`
	// Throughput value (as applies to selected measurement method)
	ThroughputMeasurementValue *int64 `json:"throughputMeasurementValue" tf:"throughput_measurement_value"`
}

type SubscriptionSpec struct {
	State *SubscriptionSpecResource `json:"state,omitempty" tf:"-"`

	Resource SubscriptionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SubscriptionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// An allowlist object
	// +optional
	Allowlist *SubscriptionSpecAllowlist `json:"allowlist,omitempty" tf:"allowlist"`
	// A cloud provider object
	CloudProvider *SubscriptionSpecCloudProvider `json:"cloudProvider" tf:"cloud_provider"`
	// A database object
	// +kubebuilder:validation:MinItems=1
	Database []SubscriptionSpecDatabase `json:"database" tf:"database"`
	// Memory storage preference: either ???ram??? or a combination of 'ram-and-flash???
	// +optional
	MemoryStorage *string `json:"memoryStorage,omitempty" tf:"memory_storage"`
	// A meaningful name to identify the subscription
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// A valid payment method pre-defined in the current account
	// +optional
	PaymentMethodID *string `json:"paymentMethodID,omitempty" tf:"payment_method_id"`
	// Encrypt data stored in persistent storage. Required for a GCP subscription
	// +optional
	PersistentStorageEncryption *bool `json:"persistentStorageEncryption,omitempty" tf:"persistent_storage_encryption"`
}

type SubscriptionStatus struct {
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

// SubscriptionList is a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Subscription CRD objects
	Items []Subscription `json:"items,omitempty"`
}

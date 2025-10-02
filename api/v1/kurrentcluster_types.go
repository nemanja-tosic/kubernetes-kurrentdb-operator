package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json:"-" or json:"fieldName"

// KurrentClusterSpec defines the desired state of KurrentCluster
type KurrentClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Size defines the number of KurrentDB nodes in the cluster
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=7
	// +kubebuilder:default=3
	Size int32 `json:"size,omitempty"`

	// Image defines the KurrentDB container image to use
	// +kubebuilder:default="docker.kurrent.io/kurrent-latest/kurrentdb:latest"
	Image string `json:"image,omitempty"`

	// ImagePullPolicy defines the image pull policy
	// +kubebuilder:validation:Enum=Always;Never;IfNotPresent
	// +kubebuilder:default=IfNotPresent
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// Resources defines the resource requirements for KurrentDB containers
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// Storage defines the storage configuration for KurrentDB data
	Storage StorageSpec `json:"storage,omitempty"`

	// TLS defines the TLS configuration for the cluster
	TLS TLSSpec `json:"tls,omitempty"`

	// Environment defines additional environment variables for KurrentDB
	Environment []corev1.EnvVar `json:"environment,omitempty"`

	// Network defines network configuration for the cluster
	Network NetworkSpec `json:"network,omitempty"`

	// NodeSelector defines node selection constraints
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Tolerations defines pod tolerations
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// Affinity defines pod affinity rules
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// ServiceAccountName defines the service account to use
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
}

// StorageSpec defines storage configuration
type StorageSpec struct {
	// Size defines the size of persistent volume for each node
	// +kubebuilder:default="10Gi"
	Size string `json:"size,omitempty"`

	// StorageClassName defines the storage class to use
	StorageClassName *string `json:"storageClassName,omitempty"`

	// AccessModes defines the access modes for the persistent volume
	// +kubebuilder:default={"ReadWriteOnce"}
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes,omitempty"`
}

// TLSSpec defines TLS configuration
type TLSSpec struct {
	// Enabled defines whether TLS is enabled
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// CertificatesSecretName defines the name of the secret containing certificates
	// If not provided, certificates will be auto-generated
	CertificatesSecretName string `json:"certificatesSecretName,omitempty"`

	// AutoGenerate defines whether to auto-generate certificates
	// +kubebuilder:default=true
	AutoGenerate bool `json:"autoGenerate,omitempty"`

	// CASecretName defines the name of the secret containing CA certificate
	// Used when autoGenerate is false
	CASecretName string `json:"caSecretName,omitempty"`
}

// NetworkSpec defines network configuration
type NetworkSpec struct {
	// ServiceType defines the type of service to create
	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer
	// +kubebuilder:default=ClusterIP
	ServiceType corev1.ServiceType `json:"serviceType,omitempty"`

	// ExternalPort defines the external port for client connections
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default=2113
	ExternalPort int32 `json:"externalPort,omitempty"`

	// InternalPort defines the internal port for client connections
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default=2113
	InternalPort int32 `json:"internalPort,omitempty"`

	// GossipPort defines the port for gossip communication
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default=2113
	GossipPort int32 `json:"gossipPort,omitempty"`

	// Annotations defines annotations to add to services
	Annotations map[string]string `json:"annotations,omitempty"`
}

// KurrentClusterStatus defines the observed state of KurrentCluster
type KurrentClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conditions represent the latest available observations of the cluster's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Phase represents the current phase of the cluster
	Phase ClusterPhase `json:"phase,omitempty"`

	// ReadyReplicas represents the number of ready replicas
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// Replicas represents the total number of replicas
	Replicas int32 `json:"replicas,omitempty"`

	// ObservedGeneration represents the generation observed by the controller
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Nodes represents the status of individual nodes
	Nodes []NodeStatus `json:"nodes,omitempty"`

	// ServiceEndpoints represents the endpoints for accessing the cluster
	ServiceEndpoints ServiceEndpoints `json:"serviceEndpoints,omitempty"`
}

// ClusterPhase represents the phase of the cluster
type ClusterPhase string

const (
	// ClusterPhaseCreating indicates the cluster is being created
	ClusterPhaseCreating ClusterPhase = "Creating"
	// ClusterPhaseRunning indicates the cluster is running
	ClusterPhaseRunning ClusterPhase = "Running"
	// ClusterPhaseUpdating indicates the cluster is being updated
	ClusterPhaseUpdating ClusterPhase = "Updating"
	// ClusterPhaseFailed indicates the cluster has failed
	ClusterPhaseFailed ClusterPhase = "Failed"
	// ClusterPhaseDeleting indicates the cluster is being deleted
	ClusterPhaseDeleting ClusterPhase = "Deleting"
)

// NodeStatus represents the status of a single node
type NodeStatus struct {
	// Name is the name of the node
	Name string `json:"name"`
	// Ready indicates if the node is ready
	Ready bool `json:"ready"`
	// Role indicates the role of the node in the cluster
	Role string `json:"role,omitempty"`
	// Message provides additional information about the node status
	Message string `json:"message,omitempty"`
}

// ServiceEndpoints represents the endpoints for accessing the cluster
type ServiceEndpoints struct {
	// External represents external endpoints
	External []string `json:"external,omitempty"`
	// Internal represents internal endpoints
	Internal []string `json:"internal,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=kc
//+kubebuilder:printcolumn:name="Size",type="integer",JSONPath=".spec.size"
//+kubebuilder:printcolumn:name="Ready",type="integer",JSONPath=".status.readyReplicas"
//+kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// KurrentCluster is the Schema for the kurrentclusters API
type KurrentCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KurrentClusterSpec   `json:"spec,omitempty"`
	Status KurrentClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KurrentClusterList contains a list of KurrentCluster
type KurrentClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KurrentCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KurrentCluster{}, &KurrentClusterList{})
}

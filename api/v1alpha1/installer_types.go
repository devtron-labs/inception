/*
Copyright 2020 Devtron Labs Pvt Ltd.

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

package v1alpha1

import (
	"github.com/argoproj/gitops-engine/pkg/health"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InstallerSpec defines the desired state of Installer
type InstallerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// URL of the BOM version to be installed
	URL string `json:"url,omitempty" protobuf:"bytes,1,opt,name=url"`
}

// InstallerStatus defines the observed state of Installer
type InstallerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CurrentSpecHash string     `json:"current_spec_hash"`
	Sync            SyncStatus `json:"sync"`
}

// SyncStatus is a comparison result of application spec and deployed application.
type SyncStatus struct {
	Resources []ResourceStatus `json:"resources,omitempty" protobuf:"bytes,1,opt,name=resources"`
	Status    SyncStatusCode   `json:"status" protobuf:"bytes,2,opt,name=status,casttype=SyncStatusCode"`
	// URL of the BOM version pulled
	URL        string               `json:"url,omitempty" protobuf:"bytes,3,opt,name=url"`
	Data       string               `json:"data,omitempty" protobuf:"bytes,4,opt,name=data"`
	Conditions []InstallerCondition `json:"conditions,omitempty" protobuf:"bytes,5,opt,name=conditions"`
	Health     HealthStatus         `json:"health,omitempty" protobuf:"bytes,6,opt,name=health"`
	History    RevisionHistories    `json:"history,omitempty" protobuf:"bytes,7,opt,name=history"`
}

// ResourceStatus holds the current sync and health status of a resource
type ResourceStatus struct {
	Group     string         `json:"group,omitempty" protobuf:"bytes,1,opt,name=group"`
	Version   string         `json:"version,omitempty" protobuf:"bytes,2,opt,name=version"`
	Kind      string         `json:"kind,omitempty" protobuf:"bytes,3,opt,name=kind"`
	Namespace string         `json:"namespace,omitempty" protobuf:"bytes,4,opt,name=namespace"`
	Name      string         `json:"name,omitempty" protobuf:"bytes,5,opt,name=name"`
	Status    SyncStatusCode `json:"status,omitempty" protobuf:"bytes,6,opt,name=status"`
	Health    *HealthStatus  `json:"health,omitempty" protobuf:"bytes,7,opt,name=health"`
	Operation string         `json:"operation,omitempty" protobuf:"bytes,8,opt,name=operation"`
}

type HealthStatus struct {
	Status  health.HealthStatusCode `json:"status,omitempty" protobuf:"bytes,1,opt,name=status"`
	Message string                  `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
}

type SyncStatusCode string

const (
	SyncStatusCodeUnknown    SyncStatusCode = "Unknown"
	SyncStatusCodeSynced     SyncStatusCode = "Synced"
	SyncStatusCodeOutOfSync  SyncStatusCode = "OutOfSync"
	SyncStatusCodeApplied    SyncStatusCode = "Applied"
	SyncStatusCodeDownloaded SyncStatusCode = "Downloaded"
)

// InstallerCondition contains details about current application condition
type InstallerCondition struct {
	// Type is an application condition type
	Type InstallerConditionType `json:"type" protobuf:"bytes,1,opt,name=type"`
	// Message contains human-readable message indicating details about condition
	Message string `json:"message" protobuf:"bytes,2,opt,name=message"`
	// LastTransitionTime is the time the condition was first observed.
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
}

// InstallerConditionType represents type of application condition. Type name has following convention:
// suffix "Error" means error condition
// suffix "Warning" means warning condition
// suffix "Info" means informational condition
type InstallerConditionType string

const (
	// InstallerConditionDeletionError indicates that controller failed to delete application
	InstallerConditionDeletionError = "DeletionError"
	// InstallerConditionInvalidSpecError indicates that application source is invalid
	InstallerConditionInvalidSpecError = "InvalidSpecError"
	// InstallerConditionComparisonError indicates controller failed to compare application state
	InstallerConditionComparisonError = "ComparisonError"
	// InstallerConditionSyncError indicates controller failed to automatically sync the application
	InstallerConditionSyncError = "SyncError"
	// InstallerConditionUnknownError indicates an unknown controller error
	InstallerConditionUnknownError = "UnknownError"
	// InstallerConditionSharedResourceWarning indicates that controller detected resources which belongs to more than one application
	InstallerConditionSharedResourceWarning = "SharedResourceWarning"
	// InstallerConditionRepeatedResourceWarning indicates that application source has resource with same Group, Kind, Name, Namespace multiple times
	InstallerConditionRepeatedResourceWarning = "RepeatedResourceWarning"
	// InstallerConditionExcludedResourceWarning indicates that application has resource which is configured to be excluded
	InstallerConditionExcludedResourceWarning = "ExcludedResourceWarning"
	// InstallerConditionOrphanedResourceWarning indicates that application has orphaned resources
	InstallerConditionOrphanedResourceWarning = "OrphanedResourceWarning"
)

// RevisionHistories is a array of history, oldest first and newest last
type RevisionHistories []RevisionHistory

func (in RevisionHistories) LastRevisionHistory() RevisionHistory {
	return in[len(in)-1]
}

func (in RevisionHistories) Trunc(n int) RevisionHistories {
	i := len(in) - n
	if i > 0 {
		in = in[i:]
	}
	return in
}

// RevisionHistory contains information relevant to an application deployment
type RevisionHistory struct {
	// Revision holds the revision of the sync
	Revision string `json:"revision" protobuf:"bytes,1,opt,name=revision"`
	// DeployedAt holds the time the deployment completed
	DeployedAt metav1.Time `json:"deployedAt" protobuf:"bytes,2,opt,name=deployedAt"`
	// ID is an auto incrementing identifier of the RevisionHistory
	ID     int64             `json:"id" protobuf:"bytes,3,opt,name=id"`
	Source ApplicationSource `json:"source,omitempty" protobuf:"bytes,4,opt,name=source"`
	// DeployStartedAt holds the time the deployment started
	DeployStartedAt *metav1.Time `json:"deployStartedAt,omitempty" protobuf:"bytes,5,opt,name=deployStartedAt"`
}

// ApplicationSource contains information about github repository, path within repository and target application environment.
type ApplicationSource struct {
	URL string `json:"url,omitempty" protobuf:"bytes,3,opt,name=url"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Installer is the Schema for the installers API
type Installer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstallerSpec   `json:"spec,omitempty"`
	Status InstallerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstallerList contains a list of Installer
type InstallerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Installer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Installer{}, &InstallerList{})
}

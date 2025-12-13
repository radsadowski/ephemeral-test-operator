/*
Copyright 2025.

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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EphemeralTestEnvironmentSpec defines the desired state
type EphemeralTestEnvironmentSpec struct {
	// GitLab MR information
	GitLabMR GitLabMRInfo `json:"gitlabMR"`

	// Application to deploy
	ApplicationTemplate ApplicationTemplate `json:"applicationTemplate"`

	// Test suite configuration
	TestSuite TestSuiteSpec `json:"testSuite"`

	// TTL for auto-cleanup
	TTL string `json:"ttl,omitempty"`
}

type GitLabMRInfo struct {
	ProjectID    int    `json:"projectId"`
	MRIID        int    `json:"mrIid"`
	SourceBranch string `json:"sourceBranch"`
	CommitSHA    string `json:"commitSha"`
}

type ApplicationTemplate struct {
	HelmChart string                `json:"helmChart"`
	Values    *apiextensionsv1.JSON `json:"values,omitempty"`
}

type TestSuiteSpec struct {
	Image            string `json:"image"`
	WorkflowTemplate string `json:"workflowTemplate"`
}

// EphemeralTestEnvironmentStatus defines the observed state
type EphemeralTestEnvironmentStatus struct {
	Phase             string      `json:"phase,omitempty"`
	TestNamespace     string      `json:"testNamespace,omitempty"`
	SetupWorkflowName string      `json:"setupWorkflowName,omitempty"`
	TestWorkflowName  string      `json:"testWorkflowName,omitempty"`
	WorkflowStartTime metav1.Time `json:"workflowStartTime,omitempty"`
	TestResults       string      `json:"testResults,omitempty"`
	Message           string      `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
//+kubebuilder:printcolumn:name="Namespace",type=string,JSONPath=`.status.testNamespace`
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// EphemeralTestEnvironment is the Schema for the ephemeraltestenvironments API
type EphemeralTestEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EphemeralTestEnvironmentSpec   `json:"spec,omitempty"`
	Status EphemeralTestEnvironmentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EphemeralTestEnvironmentList contains a list of EphemeralTestEnvironment
type EphemeralTestEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EphemeralTestEnvironment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EphemeralTestEnvironment{}, &EphemeralTestEnvironmentList{})
}

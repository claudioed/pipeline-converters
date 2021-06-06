// Copyright 2020 apirator.io
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	Provisioned          = "Provisioned"
	Error                = "Error"
	InvalidOas           = "InvalidOAS"
	WaitingAnnotations   = "WaitingAnnotations"
	IngressTag           = "ingress"
	NamespaceTag         = "namespace"
	IngressFinalizerName = "ingress.finalizers.apirator.io"
)


// APIMockSpec defines the desired state of APIMock
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".status.phase",name=Status,type=string
type APIMockSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Definition        string            `json:"definition,omitempty"`
	ServiceDefinition ServiceDefinition `json:"serviceDefinition,omitempty"`
	Watch             bool              `json:"watch,omitempty"`
	Selector          map[string]string `json:"selector,omitempty"`
	Host              string            `json:"host,omitempty"`
	Initialized       bool              `json:"initialized,omitempty"`
}


// Service Definition it will "link" the mock with created service
type ServiceDefinition struct {
	Port        int            `json:"port,omitempty"`
	ServiceType v1.ServiceType `json:"serviceType,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIMock is the Schema for the apimocks API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=apimocks,scope=Namespaced
// +kubebuilder:printcolumn:JSONPath=".status.phase",name=Status,type=string
// +kubebuilder:printcolumn:JSONPath=".metadata.annotations['apirator\\.io/cluster-ip']",name=CLUSTER-IP,type=string
// +kubebuilder:printcolumn:JSONPath=".metadata.annotations['apirator\\.io/ports']",name=PORT(S),type=string
type APIMock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   APIMockSpec   `json:"spec,omitempty"`
}



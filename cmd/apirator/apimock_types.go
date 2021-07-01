/*
Copyright 2021.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	APIMockValidated string = "Validated"
	APIMockAvailable        = "Available"
)

// APIMockSpec defines the desired state of APIMock
type APIMockSpec struct {
	//+kubebuilder:validation:Required
	Definition string   `json:"definition,omitempty"`
	Service    Service  `json:"service,omitempty"`
	Ingress    *Ingress `json:"ingress,omitempty"`
	Watch      bool     `json:"watch,omitempty"`
}

type Service struct {
	Port                  int                                     `json:"port,omitempty"`
	Type                  corev1.ServiceType                      `json:"type,omitempty"`
	ExternalTrafficPolicy corev1.ServiceExternalTrafficPolicyType `json:"externalTrafficPolicy,omitempty"`
	Annotations           map[string]string                       `json:"annotations,omitempty"`
}

type Ingress struct {
	Hostname    string            `json:"hostname"`
	Path        string            `json:"path"`
	PathType    *string           `json:"pathType,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	CertManager *bool             `json:"certManager,omitempty"`
	TLS         *TLS              `json:"tls,omitempty"`
}

type TLS struct {
	SecretName string `json:"secretName"`
}

type APIMock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIMockSpec `json:"spec,omitempty"`
}

/*
Copyright © 2021 NAME HERE apirator@apirator.io

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
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	v1alpha1 "pipeline-converters/cmd/apirator"
)

// openapi2mockCmd represents the openapi2mock command
var (
	openapi2mockCmd = &cobra.Command{
		Use:   "openapi2mock",
		Short: "It will configure a mock from the OpenAPI File",
		Long: `This commands will generate a ApiMock Custom Resource to deploy a mock in the kubernetes cluster`,
		Run: func(cmd *cobra.Command, args []string) {

			yc, err := ioutil.ReadFile(oasLocation)
			if err != nil {
				log.Fatalf("Error to read OpenAPI File  #%v ", err)
			}

			mc := v1alpha1.APIMock{
				TypeMeta:   metav1.TypeMeta{
					Kind:       "APIMock",
					APIVersion: "apirator.io/v1alpha1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:                       serviceName,
					Namespace:                  serviceNamespace,
				},
				Spec:       v1alpha1.APIMockSpec{
					Definition:        string(yc),
					ServiceDefinition: v1alpha1.ServiceDefinition{
						Port:        servicePort,
						ServiceType:  serviceTypeFromString(serviceType),
					},
					Watch:             false,
				},
			}

			apiMock, err := json.Marshal(mc)

			apiYaml,err := yaml.JSONToYAML(apiMock)

			err = ioutil.WriteFile(crLocation, apiYaml, 777)
			if err != nil {
				log.Fatalf("Error to write CR definition  #%v ", err)
			}
			fmt.Println("Custom Resource ApiMock was generated")
		},
	}
	oasLocation string
	crLocation string
	serviceType string
	servicePort int
	serviceName string
	serviceNamespace string
)

func serviceTypeFromString(st string) v1.ServiceType {
	if st == "ClusterIP" {
		return v1.ServiceTypeClusterIP
	}else if st == "NodePort" {
		return v1.ServiceTypeNodePort
	}else if st == "LoadBalancer"{
		return v1.ServiceTypeLoadBalancer
	}
	return v1.ServiceTypeClusterIP
}

func init() {
	rootCmd.AddCommand(openapi2mockCmd)
	// File location
	openapi2mockCmd.Flags().StringVar(&oasLocation,"oas_location", "/tmp/openapi.yaml", "OpenAPI File")
	openapi2mockCmd.Flags().StringVar(&crLocation,"cr_location", "/tmp/mock.yaml", "Mock File")
	// Service Definition
	openapi2mockCmd.Flags().StringVar(&serviceType,"service_type", "ClusterIP", "Service Port Type")
	openapi2mockCmd.Flags().IntVar(&servicePort,"service_port", 8080, "Service Port ")
	// k8s CR
	openapi2mockCmd.Flags().StringVar(&serviceName,"service_name", "oas_mock", "Service Name")
	openapi2mockCmd.Flags().StringVar(&serviceNamespace,"service_namespace", "mocks", "Mocks")
}


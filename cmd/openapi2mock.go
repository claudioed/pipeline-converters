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
		Long:  `This commands will generate a ApiMock Custom Resource to deploy a mock in the kubernetes cluster`,
		Run: func(cmd *cobra.Command, args []string) {

			log.Println("==================================================")
			log.Println("Start generation API Mock artifact")
			log.Println("==================================================")

			yc, err := ioutil.ReadFile(oasLocation)
			if err != nil {
				log.Fatalf("Error to read OpenAPI File  #%v ", err)
			}

			mc := v1alpha1.APIMock{
				TypeMeta: metav1.TypeMeta{
					Kind:       "APIMock",
					APIVersion: "apirator.io/v1alpha1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      serviceName,
					Namespace: serviceNamespace,
				},
				Spec: v1alpha1.APIMockSpec{
					Definition: string(yc),
					Service: v1alpha1.Service{
						Port: servicePort,
						Type: serviceTypeFromString(serviceType),
					},
					Watch: false,
				},
			}

			apiMock, err := json.Marshal(mc)

			log.Println(string(apiMock))

			apiYaml, err := yaml.JSONToYAML(apiMock)

			fn := fmt.Sprintf("%s/%s.yaml", crLocation, serviceName)

			err = ioutil.WriteFile(fn, apiYaml, 777)
			if err != nil {
				log.Fatalf("Error to write CR definition  #%v ", err)
			}

			log.Println("==================================================")
			log.Println("Destination: " + fn)
			log.Println("==================================================")

			log.Println("==================================================")
			log.Println("End generation API Mock artifact")
			log.Println("==================================================")

		},
	}
	oasLocation      string
	crLocation       string
	serviceType      string
	servicePort      int
	serviceName      string
	serviceNamespace string
)

func serviceTypeFromString(st string) v1.ServiceType {
	if st == "ClusterIP" {
		return v1.ServiceTypeClusterIP
	} else if st == "NodePort" {
		return v1.ServiceTypeNodePort
	} else if st == "LoadBalancer" {
		return v1.ServiceTypeLoadBalancer
	}
	return v1.ServiceTypeClusterIP
}

func init() {
	rootCmd.AddCommand(openapi2mockCmd)
	// File location
	openapi2mockCmd.Flags().StringVar(&oasLocation, "oas-location", "/tmp/openapi.yaml", "OpenAPI File")

	openapi2mockCmd.Flags().StringVar(&crLocation, "cr-location", "/tmp", "Mock File")
	// Service Definition
	openapi2mockCmd.Flags().StringVar(&serviceType, "service-type", "ClusterIP", "Service Port Type")
	openapi2mockCmd.Flags().IntVar(&servicePort, "service-port", 8080, "Service Port ")
	// k8s CR
	openapi2mockCmd.Flags().StringVar(&serviceName, "service-name", "oas_mock", "Service Name")
	openapi2mockCmd.Flags().StringVar(&serviceNamespace, "service-namespace", "mocks", "Mocks")
}

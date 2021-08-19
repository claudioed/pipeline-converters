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
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"pipeline-converters/cmd/backstage"
)

const (
	fileFolder = "./apis/%s.yaml"
)

// openapi2backstageCmd represents the openapi2backstage command
var (
	openapi2backstageCmd = &cobra.Command{
		Use:   "openapi2backstage",
		Short: "Convert the OpenAPI to Backstage API Entity",
		Long: `A longer description that spans multiple lines and likely contains examples
			and usage of using your command. For example:
			
			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {

			log.Println("==================================================")
			log.Println("Start generation API artifact")
			log.Println("==================================================")

			yc, err := ioutil.ReadFile(openApiLocation)
			if err != nil {
				log.Fatalf("Error to read OpenAPI File  #%v ", err)
			}

			api := backstage.API{
				ApiVersion: "backstage.io/v1alpha1",
				Kind:       "API",
				Metadata: backstage.Metadata{
					Name:        name,
					Description: "A mock for the service",
				},
				Spec: backstage.Spec{
					Type:       apiType,
					Lifecycle:  lifecycle,
					Owner:      owner,
					System:     system,
					Definition: string(yc),
				},
			}

			apiYaml, err := yaml.Marshal(api)

			fn := fmt.Sprintf("%s/%s.yaml", backstageArtifact, name)

			err = ioutil.WriteFile(fn, apiYaml, 777)
			if err != nil {
				log.Fatalf("Error to write Backstage artifact  #%v ", err)
			}

			log.Println(string(apiYaml))

			log.Println("==================================================")
			log.Println("End generation API artifact")
			log.Println("==================================================")

			log.Println("==================================================")
			log.Println("Start generation API Catalog artifact")
			log.Println("==================================================")

			cc, err := ioutil.ReadFile(apiContextLocation)
			if err != nil {
				log.Fatalf("Error to read Location File  #%v ", err)
			}
			location := &backstage.Location{}

			if err := yaml.Unmarshal(cc, location); err != nil {
				log.Fatalf("Error to parse APIs file  #%v ", err)
			}

			location.Spec.Targets = append(location.Spec.Targets, fmt.Sprintf(fileFolder, name))

			locationYaml, err := yaml.Marshal(location)

			err = ioutil.WriteFile(apiContextLocation, locationYaml, 777)
			if err != nil {
				log.Fatalf("Error to write Backstage artifact  #%v ", err)
			}

			log.Println(string(locationYaml))

			log.Println("==================================================")
			log.Println("End generation API Catalog artifact")
			log.Println("==================================================")

			log.Println("Process Completed")

		},
	}
	apiType            string
	lifecycle          string
	owner              string
	system             string
	name               string
	openApiLocation    string
	backstageArtifact  string
	apiContextLocation string
)

func init() {
	rootCmd.AddCommand(openapi2backstageCmd)
	// File location
	openapi2backstageCmd.Flags().StringVar(&name, "name", "change me. I need a name", "API Name")
	openapi2backstageCmd.Flags().StringVar(&apiType, "api-type", "external", "API documentation type")
	openapi2backstageCmd.Flags().StringVar(&lifecycle, "lifecycle", "development", "API current stage (development, retirement, production)")
	openapi2backstageCmd.Flags().StringVar(&owner, "owner", "change me. I need a owner", "API team ")
	openapi2backstageCmd.Flags().StringVar(&system, "system", "change me. I need a system", "The system name")
	openapi2backstageCmd.Flags().StringVar(&openApiLocation, "oas-location", "/tmp/openapi.yaml", "OpenAPI File Location")
	openapi2backstageCmd.Flags().StringVar(&backstageArtifact, "backstage-artifact-location", "/tmp", "Path to generate backstage artifact")
	openapi2backstageCmd.Flags().StringVar(&backstageArtifact, "api-context-location", "/tmp", "Api Context Location")
}

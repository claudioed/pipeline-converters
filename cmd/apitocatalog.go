/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

// apitocatalogCmd represents the apitocatalog command
var (
	apitocatalogCmd = &cobra.Command{
		Use:   "apitocatalog",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {

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

			location.Spec.Targets = append(location.Spec.Targets, fmt.Sprintf(fileFolder, apiName))

			locationYaml, err := yaml.Marshal(location)

			err = ioutil.WriteFile(apiContextLocation, locationYaml, 777)
			if err != nil {
				log.Fatalf("Error to write Backstage artifact  #%v ", err)
			}

			log.Println(string(locationYaml))

			log.Println("==================================================")
			log.Println("End generation API Catalog artifact")
			log.Println("==================================================")

		},
	}
	apiContextLocation string
	apiName            string
)

func init() {
	rootCmd.AddCommand(apitocatalogCmd)
	apitocatalogCmd.Flags().StringVar(&apiContextLocation, "api-context-location", "/tmp", "Api Context Location")
	apitocatalogCmd.Flags().StringVar(&apiName, "api-name", "dummy", "API Name")
}

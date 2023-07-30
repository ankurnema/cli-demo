/*
Copyright © 2023 Ankur Nema ankurnema@gmail.com

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
	"github.com/ankurnema/cli-demo/pkg/petstore"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Clinet which will be used by all sub commands for communicating with petstore.
var petStoreClient petstore.ClientWithResponses

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-demo",
	Short: "Demo command for interacting with pet store",
	Long: `cli-demo is command which can be used to interact with pet store rest server.
For more details check https://petstore.swagger.io/ 
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-demo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli-demo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cli-demo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err != nil {
			panic(err)
		}
	}

	// Lets get token and use it to initialize our client.
	token := viper.GetString("token")

	// Creating token based security provider.
	bearerToken, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		panic(err)
	}

	// Creating client which returns response back.
	client, err := petstore.NewClientWithResponses(
		"https://petstore.swagger.io/v2/",
		petstore.WithRequestEditorFn(bearerToken.Intercept),
	)

	// Since we get pointer reference back, lets get its value for later use.
	petStoreClient = *client

}

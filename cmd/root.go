// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"net/http"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var serverRoot string
var serverPort string
var urlPrefix string

// var enableLogging bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "stupidhttp",
	Short: "Quickly serve local filesystem files over http",
	Long: `Quickly serve local filesystem files over http

With no arguments, stupidhttp starts serving files under ./ over port 8080.
With '-d' arg, specify the directory to be served.
With '-p' arg, specify the port to serve on.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fs := http.FileServer(http.Dir(serverRoot))

		http.Handle(urlPrefix, http.StripPrefix(urlPrefix, fs))

		log.Infof("Serving %s over 0.0.0.0:%s... Stop with ^C", serverRoot, serverPort)
		http.ListenAndServe(":"+serverPort, nil)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.frame.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	RootCmd.Flags().StringVarP(&serverRoot, "dir", "d", "./", "root directory to be served (ex: /var/www) [default is ./]")
	// RootCmd.Flags().BoolVarP(&enableLogging, "log", "l", true, "prints usage logs to the standard output")
	RootCmd.Flags().StringVar(&serverPort, "port", "8080", "port to listen to [default is 8080)")
	RootCmd.Flags().StringVar(&urlPrefix, "prefix", "/", "prefix required (ex: /static), suffix to host:port [default is /]")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".frame" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".frame")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var dirToServe string
var serverPort string
var urlPrefix string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "simplehttp",
	Short: "Quickly serve a local filesystem directory over http",
	Long: `Quickly serve a local filesystem directory over http

With no arguments, simplehttp starts serving files under ./ over port 8080.
With '-d' arg, specify the directory to be served.
With '-p' arg, specify the port to serve on.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fs := http.FileServer(http.Dir(dirToServe))

		http.Handle(urlPrefix, http.StripPrefix(urlPrefix, fs))

		log.Infof("Serving %s over 0.0.0.0:%s... Stop with ^C", dirToServe, serverPort)
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
	RootCmd.Flags().StringVarP(&dirToServe, "dir", "d", "./", "root directory to be served (ex: /var/www) [default is ./]")

	RootCmd.Flags().StringVar(&serverPort, "port", "8080", "port to listen to [default is 8080)")
	RootCmd.Flags().StringVar(&urlPrefix, "prefix", "/", "prefix required (ex: /static), suffix to host:port [default is /]")
}

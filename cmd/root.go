/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"avatar_dl/consumer"
	"avatar_dl/producer"
	"avatar_dl/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var logFilename string
var inputFilename string
var failureFilename string
var downloadFolderName string
var workers int
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "avatar_dl",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		defer util.Elapsed()()

		file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(file)

		producer.New(inputFilename, "failure_uris.txt").Do(consumer.New(downloadFolderName), workers)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&logFilename, "logFilename", "l", "logrus.log", "")
	rootCmd.PersistentFlags().StringVarP(&inputFilename, "inputFilename", "i", "sample100.csv", "")
	rootCmd.PersistentFlags().StringVarP(&failureFilename, "failureFilename", "f", "failure_uris.txt", "")
	rootCmd.PersistentFlags().StringVarP(&downloadFolderName, "downloadFolderName", "d", "./download", "")
	rootCmd.PersistentFlags().IntVarP(&workers, "workers", "w", 10, "")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "")
}

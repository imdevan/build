// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
    "github.com/fatih/color"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "build",
	Short: "A build tool for people",
	Long: `A build tool for people`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		c := color.New(color.FgCyan).Add(color.Underline)
		c.Println("Hai! Welcome to the design tool!")
		fmt.Println("\nSince I'm in beta atm, I've made sume assumptions about your org:")

		bl := color.New(color.FgBlue).SprintFunc()
		rd := color.New(color.FgRed).SprintFunc()
		fmt.Println(bl("1. ") + "You're in whitewater")
		fmt.Println(bl("2. ") + "You care about all whitewater core ux repos")
		fmt.Println(bl("3. ") + "You have, or wouldn't mind having a forked repo of Rapid & Whitewater")

		fmt.Println("\nCommands to manage our code:")
		fmt.Println(rd("1. ") + "update-code")
		fmt.Println(rd("2. ") + "pull-request")
		fmt.Println(rd("3. ") + "push")
		fmt.Println(rd("4. ") + "clean" + bl(" *"))

		fmt.Println("\nCommands to use our server:")
		fmt.Println(rd("1. ") + "run")
		fmt.Println(rd("2. ") + "test" + bl(" *"))

		color.Blue("\n\n* Commands with a start take a while");
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.build.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".build") // name of config file (without extension)
	viper.AddConfigPath("$HOME")        // adding home directory as first search path
	viper.AutomaticEnv()                // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

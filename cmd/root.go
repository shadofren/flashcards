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
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var defaultCfgFile = "config.yaml"
var DBPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flashcards",
	Short: "Flashcards cli to add study materials and revise",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.flashcards.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.Println("initConfig")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		configPath := path.Join(home, ".config", "flashcards")
		DBPath = path.Join(configPath, "dbs")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			os.MkdirAll(configPath, 0700)
			os.Mkdir(DBPath, 0700)
		}
		// Search config in home directory with name ".flashcards" (without extension).
		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".flashcards")
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		conf, err := os.Open(defaultCfgFile)
		if err != nil {
			log.Fatalf("default config %s not found", defaultCfgFile)
		}
		viper.ReadConfig(conf)
		viper.SafeWriteConfig()
	}
}

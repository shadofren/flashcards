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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the specified plan",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := args[0]
		current := viper.GetString("Current")
		if plan == "default" {
			log.Println("Unable to delete the default plan")
			return
		} else if plan == current {
			viper.Set("Current", "default")
		}
		alls := viper.GetStringSlice("Alls")
		for i, p := range alls {
			if p == plan {
				alls[i] = alls[len(alls)-1]
				alls = alls[:len(alls)-1]
				viper.Set("Alls", alls)
				break
			}
		}
		if err := viper.WriteConfig(); err != nil {
			log.Println("Unable to update config file", err)
		}
		current = viper.GetString("Current")
		dbFile := path.Join(DBPath, current+".sqlite")
		os.Remove(dbFile)
		log.Println("Now using", viper.GetString("Current"))
	},
}

func init() {
	planCmd.AddCommand(deleteCmd)

	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

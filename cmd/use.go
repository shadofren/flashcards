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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Change to another plan",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		plan := args[0]
		allsMap := make(map[string]interface{})
		alls := viper.GetStringSlice("Alls")
		for _, p := range alls {
			allsMap[p] = nil
		}
		if _, ok := allsMap[plan]; ok {
			viper.Set("Current", plan)
		} else {
			fmt.Println("Plan doesn't exist, creating", plan)
			alls = append(alls, plan)
			viper.Set("Current", plan)
			viper.Set("Alls", alls)
			if err := viper.WriteConfig(); err != nil {
				fmt.Println("Unable to update config file", err)
			}
		}
		fmt.Println("Now using", plan)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// useCmd.PersistentFlags().String("foo", "", "A help for foo")
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

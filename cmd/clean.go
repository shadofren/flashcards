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
	"strings"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean a plan or clean all plans",
	Long: `Clean all plans: flashcards clean all
Clean a plan: flashcards clean <plan>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		plan := args[0]
		var cleanAll bool
		if strings.EqualFold(plan, "all") {
			cleanAll = true
		}
		if cleanAll {

		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/shadofren/flashcards/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// studyCmd represents the study command
var studyCmd = &cobra.Command{
	Use:   "study",
	Short: "Study mode, list the items one by one, sorted by the least familarity",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := viper.GetString("Current")
		dbFile := path.Join(DBPath, plan+".sqlite")
		database := db.Connect(dbFile)
		buf := bufio.NewReader(os.Stdin)
		for _, row := range db.ListTopRows(database, 1000) {
			fmt.Printf("Question: %s\nAnswer: %s", row.Question, row.Answer)
			buf.ReadBytes('\n')
			fmt.Println()
		}
		fmt.Println("DONE")
	},
}

func init() {
	rootCmd.AddCommand(studyCmd)

	// Here you will define your flags and configuration settings.

	// studyCmd.PersistentFlags().String("foo", "", "A help for foo")
	// studyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

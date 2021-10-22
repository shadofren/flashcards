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
	"path"

	"github.com/shadofren/flashcards/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var question string
var answer string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new question & answer into current plan",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := viper.GetString("Current")
		dbFile := path.Join(DBPath, plan+".sqlite")
		database := db.Connect(dbFile)
		row := db.Get(database, question)
		if row == nil {
			db.Insert(database, question, answer)
		} else {
			row.Answer = answer
			db.Update(database, row)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	addCmd.Flags().StringVarP(&question, "question", "q", "", "What is the question?")
	addCmd.Flags().StringVarP(&answer, "answer", "a", "", "What is the answer?")
	rootCmd.MarkFlagRequired("question")
	rootCmd.MarkFlagRequired("answer")
}

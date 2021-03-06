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
	"log"
	"os"
	"sort"
	"strconv"

	"TaskManager/taskHandler"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	doneOpt bool
	allOpt  bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your all Task",
	Long:  `List your all task`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := taskHandler.ReadFile(viper.GetString("datafile"))

		if err != nil {
			log.Printf("%v", err)
		}

		sort.Sort(taskHandler.ByPri(tasks))
		// for i, task := range tasks {
		// 	task.Position = i + 1

		// }
		for i, task := range tasks {
			if allOpt || task.Done == doneOpt {
				fmt.Fprintln(os.Stdout, strconv.Itoa(i+1)+". "+task.PrettyDone()+task.Prettier()+" "+task.Txt)

			}

		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "show done tasks")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "show all tasks")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

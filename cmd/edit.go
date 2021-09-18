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
	"TaskManager/taskHandler"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit the task",
	Long:  `edit the task by providing position and new task`,
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := taskHandler.ReadFile(viper.GetString("datafile"))

		if err != nil {
			log.Fatalln(err)
		}

		i, err := strconv.Atoi(args[0])
		newTask := args[1]

		if err != nil {
			log.Fatalln(i, "invalid Label")
		}

		if i > 0 && i < len(tasks) {
			sort.Sort(taskHandler.ByPri(tasks))
			tasks[i-1].Txt = newTask
			fmt.Fprintln(os.Stdout, "successfully edited")

			taskHandler.SaveTasks(viper.GetString("datafile"), tasks)
		} else {
			log.Println(i, "doesn't match any tasks")
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

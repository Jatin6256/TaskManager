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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"TaskManager/taskHandler"
)

// addCmd represents the add command
var priority int
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `add will create new task in the list`,
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := taskHandler.ReadFile(viper.GetString("datafile"))

		if err != nil {
			log.Printf("%v", err)
		}
		for _, y := range args {
			temp := taskHandler.Task{Txt: y}
			temp.SetPriority(priority)
			tasks = append(tasks, temp)
		}

		err = taskHandler.SaveTasks(viper.GetString("datafile"), tasks)

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "set priorty 1,2,3. default value is 2")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var cookTimer int

// cookCmd represents the cook command
var cookCmd = &cobra.Command{
	Use:   "cook",
	Short: "Cook CPU",
	Long:  `This command maxes out CPU for a specified time.  Default is 300 seconds`,
	Run: func(cmd *cobra.Command, args []string) {
		runCPU(time.Duration(cookTimer))
	},
}

func runCPU(timer time.Duration) {
	fmt.Printf("Cooking CPU for %d seconds", timer)
	done := make(chan int)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}
	time.Sleep(time.Second * timer)
	close(done)
}
func init() {
	rootCmd.AddCommand(cookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//cookCmd.PersistentFlags().String("foo", "", "A help for foo")
	cookCmd.Flags().IntVarP(&cookTimer, "seconds", "s", 300, "number of seconds between log entries")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

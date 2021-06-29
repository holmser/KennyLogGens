/*
Copyright © 2021 Chris Holmes

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

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

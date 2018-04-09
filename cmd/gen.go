// Copyright © 2018 Chris Holmes chris@holmser.net
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
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate logs to syslog",
	Long:  `This will alternately spew the lyrics to Danny's Song and Danger Zone to syslog. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func logGen() {
	backend2, err := logging.NewSyslogBackend("klog")
	if err != nil {
		log.Fatal(err)
	}
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	logging.SetBackend(backend2Formatter)

	inFile, _ := os.Open("dangerzone.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "danger") {
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
			time.Sleep(time.Second * 10)
			close(done)
			log.Error(text)
		} else {
			log.Info(text)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

var log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} ▶ %{level} %{message}`,
)

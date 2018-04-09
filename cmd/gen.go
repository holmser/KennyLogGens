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
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2, err := logging.NewSyslogBackend("klog")
	if err != nil {
		log.Fatal(err)
	}
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	logging.SetBackend(backend2Formatter, backend1Formatter)

	inFile, _ := os.Open("dangerzone.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(strings.NewReader(dangerZone))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "danger") {
			// done := make(chan int)
			// for i := 0; i < runtime.NumCPU(); i++ {
			// 	go func() {
			// 		for {
			// 			select {
			// 			case <-done:
			// 				return
			// 			default:
			// 			}
			// 		}
			// 	}()
			// }
			// time.Sleep(time.Second * 10)
			// close(done)
			log.Error(text)
		} else {
			log.Info(text)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

var log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} ▶ %{level} %{message}`,
)

var dangerZone = `	Revvin' up your engine
	Listen to her howlin' roar
	Metal under tension
	Beggin' you to touch and go
	Highway to the danger zone
	Ride into the danger zone
	Headin' into twilight
	Spreadin' out her wings tonight
	She got you jumpin' off the track
	And shovin' into overdrive
	Highway to the danger zone
	I'll take you
	Ridin' into the danger zone
	You'll never say hello to you
	Until you get it on the red line overload
	You'll never know what you can do
	Until you get it up as high as you can go
	Out along the edges
	Always where I burn to be
	The further on the edge
	The hotter the intensity
	Highway to the danger zone
	Gonna take you
	Right into the danger zone
	Highway to…`

var dannysSong = `	People smile and tell me I'm the lucky one
	And we've just begun
	Think I'm gonna have a son
	He will be like she and me, as free as a dove
	Conceived in love
	Sun is gonna shine above
	And even though we ain't got money
	I'm so in love with you honey
	And everything will bring a chain of love
	And in the mornin' when I rise
	Bring a tear of joy to my eyes
	And tell me everything is gonna be alright
	Seems as though a month ago I was Beta Chi
	Never got high
	Oh, was a sorry guy
	And now I smile and face the girl that shares my name, yeah
	Now I'm through with the game
	This boy'll never be the same
	And even though we ain't got money
	I'm so in love with you honey
	And everything will bring a chain of love
	And in the mornin' when I rise
	Bring a tear of joy to my eyes
	And tell me everything is gonna be all right
	Pisces Virgo rising is a very good sign
	Strong and kind
	And the little boy is mine
	Now I see a family where there once was none
	Now we've just begun
	Yeah we're gonna fly to the sun
	And even though we ain't got money
	I'm so in love with you honey
	And everything will bring a chain of love
	And in the mornin' when I rise
	You bring a tear of joy to my eyes
	And tell me everything is gonna be all right
	Love the girl who holds the world in a paper cup
	Drink it up
	Love her and she'll bring you luck
	And if you find she helps your mind
	Buddy, take her home
	Yeah, don't you live alone
	Try to earn what lovers own
	And even though we ain't got money
	I'm so in love with you honey
	And everything will bring a chain of love
	And in the mornin' when I rise
	You bring a tear of joy to my eyes
	And tell me everything is gonna be alright`

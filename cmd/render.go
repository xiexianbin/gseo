/*
Copyright © 2022 xiexianbin

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
	"encoding/json"
	"fmt"
	"github.com/xiexianbin/gseo/utils"
	"google.golang.org/api/searchconsole/v1"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var clicks float64
var ctr float64
var impressions float64
var position float64
var max int64

type T struct {
	Impressions int `json:"impressions"`
	Position    int `json:""`
}

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "render hugo post markdown files",
	Long: `render hugo post markdown files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr < 0 || ctr > 1 {
			fmt.Println("0 <= ctr <= 1.")
			os.Exit(1)
		}
		if max <= 0 {
			fmt.Println("max must >= 1")
			os.Exit(1)
		}

		// Open File
		fileName := utils.GetCacheFile()
		file, err := os.Open(fileName)
		if err != nil {
			_ = fmt.Errorf("%v", err)
		}
		defer file.Close()

		byteValue, _ := ioutil.ReadAll(file)
		var oldResult []*searchconsole.ApiDataRow
		_= json.Unmarshal(byteValue, &oldResult)

		newResult := utils.ParserSearchAnalyticsQuery(oldResult)
		//r, _ := json.Marshal(newResult)
		//fmt.Println(r)
		for url, item := range newResult {
			fmt.Println(fmt.Sprintf("%s", url))
			var count int64
			count = 0
			for k, v := range item {
				if v.Clicks >= clicks && v.Ctr >= ctr && v.Impressions >= impressions && v.Position >= position{
					fmt.Println(fmt.Sprintf("  - %s", k))
					count ++
				}
				if count >= max {
					break
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	renderCmd.Flags().Float64VarP(&clicks, "clicks", "k", 0, ">=clicks to render seo.")
	renderCmd.Flags().Float64VarP(&ctr, "ctr", "c", 0.3, "ctr = clicks / impressions to render seo, and 0 <= ctr <= 1.")
	renderCmd.Flags().Float64VarP(&impressions, "impressions", "i", 100, ">=impressions to render seo.")
	renderCmd.Flags().Float64VarP(&position, "position", "p", 10, ">=position to render seo.")
	renderCmd.Flags().Int64VarP(&max, "max", "m", 8, "max seo items.")
}

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
	"github.com/xiexianbin/gseo/googleapi"
	"github.com/xiexianbin/gseo/utils"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/api/searchconsole/v1"
)

var site string
var last int

// keywordCmd represents the keyword command
var keywordCmd = &cobra.Command{
	Use:   "keyword",
	Short: "show site keywords",
	Long:  "download hugo post keywords from Google Search Console API, and cache it in `./.gseo/` dir",
	Run: func(cmd *cobra.Command, args []string) {
		if site == "" {
			fmt.Println("site is unknown, use `gseo sites` get sites!")
			os.Exit(1)
		}
		if last <= 0 {
			fmt.Println("--last -l must >= 0!")
			os.Exit(1)
		}

		sc := googleapi.NewSearchConsoleAPI()
		if sc.SearchConsoleService == nil {
			fmt.Println("init search console API err.")
			os.Exit(1)
		}
		searchanalyticsqueryrequest := searchconsole.SearchAnalyticsQueryRequest{
			StartDate: utils.LastDate(last),
			EndDate: utils.TodayDate(),
			Dimensions: []string{"PAGE", "QUERY"},
		}
		rows := sc.Query(site, &searchanalyticsqueryrequest)
		if len(rows) > 0 {
			fmt.Println("Result:")
			for _, row := range rows {
				r, _ := json.Marshal(row)
				fmt.Println(string(r))
			}

			_rows, _ := json.Marshal(rows)
			// open File
			fileName := utils.GetCacheFile()
			file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
			if err != nil{
				fmt.Println("Open file err =", err)
				return
			}
			defer file.Close()

			// write to file
			n, err := file.Write(_rows)
			if err != nil{
				fmt.Println("Write file err =", err)
				return
			}
			fmt.Println("Write file", fileName, "success, n =", n)
		} else {
			fmt.Println("No Result from Google Console API.")
		}
	},
}

func init() {
	rootCmd.AddCommand(keywordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keywordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keywordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	keywordCmd.Flags().StringVarP(&site, "site", "s", "", "site url")
	keywordCmd.Flags().IntVarP(&last, "last", "l", 90, "last days")
}

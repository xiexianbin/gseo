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
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xiexianbin/golib/logger"
	"google.golang.org/api/searchconsole/v1"

	"github.com/xiexianbin/gseo/utils"
)

var contentPath string
var clicks float64
var ctr float64
var impressions float64
var position float64
var max int
var dryrun bool

type T struct {
	Impressions int `json:"impressions"`
	Position    int `json:""`
}

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "render hugo post markdown files",
	Long: `render hugo post markdown files.
default args is:
  gseo render --content PATH_OF_HUGO_CONTENT --position 10 --ctr 0 --impressions 100 --clicks 0.3 --max 8 --dryrun
`,
	Run: func(cmd *cobra.Command, args []string) {
		if contentPath == "" {
			logger.Print("content path not special.")
			os.Exit(1)
		}
		if utils.IsDir(contentPath) == false {
			logger.Print("content path not exists.")
			os.Exit(1)
		}
		if strings.HasSuffix(contentPath, "/") {
			contentPath = strings.TrimRight(contentPath, "/")
		}
		if ctr < 0 || ctr > 1 {
			logger.Print("0 <= ctr <= 1.")
			os.Exit(1)
		}
		if max < -1 || max == 0 {
			logger.Print("max must >= 1 or must = -1")
			os.Exit(1)
		}

		// Open File
		fileName := utils.GetCacheFile()
		file, err := os.Open(fileName)
		if err != nil {
			logger.Errorf("read file %s err: %v", fileName, err.Error())
			os.Exit(1)
		}
		defer file.Close()

		byteValue, _ := ioutil.ReadAll(file)
		var searchAnalyticsQueryRows []*searchconsole.ApiDataRow
		_ = json.Unmarshal(byteValue, &searchAnalyticsQueryRows)

		newResult := utils.ParserSearchAnalyticsQuery(searchAnalyticsQueryRows)
		for url, item := range newResult {
			targetKeywords := map[string]*utils.KeywordItem{}
			for k, v := range item {
				if v.Clicks >= clicks && v.Ctr >= ctr && v.Impressions >= impressions && v.Position >= position {
					targetKeywords[k] = v
				}
			}

			keywords := utils.SortKeywords(targetKeywords)

			if len(keywords) == 0 {
				continue
			} else if len(keywords) > max {
				keywords = keywords[:max]
			}

			logger.Print(url)
			for _, k := range keywords {
				logger.Printf("  - %s", k)
			}

			if dryrun == false {
				markdownFilePath, err := utils.GetMarkdownFileByURL(url, contentPath)
				if err != nil {
					logger.Printf("GetMarkdownFileByURL: %s, err: %s, skip.", url, err.Error())
					continue
				}

				err = utils.UpdateKeywords(markdownFilePath, keywords)
				if err != nil {
					logger.Printf(
						"UpdateKeywords: %s, keywords: %s, error: %s",
						markdownFilePath, strings.Join(keywords, ","), err.Error())
					break
				}

				logger.Printf("update markdownFilePath %s, keywords %v done.\n", markdownFilePath, keywords)
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
	renderCmd.Flags().StringVarP(&contentPath, "content", "", "", "hugo content path")
	renderCmd.Flags().Float64VarP(&clicks, "clicks", "k", 0, ">=clicks to render seo.")
	renderCmd.Flags().Float64VarP(&ctr, "ctr", "c", 0.3, "ctr = clicks / impressions to render seo, and 0 <= ctr <= 1.")
	renderCmd.Flags().Float64VarP(&impressions, "impressions", "i", 100, ">=impressions to render seo.")
	renderCmd.Flags().Float64VarP(&position, "position", "p", 10, ">=position to render seo.")
	renderCmd.Flags().IntVarP(&max, "max", "m", 8, "max seo items, -1 is un-limit.")
	renderCmd.Flags().BoolVarP(&dryrun, "dryrun", "r", false, "dry run mode.")
}

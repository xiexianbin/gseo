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
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	"google.golang.org/api/searchconsole/v1"

	"github.com/xiexianbin/gseo/googleapi"
	"github.com/xiexianbin/gseo/utils/logger"
)

// sitesCmd represents the sites command
var sitesCmd = &cobra.Command{
	Use:   "sites",
	Short: "site list",
	Long:  `site list.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, client := googleapi.Client()
		searchConsoleService, err := searchconsole.NewService(
			ctx,
			option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Unable to retrieve Search Console client: %v", err)
		}

		siteList := searchConsoleService.Sites.List()
		sitesListResponse, err := siteList.Do()
		logger.Debug("sitesListResponse is: %v", sitesListResponse)
		if err != nil {
			fmt.Println("Call Google Search Console API error:", err)
			return
		}

		if len(sitesListResponse.SiteEntry) > 0 {
			fmt.Println("PermissionLevel  SiteUrl")
			for _, site := range sitesListResponse.SiteEntry {
				fmt.Printf("%s        %s\n", site.PermissionLevel, site.SiteUrl)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sitesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sitesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sitesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

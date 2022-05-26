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

package googleapi

import (
	"context"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/searchconsole/v1"

	"github.com/xiexianbin/gseo/utils/logger"
)

type SearchConsoleAPI struct {
	Ctx                  context.Context
	Client               *http.Client
	SearchConsoleService *searchconsole.Service
}

func NewSearchConsoleAPI() SearchConsoleAPI {
	ctx, client := Client()
	searchConsoleService, err := searchconsole.NewService(
		ctx,
		option.WithHTTPClient(client))
	if err != nil {
		logger.Debugf("Unable to retrieve Search Console client: %v", err)
		return SearchConsoleAPI{}
	}

	return SearchConsoleAPI{
		Ctx:                  ctx,
		Client:               client,
		SearchConsoleService: searchConsoleService,
	}
}

func (sc *SearchConsoleAPI) Query(siteUrl string, searchanalyticsqueryrequest *searchconsole.SearchAnalyticsQueryRequest) []*searchconsole.ApiDataRow {
	query := sc.SearchConsoleService.Searchanalytics.Query(siteUrl, searchanalyticsqueryrequest)
	queryResponse, err := query.Do()
	logger.Debugf("queryResponse is: %v", queryResponse)
	if err != nil {
		logger.Debugf("Call Google Search Console API error: %v", err)
		return nil
	}

	return queryResponse.Rows
}

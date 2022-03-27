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

package utils

import (
    "strings"

    "google.golang.org/api/searchconsole/v1"
)

type KeywordItem struct {
    Clicks      float64 `json:"clicks"`
    Ctr         float64 `json:"ctr"`
    Impressions float64 `json:"impressions"`
    Position    float64 `json:"position"`
}

// ParserSearchAnalyticsQuery parser "PAGE" and "QUERY" Dimensions Query Response
func ParserSearchAnalyticsQuery(rows []*searchconsole.ApiDataRow) map[string]map[string]*KeywordItem {
    result := map[string]map[string]*KeywordItem{}
    for _, row := range rows {
        url := row.Keys[0]
        if strings.Contains(url, "#") {
            url = strings.Split(url, "#")[0]
        }
        if strings.Contains(url, "index.html") {
            url = strings.Split(url, "index.html")[0]
        }
        keyword := row.Keys[1]
        if keywordItems, ok := result[url]; ok {
            newKeywordItem := KeywordItem{
                Clicks: row.Clicks,
                Ctr: row.Ctr,
                Impressions: row.Impressions,
                Position: row.Position,
            }
            if keywordItem, ok := keywordItems[keyword]; ok {
                clicks := keywordItem.Clicks + row.Clicks
                impressions := keywordItem.Impressions + row.Impressions
                ctr := clicks / impressions
                position := (keywordItem.Position + row.Position) / 2
                newKeywordItem = KeywordItem{
                    Clicks: clicks,
                    Ctr: ctr,
                    Impressions: impressions,
                    Position: position,
                }
            }
            result[url][keyword] = &newKeywordItem
        } else {
            result[url] = map[string]*KeywordItem{}
            newKeywordItem := KeywordItem{
                Clicks: row.Clicks,
                Ctr: row.Ctr,
                Impressions: row.Impressions,
                Position: row.Position,
            }
            result[url][keyword] = &newKeywordItem
        }
    }

    return result
}

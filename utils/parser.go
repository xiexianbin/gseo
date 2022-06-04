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
	"sort"
	"strings"

	"google.golang.org/api/searchconsole/v1"
)

type KeywordItem struct {
	Keyword     string  `json:"-"`
	Ctr         float64 `json:"ctr"`
	Clicks      float64 `json:"clicks"`
	Position    float64 `json:"position"`
	Impressions float64 `json:"impressions"`
}

type KeywordItems []*KeywordItem

func (k KeywordItems) Len() int { return len(k) }

func (k KeywordItems) Swap(i, j int) { k[i], k[j] = k[j], k[i] }

// ByValue implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded KeywordItems value.
type ByValue struct{ KeywordItems }

// Less sort keywords KeywordItems[i]/KeywordItems[j] by Ctr > Clicks > Position > Impressions, order by value desc
func (v ByValue) Less(i, j int) bool {
	if v.KeywordItems[i].Ctr != 0 || v.KeywordItems[j].Ctr != 0 {
		return v.KeywordItems[i].Ctr < v.KeywordItems[j].Ctr
	}
	if v.KeywordItems[i].Clicks != 0 || v.KeywordItems[j].Clicks != 0 {
		return v.KeywordItems[i].Clicks < v.KeywordItems[j].Clicks
	}
	if v.KeywordItems[i].Position != 0 || v.KeywordItems[j].Position != 0 {
		return v.KeywordItems[i].Position < v.KeywordItems[j].Position
	}
	if v.KeywordItems[i].Impressions != 0 || v.KeywordItems[j].Impressions != 0 {
		return v.KeywordItems[i].Impressions < v.KeywordItems[j].Impressions
	}
	return true
}

// SortKeywords return sort keywords list by Ctr > Clicks > Position > Impressions, order by value desc
func SortKeywords(input map[string]*KeywordItem) []string {
	var keywordItemList []*KeywordItem
	for k, v := range input {
		v.Keyword = k
		keywordItemList = append(keywordItemList, v)
	}

	sort.Sort(ByValue{keywordItemList})

	var result []string
	for _, v := range keywordItemList {
		result = append(result, v.Keyword)
	}

	return result
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
				Clicks:      row.Clicks,
				Ctr:         row.Ctr,
				Impressions: row.Impressions,
				Position:    row.Position,
			}
			if keywordItem, ok := keywordItems[keyword]; ok {
				clicks := keywordItem.Clicks + row.Clicks
				impressions := keywordItem.Impressions + row.Impressions
				ctr := clicks / impressions
				position := (keywordItem.Position + row.Position) / 2
				newKeywordItem = KeywordItem{
					Clicks:      clicks,
					Ctr:         ctr,
					Impressions: impressions,
					Position:    position,
				}
			}
			result[url][keyword] = &newKeywordItem
		} else {
			result[url] = map[string]*KeywordItem{}
			newKeywordItem := KeywordItem{
				Clicks:      row.Clicks,
				Ctr:         row.Ctr,
				Impressions: row.Impressions,
				Position:    row.Position,
			}
			result[url][keyword] = &newKeywordItem
		}
	}

	return result
}

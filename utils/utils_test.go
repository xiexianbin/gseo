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
	"testing"
)

func TestTodayDate(t *testing.T) {
	t.Logf(TodayDate())
}

func TestLastDate(t *testing.T) {
	t.Logf(LastDate(90))
}

func TestSortMap(t *testing.T) {
	m := make(map[string]float64)
	m["abc 中文"] = 12.2
	m["我是 seo"] = 1
	m["tags"] = 3

	r := SortMap(m)
	t.Logf("%v\n", r)
}

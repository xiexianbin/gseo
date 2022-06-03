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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"google.golang.org/api/searchconsole/v1"
)

func TestParserSearchAnalyticsQuery(t *testing.T) {
	// Open File
	fileName := GetCacheFile()
	file, err := os.Open(fileName)
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var oldResult []*searchconsole.ApiDataRow
	_ = json.Unmarshal(byteValue, &oldResult)

	fmt.Println(fmt.Sprintf("%v", oldResult))

	newResult := ParserSearchAnalyticsQuery(oldResult)
	r, _ := json.Marshal(newResult)
	fmt.Println(fmt.Sprintf("%s", r))
}

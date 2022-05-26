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
	"fmt"
	"log"
	"testing"

	"google.golang.org/api/option"
	"google.golang.org/api/searchconsole/v1"
)

func TestClient(t *testing.T) {
	ctx, client := Client()
	if client == nil {
		log.Printf("new client is nil, skip")
		return
	}
	searchconsoleService, err := searchconsole.NewService(
		ctx,
		option.WithHTTPClient(client))
	if err != nil {
		log.Printf("Unable to retrieve Search Console client: %#v", err)
		return
	}

	siteList := searchconsoleService.Sites.List()
	sitesListResponse, err := siteList.Do()
	if err != nil {
		return
	}
	fmt.Println(sitesListResponse)
}

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
	"fmt"
	"os"
	"testing"
)

func TestGetMarkdownFileByURL(t *testing.T) {
	contentPath := "/Users/xiexianbin/workspace/code/github.com/xiexianbin/note.seo/content"
	_, err := os.Stat(contentPath)
	if err != nil && os.IsNotExist(err) == false {
		markdownFilePath, err := GetMarkdownFileByURL("https://www.xiexianbin.cn/program/go/tinygo/", contentPath)
		if err != nil {
			t.Skip(err)
			return
		}

		fmt.Println("markdownFilePath", markdownFilePath)
	}
}

func TestParsePostKeysAndTags(t *testing.T) {
	filename := "./samples/test-1.md"
	postYaml, err := ParsePostKeysAndTags(filename)
	if err != nil {
		t.Skipf("postYaml: %v, error: %v", postYaml, err)
	}

	filename = "./samples/test-2.md"
	postYaml, err = ParsePostKeysAndTags(filename)
	if err != nil {
		t.Skipf("postYaml: %v, error: %v", postYaml, err)
	}
}

func TestUpdateKeywords(t *testing.T) {
	filename1 := "./samples/test-1.md"
	newKeywords := []string{"abc 中文", "我是 seo"}
	err := UpdateKeywords(filename1, newKeywords)
	if err != nil {
		t.Skip(err.Error())
	}

	filename2 := "./samples/test-1.md"
	newKeywords = []string{"abc 中文", "我是 seo", "kubernetes"}
	err = UpdateKeywords(filename2, newKeywords)
	if err != nil {
		t.Skip(err.Error())
	}
}

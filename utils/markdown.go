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
	"io/ioutil"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/xiexianbin/golib/logger"
	"gopkg.in/yaml.v2"
)

type PostYaml struct {
	Keywords []string `json:"keywords"`
	Tags     []string `json:"tags"`
}

func GetMarkdownFileByURL(permalink, contentPath string) (path string, err error) {
	u, err := url.Parse(permalink)
	if err != nil {
		panic(err)
	}

	relURL := u.Path
	if strings.HasPrefix(relURL, "/categories/") || strings.HasPrefix(relURL, "/tags/") {
		return "", fmt.Errorf("could not find markdownd file for %s in %s", permalink, contentPath)
	}

	var markdownFilePath string
	tmpDir := filepath.Join(contentPath, relURL)
	if IsDir(tmpDir) {
		markdownFilePath = fmt.Sprintf("%s/_index.md", tmpDir)
	} else {
		tmpDir = strings.TrimRight(tmpDir, "/")
		markdownFilePath = fmt.Sprintf("%s.md", tmpDir)
	}

	if IsFile(markdownFilePath) {
		return markdownFilePath, nil
	}

	return "", fmt.Errorf("could not find markdownd file for %s in %s", permalink, contentPath)
}

func ParsePostKeysAndTags(filename string) (*PostYaml, error) {
	var postYaml PostYaml
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(buf, &postYaml)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return &postYaml, nil
}

func UpdateKeywords(filename string, keywords []string) error {
	postYaml, err := ParsePostKeysAndTags(filename)
	if err != nil {
		return err
	}

	if len(keywords) == 0 {
		return nil
	}

	var oldStr string
	if len(postYaml.Keywords) == 0 {
		oldStr = "^keywords: \\[\\]"
	} else {
		// delete line
		for _, keyword := range postYaml.Keywords {
			delLine := fmt.Sprintf("  - %s", keyword)
			// delLine can not contain like: &
			delLine = strings.Replace(delLine, "&", ".", -1)
			delLine = strings.Replace(delLine, "\"", ".", -1)
			delLine = strings.Replace(delLine, "'", ".", -1)
			delLine = strings.Replace(delLine, "/", ".", -1)

			var cmd string
			if runtime.GOOS == "darwin" {
				cmd = fmt.Sprintf("sed -i '' \"/%s/d\" %s", delLine, filename)
			} else {
				cmd = fmt.Sprintf("sed -i \"/%s/d\" %s", delLine, filename)
			}
			_, err = RunCommand(cmd)
			if err != nil {
				logger.Warnf("run [%s] occur err: %s", cmd, err.Error())
				return err
			}
		}
		oldStr = fmt.Sprintf("^keywords:")
	}

	newKeywords := postYaml.Keywords
	maxKeywordLength := 60
	for _, keyword := range keywords {
		if IsContain(postYaml.Tags, keyword) || IsContain(newKeywords, keyword) {
			continue
		}
		if len(keyword) >= maxKeywordLength {
			logger.Warnf("skip long keyword (more than %d): %s", maxKeywordLength, keyword)
			continue
		}

		// replace some key to empty
		keyword = strings.Replace(keyword, "\"", "", -1)
		keyword = strings.Replace(keyword, ":", "", -1)
		keyword = strings.Replace(keyword, "#", "", -1)
		keyword = strings.Replace(keyword, "_", "-", -1)
		keyword = strings.Replace(keyword, "[", "-", -1)
		keyword = strings.Replace(keyword, "]", "-", -1)

		newKeywords = append(newKeywords, keyword)
	}

	if len(newKeywords) == 0 {
		return nil
	}
	newStr := fmt.Sprintf("keywords:\\n  - %s", strings.Join(newKeywords, "\\n  - "))

	var cmd string
	if runtime.GOOS == "darwin" {
		cmd = fmt.Sprintf("sed -i '' \"s#%s#%s#g\" %s", oldStr, newStr, filename)
	} else {
		cmd = fmt.Sprintf("sed -i \"s#%s#%s#g\" %s", oldStr, newStr, filename)
	}

	_, err = RunCommand(cmd)
	if err != nil {
		logger.Warnf("run [%s] occur err: %s", cmd, err.Error())
		return err
	}

	return nil
}

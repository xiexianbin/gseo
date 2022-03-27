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
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"runtime"
	"strings"
)

type PostYaml struct {
	Keywords []string `json:"keywords"`
	Tags     []string `json:"tags"`
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

func UpdateKeywords(filename string, newKeywords []string) error {
	postYaml, err := ParsePostKeysAndTags(filename)
	if err != nil {
		return err
	}

	if len(newKeywords) == 0 {
		return nil
	}

	var oldStr string
	if len(postYaml.Keywords) == 0 {
		oldStr = "keywords: \\[\\]"
	} else {
		oldStr = "keywords:\\n  - " + strings.Join(postYaml.Keywords, "\\n  - ")
	}

	newStr := "keywords:\\n  - " + strings.Join(newKeywords, "\\n  - ")

	var cmd string
	if runtime.GOOS == "darwin" {
		cmd = fmt.Sprintf("sed -i '' \"s#%s#%s#g\" %s", oldStr, newStr, filename)
	} else {
		cmd = fmt.Sprintf("sed -i \"s#%s#%s#g\" %s", oldStr, newStr, filename)
	}

	_, err = RunCommand(cmd)
	if err != nil {
		return err
	}

	return nil
}

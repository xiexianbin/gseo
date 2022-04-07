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
	"bufio"
	"crypto/rand"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"sort"
	"strings"
	"time"
)

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

func ReadFromCmd(tips string) (string, error) {
	fmt.Printf("%s", tips)
	reader := bufio.NewReader(os.Stdin)
	cmdStr, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	} else {
		cmdStr = strings.TrimSuffix(cmdStr, "\n")
		return cmdStr, nil
	}
}

func EId() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	b8 := fmt.Sprintf("%x%x",
		b[4:6], b[6:8])
	return b8
}

func GetHome() string {
	home, _ := homedir.Dir()
	return home
}

func GetCacheFile() string {
	return fmt.Sprintf("%s/%s/cache-%s.json", GetHome(), DefaultConfigSubDir, TodayDate())
}

func TodayDate() string {
	return time.Now().Format("2006-01-02")
}

// LastDate last x day date
func LastDate(last int) string {
	return time.Now().AddDate(0, 0, 0-last).Format("2006-01-02")
}

func SortMap(m map[string]float64) []string {
	var r []string
	t := make(map[float64]string)
	for k, v := range m {
		t[v] = k
	}

	var keys []float64
	for k := range t {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	for _, k := range keys {
		r = append(r, t[k])
	}

	return r
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

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

package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version represents the Gobo build version.
type Version struct {
	// Major and minor version.
	Number float32

	// Increment this for bug releases
	PatchVersion int

	// HugoVersionSuffix is the suffix used in the Hugo version string.
	// It will be blank for release versions.
	Suffix string
}

func (v Version) String() string {
	return version(v.Number, v.PatchVersion, v.Suffix)
}

// CurrentVersion represents the current build version.
// This should be the only one.
var CurrentVersion = Version{
	Number:       1.0,
	PatchVersion: 1,
	Suffix:       "",
}

func version(version float32, patchVersion int, suffix string) string {
	if patchVersion > 0 {
		return fmt.Sprintf("%.2f.%d%s", version, patchVersion, suffix)
	}
	return fmt.Sprintf("%.2f%s", version, suffix)
}

// BuildVersionString creates a version string. This is what you see when
// running "gseo version".
func BuildVersionString() string {
	program := "gseo"

	version := "v" + CurrentVersion.String()

	osArch := runtime.GOOS + "/" + runtime.GOARCH

	return fmt.Sprintf("%s %s %s", program, version, osArch)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gseo.",
	Long:  "All software has versions. This is GSEO's.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(BuildVersionString())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

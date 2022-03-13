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
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/xiexianbin/gseo/utils"
)

var force bool
var clientID string
var clientSecret string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init gseo configure",
	Long:  "init gseo configure.",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		err := viper.ReadInConfig()
		if err != nil || force {
			viper.Set("name", "gseo")

			// set gseo config file
			cfgFile, err := utils.ReadFromCmd(fmt.Sprintf("Please enter gseo config path (default is %s/.gseo.yaml): ", home))
			if err != nil {
				_ = fmt.Errorf("Read gseo config file path err: %s.\n", err)
				os.Exit(1)
			}
			if cfgFile == "" {
				cfgFile = fmt.Sprintf("%s%c.gseo.yaml", home, os.PathSeparator)
			}

			_, err = os.Create(cfgFile)
			if err != nil {
				_ = fmt.Errorf("create config file %s err: %s.\n", cfgFile, err.Error())
				os.Exit(1)
			}
			viper.SetConfigFile(cfgFile)

			viper.Set("client_id", clientID)
			viper.Set("client_secret", clientSecret)


			// set gseo cache dir
			cacheDir, err := utils.ReadFromCmd(fmt.Sprintf("Please enter gseo cache dir (default is %s/.gseo/): ", home))
			if err != nil {
				_ = fmt.Errorf("Read gseo cache dir path err: %s.\n", err)
				os.Exit(1)
			}
			if cacheDir == "" {
				cacheDir = fmt.Sprintf("%s%c.gseo", home, os.PathSeparator)
			}

			err = os.MkdirAll(cacheDir, os.ModePerm)
			if err != nil {
				_ = fmt.Errorf("create config dir %s err: %s.\n", cfgFile, err.Error())
				os.Exit(1)
			}
			viper.Set("cache_dir", cacheDir)

			// write config to file
			_ = viper.MergeInConfig()
			err = viper.WriteConfig()
			if err == nil {
				fmt.Println("init config success!")
			} else {
				_ = fmt.Errorf("init config error: %v", err.Error())
			}
		} else {
			fmt.Println("gseo config is already init, if you want to re-init use `--force` flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")
	initCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Force re-init config.")
	initCmd.PersistentFlags().StringVarP(&clientID, "client-id", "i", "", "google api client id.")
	initCmd.PersistentFlags().StringVarP(&clientSecret, "client-secret", "s", "", "google api client secret.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

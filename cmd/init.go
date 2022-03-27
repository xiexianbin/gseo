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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/xiexianbin/gseo/googleapi"
	"github.com/xiexianbin/gseo/utils"
)

var force bool
var clientSecret string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init gseo configure",
	Long:  "init gseo configure.",
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName(utils.DefaultCfgFileName)
		viper.SetConfigType("yaml")
		viper.Set("name", utils.Name)
		viper.AddConfigPath(utils.DefaultConfigPath)
		err := viper.ReadInConfig()
		if err != nil || force {
			confDir := fmt.Sprintf("%s%c%s", utils.GetHome(), os.PathSeparator, utils.DefaultConfigSubDir)
			err = os.MkdirAll(confDir, os.ModePerm)
			if err != nil {
				_ = fmt.Errorf("create config dir %s err: %s.\n", confDir, err.Error())
				os.Exit(1)
			}
			viper.Set("conf_dir", confDir)

			// set google client secret
			if clientSecret == "" {
				clientSecret, err = utils.ReadFromCmd(fmt.Sprintf("Please enter Google API client_secret.json path: "))
				if err != nil {
					fmt.Println("Read Google API client_secret.json path err:", err)
					os.Exit(1)
				}
			}
			viper.Set("client_secret", clientSecret)

			cfgFile := fmt.Sprintf("%s/%s", confDir, utils.DefaultCfgFileName)
			_, err = os.Create(cfgFile)
			if err != nil {
				_ = fmt.Errorf("create config file %s err: %s.\n", cfgFile, err.Error())
				os.Exit(1)
			}

			// write config to file
			_ = viper.MergeInConfig()
			err = viper.WriteConfig()
			if err == nil {
				fmt.Println("init config success!")
			} else {
				fmt.Println("init config error:", err.Error())
				os.Exit(1)
			}
		} else {
			fmt.Println("gseo config is already init, if you want to re-init use `--force` flag")
		}

		// init google token
		_, _ = googleapi.Client()
		fmt.Println("init Google API OAuth2.0 token success!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")
	initCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Force re-init config.")
	initCmd.PersistentFlags().StringVarP(&clientSecret, "client-secret", "s", "", "google api client secret json path.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

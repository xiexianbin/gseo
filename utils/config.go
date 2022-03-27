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
	"github.com/spf13/viper"
)

func GetTokenPath() string {
	viper.SetConfigName(DefaultCfgFileName)
	viper.SetConfigType("yaml")
	viper.Set("name", Name)
	viper.AddConfigPath("$HOME/.gseo")
	_ = viper.ReadInConfig()
	return fmt.Sprintf("%s/%s", viper.Get("conf_dir"), DefaultTokenFileName)
}

func GetClientSecretPath() string {
	viper.SetConfigName(DefaultCfgFileName)
	viper.SetConfigType("yaml")
	viper.Set("name", Name)
	viper.AddConfigPath("$HOME/.gseo")
	_ = viper.ReadInConfig()
	return fmt.Sprintf("%s", viper.Get("client_secret"))
}

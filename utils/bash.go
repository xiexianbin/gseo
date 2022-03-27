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
	"os/exec"
	"runtime"
)

// RunCommand is a shared function used by check and reload
// to run the given command.
// It returns nil if the given cmd returns 0.
// The command can be run on unix and windows.
func RunCommand(cmd string) (string, error) {
    var c *exec.Cmd
    if runtime.GOOS == "windows" {
        c = exec.Command("cmd", "/C", cmd)
    } else {
        c = exec.Command("/bin/sh", "-c", cmd)
    }

    output, err := c.CombinedOutput()
    outStr := fmt.Sprintf("%s", string(output))
    return outStr, err
}

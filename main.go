// Copyright 2016 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/palantir/godel/framework/pluginapi/v2/pluginapi"
	"github.com/palantir/pkg/cobracli"

	"github.com/palantir/godel-test-plugin/cmd"
	"github.com/palantir/godel-test-plugin/gojunit/generated_src"
	"github.com/palantir/godel-test-plugin/testplugin"
)

func main() {
	if ok := pluginapi.InfoCmd(os.Args, os.Stdout, cmd.PluginInfo); ok {
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "__"+testplugin.GoJUnitReport {
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
		amalgomated.Instance().Run(testplugin.GoJUnitReport)
		return
	}
	os.Exit(cobracli.ExecuteWithDefaultParamsWithVersion(cmd.RootCmd, &cmd.DebugFlagVal, ""))
}

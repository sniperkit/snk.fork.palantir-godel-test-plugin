/*
Sniperkit-Bot
- Status: analyzed
*/

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

package v0

import (
	"github.com/palantir/pkg/matcher"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	// Tags group tests into different sets. The key is the name of the tag and the value is a
	// matcher.NamesPathsWithExcludeCfg that specifies the rules for matching the tests that are part of the tag.
	// Any test that matches the provided matcher is considered part of the tag.
	Tags map[string]matcher.NamesPathsWithExcludeCfg `yaml:"tags,omitempty"`

	// Exclude specifies the files that should be excluded from tests.
	Exclude matcher.NamesPathsCfg `yaml:"exclude,omitempty"`
}

func UpgradeConfig(cfgBytes []byte) ([]byte, error) {
	var cfg Config
	if err := yaml.UnmarshalStrict(cfgBytes, &cfg); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal test-plugin v0 configuration")
	}
	return cfgBytes, nil
}

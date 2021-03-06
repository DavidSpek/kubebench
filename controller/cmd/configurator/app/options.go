// Copyright 2018 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"flag"
)

type AppOption struct {
	InputParams string
	OutputFile  string
}

func NewAppOption() *AppOption {
	opt := AppOption{}
	return &opt
}

func (opt *AppOption) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&opt.InputParams, "input-params", "", "Structured spec for the configurator.")
	fs.StringVar(&opt.OutputFile, "output-file", "", "Path to the output file.")
}

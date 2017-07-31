// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"testing"

	"github.com/mendersoftware/go-lib-micro/log"
)

var runAcceptanceTests bool

// used for parsing '-cli-args' for urfave/cli when running acceptance tests
// this is because of a conflict between urfave/cli and regular go flags required for testing (can't mix the two)
var cliArgsRaw string

func init() {
	// disable logging thile running unit tests
	// default application settup couses to mich noice
	log.Log.Out = ioutil.Discard

	flag.BoolVar(&runAcceptanceTests, "acceptance-tests", false, "set flag when running acceptance tests")
	flag.StringVar(&cliArgsRaw, "cli-args", "", "for passing urfave/cli args (single string) when golang flags are specified (avoids conflict)")
	flag.Parse()
}

func TestRunMain(t *testing.T) {
	if !runAcceptanceTests {
		t.Skip()
	}

	// parse '-cli-args', remember about binary name at idx 0
	var cliArgs []string

	if cliArgsRaw != "" {
		cliArgs = []string{os.Args[0]}
		splitArgs := strings.Split(cliArgsRaw, " ")
		cliArgs = append(cliArgs, splitArgs...)
	}

	go doMain(cliArgs)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
}

/*
 *  Copyright (c) 2022 AlertAvert.com.  All rights reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  Author: Marco Massenzio (marco@alertavert.com)
 */

package main

import (
	"flag"
	"github.com/massenz/slf4go/logging"
)

func main() {
	logging.RootLog.Info("Program started - before any logs are available")
	var log = logging.NewLog("example")
	defer log.Trace("Exiting")

	trace := flag.Bool("trace", false, "If set it will emit trace logs")
	flag.Parse()
	if *trace {
		log.Level = logging.TRACE
	}
	log.Info("An INFO message")
	log.Debug("This will NOT be logged, unless -trace is given")

	if !*trace {
		log.Level = logging.DEBUG
	}
	log.Debug("This WILL be printed out")
	log.Info("The `trace` on exit will only be visible with the -trace option")

	logging.NullLog.Error("No one will ever see this, like, ever 😤")
}

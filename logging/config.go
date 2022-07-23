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

package logging

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
)

const (
	LogConfigFile       = "config.yaml"
	LogConfigDir        = ".slf4go"
	LogConfigDirEnvVar  = "SLF4GO_CONFIG_DIR"
	LogConfigFileEnvVar = "SLF4GO_CONFIG_FILE"
)

type LoggerConfig struct {
	Level LogLevel `yaml:"level"`
	// TODO: a placeholder for now, only uses `stdout`, will be used
	// 		 eventually to configure a file writer.
	Writer string `yaml:"writer"`
}

type LogConfig struct {
	Loggers map[string]LoggerConfig
}

// DefaultConfig is not really a var, but Go is not smart enough to figure out it's a constant
var (
	DefaultConfig = LogConfig{
		Loggers: map[string]LoggerConfig{
			"default": {Level: INFO, Writer: "stdout"},
		},
	}
	LoggersConfiguration LogConfig
)

func init() {
	var (
		configDir  string
		configFile string
	)

	configDir, found := os.LookupEnv(LogConfigDirEnvVar)
	if !found {
		// FIXME: should use os.PathSeparator, but it should not be so hard.
		configDir = os.Getenv("HOME") + "/" + LogConfigDir
	}
	configFile, found = os.LookupEnv(LogConfigFileEnvVar)
	if !found {
		configFile = LogConfigFile
	}
	configPath := configDir + "/" + configFile
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		RootLog.Warn("config file %s does not exist, using defaults", configPath)
		LoggersConfiguration = DefaultConfig
		return
	}
	cfg, err := os.Open(configPath)
	if err != nil {
		RootLog.Fatal(err)
	}
	if err = yaml.NewDecoder(cfg).Decode(&LoggersConfiguration); err != nil {
		RootLog.Fatal(err)
	}
}

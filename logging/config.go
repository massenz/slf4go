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
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

const (
	LogConfigFile       = "config.yaml"
	LogConfigDir        = ".slf4go"
	LogConfigDirEnvVar  = "SLF4GO_CONFIG_DIR"
	LogConfigFileEnvVar = "SLF4GO_CONFIG_FILE"

	DefaultLoggerName   = "default"
	DefaultLoggerLevel  = "INFO"
	DefaultLoggerWriter = "stderr"
)

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Writer string `yaml:"writer"`
}

type LogConfig struct {
	Loggers map[string]LoggerConfig `yaml:"loggers"`
}

var (
	// DefaultLoggerConfigurations is not really a var, but Go is not smart enough
	// to figure out it's a constant
	DefaultLoggerConfigurations           = LoggerConfig{Level: DefaultLoggerLevel, Writer: DefaultLoggerWriter}
	LoggersConfiguration        LogConfig = LogConfig{
		// The default logger configuration is always available, and used as
		// a fallback if no other configuration is found.
		Loggers: map[string]LoggerConfig{
			DefaultLoggerName: DefaultLoggerConfigurations,
		},
	}
)

func FindFileWithFallback(dir string, file string) (*os.File, error) {
	// First, we look in the current directory for `file`.
	if _, err := os.Stat(file); err == nil {
		return os.Open(file)
	}

	// If not found, we look in the `dir` directory.
	configPath := dir + "/" + file
	if _, err := os.Stat(configPath); err == nil {
		return os.Open(configPath)
	}
	// TODO: we could add more "fallback" locations.
	return nil, os.ErrNotExist
}

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
	// Regardless of the configuration file, we always initialize the RootLog
	defer func() {
		RootLog = NewLog("root")
	}()

	cfg, err := FindFileWithFallback(configDir, configFile)
	if errors.Is(err, os.ErrNotExist) {
		return
	} else if err != nil {
		panic(err)
	}
	if err = yaml.NewDecoder(cfg).Decode(&LoggersConfiguration); err != nil {
		panic(err)
	}
}

// GetLoggerConfig returns the configuration for the given logger; if not specified,
// it will return the default configuration.
func GetLoggerConfig(name string) LoggerConfig {
	if cfg, found := LoggersConfiguration.Loggers[name]; found {
		return cfg
	}
	return LoggersConfiguration.Loggers[DefaultLoggerName]
}

// UnmarshalLevel converts a string to a Level.
// We cannot simply implement the UnmarshalText method,
// because `LogLevel` (`uint8`)is not considered to be "local"
func UnmarshalLevel(text string) LogLevel {
	switch strings.ToUpper(text) {
	case "TRACE":
		return TRACE
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "NONE":
		return NONE
	default:
		if RootLog != nil {
			RootLog.Error("invalid log level: %s", text)
		}
		return DefaultLevel
	}
}

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
    "fmt"
    "io"
    "log"
    "os"
)

const (
    TRACE = iota
    DEBUG
    INFO
    WARN
    ERROR
    NONE // To entirely disable logging for the Logger

    DefaultLevel = INFO
    DefaultFlags = log.Lmsgprefix | log.Ltime | log.Ldate | log.Lshortfile
)

type LogLevel = int8

type Log struct {
    *log.Logger
    Level LogLevel
    Name  string
}

func (l *Log) shouldLog(level LogLevel) bool {
    return (l.Level != NONE) && (l.Level <= level)
}

func (l *Log) Trace(format string, v ...interface{}) {
    if l.shouldLog(TRACE) {
        format = l.Name + "[TRACE] " + format
        l.Output(2, fmt.Sprintf(format, v...))
    }
}

func (l *Log) Debug(format string, v ...interface{}) {
    if l.shouldLog(DEBUG) {
        format = l.Name + "[DEBUG] " + format
        l.Output(2, fmt.Sprintf(format, v...))
    }
}

func (l *Log) Info(format string, v ...interface{}) {
    if l.shouldLog(INFO) {
        format = l.Name + "[INFO] " + format
        l.Output(2, fmt.Sprintf(format, v...))
    }
}

func (l *Log) Warn(format string, v ...interface{}) {
    if l.shouldLog(WARN) {
        format = l.Name + "[WARN] " + format
        l.Output(2, fmt.Sprintf(format, v...))
    }
}

func (l *Log) Error(format string, v ...interface{}) {
    if l.shouldLog(ERROR) {
        format = l.Name + "[ERROR] " + format
        l.Output(2, fmt.Sprintf(format, v...))
    }
}

func (l *Log) Fatal(err error) {
    l.Output(2, fmt.Sprintf("[FATAL] %s", err.Error()))
    os.Exit(1)
}

func NewLog(name string) *Log {
    var writer io.Writer
    config := GetLoggerConfig(name)
    if config.Writer == "console" {
        writer = os.Stdout
    } else if config.Writer == "stderr" || config.Writer == "" {
        writer = os.Stderr
    } else {
        f, err := os.OpenFile(config.Writer, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatalf("Failed to open log file %s: %s", config.Writer, err)
        }
        writer = f
    }
    return &Log{
        Name:   name,
        Logger: log.New(writer, "", DefaultFlags),
        Level:  config.Level,
    }
}

// RootLog is the default log, it needs to be initialized after the init() function,
// so that configuration (if any) can be picked up.
var RootLog *Log

// A Loggable type is one that has a Log and exposes it to its clients
type Loggable interface {
    SetLogLevel(level LogLevel)
}

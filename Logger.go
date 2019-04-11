//
// Copyright © 2018 Anton Filatov
//
// This file is part of GoLogging project.
//
// GoLogging project is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// GoLogging project is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with GoLogging project. If not, see <https://www.gnu.org/licenses/lgpl-3.0>.
//

package logging

import (
	"fmt"
	"io"
)

// Logger is a common logging API definition
type Logger interface {
	Log(level Level, arg ...interface{})
	Logf(level Level, pattern string, arg ...interface{})
	LoggerProvider
}

// LoggerProvider can create Logger objects
type LoggerProvider interface {
	NewLogger(name ...string) (Logger, error)
}

// Factory is a basic logger constructor
type Factory interface {
	LoggerProvider
	io.Closer
}

// Level represents logs output threshold
type Level uint8

// Logging levels from least informative to most informative
const (
	OFF = Level(iota)
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
	TRACEE
	TRACEEE
	ALL = ^Level(0)
)

func (l Level) String() string {
	switch l {
	case OFF:
		return "OFF"
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case TRACEE:
		return "TRACEE"
	case TRACEEE:
		return "TRACEEE"
	case ALL:
		return "ALL"
	default:
		return fmt.Sprintf("UNKNOWN_LEVEL_%d", l)
	}
}

type logger struct {
	logger Logger
}

func (l *logger) Log(level Level, arg ...interface{}) {
	l.logger.Log(level, arg...)
}

func (l *logger) Logf(level Level, pattern string, arg ...interface{}) {
	l.logger.Logf(level, pattern, arg...)
}

func (l *logger) NewLogger(name ...string) (Logger, error) {
	return NewLogger(l.logger.NewLogger(name...))
}

// NewLogger creates a logger wrapper
func NewLogger(loggerBase Logger, err error) (Logger, error) {
	if err != nil {
		return nil, err
	}
	// Check if logger is wrapped by us
	if loggerWrap, ok := loggerBase.(*logger); ok {
		return &logger{
			logger: loggerWrap.logger,
		}, err
	}
	// Assuming loggerBase is not wrapped
	return &logger{
		logger: loggerBase,
	}, err
}
